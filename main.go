package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	bufferSize := int64(100)
	f, err := os.Open("tags.csv")
	if err != nil {
		panic(err)
	}
	fileinfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	filesize := int64(fileinfo.Size())
	fmt.Println(filesize)
	routines := filesize / bufferSize
	if remainder := filesize % bufferSize; remainder != 0 {
		routines++
	}
	fmt.Println("Total routines : ", routines)

	channel := make(chan string, 10)
	wg := &sync.WaitGroup{}

	startTime := time.Now()

	fmt.Println("Start time :: ", startTime)

	for i := int64(0); i < int64(routines); i++ {
		wg.Add(1)
		go read(i*bufferSize, f, channel, bufferSize, filesize, wg)

	}
	done := make(chan struct{})

	go readChannel(channel, done)
	wg.Wait()
	close(channel)
	endTime := time.Now()
	fmt.Println("End time :: ", endTime)
	fmt.Println("Total time (seconds) :: ", (endTime.UTC().Second() - startTime.UTC().Second()))
	<-done
}

func readChannel(channel chan string, done chan struct{}) {
	for data := range channel {
		_ = string(data)
		//fmt.Print(data)
		//do some processing
	}
	close(done)
}

func read(seek int64, file *os.File, channel chan string, bufferSize int64, filesize int64, wg *sync.WaitGroup) {
	defer wg.Done()
	var buf []byte
	if filesize < bufferSize {
		buf = make([]byte, filesize)
	} else if (filesize - seek) < bufferSize {
		buf = make([]byte, filesize-seek)
	} else {
		buf = make([]byte, bufferSize)
	}

	n, err := file.ReadAt(buf, seek)
	if err != nil {
		log.Printf("loc %d err: %v", seek, err)
		return
	}
	if n > 0 {
		channel <- string(buf[:n])
	}
}
