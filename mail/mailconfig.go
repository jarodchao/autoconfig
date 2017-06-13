package mail

import (
	"net/smtp"

	"autoconfig/type"
	"fmt"
	"net/mail"
)

var Auth LoginAuth
var Form mail.Address
var ServerAddr string

type LoginAuth struct {
	username, password string
}

func loginAuth(username, password string) smtp.Auth {
	return &LoginAuth{username, password}
}

func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// return "LOGIN", []byte{}, nil
	return "LOGIN", []byte(a.username), nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

func InitMail(m autoconfig.Mail) {
	Auth = loginAuth(m.UserName, m.Password)
	Form = mail.Address{m.SenderName, m.SenderMail}
	ServerAddr = fmt.Sprintf("%s:%s", m.Server, m.Port)


}
