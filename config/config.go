package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	DatabasePSQL struct {
		Host     string `env:"DATABASE_HOST"`
		Port     int    `env:"DATABASE_PORT,default=5432"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		DbName   string `env:"DATABASE_DB_NAME"`
	}
	Server struct {
		HTTPPort int32 `env:"HTTPPORT,default=3000"`
	}
	DataBaseMongo struct {
		Host string `env:"MONGO_HOST"`
		Port int    `env:"MONGO_PORT,default=27017"`
	}
}

var c config

func ReadConfig() error {
	ctx := context.Background()
	err := envconfig.Process(ctx, &c)
	return err
}

func HTTPListener() string {
	return fmt.Sprintf(":%d", c.Server.HTTPPort)
}

func MongoConn() string {
	ReadConfig()
	return fmt.Sprintf("mongodb://%s:%d/",
		c.DataBaseMongo.Host,
		c.DataBaseMongo.Port,
	)
}
