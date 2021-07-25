package foo

import (
	"fmt"
	"net/http"
)

func getSomething(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Foo said: %s", "Hello world!")))
}
