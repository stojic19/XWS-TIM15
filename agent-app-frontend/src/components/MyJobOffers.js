import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import JobOffersList from './JobOffersList';

const MyJobOffers = () => {

    const [jobOffers, setJobOffers] = useState([]);
    const [loading, setLoading] = useState(true);

    const fetchJobOffers = async () => {
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'api/Companies/User/' + localStorage.getItem('user_id'))
            .then(res => {
                let data = res.data;
                console.log(data);
                let jobOffers = Array.from(res.data[0].jobOffers)
                setJobOffers(jobOffers);
                setLoading(false);
            }).catch(err => {
                console.log(err)
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    };

    useEffect(() => {
        fetchJobOffers();
    }, [])

    return (
        <div>
            {loading && <h3>Loading...</h3>}
            {!loading && jobOffers && <JobOffersList jobOffers={jobOffers} />}
        </div>
    );
}
export default MyJobOffers;