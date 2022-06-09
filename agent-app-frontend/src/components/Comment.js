import { useEffect, useState } from "react";
import axios from "axios";

const Comment = (props)=>{

    const[firstName, setFirstName] = useState('')
    const[lastName, setLastName] = useState('')

    useEffect(() => {
        const getUserNameById = async () => {
            axios.get(axios.defaults.baseURL + 'api/Users/' + props.comment.userId)
                    .then(res => {
                        //console.log(res.data)
                        setFirstName(res.data.personalInfo.firstName)
                        setLastName(res.data.personalInfo.lastName)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        getUserNameById();
    }, []);

    return(
        <li key={props.index} class="list-group-item">
            {firstName} {lastName} : {props.comment.content}</li>
    );
}

export default Comment;