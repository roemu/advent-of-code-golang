package utils

import (
	"os"

	"github.com/charmbracelet/log"
)

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("error when reading file", err)
	}

	return string(data)
 }

