// 包名和文件名一致就行
package utils

import "fmt"

type han_count struct {

	// 声明一个字段，保存用户输入的选项
	key string
	// 声明一个变量，是否退出 for 循环
	loop bool

	// 定义账户的余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对 details 进行拼接处理
	details string

	// 定义一个字段，记录是否有收支的行为
	flag bool
}

// 给该结构体绑定相应的方法

// 将显示明细写成一个方法
func (this *han_count) showDetails(){
	fmt.Println("-----当前收支明细记录-----")
	if this.flag{
		fmt.Println(this.details)
	}else {
		fmt.Println("当前没有收支明细，请来一笔吧！")
	}
}

// 将登记收入写成一个方法
func (this *han_count) income(){
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	// 修改账户余额
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	// 将这个收入情况，拼接到 details 变量
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true

}

// 将登记支出写成一个方法
func (this *han_count) pay(){
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)

	if this.money > this.balance{

		fmt.Println("余额金额不足")

		//break
	}

	// 修改账户余额
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	// 将这个收入情况，拼接到 details 变量
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag = true

}

// 将退出系统写成一个方法
func (this *han_count) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""

	for {
		fmt.Scanln(&choice)

		if choice == "y" || choice == "n"{

			break
		}else {

			fmt.Println("你的输入有误，请重新输入 y/n")
		}

	}

	if choice == "y"{

		this.loop = false
	}

}

// 编写要给工厂模式的构造方法，返回一个 *han_count 实例
func NewHanCount() *han_count{


	return &han_count{
		key : "",

		loop : true,
		balance : 1000.0,

		money : 0.0,
		// 每次收支的说明
		note : "",

		details : "收支\t账户金额\t收支金额\t说明",

		flag : false,
	}


}


// 显示主菜单
func (this *han_count) MainMenu(){

	for {
		fmt.Println("\n-----家庭收支记账软件-----")
		fmt.Println("     1 收支明细         ")
		fmt.Println("     2 登记收入         ")
		fmt.Println("     3 登记支出         ")
		fmt.Println("     4 退出         ")

		fmt.Print("请选择（1-4）：")

		this.key = ""

		fmt.Scanln(&this.key)

		switch this.key {

		case "1":
			this.showDetails()

		case "2":
			this.income()

		case "3":
			this.pay()

		case "4":

			this.exit()

		default:
			fmt.Println("请输入正确的选项！")

		}

		if !this.loop{
			break

		}


	}

}











