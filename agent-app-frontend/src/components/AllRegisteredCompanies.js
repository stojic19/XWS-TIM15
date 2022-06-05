import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import CompaniesList from './CompaniesList';

const AllRegisteredCompanies = () =>{

    const[companies, setCompanies] = useState([]);
    const[loading, setLoading] = useState(true);

    const fetchRequests = async () => {
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'api/Companies/Registered')
        .then(res => {
            console.log(res);
            let companies = Array.from(res.data)
            setCompanies(companies);
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
            {!loading && companies && <CompaniesList companies={companies} admin={false} />}
        </div>
    );
}
export default AllRegisteredCompanies;