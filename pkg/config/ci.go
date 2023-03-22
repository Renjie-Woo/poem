package config

import (
	"doraemon/poem/pkg/consts"
	"doraemon/poem/pkg/data/ci"
	"doraemon/poem/pkg/model"
)

var ciConfig Config

func init() {
	ciConfig = Config{
		Filepath:     consts.CiDirPath,
		MaxFileCount: consts.CiFilesCount,
		Data:         model.Knowledge(new(model.CiJi)),
		ReadFile:     ci.ReadFile,
	}
}

func GetCiConfig() Config {
	ciConfig.Data = nil
	ciConfig.Data = model.Knowledge(new(model.CiJi))
	return ciConfig
}
