package controller

import "net/http"

func GetTrivia(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Trivia based on date"))
}