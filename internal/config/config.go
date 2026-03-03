package config

import (
	"errors"
	"flag"
	"strings"
	"sync"
)

type config struct {
	ServerAddress string
	BaseURL       string
}

var (
	cfg  config
	once sync.Once
)

func Get() config {
	once.Do(func() {
		cfg.BaseURL = "http://localhost:8080"
		flag.StringVar(&cfg.ServerAddress, "a", ":8080", "address and port to run server")
		// flag.StringVar(&cfg.B, "b", "http://localhost:8080", "address and port before short url")
		flag.Func("b", `address and port before short url (default "http://localhost:8080")`, func(s string) error {
			if !strings.HasPrefix(s, "http") {
				return errors.New("empty protocol")
			}
			cfg.BaseURL = s
			return nil
		})

		flag.Parse()
	})
	return cfg
}
