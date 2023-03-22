package config

import (
	"doraemon/poem/pkg/consts"
	"doraemon/poem/pkg/data/shi"
	"doraemon/poem/pkg/model"
)

var shiConfig Config

func init() {
	shiConfig = Config{
		Filepath:     consts.ShiDirPath,
		MaxFileCount: consts.ShiFilesCount,
		Data:         model.Knowledge(new(model.ShiJi)),
		ReadFile:     shi.ReadFile,
	}
}

func GetShiConfig() Config {
	shiConfig.Data = nil
	shiConfig.Data = model.Knowledge(new(model.ShiJi))
	return shiConfig
}
