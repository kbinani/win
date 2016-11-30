package internal

import (
	"archive/zip"
	"compress/flate"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func LoadCache(cacheFilePath string) (map[string][]Func, error) {
	funcs := make(map[string][]Func)
	if !IsFileExist(cacheFilePath) {
		return funcs, fmt.Errorf("Cache file not found: %s", cacheFilePath)
	}
	fp, err := os.Open(cacheFilePath)
	if err != nil {
		return funcs, fmt.Errorf("Cannot open cache: %s", cacheFilePath)
	}
	defer fp.Close()
	r, err := zip.NewReader(fp, FileSize(cacheFilePath))
	if err != nil {
		return funcs, err
	}
	if len(r.File) != 1 {
		return funcs, fmt.Errorf("%s does not contains 'cache.json' file entry", cacheFilePath)
	}
	file := r.File[0]
	f, _ := file.Open()
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return funcs, err
	}
	if err := json.Unmarshal(bytes, &funcs); err != nil {
		return funcs, err
	}

	return funcs, nil
}

func SaveCache(cache map[string][]Func, cacheFilePath string) {
	fp, err := os.Create(cacheFilePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	w := zip.NewWriter(fp)
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})
	defer w.Close()
	f, err := w.Create("cache.json")
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(f)
	enc.Encode(&cache)
}
