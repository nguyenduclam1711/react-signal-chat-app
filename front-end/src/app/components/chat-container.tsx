import MessageViewer from "./message-viewer";
import ChatInput from "./chat-input";
import { useChatContext } from "./chat-context/hooks";

const ChatContainer = () => {
  const { messages } = useChatContext();
  return (
    <div className="flex flex-col gap-2 justify-between h-screen">
      <div className="flex-1 overflow-y-auto">
        {messages.length > 0 && (
          <div className="px-2">
            {messages.map((message, index) => {
              return (
                <MessageViewer key={`message-${index}`} {...message} />
              );
            })}
          </div>
        )}
      </div>
      <ChatInput />
    </div>
  );
};

export default ChatContainer;
