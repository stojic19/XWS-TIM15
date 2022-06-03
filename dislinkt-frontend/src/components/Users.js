import axios from "axios";
import { useEffect, useState } from "react"
import Swal from "sweetalert2";
import UserList from "./UserList";

const Users = () => {
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchUsers = async () => {
            setLoading(true);
            axios.get(axios.defaults.baseURL + 'users')
            .then(res => {
                let users = Array.from(res.data.users)
                setUsers(users);
                setLoading(false);
                console.log(users);
            }).catch(err =>{
                console.log(err)
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
        };
        fetchUsers();
    }, [])

    return (
        <div>
            {loading && <h3>Loading...</h3>}
            {!loading && users && <UserList users={users} />}
        </div>
    );
}

export default Users;