package main

import(
  "net/http"
)

func main() {
  http.HandleFunc("/routes", ServeRoutes)

  http.ListenAndServe(":3002", nil)
}
