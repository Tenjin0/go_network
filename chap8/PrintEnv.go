package main

import (
	"fmt"
	"net/http"
	"os"
)

func PrintEnv() {

	fileServer := http.FileServer(http.Dir("/var/www"))
	http.Handle("/", fileServer)

	http.HandleFunc("/cgi-bin/printenv", func(writer http.ResponseWriter, req *http.Request) {
		env := os.Environ()
		fmt.Println(env)

		writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
		for _, v := range env {
			writer.Write([]byte(v + "\n"))
		}
		writer.Write([]byte("</pre>"))
	})

	err := http.ListenAndServe(":8000", nil)
	checkError(err)

}
