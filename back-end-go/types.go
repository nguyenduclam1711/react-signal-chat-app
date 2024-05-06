package main

const (
	PAYLOAD_TYPE_MESSAGE   = "PAYLOAD_TYPE_MESSAGE"
	PAYLOAD_TYPE_CONNECTED = "PAYLOAD_TYPE_CONNECTED"
)

type MessageObject struct {
	Type     string `json:"type"`
	UserId   string `json:"userId"`
	SocketId string `json:"socketId"`
	Data     string `json:"data"`
}

func newMessageObject() MessageObject {
	return MessageObject{
		Type: PAYLOAD_TYPE_MESSAGE,
	}
}

type ConnectedObject struct {
	Type     string `json:"type"`
	UserId   string `json:"userId"`
	SocketId string `json:"socketId"`
}

func newConnectedObject() ConnectedObject {
	return ConnectedObject{
		Type: PAYLOAD_TYPE_CONNECTED,
	}
}
