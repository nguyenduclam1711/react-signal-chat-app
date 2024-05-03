import { Server } from "socket.io";
import express from "express";
import http from "http";

const app = express();
const server = http.createServer(app);

const io = new Server(server, {
  // options
  cors: {
    origin: "*",
  },
});

io.of("/chat").on("connection", (socket) => {
  console.log("New connection", socket.id);
});

app.get("/", (req, res) => {
  res.send("hello world");
});

server.listen(3000);
