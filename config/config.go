package config

import (
	"fmt"

	"github.com/antoniomarfa/hexatools/api/utils"
)

type Async struct {
	Run      bool
	Interval utils.Duration
}

type Config struct {
	// set in flags
	Version     string
	Environment string
	Port        int
	Database    string
	DSN         string
	// set in json config files
	config
}

type config struct {
	PostgresMigrationsDir string
	JWTSecret             string
	Timeout               utils.Duration
	Async                 Async

	// B2 Configuration
	B2KeyID       string
	B2Application string
	B2Bucket      string
	B2Region      string
	B2Endpoint    string
}

// ReadConfig from the project´s JSON config files.
// Default values are specified in the default configuration file, config/config.json
// and can be overrided with values specified in the environment configuration files, config/config.{env}.json.
func ReadConfig(version, env string, port int, database, dsn string) (Config, error) {
	var c Config
	c.Version = version
	c.Environment = env
	c.Port = port
	c.Database = database
	c.DSN = dsn
	// Asignar manualmente los valores a los campos del struct 'config'

	var cfg config
	/*
		configPath := "./config"

		if err := utils.LoadJSON(path.Join(configPath, "config.json"), &cfg); err != nil {
			return c, fmt.Errorf("error parsing configuration, %s", err)
		}

		if err := utils.LoadJSON(path.Join(configPath, "config.local.json"), &cfg); err != nil {
			//		if err := utils.LoadJSON(path.Join(configPath, "config."+env+".json"), &cfg); err != nil {
			return c, fmt.Errorf("error parsing environment configuration, %s", err)
		}
	*/
	configJson := `{
		"Address": "http://localhost",
		"PostgresMigrationsDir": "infrastructure/postgres/migrations",
		"JWTSecret": "CTeemck6Gg",
		"Timeout": "5s",
		"Async": {
			"Run": true,
			"Interval": "2m"
		}
	}`

	configlocaljson := `{
    "Timeout": "2m"
    }`

	if err := utils.ViewJSON(configJson, &cfg); err != nil {
		return c, fmt.Errorf("error parsing configuration, %s", err)
	}

	if err := utils.ViewJSON(configlocaljson, &cfg); err != nil {
		//		if err := utils.LoadJSON(path.Join(configPath, "config."+env+".json"), &cfg); err != nil {
		return c, fmt.Errorf("error parsing environment configuration, %s", err)
	}

	// Asignar la configuración cargada al campo config
	c.config = cfg

	c.config.B2KeyID = "da435c192e07"                                     //os.Getenv("B2_KEY_ID")
	c.config.B2Application = "00541a30c28c056a2403e4a8fb27a4fdcce45331fb" //os.Getenv("B2_APPLICATION_KEY")
	c.config.B2Bucket = "tourmanagerdocument"                             //os.Getenv("B2_BUCKET")
	c.config.B2Region = "us-east-005"                                     //os.Getenv("B2_REGION")
	c.config.B2Endpoint = "https://s3.us-east-005.backblazeb2.com"        //os.Getenv("B2_ENDPOINT")

	return c, nil
}
