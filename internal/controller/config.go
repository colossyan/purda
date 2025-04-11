package controller

import "os"

type Config struct {
	Token string
}

func getConfig() Config {
	return Config{
		Token: os.Getenv("GITHUB_TOKEN"),
	}
}
