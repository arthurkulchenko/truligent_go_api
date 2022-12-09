package services

import(
	"fmt"
)

func (c *Client) OmsPingCall() string {
	connection := fmt.Sprintf("%v", c.dbConn)
	fmt.Println(connection)
	return "Pong"
}
