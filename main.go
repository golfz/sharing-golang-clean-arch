package main

import (
	"demo/go-clean-demo/fakedb"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
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
	Speed    string  `json:"speed"`
}

type errorMessage struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func main() {
	log.Println("GPS service is starting...")

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/gps-location", addNewGPSLocation).Methods("POST")

	log.Println("GPS service is on 8989")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	headersOk := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Accept", "Content-Type"})

	log.Fatal(http.ListenAndServe(":8989", handlers.CORS(originsOk, headersOk, methodsOk)(apiRouter)))
}



func addNewGPSLocation(w http.ResponseWriter, r *http.Request) {

	reqData := requestData{}

	errReqData := json.NewDecoder(r.Body).Decode(&reqData)
	if errReqData != nil {
		sendResponse(w, http.StatusInternalServerError, errorMessage{
			ErrorCode: http.StatusBadRequest,
			ErrorMsg:  "request body mismatched",
		})
		return
	}

	t, errTime := time.Parse("2006-01-02 15:04:05Z07:00", reqData.Datetime)
	if errTime != nil {
		sendResponse(w, http.StatusInternalServerError, errorMessage{
			ErrorCode: http.StatusBadRequest,
			ErrorMsg:  "time-format mismatched",
		})
		return
	}

	if (reqData.Lat < -90 || 90 < reqData.Lat) || (reqData.Long < -180 || 180 < reqData.Long) {
		sendResponse(w, http.StatusInternalServerError, errorMessage{
			ErrorCode: http.StatusBadRequest,
			ErrorMsg:  "Lat or Long is not corrected",
		})
		return
	}

	db := fakedb.InitDBConnection()

	locationModel := fakedb.InitLocationModel(db);

	locationModel.AddNewLocation(fakedb.Location{
		Time: t,
		Lat:  reqData.Lat,
		Long: reqData.Long,
	})

	locationList := locationModel.GetAll()

	resp := []responseData{}

	for _, v := range locationList {
		kmh := float64(v.GetSpeedMPH()) * 1.60934
		speed := fmt.Sprintf("%d km/h", int64(kmh))

		resp = append(resp, responseData{
			Id:       *v.Id,
			Datetime: v.Time.Format(time.RFC1123),
			Lat:      v.Lat,
			Long:     v.Long,
			Speed:    speed,
		})
	}

	sendResponse(w, http.StatusInternalServerError, resp)
	return
}
