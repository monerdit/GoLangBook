package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// echo running time tested for args "I love goLang a lot, and go programming book is great book to learn goLang"
// echo3 was the fastest

func main() {
	start := time.Now()
	// echo1()
	// echo1 running time: 35.17 - 38.62 microseconds
	// echo2()
	// echo2 running time: 27 - 32 microseconds
	echo3()
	// echo3 running time: 26 - 29 microseconds
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("echo execution time: ", elapsed)
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))

}
