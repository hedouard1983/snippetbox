package main
import (
        "fmt"
	"log"
        "net/http"
	"strconv"
)
// Defin a home handler function which writes a byte slice containing
// "Hello from SnippetBox" as the body response
func home(w http.ResponseWriter, r *http.Request) {
	//CHeck if the current request URL PATH exactly matches "/"
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}
	w.Write ([]byte("Hello from Snippetbox"))
}

func  showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound()
	}
	w.Write ([]byte("Displays a specific Snippet...."))
}
	//Test Changes

func createSnippet(w http.ResponseWriter, r *http.Request){
	// Us r.Method to check whether the request is using POST or not
	if  r.Method != http.MethodPost{
	w.Header().Set("Allow", http.MethodPost)
	http.Error(w, "Method Now Allowed", 405)
	return
	}
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
