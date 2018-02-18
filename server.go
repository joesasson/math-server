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

func (o *operands) setAnswer(answer float64){
  o.Answer = answer
}

func getParams(r *http.Request) operands {
  params := r.URL.Query()
  x, _ := strconv.ParseFloat(params["x"][0], 64)
  y, _ := strconv.ParseFloat(params["y"][0], 64)
  action := r.URL.Path
  fmt.Println(action)
  return operands{x, y, 0, action}
}

func addHandler(w http.ResponseWriter, r *http.Request){
  operands := getParams(r)
  operands.setAnswer(operands.X + operands.Y)
  j, _ := json.Marshal(operands)
  w.Header().Set("Content-Type", "application/json")
  w.Write(j)
}

func subtractHandler(w http.ResponseWriter, r *http.Request){
  operands := getParams(r)
  answer := operands.X - operands.Y
  fmt.Fprintln(w, answer)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request){
  operands := getParams(r)
  answer := operands.X * operands.Y
  fmt.Fprintln(w, answer)
}

func divideHandler(w http.ResponseWriter, r *http.Request){
  operands := getParams(r)
  answer := operands.X / operands.Y
  fmt.Fprintln(w, answer)
}



func main(){
  http.HandleFunc("/add", addHandler)
  http.HandleFunc("/subtract", subtractHandler)
  http.HandleFunc("/multiply", multiplyHandler)
  http.HandleFunc("/divide", divideHandler)
  http.ListenAndServe(":3000", nil)
}
