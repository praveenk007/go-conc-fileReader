package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../tags.csv")
	if err != nil {
		panic(err)
	}
	fileinfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	filesize := int(fileinfo.Size())
	fmt.Println(filesize)
	buf := make([]byte, 100)
	n, err := f.ReadAt(buf, 10)
	if err != nil {
		log.Printf("err: %v", err)
		//close(channel)
		return
	}
	fmt.Println(string(buf[:n]))
}
