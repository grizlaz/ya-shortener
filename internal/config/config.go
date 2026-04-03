package config

import (
	"errors"
	"flag"
	"os"
	"strings"
	"sync"
)

type config struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseDSN     string
}

var (
	cfg  config
	once sync.Once
)

func Get() config {
	once.Do(func() {
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
	})
	return cfg
}

func checkBaseURL(url string) error {
	if !strings.HasPrefix(url, "http") {
		return errors.New("empty protocol for base url")
	}
	return nil
}
