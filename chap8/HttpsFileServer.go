package main

import "net/http"

func HttpsFileServer() {

	fileServer := http.FileServer(http.Dir("/var/www"))
	err := http.ListenAndServeTLS(":8000", "certs/patrice.petit.name.pem", "certs/private.pem", fileServer)
	checkError(err)

}
