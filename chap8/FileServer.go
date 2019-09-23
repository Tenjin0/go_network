package main

import "net/http"

func FileServer() {

	fileServer := http.FileServer(http.Dir("/var/www"))

	err := http.ListenAndServe(":8000", fileServer)

	checkError(err)
}
