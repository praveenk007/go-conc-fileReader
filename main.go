package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("tags.csv")
	if err != nil {
		panic(err)
	}
	//f.Seek(6, 0)
	channel := make(chan string, 20)
	go read(f, channel)
	for {
		v, ok := <-channel
		if ok == false {
			break
		}
		fmt.Println(v)
	}
}

func read(file *os.File, channel chan string) {
	buf := make([]byte, 32*1024)
	for {
		n, err := file.Read(buf)

		if n > 0 {
			channel <- string(buf[:n]) // your read buffer.
		}

		if err == io.EOF {
			close(channel)
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			close(channel)
			break
		}
	}
}
