package domain

//它这个命名不好，customer其实只是提供接口，而这个类是实现接口，操作数据，真实的情景是这里应该实际会从数据库获得数据，而不是模拟使用数据

//通过这个类存入数据切片
type CustomerRepositoryStub struct {
	customers []Customer
}

//类有获取数据切片的方法并返回这个切片和错误
// implement the logic of ICustomerRepository interface, which means CustomerRepositoryStub inherit from the interface, it's child class
func (s CustomerRepositoryStub) FindAll(string) ([]Customer, error) {
	return s.customers, nil
}

//创建一个模拟的数据并返回
func CreateCustomerRepositoryStubData() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Julian", City: "A", Zipcode: "1016", Birthday: "1981-4-24", Status: "0"},
		{Id: "1002", Name: "M", City: "A", Zipcode: "1016", Birthday: "1981-4-22", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
