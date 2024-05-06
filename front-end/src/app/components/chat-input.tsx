import { getMessagePayload } from "@/utils/messages";
import { mockUserId } from "@/utils/socket-connect";
import { KeyboardEventHandler, useRef } from "react";
import { useChatContext } from "./chat-context/hooks";

const ChatInput = () => {
  const ref = useRef<HTMLInputElement>(null);
  const { socketId, socket } = useChatContext();

  const onSendMessage: KeyboardEventHandler<HTMLInputElement> = (e) => {
    if (!ref.current || !socketId || !socket) {
      return;
    }
    if (e.key === "Enter") {
      const messagePayload = getMessagePayload({
        userId: mockUserId,
        data: ref.current.value,
        socketId,
      });
      socket.send(JSON.stringify(messagePayload));
      ref.current.value = "";
    }
  };

  return (
    <div className="h-13 px-2">
      <input
        onKeyUp={onSendMessage}
        ref={ref}
        className="w-full rounded border-gray-50 border-2 bg-transparent h-full p-2"
      />
    </div>
  );
};

export default ChatInput;
