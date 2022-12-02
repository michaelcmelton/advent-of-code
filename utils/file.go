package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetInput(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("error opening file: (%w)", err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return "", fmt.Errorf("error reading file: (%w)", err)
	}

	return string(bytes), nil
}
