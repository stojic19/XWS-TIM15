import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import JobOffersList from './JobOffersList';
import { useParams } from "react-router-dom";

const JobOffersByCompanyId = () => {
    const { id } = useParams();
    const [jobOffers, setJobOffers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [name, setName] = useState('');

    const fetchJobOffers = async () => {
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'api/Companies/' + id)
            .then(res => {
                console.log(res.data);
                let jobOffers = Array.from(res.data.jobOffers);
                setJobOffers(jobOffers);
                setName(res.data.companyInfo.name);
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
            {!loading && jobOffers.length!=0 && <h3 style={{textAlign: 'center'}}>{name}'s job offers</h3>}
            {!loading && jobOffers.length==0 && <h3 style={{textAlign: 'center'}}>No {name}'s offers for showing.</h3>}
            {loading && <h3 style={{textAlign: 'center'}}>Loading...</h3>}
            {!loading && jobOffers && <JobOffersList jobOffers={jobOffers} hideTitle={true}/>}
        </div>
    );
}
export default JobOffersByCompanyId;