package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Port                  string `mapstructure:"PORT"`
	MongoConnectionString string `mapstructure:"MONGO_CONNECTION_STRING"`
	JWTSecret             string `mapstructure:"JWT_SECRET"`
	SquareUpUrl           string `mapstructure:"SQUARE_UP_BASE_URL"`
	SquareUpToken         string `mapstructure:"SQUARE_UP_TOKEN"`
}

var Env *Config

func Load() {
	Env = load()
}

func setDefaults() {
	viper.SetDefault("PORT", "9001")
	viper.SetDefault("SQUARE_UP_BASE_URL", "https://connect.squareupsandbox.com/v2")
}

func load() *Config {
	viper.AutomaticEnv()

	//viper.SetConfigFile(".env")
	//_ = viper.ReadInConfig()

	setDefaults()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}
	cfg.MongoConnectionString = os.Getenv("MONGO_CONNECTION_STRING")
	return &cfg
}
