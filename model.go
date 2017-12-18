package main

import (
  "encoding/json"
  "io/ioutil"
)

type Routes struct {
  Origin         string  `json:"origin"`
  Lat            string  `json:"lat"`
  Lng            string  `json:"lng"`
  Destination    string  `json:"destination"`
  Distance       string  `json:"distance"`
}

func getRoutes(file string) []Routes {
  raw, err := ioutil.ReadFile(file)
  if err != nil {
      // fmt.Println(err.Error())
      // os.Exit(1)
  }

  var routes []Routes
  json.Unmarshal(raw, &routes)
  return routes
}

func saveRoutes(route Routes){
  routes := getRoutes("./data/routes.json")
  routes = append(routes, route)

  routesJson, _ := json.Marshal(routes)

  ioutil.WriteFile("data/routes.json", routesJson, 0644)
}
