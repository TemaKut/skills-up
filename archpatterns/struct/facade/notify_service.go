package main

import "fmt"

type NotifyService struct{}

func NewNotifyService() *NotifyService {
	return &NotifyService{}
}

func (n *NotifyService) SendSuccessRegisterEmail(email string) error {
	println("SendSuccessRegisterEmail")

	return nil
}

func (n *NotifyService) SendFailRegisterEmail(email string) error {
	return fmt.Errorf("error from message broker")
}
