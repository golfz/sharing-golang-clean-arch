package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type requestData struct {
	Datetime string  `json:"datetime"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
}

type responseData struct {
	Id       int64   `json:"id"`
	Datetime string  `json:"datetime"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
}

type errorMessage struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}


func main() {
var db *LocationCollection
	log.Println("GPS service is starting...")

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/gps-location", addNewGPSLocation).Methods("POST")

	log.Println("GPS service is on 8989")

	headersOk := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Accept", "Content-Type"})
	// ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8989", handlers.CORS(originsOk, headersOk, methodsOk)(apiRouter)))
}

func sendResponse(w http.ResponseWriter, statusCode int, output interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if output != nil {
		json.NewEncoder(w).Encode(output)
	}
}

func addNewGPSLocation(w http.ResponseWriter, r *http.Request) {

	sendResponse(w, http.StatusInternalServerError, errorMessage{
		ErrorCode: http.StatusInternalServerError,
		ErrorMsg:  "cannot read inserted-data",
	})
	return
}
