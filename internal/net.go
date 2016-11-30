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
	_, e := io.Copy(fp, res.Body)
	if e != nil {
		return e
	}
	return nil
}
