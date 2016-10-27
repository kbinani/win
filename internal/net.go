package internal

import (
	"io"
	"net/http"
	"os"
)

func Wget(url string, destFilePath string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fp, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer fp.Close()
	io.Copy(fp, res.Body)
	return nil
}
