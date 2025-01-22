package config

import "os"

type Config struct {
    Server struct {
        Port string
    }
    Database struct {
        DSN string
    }
}

func Load() *Config {
    cfg := &Config{}
    cfg.Server.Port = os.Getenv("SERVER_PORT")
    cfg.Database.DSN = os.Getenv("DATABASE_DSN")
    return cfg
}
