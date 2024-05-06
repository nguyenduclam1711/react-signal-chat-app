import { useContext } from "react";
import { ChatContext } from ".";

export const useChatContext = () => {
  return useContext(ChatContext);
};
