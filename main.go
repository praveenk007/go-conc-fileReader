package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("tags.csv")
	if err != nil {
		panic(err)
	}
	//f.Seek(6, 0)
	r := bufio.NewReader(f)

	b, pref, err := r.ReadLine()
	fmt.Println(string(b))
	fmt.Println(pref)

	b1, pref1, _ := r.ReadLine()
	fmt.Println(string(b1))
	fmt.Println(pref1)
}
