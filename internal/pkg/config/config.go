package config

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

//Config contains configuration for db connection
type Postgres struct {
	Host     string `env:"HOST,default=localhost"`
	Database string `env:"DATABASE,required"`
	Username string `env:"USERNAME,required"`
	Password string `env:"PASSWORD,required"`
	Port     uint16 `env:"PORT,default=5432"`
}

//Config represents configuration of the application
type Config struct {
	//http server will start on this port
	Port     uint16    `env:"PORT,default=8080"`
	Postgres *Postgres `env:",prefix=PG_"`
}

//NewConfig will parse configuration from env variables or .env file
func NewConfig(ctx context.Context, logger log.Logger) (Config, error) {
	file := flag.String("c", "", "load config from .env file")
	flag.Parse()

	if *file != "" {
		if err := godotenv.Load(*file); err != nil {
			return Config{}, fmt.Errorf("load config from file: %w", err)
		}
	}

	var c = Config{
		Postgres: &Postgres{},
	}

	if err := envconfig.Process(ctx, &c); err != nil {
		return Config{}, fmt.Errorf("env config parsing: %w", err)
	}

	return c, nil
}
