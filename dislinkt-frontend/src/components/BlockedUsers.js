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
    const buttonClickChanger = () => setButtonClick(buttonClick + 1);


    const filterUsers = (users) =>{
        const uniqueIds = [];

        const unique = users.filter(element => {
        const isDuplicate = uniqueIds.includes(element.id);

        if (!isDuplicate) {
            uniqueIds.push(element.id);

            return true;
        }

        return false;
        });

        console.log(unique)
        setUsers(unique)
    }

    const fetchUsers = async (blockedList) => {
        setUsers()
        setLoading(true);
        blockedList.forEach(user => {
            axios.get(axios.defaults.baseURL + 'users/' + user.id)
                .then(res => {
                    let newUsers = users
                    newUsers.push(res.data.user);
                    setUsers(Array.from(newUsers));
                    filterUsers(users);

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
                //console.log(blocked)
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


    return (
        <div>
            {<h2 style={{ textAlign: "center" }}>Blocked users</h2>}
            {loading && <h3>Loading...</h3>}
            {!loading && users && <UserList users={users} buttonClickChanger={buttonClickChanger} blocked={true} />}
            {!users && <h3 style={{ textAlign: "center" }}>No blocked users.</h3>}
        </div>
    );
}
export default BlockedUsers;