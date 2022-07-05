package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	product, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Running at:")
	fmt.Println(GetOutboundIP(), ":8000 !!!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, product)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
