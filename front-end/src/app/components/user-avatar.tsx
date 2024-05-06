
type UserAvatarProps = {
  userId: string,
}
const UserAvatar = (props: UserAvatarProps) => {
  const { userId } = props;

  return (
    <div className="text-gray-400 italic text-sm mb-2">
      User {userId}
    </div>
  );
};

export default UserAvatar;
