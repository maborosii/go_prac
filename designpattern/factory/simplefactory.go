package factory

// 简单工厂模式
type UserType int

const (
	FrontUser = iota
	AdminUser
)

func CreateUser(t UserType) ICreateUser {
	switch t {
	case FrontUser:
		return NewUser()
	case AdminUser:
		return NewAdmin()
	default:
		return NewUser()
	}
}
