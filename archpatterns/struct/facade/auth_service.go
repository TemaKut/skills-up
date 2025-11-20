package main

import "fmt"

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) UserRegister(email string, password string) error {
	return fmt.Errorf("error from db")
}
