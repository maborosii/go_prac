package main

import (
	"fmt"

	"github.com/maborosii/designpattern/factory"
)

func main() {
	// user := factory.NewUser()
	// admin := factory.NewAdmin()

	// 空接口的断言获取其类型
	// 将工厂方法与业务逻辑(1, "aaa")解耦，保证工厂方法的通用性，抽象性
	user := factory.CreateUser(factory.AdminUser)(1, "aaa").(*factory.Admin)
	fmt.Println(user)
}
