package payment

import (
	"log"
	msasubmoduleci "user/pkg/proto/userpb/userpb"
)

func main() {
	user := msasubmoduleci.User{
		Id:        "1",
		LastName:  "test",
		FirstName: "test",
		Age:       31,
	}
	log.Println(user.Id, user.FirstName, user.LastName, user.Age)
}
