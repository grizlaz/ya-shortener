package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/grizlaz/ya-shortener/internal/logger"
	"go.uber.org/zap"
)

type config struct {
	ServerAddress   string `json:"server_address"`
	BaseURL         string `json:"base_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	TokenExp        time.Duration
	SecretKey       []byte
	AuditFilePath   string
	AuditURL        string
	EnableHTTPS     bool `json:"enable_https"`
	ConfigPath      string
}

var (
	cfg  config
	once sync.Once
)

const baseURLDefault = "http://localhost:8080"
const serverAddressDefault = ":8080"

func Get() config {
	once.Do(func() {
		cfg.SecretKey = []byte("supersecretkey")
		cfg.TokenExp = time.Hour * 3
		cfg.BaseURL = baseURLDefault
		flag.StringVar(&cfg.ServerAddress, "a", serverAddressDefault, "address and port to run server")
		flag.Func("b", fmt.Sprintf("address and port before short url (default \"%s\")", baseURLDefault), func(s string) error {
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
		flag.StringVar(&cfg.ConfigPath, "config", "", "path to json config")
		flag.StringVar(&cfg.ConfigPath, "c", "", "path to json config")
		flag.Parse()
		if cfg.ConfigPath != "" {
			parseConfigFile(cfg.ConfigPath)
		}
		parseEnvVars()
		logger.Log.Debug("config", zap.Any("cfg", cfg))
	})
	return cfg
}

func parseConfigFile(path string) {
	fileCfg := &config{}
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Log.Fatal("error reading config file", zap.Error(err))
	}
	err = json.Unmarshal(data, fileCfg)
	if err != nil {
		logger.Log.Fatal("error parsing config file", zap.Error(err))
	}
	if fileCfg.ServerAddress != "" && cfg.ServerAddress == serverAddressDefault {
		cfg.ServerAddress = fileCfg.ServerAddress
	}
	if fileCfg.BaseURL != "" && cfg.BaseURL == baseURLDefault {
		if err := checkBaseURL(fileCfg.BaseURL); err != nil {
			logger.Log.Fatal("error parsing baseURL from config file", zap.Error(err))
		}
		cfg.BaseURL = fileCfg.BaseURL
	}
	if fileCfg.FileStoragePath != "" && cfg.FileStoragePath == "" {
		cfg.FileStoragePath = fileCfg.FileStoragePath
	}
	if fileCfg.DatabaseDSN != "" && cfg.DatabaseDSN == "" {
		cfg.DatabaseDSN = fileCfg.DatabaseDSN
	}
	if fileCfg.EnableHTTPS && !cfg.EnableHTTPS {
		cfg.EnableHTTPS = fileCfg.EnableHTTPS
	}
}

func parseEnvVars() {
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		if err := checkBaseURL(envBaseURL); err != nil {
			logger.Log.Fatal("error parsing BASE_URL from env", zap.Error(err))
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
		v, err := strconv.ParseBool(envEnableHTTPS)
		if err != nil {
			logger.Log.Fatal("error parsing ENABLE_HTTPS from env", zap.Error(err))
		}
		cfg.EnableHTTPS = v
	}
	if envPathToConfig := os.Getenv("CONFIG"); envPathToConfig != "" {
		cfg.ConfigPath = envPathToConfig
	}
}

func checkBaseURL(url string) error {
	if !strings.HasPrefix(url, "http") {
		return errors.New("empty protocol for base url")
	}
	return nil
}
