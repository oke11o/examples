package env

import (
	"github.com/spf13/viper"
	"os"
)

func init() {
	os.Setenv("TITLE", "app-title")
	os.Setenv("DB_NAME", "database-name")
	os.Setenv("SPF_TITLE", "spf-app-title")
	os.Setenv("SPF_DB_NAME", "spf-database-name")
}

type Config struct {
	Title string
	Db    DbConfig
}

type DbConfig struct {
	Name string
}

func ReadConfigGetEnv() Config {
	cfg := Config{}
	cfg.Title = os.Getenv("TITLE")
	cfg.Db.Name = os.Getenv("DB_NAME")
	return cfg
}

func ReadConfigLookupEnv() Config {
	cfg := Config{}
	if title, ok := os.LookupEnv("TITLE"); ok {
		cfg.Title = title
	}
	if dbName, ok := os.LookupEnv("DB_NAME"); ok {
		cfg.Db.Name = dbName
	}
	return cfg
}

func ReadConfigViper() Config {
	cfg := Config{}

	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("title")
	viper.BindEnv("db_name")

	title := viper.Get("title") // spf-database-name
	cfg.Title = title.(string)

	dbName := viper.Get("db_name") // spf-database-name
	cfg.Db.Name = dbName.(string)

	return cfg
}

func ReadConfigViperMarshal() Config {
	cfg := Config{}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("title")
	viper.BindEnv("db_name")
	viper.Unmarshal(&cfg)
	// cfg.Title == "spf-app-title"
	// cfg.Db.Name == ""

	return cfg
}
