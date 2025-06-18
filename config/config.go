package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HOST   string
	DNS    string
	SECRET string
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("couldnt load env")
	}
}

func SetupEnv() (Config, error) {
	host := os.Getenv("HOST")
	if len(host) <= 0 {
		return Config{}, errors.New("couldnt load host")
	}

	dns := os.Getenv("DNS")
	if len(dns) <= 0 {
		return Config{}, errors.New("couldnt load dns")
	}

	secret := os.Getenv("SECRET")
	if len(secret) <= 0 {
		return Config{}, errors.New("couldnt load secret")
	}

	return Config{HOST: host, DNS: dns, SECRET: secret}, nil
}
