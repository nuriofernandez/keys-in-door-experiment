package httpKeys

import (
	"fmt"
	"net/http"
)

var KeysInDoor bool

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"keysInDoor\":%t}\n", KeysInDoor)
}

func Start() {
	http.HandleFunc("/status", hello)
	http.ListenAndServe(":8090", nil)
}
