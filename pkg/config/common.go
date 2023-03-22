package config

import "doraemon/poem/pkg/model"

type Config struct {
	Filepath     string
	MaxFileCount int
	Data         model.Knowledge
	ReadFile     func(string) ([]byte, error)
}
