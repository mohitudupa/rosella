package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

const (
	FileRepositoryBackend = "fileRepositoryBackend"
)

type FileRepositoryConfig struct {
	Path string `yaml:"path"`
}

type ServerConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	ReadOnly bool   `yaml:"readOnly" json:"readOnly"`
}

type ApplicationConfig struct {
	Server            ServerConfig         `yaml:"server" json:"server"`
	RepositoryBackend string               `yaml:"repositoryBackend" json:"repositoryBackend"`
	FileRepository    FileRepositoryConfig `yaml:"fileRepository" json:"fileRepository"`
}

func NewApplicationConfig(configPath string) (*ApplicationConfig, error) {
	ac := ApplicationConfig{Server: ServerConfig{}, RepositoryBackend: "", FileRepository: FileRepositoryConfig{}}

	ac.Server.Port = "8080"

	jsonFile, err := os.ReadFile(configPath)
	if err != nil {
		return &ac, errors.New("failed to locate config file at " + configPath)
	}

	err = json.Unmarshal(jsonFile, &ac)
	if err != nil {
		return &ac, errors.New("failed to load config file at " + configPath)
	}

	log.Printf("INFO: finished loading config file at " + configPath)
	return &ac, nil
}
