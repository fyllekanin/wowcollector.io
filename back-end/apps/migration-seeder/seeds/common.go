package seeds

import (
	"io"
	"log"
	"os"
)

func GetBytesFromFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening resource file %s", fileName)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file %s", fileName)
	}

	return byteValue, nil
}
