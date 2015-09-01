package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var home = os.Getenv("HOME")
var testJson = strings.Join([]string{home, "/.testSvc/test.json"}, "")
var route = flag.String("r", "/", "The local route.")
var port = flag.String("p", "1234", "The local port.")
var verbose = flag.Bool("v", false, "Verbosity.")

func response(resp http.ResponseWriter, request *http.Request) {
	log.Debug(request)
	json, err := ioutil.ReadFile(testJson)
	if err != nil {
		panic(err)
	}
	log.Debug("Sending :\n", string(json))
	resp.Write([]byte(json))
}

func main() {
	flag.Parse()
	if *verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("Loglevel: Debug")
	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("Loglevel: Info")
	}
	log.Info("testSvc starting up: ", "http://localhost:", *port, *route)
	http.HandleFunc(*route, response)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
