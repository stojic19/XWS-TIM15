import { useState } from 'react';
import '../css/search.css'
import axios from 'axios';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import JobOffersList from './JobOffersList';

const RecommendedJobOffers = () => {

    const [jobOffers, setJobOffers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [searchTerm, setSearchTerm] = useState('');

    const fetchJobOffers = async () => {
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'job_offers/recommended/' + localStorage.getItem('user_id'))
            .then(res => {
                console.log(res.data)
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

    useEffect(() => {
        //fetchJobOffers();
        searchJobOffers();
    }, [searchTerm])

    const searchJobOffers = async () => {

        setLoading(true);
        axios.get(axios.defaults.baseURL + 'job_offers/recommended/' + localStorage.getItem('user_id'))
            .then(res => {
                let jobOffers = Array.from(res.data.jobOffers)
                let offers = []
                jobOffers.forEach(job => {
                    if(job.position.toLowerCase().includes(searchTerm.toLowerCase())){
                        offers.push(job)
                    }
                })
                setJobOffers(offers);
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

    return (
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
            {!loading && jobOffers && <JobOffersList jobOffers={jobOffers} />}
        </div>
    );
}
export default RecommendedJobOffers;