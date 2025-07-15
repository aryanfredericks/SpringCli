package utils

import (
	"io"
	"net/http"
	"os"
)

func SaveZipFile(zipFileName string, resp *http.Response) {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		panic("Failed to create zip file: " + err.Error())
	}
	defer zipFile.Close()

	_, err = io.Copy(zipFile, resp.Body)
	if err != nil {
		panic("❌ Failed to save zip file: " + err.Error())
	}
}
