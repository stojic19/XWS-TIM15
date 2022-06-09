import { useParams } from 'react-router';
import '../css/userProfile.css'
import "../css/userCard.css"
import Swal from "sweetalert2";
import { useEffect, useState } from "react";
import Comment from './Comment';

import axios from 'axios';

const JobOfferProfile = () =>{
    const {companyId, offerId} = useParams()
    const [jobOffers, setJobOffers] = useState([]);
    const [jobOffer, setJobOffer] = useState({});
    const [comments, setComments] = useState([]);

    const fetchJobOffers = async () => {

        axios.get(axios.defaults.baseURL + 'api/Companies/' + companyId)
            .then(res => {
                let offers = Array.from(res.data.jobOffers);
                setJobOffers(offers);
                setComments(res.data.comments)
                console.log(res.data)

            }).catch(err => {
                console.log(err)
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    };

    const getOffer = () =>{
        let offers = jobOffers
        for (let i = 0; i < offers.length; i++) {
            if(offers[i].id == offerId){
                setJobOffer(offers[i])
            }
          }
    }


    useEffect(() => {
        fetchJobOffers();
        getOffer()
    }, [])

    return (
        <div className="container">
            <div className="row">
                <div className="col-md-4">
                    <div className="row">
                        <div className="col-12 bg-white p-0 px-3 py-3 mb-5">
                            <div className="d-flex flex-column align-items-center">
                                <img src={require('../images/job-avatar.png')} />
                                <p className="fw-bold h4 mt-3">pozicija</p>
                                
                            </div>
                        </div>
                       

                    </div>
                </div>
                
                    <div className="col-md-8 ps-md-4">
                        <div className="row">
                            <div className="col-12 bg-white p-0 px-2 pb-3 mb-3">
                                <br></br>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <h4>About</h4>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Requirements:</p>
                                    <p className="py-2 text-muted">req</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Description:</p>
                                    <p style={{marginLeft: "10%"}} className="py-2 text-muted">opis</p>
                                </div>
                                <br></br>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <h4>Comments</h4>
                                </div>
                                <ul className="list-group list-group-flush">
                                {
                                    comments.map((comment, index) => {
                                        return (
                                            <Comment index={index} comment={comment}></Comment>
                                        );
                                    })
                                }
                            </ul>
                            <br></br>
                            <input className="form-control form-control-sm" type="text" placeholder="Add comment..."
                                onKeyPress={(ev) => {
                                    if (ev.key === "Enter") {
                                        ev.preventDefault();
                                        addComment(ev.target.value);
                                    }
                                }} />
                                <br></br>
                            </div>
                        </div>
                    </div>
                
            </div>
        </div>
    );
}
export default JobOfferProfile;