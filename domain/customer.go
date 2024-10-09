package domain

import (
	dto "bankserver/DTO"
	"bankserver/errs"
)

//数据结构
type Customer struct {
	Id       string `db:"customer_id"`
	Name     string
	City     string
	Zipcode  string
	Birthday string
	Status   string //is customer or not
}

func(c Customer)ToDto() dto.CustomerRes{
	return dto.CustomerRes{
		Id: c.Id,
		Name: c.Name,
		Status: c.Status,
		Birthday: c.Birthday,
		Zipcode: c.Zipcode,
		City: c.City,
	}
}


//获取数据的接口，如果谁继承这个接口就可以返回切片的数据，需要什么数据返回什么数据
type ICustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string)(*Customer, *errs.AppError)
}


