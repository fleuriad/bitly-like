package main

//Include for the program
import (
  "github.com/gorilla/mux" // library for the redirections on the server
  "log" //library to print logs 
  "net/http" //libray used for the creation of the server
  "encoding/json" //library for handling json
)


type Request struct {
    Origin	string 'json:"origin"'
    Token	string 'json:"fruits"'
}

func main() {
	
  rtr := mux.NewRouter() //we create the variable rtr which going to handle each redirections of the API
  
  rtr.HandleFunc("/v1/shortlinks/",shortURL).Methods("POST"); 
  
  rtr.HandleFunc("/v1/{shorturl:[a-z]}",shortURL).Methods("POST");
  
  rtr.HandleFunc("/v1/admin/{value:[a-z0-9]+}",Monitoring).Methods("GET")
  
  http.Handle("/", rtr)

  log.Println("Listeningg...")
  http.ListenAndServe(":8080", nil)
}


func shortURL(w http.ResponseWriter, r *http.Request) {
	
	
	decoder := json.NewDecoder(r.Body)
	log.Println(decoder)
	//Case where custom doesn't exist
	//Case where the value custom exist in the request

}

/*The Monitoring function is used to see how many times a shortlink had been used
 * The function wait in entry :
 * value : the value of the short url
 * This function return a JSON code with inside :
 * creation_timestamp
 * origin : the real url
 * token : the short url
 * count : number of times the short link has been used
 * */
func Monitoring(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	value := params["value"]
	w.Write([]byte("Shortlink : " + value))

}

func redirection(w http.ResponseWriter, r *http.Request) {
	
}
