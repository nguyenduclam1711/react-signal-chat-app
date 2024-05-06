"use client";

import { mockUserId, socket } from "@/utils/socket-connect";

export default function Home() {
  const onSendMessage = () => {
    socket.send(`Fake message from ${mockUserId}`);
  };

  return (
    <main>
      <div>Chat App</div>
      <button onClick={onSendMessage}>Send a mock message</button>
    </main>
  );
}
