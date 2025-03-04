package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":7000", nil)
	fmt.Println("LISTENING ON 7000")
}

func handleHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("HOLA AMIGO"))
}
