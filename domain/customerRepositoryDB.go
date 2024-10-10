package domain

import (
	"bankserver/errs"
	"bankserver/logger"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct{
	client *sqlx.DB
}

//using sql to get the corresponding data
// func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
// 	var rows *sql.Rows
// 	var err error
// 	customers:=make([]Customer,0)
// 	if status==""{
// 		fmt.Println("here")
// 		findAllSql:="select customer_id, name, city, zipcode, birthday, status from customers"
// 		rows,err=d.client.Query(findAllSql)
// 	}else{
// 		fmt.Println("here1")
// 		findAllSql:="select customer_id, name, city, zipcode, birthday, status from customers where status = $1"
// 		rows,err=d.client.Query(findAllSql,status)
// 	}	
// 	if err!=nil {
// 		// log.Println("Error while querying customer table"+ err.Error())
// 		logger.Error("Error while querying customer table"+ err.Error())
// 		return nil,errs.NewUnexpectedError(err.Error())
// 	}
// 	err=sqlx.StructScan(rows,&customers)
// 	if err != nil {
// 		logger.Error("Error while scanning customers"+err.Error())
// 		return nil, errs.NewUnexpectedError("Unexpected error")
// 	}
// 	// for rows.Next(){
// 	// 	var c Customer
// 	// 	err:=rows.Scan(&c.Id,&c.Name,&c.City,&c.Zipcode,&c.Birthday,&c.Status)
// 	// 	if err!=nil {
// 	// 		log.Println("Error while scanning customers"+ err.Error())
// 	// 		return nil,err
// 	// 	}
// 	// 	customers = append(customers, c)
// 	// }
// 	return customers,nil
// }

func(d CustomerRepositoryDB)FindAll(status string)([]Customer, *errs.AppError){
	customers:=make([]Customer,0)

	var findAllSql string
	var err error
	if status=="" {
		findAllSql="select customer_id, name, city, zipcode, birthday, status from customers"
		err=d.client.Select(&customers,findAllSql)
	}else{
		findAllSql="select customer_id, name, city, zipcode, birthday, status from customers where status = $1"
		err=d.client.Select(&customers,findAllSql,status)
	}

	if err!=nil {
		//log and return error
		logger.Error("Error while querying customer table"+ err.Error())
		return nil,errs.NewUnexpectedError("Unexpected database error")
	}
	return customers, nil
}

func(d CustomerRepositoryDB)FindById(id string)(*Customer,*errs.AppError){
	var c Customer
	customerSql:="select customer_id, name, city, zipcode, birthday, status from customers where customer_id = $1"
	err:=d.client.Get(&c,customerSql,id)

		if err!=nil {
		if err==sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}else{
			logger.Error("Error while scanning customer"+ err.Error())
			return nil,errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func InitCustomerRepositoryDB(db *sqlx.DB)CustomerRepositoryDB{
	return CustomerRepositoryDB{db}
}

// func(d CustomerRepositoryDB)FindById(id string)(*Customer, *errs.AppError){
// 	customerSql:="select customer_id, name, city, zipcode, birthday, status from customers where customer_id = $1"
// 	row:=d.client.QueryRow(customerSql,id)
// 	// fmt.Println(row)
// 	var c Customer
// 	err:= row.Scan(&c.Id,&c.Name,&c.City,&c.Zipcode,&c.Birthday,&c.Status)
// 	if err!=nil {
// 		if err==sql.ErrNoRows {
// 			return nil, errs.NewNotFoundError("Customer not found")
// 		}else{
// 			// log.Println("Error while scanning customer"+ err.Error())
// 			logger.Error("Error while scanning customer"+ err.Error())
// 			return nil,errs.NewUnexpectedError("Unexpected database error")
// 		}
// 	}
// 	return &c, nil
// }

//-------
//actually, here has two responsibilities, connect the db server and create repositoryDB object
/*func ConnectDB() CustomerRepositoryDB{
	err:=godotenv.Load(".env")
	if err!=nil{
		logger.Error("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo:=fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	client,err:=sqlx.Open("postgres",psqlInfo)
	
	if err!=nil {
		panic(err)
	}
	err=client.Ping()
	if err!=nil {
		panic(err)
	}

	client.SetConnMaxIdleTime(time.Minute*3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	fmt.Println("DB connect successfully")
	return CustomerRepositoryDB{client}
}*/
//------