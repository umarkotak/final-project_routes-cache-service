package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)

func ServeRoutes(w http.ResponseWriter, r *http.Request) {
  switch r.Method{
    case "GET":
      routes := getRoutes("data/routes.json")
      w.Header().Set("Content-Type", "application/json")

      if r.FormValue("origin") != "" && r.FormValue("destination") != "" {
        origin := r.FormValue("origin")
        destination := r.FormValue("destination")

        var match_route Routes
        for _, route := range routes {
          if route.Origin == origin && route.Destination == destination {
            match_route = route
            break
          }
        }

        data, err := json.Marshal(match_route)
        if err != nil {
          fmt.Fprintf(w, "error")
        }
        fmt.Fprintf(w, string(data))
      } else {
        data, err := json.Marshal(routes)
        if err != nil {
          fmt.Fprintf(w, "error")
        }
        fmt.Fprintf(w, string(data))
      }
      
    case "POST":
      route := Routes{
        r.FormValue("origin"),
        r.FormValue("lat"),
        r.FormValue("lng"),
        r.FormValue("destination"),
        r.FormValue("distance"),
      }

      saveRoutes(route)
  }
}
