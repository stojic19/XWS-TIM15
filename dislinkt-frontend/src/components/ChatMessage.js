const ChatMessage = (props) => {
    const { text, senderId, timeSent } = props.message;

    const messageClass = senderId === localStorage.getItem('user_id') ? 'sent' : 'received';
  
    return (<>
      <div className={`chatmessage ${messageClass}`}>
        <img src={require('../images/user-avatar.png')} />
        <p>{text}</p>
      </div>
    </>)
}

export default ChatMessage