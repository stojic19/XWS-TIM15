import { useEffect, useState } from "react";
import { Dropdown } from "react-bootstrap";
import { useNavigate } from "react-router-dom"
import axios from "axios";
import Swal from "sweetalert2";
import NotificationFollowerName from "./NotificationFollowerName";

const NotificationsList = () =>{

    const history = useNavigate()
    const [notifications, setNotifications] = useState();
    const [follower, sentFollower] = useState();

    const fetchNotifications = async () => {

        axios.get(axios.defaults.baseURL + 'notifications/user/' + localStorage.getItem('user_id'))
            .then(res => {
                //console.log(res.data)
                let notifications = Array.from(res.data)
                setNotifications(notifications)
            }).catch(err => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    }

    useEffect(() => {
        fetchNotifications();
    }, [])

    const openNotification = (notification) =>{
        switch(notification.type){
            case 'message':
                history('/chat/' + notification.messagesId);
                window.location.reload(true);
                break;
            case 'profile':
                history('/post/' + notification.postId);
                window.location.reload(true);
                break;
            case 'post':
                history('/post/' + notification.postId);
                window.location.reload(true);
        }
    }

    return (
        <div>
            <Dropdown>
                <Dropdown.Toggle variant="light" id="dropdown-basic">
                    <img src={require("../images/notification.png")}  style={{ height: "20px", width:"20px" }}/>
                </Dropdown.Toggle>
                <Dropdown.Menu>
                    {notifications &&
                    notifications.map((n, index) => {
                        return (
                            <div key={index}>
                            <Dropdown.Item onClick={() => openNotification(n)}>
                                <NotificationFollowerName followerId={n.followerId}></NotificationFollowerName>
                                {
                                    n.type == 'post' ? (n.action == 'like' ? ' liked your post.' : ' commented on your post.') : (n.type == 'message' ? ' send you a message.' : ' shared a new post.')
                                }
                            </Dropdown.Item>
                            <Dropdown.Divider />
                            </div>
                        );
                    })}
                    {notifications && notifications.length==0 && <Dropdown.Item>No notifications to show.</Dropdown.Item>}
                </Dropdown.Menu>
            </Dropdown>
        </div>
    );
}

export default NotificationsList;