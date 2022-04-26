package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "regexp"
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

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("../Frontend/rucikawavin/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

}
