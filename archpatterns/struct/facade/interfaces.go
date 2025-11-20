package main

type IUserRegisterFacade interface {
	Register(email string, password string) error
}

type IAuthService interface {
	UserRegister(email string, password string) error
}

type INotifyService interface {
	SendSuccessRegisterEmail(email string) error
	SendFailRegisterEmail(email string) error
}
