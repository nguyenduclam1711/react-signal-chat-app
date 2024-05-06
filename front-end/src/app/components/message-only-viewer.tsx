import dayjs from "dayjs";

type MessageViewerProps = {
  data: string,
  createdTime: string,
};

const MessageOnlyViewer = (props: MessageViewerProps) => {
  const { data, createdTime } = props;

  return (
    <div className="rounded-lg bg-sky-900 p-2 max-w-100">
      <p className="mb-1 break-words text-sm">
        {data}
      </p>
      <p className="italic text-xs text-gray-300">
        {dayjs(createdTime).format("MM/DD/YYYY HH:mm:ss")}
      </p>
    </div>
  );
};

export default MessageOnlyViewer;
