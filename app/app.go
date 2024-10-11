package app

import (
	"bankserver/domain"
	"bankserver/logger"
	"bankserver/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)


func sanityCheck(){
	err:=godotenv.Load(".env")
	if err!=nil{
		logger.Error("Error loading .env file")
	}
	if os.Getenv("SERVER_ADDRESS")=="" || os.Getenv("SERVER_PORT")==""{
		log.Fatal("Environment variable not defined")	
	}
}

func connectDB()*sqlx.DB{

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo:=fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	psqDB,err:=sqlx.Open("postgres",psqlInfo)
	
	if err!=nil {
		panic(err)
	}
	err=psqDB.Ping()
	if err!=nil {
		panic(err)
	}

	psqDB.SetConnMaxIdleTime(time.Minute*3)
	psqDB.SetMaxOpenConns(10)
	psqDB.SetMaxIdleConns(10)

	fmt.Println("DB connect successfully")

	return psqDB
}

func Start() {
	sanityCheck()
	
	// mux:=http.NewServeMux() //it's similar to Router in express.js
	router:=mux.NewRouter()

	//inject the service
	// handlers:=CustomerHandlers{service.RegisterCustomerService(domain.CreateCustomerRepositoryStubData())}
	
	//这个db继承了customer 的interface接口，所以就有了findall方法
	psqdb:=connectDB()
	customerDB:=domain.InitCustomerRepositoryDB(psqdb)
	accountDB:=domain.InitAccountRepositoryDB(psqdb)
	ch:=CustomerHandlers{service.RegisterCustomerService(customerDB)}
	ah:=AccountHandler{service.RegisterAccountService(accountDB)}
	//routes
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet) //[a-zA-Z0-9]

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account",ah.createAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}",ah.makeTransaction).Methods(http.MethodPost)
	//start server
	address:=os.Getenv("SERVER_ADDRESS")
	port:=os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",address,port),router))
}

