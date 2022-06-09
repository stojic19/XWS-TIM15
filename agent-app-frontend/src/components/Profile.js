import '../css/userProfile.css'
import "../css/userCard.css"
import { useNavigate } from 'react-router-dom';
import Swal from "sweetalert2";
import { useEffect, useState } from "react";

import axios from 'axios';
import JobOfferPostList from './JobOfferPostList';

const Profile = (props) => {
    const [jobOffers, setJobOffers] = useState();

    const fetchJobOffers = async () => {

        axios.get(axios.defaults.baseURL + 'api/Companies/' + props.company.id)
            .then(res => {
                let jobOffers = Array.from(res.data.jobOffers);
                setJobOffers(jobOffers);
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
        <div className="container">
            <div className="row">
                <div className="col-md-4">
                    <div className="row">
                        <div className="col-12 bg-white p-0 px-3 py-3 mb-5">
                            <div className="d-flex flex-column align-items-center">
                                <img src={require('../images/company-avatar.jpg')} />
                                <p className="fw-bold h4 mt-3">{props.company.companyInfo.name}</p>
                                <p className="fw-bold text-muted">Owner:</p> 
                                <p className="text-muted">{props.company.owner.personalInfo.firstName} {props.company.owner.personalInfo.lastName}</p> 
                                <p className="text-muted mb-3">@{props.company.owner.username}</p>
                                <p className="text-muted mb-3">{props.company.owner.personalInfo.email}</p>
                                
                            </div>
                        </div>
                       

                    </div>
                </div>
                
                    <div className="col-md-8 ps-md-4">
                        <div className="row">
                            <div className="col-12 bg-white p-0 px-2 pb-3 mb-3">
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <h4>About</h4>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Email:</p>
                                    <p className="py-2 text-muted">{props.company.companyInfo.email}</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Phone number:</p>
                                    <p className="py-2 text-muted">{props.company.companyInfo.phoneNumber}</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Address:</p>
                                    <p className="py-2 text-muted">{props.company.companyInfo.address}</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Culture:</p>
                                    <p className="py-2 text-muted">{props.company.companyInfo.culture}</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Description:</p>
                                    <p style={{marginLeft: "10%"}} className="py-2 text-muted">{props.company.companyInfo.description}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <>{ jobOffers && <JobOfferPostList offers={jobOffers}></JobOfferPostList>}</>
            </div>
        </div>
    );
}

export default Profile;