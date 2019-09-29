package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/Tenjin0/go_network/chap10/flashcard"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}

func listFlashCards(rw http.ResponseWriter, req *http.Request) {

	flashCardsNames := flashcard.Gets()
	t, err := template.ParseFiles("chap10/html/listFlashcards.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(rw, flashCardsNames)
}
func showFlashCards(rw http.ResponseWriter, cardname string, order string, half string) {
	fmt.Println("Loading card name", cardname)
	aFlashcard := new(flashcard.FlashCards)
	err := flashcard.LoadJSON(cardname, aFlashcard)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonByte, err := json.Marshal(aFlashcard)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write(jsonByte)
}

func listWords(rw http.ResponseWriter, cardname string) {

}

func manageFlashCards(rw http.ResponseWriter, req *http.Request) {

	cardname := req.FormValue("flashcardSets")
	order := req.FormValue("order")
	action := req.FormValue("submit")
	half := req.FormValue("half")

	fmt.Println("cardname", cardname, "action", action)
	if action == "Show cards in set" {
		showFlashCards(rw, cardname, order, half)
	} else if action == "List words in set" {
		listWords(rw, cardname)
	}
}

func server() {

	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: make run target=FlashCardServer host=:port\n")
		os.Exit(1)
	}

	port := os.Args[1]

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
func main() {
	server()
}
