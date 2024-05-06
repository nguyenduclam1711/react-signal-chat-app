import { v4 as uuidv4 } from "uuid";

export const mockUserId = uuidv4();

const socketUrl = `ws://localhost:3000/ws/${mockUserId}`;

export const socket = new WebSocket(socketUrl);

if (socket) {
  socket.onopen = function (e) {
    console.log("Connected");
  };

  socket.onerror = function (this, e) {
    console.log("Connection error");
  };

  socket.onclose = function (e) {
    console.log("Connection close");
  };

  socket.onmessage = function (e) {
    console.log(`Received message: ${e.data}`);
  };
}
