package main
import (
       "fmt"
       "net/http"
       "log")
func homepage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is a page in Golang. Golang is super cool!")
}
func cpppage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "C++ is super cool too!")
}
func main() {
    http.HandleFunc("/", homepage)
    http.HandleFunc("/cpp/", cpppage)
    _, err := http.ListenAndServe(":7258", nil)
    if err != nil {
        log.Fatal(err)
    }
}
    
