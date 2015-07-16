package main

import (
  "log"
  "net/http"
  "fmt"
)

type String string

type Struct struct {
  Greeting string
  Punct    string
  Who      string
}

func (s String) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request) {
  fmt.Fprint(w, "I'm a frayed knot.")
}

func (s Struct) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request) {
  fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func main() {
  http.Handle("/string", String("I'm a frayed knot."))
  http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

  err := http.ListenAndServe("localhost:4000", nil)
  if err != nil {
    log.Fatal(err)
  }
}
