"use client";

import { BackEndMessageType, PAYLOAD_TYPE_CONNECTED, PAYLOAD_TYPE_MESSAGE } from "@/utils/messages";
import { MAX_RETRY_TIMES, socketUrl } from "@/utils/socket-connect";
import { ReactElement, createContext, useEffect, useState } from "react";

type ChatContextType = {
  messages: BackEndMessageType[],
  socket: null | WebSocket,
  socketId: null | string,
}

export const ChatContext = createContext<ChatContextType>({
  messages: [],
  socket: null,
  socketId: null,
});

type ChatContextProviderProps = {
  children: ReactElement,
}

const ChatContextProvider = (props: ChatContextProviderProps) => {
  const { children } = props;
  const [messages, setMessages] = useState<BackEndMessageType[]>([]);
  const [socket, setSocket] = useState<null | WebSocket>(null);
  const [socketId, setSocketId] = useState<null | string>(null);
  const [retryTimes, setRetryTimes] = useState(0);

  const onAddMessage = (payload: BackEndMessageType) => {
    setMessages(messages => [...messages, payload]);
  };

  useEffect(() => {
    if (!socket && retryTimes < MAX_RETRY_TIMES) {
      const newSocket = new WebSocket(socketUrl);

      newSocket.onopen = function(e) {
        console.log("Connected");
        setSocket(newSocket);
        setRetryTimes(0);
      };

      newSocket.onerror = function(this, e) {
        console.log("Connection error", e);
        setRetryTimes(retryTimes + 1);
      };

      newSocket.onclose = function(e) {
        console.log("Connection close");
      };

      newSocket.onmessage = function(e) {
        console.log(`Received message: ${e.data}`);
        let payload: null | Record<string, any> = null;
        try {
          payload = JSON.parse(e.data);
        } catch (error) {
          console.log("Cannot parse message payload: ", error);
          return;
        }
        if (!payload || !payload.type) {
          console.log("Payload doesn't have type");
          return;
        }
        switch (payload.type) {
          case PAYLOAD_TYPE_CONNECTED: {
            // handle payload type connected
            setSocketId(payload.socketId);
            break;
          }
          case PAYLOAD_TYPE_MESSAGE: {
            // handle payload type message
            onAddMessage(payload as BackEndMessageType);
            break;
          }
        }
      };
    }
  }, [socket, retryTimes]);

  return (
    <ChatContext.Provider value={{
      messages,
      socket,
      socketId,
    }}>
      {children}
    </ChatContext.Provider>
  );
};

export default ChatContextProvider;
