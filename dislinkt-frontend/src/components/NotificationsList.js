import { useEffect, useState } from "react";
import { Dropdown } from "react-bootstrap";
import axios from "axios";
import Swal from "sweetalert2";

const NotificationsList = () =>{

    const [notifications, setNotifications] = useState();

    const fetchNotifications = async () => {
        axios.get('http://localhost:8081/notifications/user/0a93d6c1-ef32-4287-b7db-8ad566481d53')
            .then(res => {
                console.log(res.data)
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


    return (
        <div>
            <Dropdown>
                <Dropdown.Toggle variant="light" id="dropdown-basic">
                    <img src={require("../images/notification.png")}  style={{ height: "20px", width:"20px" }}/>
                </Dropdown.Toggle>

                <Dropdown.Menu>
                    <Dropdown.Item href="#/action-1">Milan Mikic liked your post. (3s ago)</Dropdown.Item>
                    <Dropdown.Item href="#/action-2">Milan Mikic commented on your post. (30s ago)</Dropdown.Item>
                    <Dropdown.Item href="#/action-3">Milos Trivic send you a message. (1h ago)</Dropdown.Item>
                    <Dropdown.Divider />
                    <Dropdown.Item eventKey="4">Show all</Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>
        </div>
    );
}

export default NotificationsList;