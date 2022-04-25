package main

import (
	"fmt"
	"os"
	"regexp"
)

func readVirus(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func readHuman(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// var flag string
	var virus = readVirus("D:/KULIAH/SEMESTER 4/Strategi Algoritma/Tugas/Tugas Besar/Tubes 3/Tubes3_13520074/src/Backend/virus.txt")
	var human = readHuman("D:/KULIAH/SEMESTER 4/Strategi Algoritma/Tugas/Tugas Besar/Tubes 3/Tubes3_13520074/src/Backend/human.txt")
	var regex, err = regexp.Compile("[AGCT]")

	if err != nil {
		fmt.Println(err.Error())
	}

	if len(regex.FindAllString(virus, -1)) == len(virus) && len(regex.FindAllString(human, -1)) == len(human) {
		fmt.Println("Virus and human is valid")
		fmt.Println(len(regex.FindAllString(virus, -1)))
		fmt.Println(len(virus))
		fmt.Println(len(regex.FindAllString(human, -1)))
		fmt.Println(len(human))
	} else {
		fmt.Println("Virus and human is invalid")
		fmt.Println(len(regex.FindAllString(virus, -1)))
		fmt.Println(len(virus))
		fmt.Println(len(regex.FindAllString(human, -1)))
		fmt.Println(len(human))
	}

	// KMPSearch(virus, human, &flag)
	// fmt.Printf("%s", flag)
	// search(human, virus, &flag)
	// fmt.Printf("%s", flag)
}
