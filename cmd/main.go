package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"tourmanager/app/api"

	"tourmanager/config"

	"github.com/hashicorp/go-multierror"
	"github.com/joho/godotenv"
)

// @title Go Hexagonal API
// @description Powered by scv-go-tools - https://github.com/sergicanet9/scv-go-tools

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {

	var opts struct {
		Version     string `long:"ver" description:"Version" required:"true"`
		Environment string `long:"env"` // description:"environments" choices:"local" choices:"dev" requireds:"true"`
		Port        int    `long:"port" description:"Running port" required:"true"`
		Database    string `long:"db"` //  description:"the database adapter to use" choice:"mongo" choice:"postgres" required:"true"`
		DSN         string `long:"dsn" description:"DSN of the selected database" required:"true"`
	}
	env := os.Getenv("ENVIRONMENT")

	if env == "" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	opts.Version = os.Getenv("VERSION")
	opts.Environment = os.Getenv("ENVIRONMENT")
	opts.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	opts.Database = os.Getenv("DATABASE")
	opts.DSN = os.Getenv("DSN")

	cfg, err := config.ReadConfig(opts.Version, opts.Environment, opts.Port, opts.Database, opts.DSN) // , "config")
	if err != nil {
		log.Fatal(fmt.Errorf("no se puede analizar el archivo de configuración ENV %s: %w", opts.Environment, err))
	}

	var g multierror.Group
	ctx, cancel := context.WithCancel(context.Background())

	a := api.New(ctx, cfg)
	g.Go(a.Run(ctx, cancel))

	/*
		if cfg.Async.Run {
			async := async.New(cfg)
			g.Go(async.Run(ctx, cancel))
		}
	*/
	if err := g.Wait().ErrorOrNil(); err != nil {
		log.Fatal(err)
	}
}
