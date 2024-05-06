import { BackEndMessageType } from "@/utils/messages";
import { mockUserId } from "@/utils/socket-connect";
import { useMemo } from "react";
import UserAvatar from "./user-avatar";
import MessageOnlyViewer from "./message-only-viewer";

type MessageViewerProps = BackEndMessageType;

const MessageViewer = (props: MessageViewerProps) => {
  const { userId, data, createdTime } = props;
  const isCurrentUser = userId === mockUserId;

  const className = useMemo(() => {
    let result = "flex mb-4";
    if (isCurrentUser) {
      result += " justify-end";
    }
    return result;
  }, [isCurrentUser]);

  return (
    <div className={className}>
      <div>
        <UserAvatar userId={userId} />
        <MessageOnlyViewer data={data} createdTime={createdTime} />
      </div>
    </div>
  );
};

export default MessageViewer;
