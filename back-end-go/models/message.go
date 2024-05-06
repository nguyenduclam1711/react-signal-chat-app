package models

import "time"

const (
	PAYLOAD_TYPE_MESSAGE   = "PAYLOAD_TYPE_MESSAGE"
	PAYLOAD_TYPE_CONNECTED = "PAYLOAD_TYPE_CONNECTED"
)

type FrontendMessageObject struct {
	Type     string `json:"type"`
	UserId   string `json:"userId"`
	SocketId string `json:"socketId"`
	Data     string `json:"data"`
}

type BackendMessageObject struct {
	FrontendMessageObject
	CreatedTime time.Time `json:"createdTime"`
}

func NewMessageObject() BackendMessageObject {
	return BackendMessageObject{
		FrontendMessageObject: FrontendMessageObject{
			Type: PAYLOAD_TYPE_MESSAGE,
		},
		CreatedTime: time.Now(),
	}
}

type ConnectedObject struct {
	Type     string `json:"type"`
	UserId   string `json:"userId"`
	SocketId string `json:"socketId"`
}

func NewConnectedObject() ConnectedObject {
	return ConnectedObject{
		Type: PAYLOAD_TYPE_CONNECTED,
	}
}
