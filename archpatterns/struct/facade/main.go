package main

func main() {
	userRegisterFacade := NewUserRegisterFacade(NewAuthService(), NewNotifyService())

	if err := userRegisterFacade.Register("someemail@email.ru", "123"); err != nil {
		println(err.Error())
	}
}
