package bar

import (
	"fmt"
	"net/http"
)

func getSomething(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Bar said: %s", "Hello world!")))
}
