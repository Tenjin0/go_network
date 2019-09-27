package main

import (
	"fmt"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}

func listFlashCards(rw http.ResponseWriter, req *http.Request) {

}

func manageFlashCards(rw http.ResponseWriter, req *http.Request) {

}

func main() {

	if len(os.Args) != 1 {
		fmt.Fprint(os.Stderr, "Usage: make run target=FlashCardServer host=:port\n")
		os.Exit(1)
	}

	port := os.Args[0]

	fileServer := http.StripPrefix("/jscript/", http.FileServer(http.Dir("jscript")))
	http.Handle("/jscript/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
	http.Handle("/html/", fileServer)

	http.HandleFunc("/", listFlashCards)
	http.HandleFunc("/flashcards.html", listFlashCards)
	http.HandleFunc("/flashcardSets", manageFlashCards)

	err := http.ListenAndServe(port, nil)
	checkError(err)

}
