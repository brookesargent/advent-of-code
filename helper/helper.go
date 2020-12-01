package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadInputTxt(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %s", err)
	}
	lines := strings.Split(string(b), "\n")
	return lines, nil
}