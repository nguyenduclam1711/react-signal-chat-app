"use client";
import { io } from "socket.io-client";

// make the connection through socket io
const socket = io("localhost:3000/chat");

socket.on("connect", () => {
  console.log("Connected to chat");
});

socket.on("disconnect", (reason, details) => {
  console.log("Disconnected from chat");
  console.log("Reason: ", reason, "Details: ", details);
});

socket.on("connect_error", (error) => {
  console.log("Connection error", error.message);
});

export default function Home() {
  return (
    <main>
      <div>Chat App</div>
    </main>
  );
}
