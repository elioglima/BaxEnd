package views

import (
	"io/ioutil"
	"log"
	"os"
)

func getHash() ([]byte, error) {
	fpwdr := "files/config/pwdr.json"
	file, err := os.Open(fpwdr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte(""), err
	}

	return b, nil
}
