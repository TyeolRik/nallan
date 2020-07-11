package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) {
	inFile, err := os.Open(path)
	check(err)
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	count := 0
	for scanner.Scan() {
		eachData := regexp.MustCompile("[\t\n]")
		splitedDataArray := eachData.Split(scanner.Text(), -1)
		for _, item := range splitedDataArray {
			_, err = strconv.Atoi(item)
		}
		count++
		if count != 0 && count%1000000 == 0 {
			fmt.Println("Status:", count)
		}
	}
}

func readDiskFiles() {
	// This environment is SSD
	// SSD: Intel SSD 530 series 480GB (INTEL SSDSC2BW480H6)
	fmt.Println("Read From Disk")
	readFile("/root/nallan/tmpfs_test/tmpfs_test/430mb_textfile.txt")
}

func readTmpfsFiles() {
	// RAM: 16GB | SAMSUNG PC3L-12800S x2 (M471B1G73EB0-YK01, M471B1G73EB0-YK01)
	fmt.Println("Read From tmpfs(RAM)")
	readFile("/ramdisk/430mb_textfile.txt")
}

func main() {

	for i := 0; i < 10; i++ {
		startDisk := time.Now()
		readDiskFiles()
		duration := time.Since(startDisk)
		fmt.Println(duration)

		starttmpfs := time.Now()
		readTmpfsFiles()
		duration = time.Since(starttmpfs)
		fmt.Println(duration)
	}

}
