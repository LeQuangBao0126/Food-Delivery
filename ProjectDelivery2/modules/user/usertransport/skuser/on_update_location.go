package skuser

import (
	"ProjectDelivery2/common"
	"ProjectDelivery2/component"
	"fmt"
	"github.com/googollee/go-socket.io"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
func OnUserUpdateLocation ( appContext component.AppContext ,requester common.Requester ) func(s  socketio.Conn, location LocationData) {
	return func(s  socketio.Conn,  location LocationData){
		fmt.Println("User update location ",requester.GetUserId() ,"   location %v", location)
	}
}