package service

import "mygithub/smallproject/a03客户关系管理系统/model"

//二: 业务层
type CustomerService struct {
	customers   []model.Customer //客户切片	集合
	customerNum int              //记录客户数量
}

//构造函数	返回service实例
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	//先实例化Coustomer对象 给一个默认对象
	//customer := model.Customer{}
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	//将实例对象添加到集合
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户集合
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	//将用户添加到集合
	this.customers = append(this.customers, customer)
	return true
}

//根据id查询客户 返回客户在集合里的下标
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	//返回-1 即不存
	return index
}

//删除客户
func (this *CustomerService) Delete(id int) bool {
	//先查询集合里是否有该客户
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	//集合里移除此元素
	//如index=2, list的数据:[0 2),即舍弃了2.	然后从index+1开始,将元素添加进list 	"..."表示重复添加
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
