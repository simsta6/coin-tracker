package apicalls

import (
	"io"
	"net/http"
	"os"
)

//DownloadFile method downloads file and writes it in directory
func DownloadFile(url string, filePath string) (err error) {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	return err
}
