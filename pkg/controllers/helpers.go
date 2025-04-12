package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// DecodeF parses a form from a Route and populates dst with PostForm map values
func DecodeF(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
