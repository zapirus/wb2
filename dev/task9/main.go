package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const (
	Dir = "html"
)

func GetWithClient(cl *http.Client, fileName string) func(string) error {
	return func(url string) error {
		r, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		w, err := cl.Do(r)
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}(w.Body)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Printf("status: %d", w.StatusCode)
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return err
		}

		err = writeHtml(fileName, p)
		if err != nil {
			return err
		}

		return nil
	}
}

func writeHtml(fileName string, p []byte) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(dir, Dir, fileName), os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	defer f.Close()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(f.Name(), p, fs.ModePerm)
}

var fileName = flag.String("O", "index.html", "файл для сохранения данных")

func main() {

	flag.Parse()
	path := flag.Arg(0)
	transport := &http.Transport{}
	cl := &http.Client{Transport: transport}

	DownPage := GetWithClient(cl, *fileName)
	err := DownPage(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}
