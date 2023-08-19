package main

import (
	"fmt"
	"net/http"
    "log"
    "html/template"
)


// Simple hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Hello!")
}


// Reads the completed form
func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful")
    name := r.FormValue("name")
    address := r.FormValue("address")

    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}


// For now, redirects to form.html
func formWriter(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/formwrite" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    //fmt.Fprintf(w, "Hello!")
    //http.ServeFile(w, r, "form.html")
    //http.ServeFile(w, r, "form.html")
    //if err := r.ParseForm(); err != nil {
    //    fmt.Fprintf(w, "ParseForm() err: %v", err)
    //    return
    //}
    //fmt.Fprintf(w, "POST request successful")
    //name := r.FormValue("name")
    //address := r.FormValue("address")

    //fmt.Fprintf(w, "Name = %s\n", name)
    //fmt.Fprintf(w, "Address = %s\n", address)
    
    http.Redirect(w, r, "/form.html", http.StatusSeeOther)
}


// Simple wiki
func wikiHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/wiki" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    title := r.URL.Path[len("/view/"):]

}


func main() {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/formwrite", formWriter)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

    //if err := http.ListenAndServeTLS(":443", "full_cert.crt", "private_key.key", nil); err != nil {
    //    log.Fatal("ListenAndServeTLS: ", err)
    //}
}


