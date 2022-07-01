import { useEffect, useState } from "react"
import axios from "axios";
import Swal from "sweetalert2";

const PostOwner = (props) =>{
    
    const [user, setUser] = useState({});

    const fetchUser = async () => {

        axios.get(axios.defaults.baseURL + 'users/' + props.id)
            .then(res => {
                setUser(res.data.user);
                console.log(user)
            }).catch(err => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    };

    useEffect(() => {
        fetchUser();
    }, []);

    return (
        <div>
            <h3>{user.name}</h3>
        </div>
    )
}

export default PostOwner;