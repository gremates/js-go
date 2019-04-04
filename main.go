package main

import (
    "fmt"
	//"time"
	//"io"
    "log"
    "net/http"
	"encoding/json"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	//"golang.org/x/crypto/acme/autocert"
	//"github.com/mholt/certmagic"
)

/*type Child struct {
    CID        	string   `json:"cid,omitempty"`
    CName 		string   `json:"cname,omitempty"`
}*/

type Industry struct {
    ID        	string   `json:"id,omitempty"`
    Name 		string   `json:"name,omitempty"`
    Children   *Industry 	 `json:"children,omitempty"`
}

var Industries []Industry

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Schemes("https")

	Industries = append(Industries, Industry{ID: "11", Name: "Agricultura", Children: &Industry{ID: "111", Name: "Siembra de soya", Children: &Industry{ID: "1111", Name: "Siembra de soya"}}})
	Industries = append(Industries, Industry{ID: "97", Name: "Otros", Children: &Industry{ID: "971", Name: "Otros perecederos", Children: &Industry{ID: "972", Name: "Otros servicios", Children: &Industry{ID: "9721", Name: "Servicios varios"}}}})
	
	//Industries = Industry{[{ID: "11",Name: "Agricultura",Children: [{ID: "111",Name: "Siembra de soya", Children: [{ID: "1111", Name: "Siembra de soya"}]}]}, {"id": "97", "name": "Otros", "children": [{ "id" "971","name": "Otros perecederos"}, {"id" "972","name": "Otros servicios","children": [{	"id": "9721","name": "Servicios varios"}] }] }]}
	
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/industries", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/industries/{id:[0-9]+}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/createindustries/{id}", CreateEntry).Methods("POST")
    myRouter.HandleFunc("/deleteindustries/{id}", DeleteEntry).Methods("POST")	
	//myRouter.HandleFunc("/industries/{id:[0-9]+}/{sort:(?:asc|desc)", returnSortedArticle)
    
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Upgrade-Insecure-Requests", "Content-Type: text/json; charset=utf-8"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"})

	//log.Fatal(http.ListenAndServe(":80", handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)))
	port := os.Getenv("PORT")
       if port == "" {
               port = "8081"
       }
       log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)))
	
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	
    fmt.Println("Endpoint Hit: returnAllArticles")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//io.WriteString(w, `{"status":"ok"}`)
	json.NewEncoder(w).Encode(Industries)
	//b, _ := json.Marshal(Industries)
	//fmt.Fprint(w, string(b)) 
	
}


func returnSingleArticle(w http.ResponseWriter, r *http.Request){

    vars := mux.Vars(r)
    //key := vars["id"]
	fmt.Println("Endpoint Hit: returnSingleArticle")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//io.WriteString(w, `{"status":"ok"}`)
    //fmt.Fprintf(w, "Key: " + key)
	
	for _, item := range Industries {
	
		if item.ID == vars["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
		
		if item.Children.ID != "" {
		if item.Children.ID == vars["id"] {
			json.NewEncoder(w).Encode(item.Children)
			return
		}
		}
		
		if item.Children.Children.ID != ""  {
		if item.Children.Children.ID == vars["id"] {
			json.NewEncoder(w).Encode(item.Children.Children)
			return
		}
		}
		
	}
	
	json.NewEncoder(w).Encode(&Industry{})
}

// create a new item
func CreateEntry(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: CreateEntry")

    params := mux.Vars(r)
    var person Industry
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
	fmt.Println(person.ID)
    Industries = append(Industries, person)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//io.WriteString(w, `{"status":"ok"}`)
    json.NewEncoder(w).Encode(Industries)
}

// Delete an item
func DeleteEntry(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: DeleteEntry")

    params := mux.Vars(r)
    for index, item := range Industries {
        if item.ID == params["id"] {
            Industries = append(Industries[:index], Industries[index+1:]...)
            break
        }

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

		//io.WriteString(w, `{"status":"ok"}`)
        json.NewEncoder(w).Encode(Industries)
    }
}


func main() {
    handleRequests()

}