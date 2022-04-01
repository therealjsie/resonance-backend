package internals

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func ReadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	config := Config{
		Host:     viper.GetString("postgresql.host"),
		Port:     viper.GetInt("postgresql.port"),
		Username: viper.GetString("postgresql.username"),
		Password: viper.GetString("postgresql.password"),
		Database: viper.GetString(("postgresql.database"))}
	return config
}

func (c Config) DatabaseConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Username, c.Password, c.Host, strconv.Itoa(c.Port), c.Database)
}
