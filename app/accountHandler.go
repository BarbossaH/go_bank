package app

import (
	dto "bankserver/DTO"
	"bankserver/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.IAccountService
}

func(handler AccountHandler)createAccount(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	customerId:=vars["customer_id"]
	var req dto.CreateAccountReq
	err:=json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(req)
	if err!=nil {
		writeRs(w,http.StatusBadRequest,err.Error())
	}else{
		req.CustomerId=customerId
		account,appError:=handler.service.CreateAccount(req)
		if appError!=nil {
			writeRs(w, appError.Code,appError.Message)
		}else{
			writeRs(w,http.StatusCreated,account)
		}
	}
}