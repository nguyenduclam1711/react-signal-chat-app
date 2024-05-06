package modules

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
)

func removeUserAndSocketFromMapping(mapUserIdToSocket, mapSocketToUserId map[string]string, payload *socketio.EventPayload) {
	socketId := payload.Kws.UUID
	userId := mapSocketToUserId[socketId]

	delete(mapUserIdToSocket, userId)
	delete(mapSocketToUserId, socketId)
}

func CreateChatModule(app *fiber.App) {
	// save client ids
	// map userId and socketId
	mapUserIdToSocket := map[string]string{}
	mapSocketToUserId := map[string]string{}

	// Setup the middleware to retrieve the data sent in first GET request
	app.Use(func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	socketio.On(socketio.EventDisconnect, func(payload *socketio.EventPayload) {
		fmt.Println("On Disconnect", payload.Kws.UUID)
		removeUserAndSocketFromMapping(mapUserIdToSocket, mapSocketToUserId, payload)
	})

	socketio.On(socketio.EventClose, func(payload *socketio.EventPayload) {
		fmt.Println("On Close", payload.Kws.UUID)
		removeUserAndSocketFromMapping(mapUserIdToSocket, mapSocketToUserId, payload)
	})

	socketio.On(socketio.EventMessage, func(payload *socketio.EventPayload) {
		fmt.Println("On message from", payload.Kws.UUID, " message ", string(payload.Data))
		channels := []string{}
		var currUserId string
		for socket, userId := range mapSocketToUserId {
			if socket == payload.Kws.UUID {
				currUserId = userId
			}
			channels = append(channels, socket)
		}

		if currUserId == "" {
			fmt.Println("Cant find userId for socket ", payload.Kws.UUID)
			return
		}

		messageObj := models.NewMessageObject()
		err := json.Unmarshal(payload.Data, &messageObj)
		if err != nil {
			fmt.Println("Cant parse json from user ", currUserId)
			return
		}
		payloadMessage, payloadMessageErr := json.Marshal(messageObj)
		if payloadMessageErr != nil {
			fmt.Println("Cant encode to string", messageObj)
			return
		}
		payload.Kws.EmitToList(channels, payloadMessage, socketio.TextMessage)
	})

	app.Get("/ws/:id", socketio.New(func(kws *socketio.Websocket) {
		userId := kws.Params("id")
		mapUserIdToSocket[userId] = kws.UUID
		mapSocketToUserId[kws.UUID] = userId

		connectedObj := models.NewConnectedObject()
		connectedObj.UserId = userId
		connectedObj.SocketId = kws.UUID
		data, err := json.Marshal(connectedObj)
		if err != nil {
			fmt.Println("Cant encode json from user ", userId)
			return
		}
		kws.Emit(data)
	}))
}
