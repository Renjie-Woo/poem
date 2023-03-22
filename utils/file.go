package utils

import (
	"embed"
	"fmt"
)

func ReadFile(f embed.FS, filepath string) ([]byte, error) {
	content, err := f.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("read file %s failed with err : %v", filepath, err)
	}

	return content, nil
}
