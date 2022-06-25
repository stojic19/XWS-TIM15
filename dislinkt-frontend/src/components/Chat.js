import { useRef, useState, useEffect } from 'react';
import ChatMessage from './ChatMessage';
import '../css/chat.css'
import { useParams } from "react-router-dom";
import axios from 'axios';
import Swal from 'sweetalert2';

const Chat = () => {

    const { id } = useParams();
    const dummy = useRef();
    //const query = messagesRef.orderBy('createdAt').limit(25);
    const [messages, setMessages] = useState();
    //useCollectionData(query, { idField: 'id' });
    const [formValue, setFormValue] = useState('');
    const [loading, setLoading] = useState(true);

    const fetchChat = async () => {
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'chat/' + id)
            .then(res => {
                console.log(res)
                let chats = Array.from(res.data.chats)
                let messages = []
                chats.every((chat)=>{
                    if(chat.firstParticipantId===localStorage.getItem('user_id')||chat.secondParticipantId===localStorage.getItem('user_id')){
                        if(chat.firstParticipantId===id||chat.secondParticipantId===id){
                            messages = Array.from(chat.messages);
                            return false;
                        }
                        return true;
                    }
                    return true;
                })
                setMessages(messages.sort());
                setLoading(false);
            }).catch(err => {
                console.log(err)
                if (err.response.status != 404)
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: err.data,
                    });
            });
    };

    useEffect(() => {
        if(formValue==='')
        fetchChat();
    }, [formValue])
    const sendMessage = async (e) => {
        e.preventDefault();

        const data = {
            senderId: localStorage.getItem('user_id'),
            receiverId: id,
            text: formValue
        }

        axios.post(axios.defaults.baseURL + 'chat', data)
            .then(() => {
                /*Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });*/

                setFormValue('');
                dummy.current.scrollIntoView({ behavior: 'smooth' });
            }).catch(err => {
                console.log(err)
            });


    }

    return (
        <div className="chatApp">
            <main>

                {messages && messages.map(msg => <ChatMessage key={msg.timeSent} message={msg} />)}

                <span ref={dummy}></span>

            </main>

            <form onSubmit={sendMessage}>

                <input value={formValue} onChange={(e) => setFormValue(e.target.value)} placeholder="say something nice" />

                <button type="submit" disabled={!formValue}>🕊️</button>

            </form>
        </div>
    )
}

export default Chat