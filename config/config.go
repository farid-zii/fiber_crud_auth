package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUser 		string
	DBPassword 	string
	DBName	 	string
	DBHost	 	string
}

func GetConfig() *Config {
	return &Config{
		DBUser:		os.Getenv("DB_USER"),
		DBPassword:	os.Getenv("DB_PASSWORD"),
		DBName:		os.Getenv("DB_NAME"),
		DBHost:		os.Getenv("DB_HOST"),
	}
}

func (c *Config) GetDBConnectionString()string{
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=local",c.DBUser,c.DBPassword,c.DBHost,c.DBName)
}