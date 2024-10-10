package service

import (
	dto "bankserver/DTO"
	"bankserver/domain"
	"bankserver/errs"
)

//supply all kinds of interface
type ICustomerService interface {
	GetAllCustomers(string)([]dto.CustomerRes,*errs.AppError)
	GetCustomer(string)(*dto.CustomerRes,*errs.AppError)
}

//返回ICustomerRepository的接口，目前只有獲取全部用戶的接口，誰繼承了這個接口都可以獲得其對象

//这个service结构其实可以包含很多service，我也不知道后面会包含什么
type DefaultCustomerService struct {
	repo domain.ICustomerRepository
}

//调用数据的真实提供的接口，比如查找数据，增加数据等等CRUD操作
func(s DefaultCustomerService)GetAllCustomers(status string)([]dto.CustomerRes,*errs.AppError){
	cs,err:=s.repo.FindAll(status)
	customers := make([]dto.CustomerRes, 0, len(cs))
	if err!=nil{
		return nil,err
	}
	for _, v := range cs {
		c:=v.CustomerToDto()
		customers =append(customers, c)
	}
	return customers,nil
}

func(s DefaultCustomerService)GetCustomer(id string)(*dto.CustomerRes,*errs.AppError){
	// return s.repo.FindById(id)
	c,err:=s.repo.FindById(id)
	if err!=nil{
		return nil,err
	}
	res:=c.CustomerToDto()

	return &res, nil
}
//谁只要是实现了ICustomerRepository这个可以操作数据的接口，就可以注册使用这个service
func RegisterCustomerService(repository domain.ICustomerRepository)DefaultCustomerService{
	return DefaultCustomerService{repository}
}

