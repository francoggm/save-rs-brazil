package config

import "os"

type Config struct {
	Port       string
	Mode       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DB         string
}

func New() *Config {
	return &Config{
		Port:       os.Getenv("PORT"),
		Mode:       os.Getenv("MODE"),
		DBUser:     os.Getenv("DBUSER"),
		DBPassword: os.Getenv("DBPASSWORD"),
		DBHost:     os.Getenv("DBHOST"),
		DBPort:     os.Getenv("DBPORT"),
		DB:         os.Getenv("DB"),
	}
}
