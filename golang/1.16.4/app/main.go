package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "I Love You! Hana!")
	if err != nil {
		log.Fatalln("handle func err:", err)
	}
	fmt.Println(time.Now().Format("2006-01-02  15:04:05"))
}

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalln("server listen err:", err)
	}
}
