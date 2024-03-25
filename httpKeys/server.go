package httpKeys

import (
	"fmt"
	"github.com/nuriofernandez/keys-in-door-experiment/keyscheck"
	"net/http"
)

func status(w http.ResponseWriter, req *http.Request) {
	keysInDoor, err := keyscheck.AreKeysThere()
	if err != nil {
		fmt.Fprintf(w, "{\"error\":1}\n")
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "{\"keysInDoor\":%t}\n", keysInDoor)
}

func Start() {
	http.HandleFunc("/status", status)
	http.ListenAndServe(":8090", nil)
}
