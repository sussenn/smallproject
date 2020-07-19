package main

import (
	"fmt"
	"mygithub/smallproject/a03客户关系管理系统/model"
	"mygithub/smallproject/a03客户关系管理系统/service"
)

//三: 调用层
type customerView struct {
	key             string                   //接收客户端输入
	loop            bool                     //决定是否结束 for
	customerService *service.CustomerService //类似Java注入service
}

//显示所有客户信息
func (this *customerView) list() {
	//获取当前所有用户信息. 使用service的List()
	customers := this.customerService.List()
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	//遍历输出客户信息
	for i := 0; i < len(customers); i++ {
		//类似Java实体类的get()set()
		fmt.Println(customers[i].GetInfo())
	}

}

//控制台输入, 添加客户信息
func (this *customerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	//将接收的数据封装到Customer对象.	注意:id是系统生成,不需要用户输入
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用service的Add()方法
	if this.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

//删除客户信息
func (this *customerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		//放弃当前操作
		return
	}
	fmt.Println("确认是否删除(Y/N)：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("你的输入有误，请输入y/n:")
	}
	if choice == "y" || choice == "Y" {
		//调用customerService 的 Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在------")
		}
	}
}

//退出系统
func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

//显示菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("1添加客户")
		fmt.Println("2修改客户")
		fmt.Println("3删除客户")
		fmt.Println("4客户列表")
		fmt.Println("5退出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("退出系统...")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//实例化CustomerService, 用于存储客户信息
	//service在当前包已调用各个方法,这里将它实例化即运行各个方法
	customerView.customerService = service.NewCustomerService()
	//显示菜单
	customerView.mainMenu()
	//========================================================
	//"方法": 需要传参当前对象(结构体),为了能调用当前对象里的各个方法 	对象.方法名
	//"函数": 包名.函数名
}
