export const PAYLOAD_TYPE_MESSAGE = "PAYLOAD_TYPE_MESSAGE";
export const PAYLOAD_TYPE_CONNECTED = "PAYLOAD_TYPE_CONNECTED";

export type MessageType = {
  type: string,
  userId: string,
  socketId: string,
  data: string,
};

export type BackEndMessageType = MessageType & {
  createdTime: string,
}

export const getMessagePayload = ({
  userId,
  socketId,
  data,
}: {
  userId: string,
  socketId: string,
  data: string,
}): MessageType => {
  return {
    type: PAYLOAD_TYPE_MESSAGE,
    userId,
    socketId,
    data,
  };
};
