package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Welcome struct {
	Greet string
}

func Hello(w http.ResponseWriter, r *http.Request) {

	hi := Welcome{Greet: "Oh shit, waddup"}
	json.NewEncoder(w).Encode(hi)

	fmt.Print("Oh shit, waddup")
}
