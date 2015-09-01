package main

import (
	//"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var home = os.Getenv("HOME")
var testJson = strings.Join([]string{home, "/.testSvc/test.json"}, "")

func response(rw http.ResponseWriter, request *http.Request) {
	json, _ := ioutil.ReadFile(testJson)
	log.Info(string(json))
	rw.Write([]byte(json))
}

func main() {
	http.HandleFunc("/testSvc", response)
	http.ListenAndServe(":12345", nil)
}
