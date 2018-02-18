package main

import (
  "fmt"
  "net/http"
  "strconv"
  "encoding/json"
)

type operands struct {
  X float64 `json:x`
  Y float64
  Answer float64
  Action string
  // cached bool
}

func mathHandler(w http.ResponseWriter, r *http.Request){
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
  }
  operands.setAnswer(answer)
  j, _ := json.Marshal(operands)
  w.Header().Set("Content-Type", "application/json")
  w.Write(j)
}

func (o *operands) setAnswer(answer float64){
  o.Answer = answer
}

func getParams(r *http.Request) operands {
  params := r.URL.Query()
  x, _ := strconv.ParseFloat(params["x"][0], 64)
  y, _ := strconv.ParseFloat(params["y"][0], 64)
  action := r.URL.Path[1:]
  return operands{x, y, 0, action}
}


func main(){
  http.HandleFunc("/", mathHandler)
  http.ListenAndServe(":3000", nil)
}
