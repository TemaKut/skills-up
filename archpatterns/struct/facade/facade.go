package main

import "fmt"

type UserRegisterFacade struct {
	authService   IAuthService
	notifyService INotifyService
}

func NewUserRegisterFacade(authService IAuthService, notifyService INotifyService) *UserRegisterFacade {
	return &UserRegisterFacade{
		authService:   authService,
		notifyService: notifyService,
	}
}

func (u *UserRegisterFacade) Register(email string, password string) error {
	var userRegisterErr error

	if err := u.authService.UserRegister(email, password); err != nil {
		userRegisterErr = fmt.Errorf("error user register. %w", err)

		if err := u.notifyService.SendFailRegisterEmail(email); err != nil {
			userRegisterErr = fmt.Errorf("%w +++ error notify fial register. %w", userRegisterErr, err)
		}
	}

	if userRegisterErr != nil {
		return userRegisterErr
	}

	if err := u.notifyService.SendSuccessRegisterEmail(email); err != nil {
		return fmt.Errorf("error send success email. %w", err)
	}

	return nil
}
