package app

import (
	dto "bankserver/DTO"
	"bankserver/service"
	"encoding/json"
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

//make transaction

func(h AccountHandler)makeTransaction(w http.ResponseWriter,r *http.Request){
	//in this handler function, we only need to deal with api-related things and call the service
	vars:=mux.Vars(r)
	accountId:=vars["account_id"]
	customerId:=vars["customer_id"]

	var reqDto dto.TransactionReq
	//check the data in the body
	if err:=json.NewDecoder(r.Body).Decode(&reqDto);err!=nil {
		writeRs(w, http.StatusBadRequest,err.Error())
	}else{
		reqDto.AccountId = accountId
		reqDto.CustomerId=customerId
		//based on reqDto to call the service
		resDto,err:=h.service.MakeTransaction(reqDto)
		if err!=nil {
			writeRs(w,err.Code,err.AsMessage())
		}else{
			writeRs(w,http.StatusOK,resDto)
		}
	}
	
}

