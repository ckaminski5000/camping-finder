package util

import (
	"io"
	"os"
	"time"
)

func ReadFileBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content []byte
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if n == 0 {
			break
		}

		content = append(content, buffer[:n]...)
	}

	return content, nil
}

func ValidateAndReturnDate(date string) (bool, string) {
	expectedFormat := "2006-01-02" // YYYY-MM-DD

	// Parse the date parameter with the expected format
	parsedDate, err := time.Parse(expectedFormat, date)
	if err != nil {
		return false, ""
	}

	formattedDate := parsedDate.Format("2006-01-02T15:04:05Z")
	return true, formattedDate
}

func IsEmpty(str string) bool {
	return str == ""
}
