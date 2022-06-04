import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import UserList from './UserList';

const PublicProfileSearch = () =>{

    const[users, setUsers] = useState('');
    const [loading, setLoading] = useState(true);
    const[searchTerm, setSearchTerm] = useState('');

    const fetchUsers = async () => {
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'users/searchPublicUsers/' + searchTerm)
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

    useEffect(() => {
        fetchUsers();
    }, [searchTerm])

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
            {!loading && users && <UserList users={users} />}
        </div>
    );
}
export default PublicProfileSearch