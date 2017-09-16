package hstore

import (
	"bufio"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addString(filepath string, data string) {

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	check(err)

	defer f.Close()

	n3, err := f.WriteString(data)
	log.Printf("wrote %d bytes\n", n3)

	f.Sync()

}

func addBuffer(filepath string, data string) {
	f, err := os.Create(filepath)
	check(err)

	defer f.Close()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	log.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
