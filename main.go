package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gobuffalo/packr"
	"github.com/joemarct/neutrino/backend"
	"github.com/phayes/freeport"
	"github.com/zserge/webview"
)

func runWebApp(port int) {
	box := packr.NewBox("./webapp/dist")
	http.Handle("/", http.FileServer(box))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func main() {
	webappPort, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	backendPort := 4564
	go runWebApp(webappPort)
	go backend.RunServer(backendPort)
	webview.Open("Example Application",
		"http://localhost:"+strconv.Itoa(webappPort), 800, 600, true)
}