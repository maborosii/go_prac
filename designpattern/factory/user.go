package factory

type ICreateUser func(id int, name string) interface{}

type User struct {
	Id       int
	Username string
}

func NewUser() ICreateUser {
	return func(id int, name string) interface{} {
		return &User{Id: id, Username: name}
	}
}

type Admin struct {
	Id        int
	Adminname string
	Role      string
}

func NewAdmin() ICreateUser {
	return func(id int, name string) interface{} {
		return &Admin{Id: id, Adminname: name, Role: "admin"}
	}
}
