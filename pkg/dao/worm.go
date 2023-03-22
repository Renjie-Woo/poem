package dao

import (
	"doraemon/poem/pkg/config"
	errors "doraemon/poem/pkg/error"
	"encoding/json"
	"fmt"
	"time"
)

func GetKnowledge() string {
	current := time.Now().Unix()

	var cfg config.Config

	if current%2 == 0 {
		cfg = config.GetShiConfig()
	} else {
		cfg = config.GetCiConfig()
	}

	fileIndex := current % int64(cfg.MaxFileCount)
	filename := fmt.Sprintf(cfg.Filepath, fileIndex)
	content, err := cfg.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return errors.LoadDataError
	}

	err = json.Unmarshal(content, &cfg.Data)
	if err != nil {
		fmt.Println(err)
		return errors.LoadDataError
	}

	return cfg.Data.Get(int(current))
}
