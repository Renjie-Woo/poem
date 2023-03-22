package ci

import (
	"doraemon/poem/utils"
	"embed"
)

//go:embed *
var f embed.FS

func ReadFile(filepath string) ([]byte, error) {
	return utils.ReadFile(f, filepath)
}
