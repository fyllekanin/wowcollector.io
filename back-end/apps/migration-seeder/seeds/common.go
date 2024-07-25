package seeds

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
)

func GetBytesFromFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("Error opening resource file %s", fileName))
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("Error reading file %s", fileName))
	}

	return byteValue, nil
}
