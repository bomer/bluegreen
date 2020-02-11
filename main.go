package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
)

var port int = 8000

func echoColour(w http.ResponseWriter, req *http.Request) {
	if port == 8000 {
		fmt.Fprintf(w, "blue")
	} else if port == 8001 {
		fmt.Fprintf(w, "green")
	} else {
		fmt.Fprintf(w, "DIE")
		panic("Ya done mesesd up")
	}

}

func checkPortIsFree(p int) bool {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(p))

	if err != nil {
		fmt.Sprintf("Can't listen on port %d: %s", p, err)
		return false
	}

	ln.Close()
	fmt.Printf("TCP Port %d is available", port)
	return true
}

func main() {
	println("Hello Green")
	println("Running on port, %d", port)
	http.HandleFunc("/", echoColour)

	//if port in use, increase by 1
	if !checkPortIsFree(port) {
		port++
		if !checkPortIsFree(port) {
			panic("Port 8001 and 8001 in use. shutting down")
		}
	}

	serveOn := ":" + strconv.Itoa(port)
	http.ListenAndServe(serveOn, nil)

}
