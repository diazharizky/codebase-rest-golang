package bar

import (
	"github.com/go-chi/chi"
)

func Handler() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getSomething)
	return r
}
