import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import UserList from './UserList';

const BlockedUsers = (props) => {

    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [buttonClick, setButtonClick] = useState(0);
    const [blocked, setBlocked] = useState();
    const [blockers, setBlockers] = useState();
    const buttonClickChanger = () => setButtonClick(buttonClick + 1);


    const uniqueId = (id) => {
        let goodToGo = true
        if (users)
            users.every((user) => {
                if (id === user.id) {
                    goodToGo = false;
                    return false;
                }
                return true;
            })
        return goodToGo;
    }

    const fetchUsers = async (blocked) => {
        setUsers();
        setLoading(true);
        blocked.forEach(user => {
            axios.get(axios.defaults.baseURL + 'users/' + user.id)
                .then(res => {
                    let newUsers = users
                    if (uniqueId()) {
                        newUsers = newUsers.concat(res.data.user);
                        setUsers(Array.from(newUsers));
                    }
                    //console.log(users);
                }).catch(err => {
                    console.log(err)
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: err.data,
                    });
                });
        })
        setLoading(false);
    };

    const fetchBlocked = async () => {
        let id = localStorage.getItem('user_id');
        setBlocked();
        axios.get(axios.defaults.baseURL + 'followers/blocked/' + id)
            .then(res => {
                let blocked = res.data.ids;
                //console.log(res);
                setBlocked(blocked);
                fetchUsers(blocked);
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
        fetchBlocked();
    }, []);

    /*useEffect(() => {
        fetchUsers();
    }, [])*/

    useEffect(() => {
        fetchBlocked();
    }, [buttonClick])

    const getUsers = () => {
        console.log(users);
        let filteredUsers = []
        let goodToGo
        users.forEach((user) => {
            goodToGo = true
            filteredUsers.every((checkUser) => {
                if (checkUser.id === user.id) {
                    goodToGo = false;
                    return false;
                }
                return true;
            })
            if (goodToGo)
                filteredUsers = filteredUsers.concat(user);
        })
        return Array.from(filteredUsers);
    }

    return (
        <div>
            {loading && <h3>Loading...</h3>}
            {!loading && users && <UserList users={getUsers()} buttonClickChanger={buttonClickChanger} blocked={true} />}
            {!users && <h3 style={{ textAlign: "center" }}>No blocked users.</h3>}
        </div>
    );
}
export default BlockedUsers;