import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import CompaniesList from './CompaniesList';

const Requests = () =>{

    const[requests, setRequests] = useState([]);
    const[loading, setLoading] = useState(true);

    const fetchRequests = async () => {
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'api/Companies/NotRegistered')
        .then(res => {
            console.log(res);
            let companies = Array.from(res.data)
            setRequests(companies);
            setLoading(false);
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
        fetchRequests();
    }, []);

    return(
        <div>
            {loading && <h3>Loading...</h3>}
            {!loading && requests && <CompaniesList companies={requests} admin={true} />}
        </div>
    );
}
export default Requests;