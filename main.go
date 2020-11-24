package main
import (
        "log"
        "net/http"
)
// Defin a home handler function which writes a byte slice containing
// "Hello from SnippetBox" as the body response
func home(w http.ResponseWriter, r *http.Request) {
        w.Write ([]byte("Hello from Snippetbox"))
}

func  showSnippet(w http.ResponseWriter, r *http.Request){
        w.Write ([]byte("Displaye a specific Snippet...."))
}

func createSnippet(w http.ResponseWriter, r *http.Request){
        w.Write ([]byte ("Create a new Snippet...."))
}
func main(){
        // use the http.NewServMux() function to initialize a new servemux, then
        // register the home function as the handler for the "/" URL patter

        mux := http.NewServeMux()
        mux.HandleFunc("/", home)
        mux.HandleFunc("/snippet", showSnippet)
        mux.HandleFunc("/snippet/create", createSnippet)

        // Use the http.ListenAndServer() function to start a new web server. We pass in
        // two params

        log.Println("Starting server on :4000")
        err := http.ListenAndServe(":4000", mux)
        log.Fatal(err)

}