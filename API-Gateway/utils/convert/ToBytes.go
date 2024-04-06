package convert

import (
	"io"
	"mime/multipart"
)

func ConvertMultipartFileToBytes(file *multipart.FileHeader) ([]byte, error) {
	// Open the uploaded file
	uploadedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer uploadedFile.Close()

	// Read the content of the file into a byte slice
	content, err := io.ReadAll(uploadedFile)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func ConvertAllMultipartFilesToBytes(files []*multipart.FileHeader) ([][]byte, error) {
	var contents [][]byte

	// Iterate over each uploaded file
	for _, file := range files {
		// Convert each file to bytes
		content, err := ConvertMultipartFileToBytes(file)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}

	return contents, nil
}
