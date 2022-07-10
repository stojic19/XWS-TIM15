import { useState, useEffect } from "react";
import axios from "axios";

const NotificationUserName = (props) =>{

    const [name, setName] = useState();

    useEffect(() => {
        fetchFollower();
    }, [])

    const fetchFollower = () =>{
        axios.get(axios.defaults.baseURL + 'users/' + props.followerId)
            .then(res => {
                //console.log(res)
                setName(res.data.user.name)
            });

    }

    return(
        <>{name}</>
    );
}

export default NotificationUserName;