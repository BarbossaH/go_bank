package app

import (
	"bankserver/domain"
	"bankserver/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// mux:=http.NewServeMux() //it's similar to Router in express.js
	router:=mux.NewRouter()

	//inject the service
	// handlers:=CustomerHandlers{service.RegisterCustomerService(domain.CreateCustomerRepositoryStubData())}
	
	//这个db继承了customer 的interface接口，所以就有了findall方法
	db:=domain.ConnectDB()
	handlers:=CustomerHandlers{service.RegisterCustomerService(db)}
	
	//routes
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", handlers.getAllCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.getCustomer).Methods(http.MethodGet) //[a-zA-Z0-9]


	//start server
	log.Fatal(http.ListenAndServe("localhost:8000",router))
}

