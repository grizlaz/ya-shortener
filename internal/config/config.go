package config

import (
	"errors"
	"flag"
	"os"
	"strings"
	"sync"
	"time"
)

type config struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseDSN     string
	TokenExp        time.Duration
	SecretKey       []byte
	AuditFilePath   string
	AuditURL        string
	EnableHTTPS     bool
}

var (
	cfg  config
	once sync.Once
)

func Get() config {
	once.Do(func() {
		cfg.SecretKey = []byte("supersecretkey")
		cfg.TokenExp = time.Hour * 3
		cfg.BaseURL = "http://localhost:8080"
		flag.StringVar(&cfg.ServerAddress, "a", ":8080", "address and port to run server")
		flag.Func("b", `address and port before short url (default "http://localhost:8080")`, func(s string) error {
			if err := checkBaseURL(s); err != nil {
				return err
			}
			cfg.BaseURL = s
			return nil
		})
		flag.StringVar(&cfg.FileStoragePath, "f", "storage.txt", "storage path")
		flag.StringVar(&cfg.DatabaseDSN, "d", "", "DSN for db")
		flag.StringVar(&cfg.AuditFilePath, "audit-file", "", "path to audit file")
		flag.StringVar(&cfg.AuditURL, "audit-url", "", "audit url")
		flag.BoolVar(&cfg.EnableHTTPS, "s", false, "enable tls")

		flag.Parse()
		if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
			if err := checkBaseURL(envBaseURL); err != nil {
				panic(err)
			}
			cfg.BaseURL = envBaseURL
		}
		if envServerAddress := os.Getenv("SERVER_ADDRESS"); envServerAddress != "" {
			cfg.ServerAddress = envServerAddress
		}
		if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
			cfg.FileStoragePath = envFileStoragePath
		}
		if envDBDSN := os.Getenv("DATABASE_DSN"); envDBDSN != "" {
			cfg.DatabaseDSN = envDBDSN
		}
		if envAuditFilePath := os.Getenv("AUDIT_FILE"); envAuditFilePath != "" {
			cfg.AuditFilePath = envAuditFilePath
		}
		if envAuditURL := os.Getenv("AUDIT_URL"); envAuditURL != "" {
			cfg.AuditURL = envAuditURL
		}
		if envEnableHTTPS := os.Getenv("ENABLE_HTTPS"); envEnableHTTPS != "" {
			cfg.EnableHTTPS = true
		}
	})
	return cfg
}

func checkBaseURL(url string) error {
	if !strings.HasPrefix(url, "http") {
		return errors.New("empty protocol for base url")
	}
	return nil
}
