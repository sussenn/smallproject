package main

import "fmt"

func main() {
	//用于接收控制台输入
	key := ""
	//用于控制是否退出 for
	loop := true
	//账户余额
	balance := 1000.0
	//收支金额
	money := 0.0
	//收支说明
	note := ""
	//收支说明的字符串拼接
	details := "收支\t账户余额\t收支金额\t说 明"
	//用于记录记账本是否有信息记录
	flag := false
	//显示菜单
	for {
		fmt.Println("-----------------家庭收支记账本-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4  退出")
		fmt.Print("请选择(1-4): ")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("-----------------当前收支明细-----------------")
			if flag {
				fmt.Println(details) //收支说明
			} else {
				fmt.Println("暂无记录...")
			}
		case "2":
			fmt.Println("===登记收入===")
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			//将收入说明拼接到details	//收入	1100	100		抢红包
			details += fmt.Sprintf("\n收入\t%v\t%v\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("===登记支出===")
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足!")
				break
			}
			balance -= money
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("确认退出? y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("输入有误!请输入y 或 n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确选项!")
		}
		if !loop {
			break
		}
	}
	fmt.Println("已退出系统...")
}
