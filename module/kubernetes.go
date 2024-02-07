package kubernetes

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello user, Welcome at kubernetes lib")))
}
