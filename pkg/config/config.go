package config

import (
	"os"
	"strings"
)

type config struct {
	DiscordHook string
	Namespaces  []string
}

func GetConfig() *config {
	return &config{
		DiscordHook: os.Getenv("DISCORDHOOK"),
		Namespaces:  strings.Split(os.Getenv("NAMESPACES"), ","),
	}
}
