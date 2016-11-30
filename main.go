package main

import(
    "fmt"
    "net/http"
    
)

func main(){

    http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
        fmt.Fprint(w,"Hello goapi")
    })
    http.ListenAndServe(":8081",nil)
}