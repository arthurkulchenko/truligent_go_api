package services

import(
	"fmt"
)

func (u *User) OmsPingCall() string {
	fmt.Println("Ping... Pong!")
	return "Pong"
}
