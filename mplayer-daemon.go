package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tamentis/go-mplayer"
)

var logger *log.Logger

func errorHandler(err error) {
	logger.Fatalf(err.Error())
}

func handler(w http.ResponseWriter, r *http.Request) {
	command, _ := ioutil.ReadAll(r.Body)
	logger.Println(string(command))
	mplayer.SendCommand(string(command))
}

func main() {
	logFile, err := os.Create("/tmp/mplayer-daemon.log")
	if err != nil {
		panic(err)
	}

	logger = log.New(logFile, "", log.Lshortfile)

	mplayer.StartSlave(errorHandler)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
