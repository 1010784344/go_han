package main

import (
	"fmt"
	"han_count/model"
	"han_count/service"
)

type customerView struct {

	key string  // 接收用户的收入
	loop bool	// 表示是否循环的显示主菜单
	customerService *service.CustomerService // 接收业务逻辑处理对象
}



// 得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add(){

	fmt.Println("\n----添加客户-------")

	fmt.Println("姓名：")
	name:=""
	fmt.Scanln(&name)

	fmt.Println("性别：")
	gender:=""
	fmt.Scanln(&gender)

	fmt.Println("年龄：")
	age:= 0
	fmt.Scanln(&age)

	fmt.Println("电话：")
	phone:=""
	fmt.Scanln(&phone)

	fmt.Println("邮箱：")
	email:=""
	fmt.Scanln(&email)

	// 构建一个新的 customer 实例
	// 注意 id 号，没有让用户输入，id 是唯一的，需要系统进行分配
	customer := model.NewCustomer2(name,gender,age,phone,email)

	// 调用
	if this.customerService.Add(customer){
		fmt.Println("\n----添加完成-------")
	}else {
		fmt.Println("\n----添加失败-------")
	}


}


// 显示所有的客户信息
func (this *customerView) list(){

	// 首先获取到当前所有的客户的信息（在切片中）
	customers := this.customerService.List()
	// 显示
	fmt.Println("\n----客户列表-------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0;i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}

	fmt.Println("----客户列表完成----")

}



// 得到用户的输入id，删除该 id 对应的客户
func (this *customerView) delete() {

	fmt.Println("\n----删除客户-------")

	fmt.Println("请选择待删除客户编号(-1退出)： ")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return//放弃删除操作
	}


	fmt.Println("确认是否删除(y/n)")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y"{

		// 调用 customerService 的 Delete 方法
		if this.customerService.Delete(id){

			fmt.Println("----删除完成-------")
		}else {

			fmt.Println("----删除失败，输入的id号不存在 -------")
		}

	}


}



// 退出软件
func (this *customerView) exit() {

	fmt.Println("确认是否退出（Y/N）")
	isexit := ""


	for{
		fmt.Scanln(&isexit)

		if isexit == "Y" || isexit == "y" || isexit == "N" || isexit == "n"{

			break
		}

		fmt.Println("你的输入有误，确认是否退出（Y/N）")

	}

	if isexit == "Y" || isexit == "y"{

		this.loop = false
	}

}


// 显示主菜单
func (this *customerView) mainmenu(){

	for{
		fmt.Println("\n----客户信息管理软件----")
		fmt.Println("    1 添加客户")
		fmt.Println("    2 修改客户")
		fmt.Println("    3 删除客户")
		fmt.Println("    4 客户列表")
		fmt.Println("    5 退   出")

		fmt.Println("    请选择（1-5）")

		fmt.Scanln(&this.key)

		switch this.key {

		case "1":
			//fmt.Println("添加客户")
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			//fmt.Println("删除客户")
			this.delete()
		case "4":
			//fmt.Println("客户列表")
			this.list()
		case "5":
			//this.loop = false
			this.exit()

		default:
			fmt.Println("你的输入有误，请重新输入！")

		}

		// 跳出 for 循环
		if !this.loop{

			break
		}

	}


	fmt.Println("你退出了客户关系管理系统！")

}


func main()  {

	// 在 main 函数中，创建一个 customerView ，并运行显示主菜单
	customerview := customerView{

		key:"",
		loop:true,
	}

	// 完成对 customerview 结构体的 customerService 字段的初始化
	customerview.customerService = service.NewCustomerService()


	// 显示主菜单
	customerview.mainmenu()
}
