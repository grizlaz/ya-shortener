package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/grizlaz/ya-shortener/internal/audit"
	"github.com/grizlaz/ya-shortener/internal/config"
	"github.com/grizlaz/ya-shortener/internal/handler"
	"github.com/grizlaz/ya-shortener/internal/logger"
	"github.com/grizlaz/ya-shortener/internal/repository"
	"github.com/grizlaz/ya-shortener/internal/service"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func main() {
	setDefaultBuildInfo()
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	var err error
	if err = logger.Initialize("info"); err != nil {
		panic(err)
	}
	cfg := config.Get()

	var shorteningStorage service.Storage
	var db *sql.DB
	if cfg.DatabaseDSN != "" {
		db, err = sql.Open("pgx", cfg.DatabaseDSN)
		if err != nil {
			logger.Log.Sugar().Fatalf("error init db: %w", err)
		}
		defer func() {
			if derr := db.Close(); derr != nil {
				logger.Log.Sugar().Errorf("error defer db.Close(): %v", derr)
			}
		}()
		shorteningStorage, err = repository.NewPostgresDB(db)
	} else {
		shorteningStorage, err = repository.NewInFile(cfg.FileStoragePath, false)
	}
	if err != nil {
		logger.Log.Sugar().Fatalf("error init storage: %w", err)
	}
	auditService, err := initAudit()
	if err != nil {
		logger.Log.Sugar().Fatalf("error init audit service: %w", err)
	}

	shortener := service.NewService(context.Background(), shorteningStorage)
	srv := handler.NewServer(shortener, cfg.BaseURL, db, auditService)
	if cfg.EnableHTTPS {
		cert, key, err := generateCrt()
		if err != nil {
			logger.Log.Sugar().Fatalf("error creating cert: %v", err)
		}
		defer os.Remove(cert)
		defer os.Remove(key)
		err = http.ListenAndServeTLS(cfg.ServerAddress, cert, key, srv)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Sugar().Fatalf("error running server: %v", err)
		}
	} else {
		err = http.ListenAndServe(cfg.ServerAddress, srv)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Sugar().Fatalf("error running server: %v", err)
		}
	}
}

func generateCrt() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", err
	}
	publicKey := privateKey.Public()
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"iter24"},
			Country:      []string{"RU"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, publicKey, privateKey)
	if err != nil {
		return "", "", err
	}

	var certPEM bytes.Buffer
	err = pem.Encode(&certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		return "", "", err
	}

	var privateKeyPEM bytes.Buffer
	err = pem.Encode(&privateKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return "", "", err
	}

	tmp, err := os.MkdirTemp("", "tmp")
	if err != nil {
		return "", "", err
	}

	crtName := filepath.Join(tmp, "cert.pem")
	if err = os.WriteFile(crtName, certPEM.Bytes(), 0644); err != nil {
		return "", "", err
	}
	keyName := filepath.Join(tmp, "key.pem")
	if err = os.WriteFile(keyName, privateKeyPEM.Bytes(), 0644); err != nil {
		return "", "", err
	}

	return crtName, keyName, nil
}

func setDefaultBuildInfo() {
	if buildVersion == "" {
		buildVersion = "N/A"
	}
	if buildDate == "" {
		buildDate = "N/A"
	}
	if buildCommit == "" {
		buildCommit = "N/A"
	}
}

func initAudit() (*audit.Audit, error) {
	cfg := config.Get()
	auditService := audit.NewAudit()
	if cfg.AuditFilePath != "" {
		fileStorage, err := repository.NewInFile(cfg.AuditFilePath, true)
		if err != nil {
			return nil, err
		}
		auditInFile := audit.NewObserver("file", fileStorage)
		auditService.Register(auditInFile)
	}
	if cfg.AuditURL != "" {
		auditClient := repository.NewAuditClient(cfg.AuditURL)
		httpAudit := audit.NewObserver("http", auditClient)
		auditService.Register(httpAudit)
	}
	return auditService, nil
}
