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
                let jobOffersTemp = [];
                res.data.map((company) => {
                    company.jobOffers.map((jobOffer) => {
                        jobOffersTemp = jobOffersTemp.concat({
                            companyId: company.id,
                            description: jobOffer.description,
                            id: jobOffer.id,
                            isActive: jobOffer.isActive,
                            position: jobOffer.position,
                            requirements: jobOffer.requirements,
                            timeOfCreation: jobOffer.timeOfCreation,
                        })
                    })
                });
                jobOffersTemp = Array.from(jobOffersTemp)
                setJobOffers(jobOffersTemp);
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
            {!loading && jobOffers && <JobOffersList jobOffers={jobOffers} myJobOffers={true} />}
        </div>
    );
}
export default MyJobOffers;