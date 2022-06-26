import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import UserList from './UserList';

const AllProfiles = (props) =>{

    const[users, setUsers] = useState('');
    const [loading, setLoading] = useState(true);
    const[follows, setFollows] = useState([]);
    const[followRequests, setFollowRequests] = useState();
    const[buttonClick, setButtonClick] = useState(0);
    const [blocked, setBlocked] = useState();
    const [blockers, setBlockers] = useState();
    const buttonClickChanger = () => setButtonClick(buttonClick+1);

    const fetchUsers = async () => {
        setLoading(true);
        setUsers();
        axios.get(axios.defaults.baseURL + 'users')
        .then(res => {
            let users = Array.from(res.data.users)
            setUsers(users);
            setLoading(false);
            //console.log(users);
        }).catch(err =>{
            console.log(err)
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: err.data,
            });
        });
    };
    const fetchFollowRequests = async () => {
        let id = localStorage.getItem('user_id');
        setFollowRequests();
        axios.get(axios.defaults.baseURL + 'followers/followRequests/' + id)
            .then(res => {
                let followRequests = res.data.followRequests;
                setFollowRequests(followRequests);
                //console.log(res);
            }).catch(err => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    };
    const fetchFollows = async () => {
        let id = localStorage.getItem('user_id');
        setFollows();
        axios.get(axios.defaults.baseURL + 'followers/follows/' + id)
            .then(res => {
                let follows = res.data.follows;
                //console.log(res);
                setFollows(follows);
            }).catch(err => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    };
    const fetchBlocked = async () => {
        let id = localStorage.getItem('user_id');
        setBlocked();
        axios.get(axios.defaults.baseURL + 'followers/blocked/' + id)
            .then(res => {
                let blocked = res.data.ids;
                //console.log(res);
                setBlocked(blocked);
            }).catch(err => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    };
    const fetchWhoBlockedMe = async () => {
        let id = localStorage.getItem('user_id');
        setBlockers();
        axios.get(axios.defaults.baseURL + 'followers/blockers/' + id)
            .then(res => {
                let blockers = res.data.ids;
                //console.log(res);
                setBlockers(blockers);
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
        if(props.displayFollowButtons){
            fetchFollowRequests();
            fetchFollows();
            fetchBlocked();
            fetchWhoBlockedMe();
        }
    }, []);

    useEffect(() => {
        fetchUsers();
    }, [])

    useEffect(() => {
        //console.log(buttonClick)
        fetchUsers();
        fetchFollowRequests();
        fetchFollows();
    }, [buttonClick])

    const getUsers = () => {
        if (blocked && blockers) {
            let userId = localStorage.getItem('user_id');
            let filteredUsers = []
            let goodToGo
            users.forEach((user) => {
                goodToGo = true
                blocked.every((checkUser) => {
                    if (checkUser.id === user.id) {        
                        goodToGo = false;
                        return false;
                    }
                    return true;
                })
                if (goodToGo) {
                    blockers.every((checkUser) => {
                        if (checkUser.id === user.id) {
                            goodToGo = false;
                            return false;
                        }
                        return true;
                    })
                    if (goodToGo)
                        filteredUsers = filteredUsers.concat(user);
                }
            })
            return filteredUsers;
        }
        return users;
    }

    return(
        <div>
            {loading && <h3>Loading...</h3>}
            {!loading && getUsers() && follows && followRequests && <UserList users={getUsers()} displayFollowButtons={props.displayFollowButtons} follows={follows} followRequests={followRequests} buttonClickChanger={buttonClickChanger} />}
        </div>
    );
}
export default AllProfiles;