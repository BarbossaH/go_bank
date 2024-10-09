package app

import (
	"bankserver/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.ICustomerService
}

func(ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	fmt.Println(status)
	customers,err:=ch.service.GetAllCustomers(status)

	if err!=nil{
		writeRs(w,err.Code,err.AsMessage())

	}else{
		writeRs(w, http.StatusOK,customers)
	}
	// if r.Header.Get("Content-Type")=="application/xml" {
	// 	//if I don't add this line of code, it will return content type as plan/text
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// }else{
	// 	w.Header().Add("Content-Type","application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }
}

func(ch* CustomerHandlers) getCustomer(w http.ResponseWriter,r* http.Request){
	vars := mux.Vars(r)
	id:=vars["customer_id"]
	customer,err:=ch.service.GetCustomer(id)
	if err!=nil{
		writeRs(w,err.Code,err.AsMessage())

	}else{
		writeRs(w, http.StatusOK,customer)
	}
}

func writeRs(w http.ResponseWriter, code int, data interface{}){
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	if err:=json.NewEncoder(w).Encode(data);err!=nil{
		panic(err)
	}
}