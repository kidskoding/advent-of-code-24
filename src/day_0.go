package src

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Day0(filename string) {
	filepath := filename

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
