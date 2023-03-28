package main

import (
	"fmt"
	"shuser/models"
)

func main() {
	models.Init()
	for {
		//输入需要删除的账号
		fmt.Println("请输入需要删除的账号")
		var res string
		fmt.Scanln(&res)
		fmt.Println("y确认删除，n撤销删除")
		var x string
		fmt.Scanln(&x)
		switch x {
		case "y":
			//查找账号并删除
			user := &models.User{}
			//Unscoped()数据库彻底删除
			tx := models.DB.Table("user").Where("name = ?", &res).Unscoped().Delete(&user)
			if tx.RowsAffected == 0 {
				fmt.Println("未找到该账户，删除失败")
			} else {
				fmt.Println("删除成功")
			}
		case "n":
			fmt.Println("已撤销删除")
		default:
			fmt.Println("操作错误，请重试")
		}
		fmt.Println("=·=·=·=·=·=·=·=·=·=分隔符=·=·=·=·=·=·=·=·=·=")
	}
}
