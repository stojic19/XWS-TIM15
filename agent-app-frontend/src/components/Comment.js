import { useEffect, useState } from "react";
import axios from "axios";

const Comment = (props)=>{

    const[name, setName] = useState('')

    useEffect(() => {
        const getUserNameById = async () => {
            axios.get(axios.defaults.baseURL + '/api/Users/' + props.comment.userId)
                    .then(res => {
                        console.log(res.data)
                        setName(res.data.user.name)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        getUserNameById();
    }, []);

    return(
        <li key={props.index} class="list-group-item">
            {name} : {props.comment.content}</li>
    );
}

export default Comment;