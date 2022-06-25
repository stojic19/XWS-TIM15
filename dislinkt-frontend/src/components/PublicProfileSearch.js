import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import UserList from './UserList';

const PublicProfileSearch = (props) =>{

    const[users, setUsers] = useState('');
    const [loading, setLoading] = useState(true);
    const[searchTerm, setSearchTerm] = useState('');
    const[follows, setFollows] = useState([]);
    const[followRequests, setFollowRequests] = useState();
    const[buttonClick, setButtonClick] = useState(0);

    const buttonClickChanger = () => setButtonClick(buttonClick+1);

    const fetchUsers = async () => {
        setLoading(true);
        setUsers();
        axios.get(axios.defaults.baseURL + 'users/searchPublicUsers/' + searchTerm)
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
    useEffect(() => {   
        if(props.displayFollowButtons){
            fetchFollowRequests();
            fetchFollows();
        }
    }, []);

    useEffect(() => {
        fetchUsers();
    }, [searchTerm])

    useEffect(() => {
        //console.log(buttonClick)
        fetchUsers();
        fetchFollowRequests();
        fetchFollows();
    }, [buttonClick])

    return(
        <div>
	    <div className="col-12">
    	    <div id="custom-search-input">
                <div className="input-group">
                    <input type="text" className="search-query form-control" placeholder="Search" 
                    onKeyPress={(ev) => {
                        if (ev.key === "Enter") {
                        ev.preventDefault();
                        setSearchTerm(ev.target.value)
                        }
                    }} />
                    <span className="input-group-btn">
                        <button type="button" disabled>
                            <span className="fa fa-search"></span>
                        </button>
                    </span>
                </div>
            </div>
        </div>
            {loading && <h3>Loading...</h3>}
            {!loading && users && follows && followRequests && <UserList users={users} displayFollowButtons={props.displayFollowButtons} follows={follows} followRequests={followRequests} buttonClickChanger={buttonClickChanger}/>}
        </div>
    );
}
export default PublicProfileSearch