// multiplexer

package main

import (
	"fmt"
	"log"
	"net/http"
)

type pounds float32

func (p pounds) String() string {
	return fmt.Sprintf("%.2f", p)
}

type database map[string]pounds

func (d database) foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo:%s\n", d["foo"])
}

func (d database) bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar:%s\n", d["bar"])
}

func (d database) baz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "baz:%s\n", d["baz"])
}

func main() {
	db := database{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	mux := http.NewServeMux()

	mux.Handle("/foo", http.HandlerFunc(db.foo))
	mux.Handle("/bar", http.HandlerFunc(db.bar))

	mux.HandleFunc("/baz", db.baz)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
