import ChatContainer from "./components/chat-container";
import ChatContextProvider from "./components/chat-context";

export default function Home() {
  return (
    <main>
      <ChatContextProvider>
        <ChatContainer />
      </ChatContextProvider>
    </main>
  );
}
