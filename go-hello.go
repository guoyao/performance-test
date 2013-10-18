package main   
  
import (  
    "fmt"  
    "net/http"  
)  
  
func sayhelloName(w http.ResponseWriter, r *http.Request){  
    fmt.Fprintf(w, "hello world!")  
}  
  
func main() {  
    http.HandleFunc("/", sayhelloName)
    fmt.Println("Server running at http://127.0.0.1:9090/")  
    http.ListenAndServe(":9090", nil)
}  