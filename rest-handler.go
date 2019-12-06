package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getFile(w http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["fileName"]

	fmt.Printf("FileName %s\n", fileName)

	var theMap interface{}

	reqBoyd, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)

	} else {

		// Body is expected to be JSON, so just
		// iterate

		json.Unmarshal(reqBoyd, &theMap)

		if theMap == nil {
			http.Error(w, "Body doesn't contain JSON", http.StatusBadRequest)
			return
		}

		m := theMap.(map[string]interface{})

		// for k, v := range m {
		// 	fmt.Printf("Key: '%s' - Value: '%s'\n", k, v)
		// }

		theFile := m["file"]

		fmt.Printf("theFile in the map: '%s'\n", theFile)

		if theFile != nil {
			// fmt.Printf("The file: %s\n", theFile)

			contents, err := ioutil.ReadFile(fmt.Sprintf("%s", theFile))

			if err == nil {
				fmt.Fprintf(w, "%s", contents)
			} else {

				s := fmt.Sprintf("Unable to open file '%s'. -- %+v", theFile, err)

				http.Error(w, s, http.StatusNotFound)
			}

		} else {
			http.Error(w, "File parameter missing", http.StatusNotFound)
		}
	}

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home. Method %s. RequestURI: %s. RemoteAddr: %s. Agent: %s", r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

	// r.
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/getFile/{fileName}", getFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
