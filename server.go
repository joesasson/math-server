package main

import (
  "fmt"
  "net/http"
  "strconv"
  "encoding/json"
  "log"
)

type operands struct {
  X float64
  Y float64
  Answer float64
  Action string
  Cached bool
}

func (o *operands) setAnswer(answer float64){
  o.Answer = answer
}

func mathHandler(w http.ResponseWriter, r *http.Request){
  // Caching
  w.Header().Set("cache-control", "public, max-age=60")

  operands := getParams(r)
  x, y := operands.X, operands.Y
  answer := 0.0
  switch operands.Action {
  case "add":
    answer = x + y
  case "subtract":
    answer = x - y
  case "multiply":
    answer = x * y
  case "divide":
    answer = x / y
  default:
    fmt.Fprintf(w, "Not a valid operation")
    return
  }
  operands.setAnswer(answer)
  j, _ := json.Marshal(operands)
  w.Header().Set("Content-Type", "application/json")
  w.Write(j)
}

func getParams(r *http.Request) operands {
  params := r.URL.Query()
  x, _ := strconv.ParseFloat(params["x"][0], 64)
  y, _ := strconv.ParseFloat(params["y"][0], 64)
  action := r.URL.Path[1:]
  return operands{x, y, 0, action, false}
}


func main(){
  http.HandleFunc("/", mathHandler)
  log.Println("Test the add action with http://localhost:3000/add?x=5&y=106")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
