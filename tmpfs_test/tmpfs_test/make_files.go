package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const arraysize = 50000000

var globalarray [arraysize]int

func shuffle(from int, to int) {
	temp := globalarray[from]
	globalarray[from] = globalarray[to]
	globalarray[to] = temp
}

func writeonfile() {
	// 출력파일 생성
	fo, err := os.Create("/root/nallan/tmpfs_test/output.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	for i := 0; i < arraysize; i++ {
		fo.WriteString(fmt.Sprintf("%08d", globalarray[i]))
		if i != 0 && i%10 == 9 && i != arraysize-1 {
			fo.WriteString("\n")
		} else {
			fo.WriteString("\t")
		}
	}
}

func main_deprecated() {
	start := time.Now()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < arraysize; i++ {
		globalarray[i] = i
	}
	for i := 0; i < arraysize/2; i++ {
		shuffle(i, r1.Intn(arraysize))
	}

	writeonfile()

	duration := time.Since(start)
	fmt.Println(duration)
}
