package httphandler

import (
	"github.com/go-chi/chi"

	"github.com/diazharizky/codebase-rest-golang/pkg/bar"
	"github.com/diazharizky/codebase-rest-golang/pkg/foo"
)

func Handler() chi.Router {
	r := chi.NewRouter()
	r.Mount("/foo", foo.Handler())
	r.Mount("/bar", bar.Handler())
	return r
}
