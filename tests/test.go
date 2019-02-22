package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../tags-c.csv")
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
	buf := make([]byte, 5)
	n, err := f.ReadAt(buf, 1)
	if err != nil {
		log.Printf("err: %v", err)
		//close(channel)
		return
	}
	fmt.Println(string(buf[:n]))

	buf1 := make([]byte, 5)
	n1, err1 := f.ReadAt(buf1, 10)
	if err1 != nil {
		log.Printf("err: %v", err1)
		//close(channel)
		return
	}
	fmt.Println(string(buf1[:n1]))
}
