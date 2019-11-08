package main

import (
	"demo/go-clean-demo/controller"
	"demo/go-clean-demo/presenter"
	"demo/go-clean-demo/usecase"
	"demo/go-clean-demo/view"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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

	v := view.InitJsonResponseView(w)

	pSuccess := presenter.InitLocationPresenter(v)
	pError := presenter.InitErrorPresenter(v)

	uc := usecase.InitLocationUseCase()

	ctrl := controller.InitLocationController(r, pSuccess, pError)
	ctrl.AddLocation(uc)
}
