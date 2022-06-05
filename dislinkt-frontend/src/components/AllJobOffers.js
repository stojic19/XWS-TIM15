import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import JobOffersList from './JobOffersList';

const AllJobOffers = () => {

    const [jobOffers, setJobOffers] = useState([]);
    const [loading, setLoading] = useState(true);

    const fetchJobOffers = async () => {
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'job_offers')
            .then(res => {
                let jobOffers = Array.from(res.data.jobOffers)
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
export default AllJobOffers;