package kernel

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

const (
	filePerms uint32 = 0777
)

func ErrorHandler(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func serialize(input Dictionary) ([]byte, error) {
	return json.Marshal(&input)
}

func deserialize(input []byte) *Dictionary {
	temp := Dictionary{}
	err := json.Unmarshal(input, &temp)
	ErrorHandler(err)
	return &temp
}

func Save(filename string) {
	if filename == "" {
		fmt.Printf("No file specified")
		return
	}
	var file *os.File
	data, err := serialize(Store)
	ErrorHandler(err)
	if _, err := os.Stat(filename); err != nil {
		file, err = os.Create(filename)
		ErrorHandler(err)
	} else {
		file, err = os.OpenFile(filename, os.O_RDWR, fs.FileMode(filePerms))
		ErrorHandler(err)
	}
	defer file.Close()
	count, err := file.Write(data)
	ErrorHandler(err)
	fmt.Printf("Successfully wrote %d bytes to %s\n", count, filename)
}

func Load(filename string) {
	stream, err := ioutil.ReadFile(filename)
	ErrorHandler(err)
	content := deserialize(stream)
	Store = *content
	fmt.Printf("Successful read from %s\n", filename)
}
