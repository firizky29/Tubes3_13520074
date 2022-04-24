package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readVirus() string {
	data, err := os.ReadFile("D:/KULIAH/SEMESTER 4/Strategi Algoritma/Tugas/Tugas Besar/Tubes 3/Tubes3_13520074/src/Backend/virus.txt")
	check(err)
	return string(data)
}

func readHuman() string {
	data, err := os.ReadFile("D:/KULIAH/SEMESTER 4/Strategi Algoritma/Tugas/Tugas Besar/Tubes 3/Tubes3_13520074/src/Backend/human.txt")
	check(err)
	return string(data)
}

func findLPS(virus *string, m int, lps []int) {
	var len = 0
	var i = 1
	lps[0] = 0

	for i < m {
		if (*virus)[i] == (*virus)[len] {
			len++
			lps[i] = len
			i++
		} else {
			if len != 0 {
				len = lps[len-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
}

func KMPSearch(virus string, human string, flag *string) {
	*flag = "not found"
	var m = len(virus)
	var n = len(human)
	var lps = make([]int, m)
	var i = 0
	var j = 0

	findLPS(&virus, m, lps)

	for i < n {
		if virus[j] == human[i] {
			j++
			i++
		}
		if j == m {
			*flag = "found"
			j = lps[j-1]
		} else if i < n && virus[j] != human[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i = i + 1
			}
		}
	}

}

func main() {
	var flag string
	KMPSearch(readVirus(), readHuman(), &flag)
	fmt.Printf("%s", flag)
}
