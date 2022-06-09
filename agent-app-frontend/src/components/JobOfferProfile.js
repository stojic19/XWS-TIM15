import { useParams } from 'react-router';
import '../css/userProfile.css'
import "../css/userCard.css"
import Swal from "sweetalert2";
import { useEffect, useState } from "react";
import Comment from './Comment';
import ReactStars from 'react-stars'

import axios from 'axios';
import { isLabelWithInternallyDisabledControl } from '@testing-library/user-event/dist/utils';

const JobOfferProfile = () =>{
    const {companyId, offerId} = useParams()
    const [jobOffers, setJobOffers] = useState([]);
    const [jobOffer, setJobOffer] = useState({});
    const [comments, setComments] = useState([]);
    const [ratingDisabled, setRating] = useState(false);
    const [ratingValue, setRatingValue] = useState(0);

    const fetchJobOffers = async () => {

        axios.get(axios.defaults.baseURL + 'api/Companies/' + companyId)
            .then(res => {
                let offers = Array.from(res.data.jobOffers);
                setJobOffers(offers);
                setComments(res.data.comments)
                //console.log(res.data)

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

    const addComment = async (newContent) => {
        let newComment = {
            content: newContent,
            companyId: companyId

        }
        console.log(newComment)
        axios.post(axios.defaults.baseURL + 'api/Companies/Comment', newComment)
            .then(res => {
                //console.log(res.data)
                window.location.reload()
            }).catch(err => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    }
    const ratingChanged = (newRating) => {
        setRatingValue(newRating)
      }

      const addGrade = async () => {
        let newGrade = {
            value: ratingValue | 0,
            companyId: companyId
        }
        console.log(newGrade)

        axios.put(axios.defaults.baseURL + 'api/Companies/Grade', newGrade)
            .then(res => {
                //console.log(res.data)
                //window.location.reload()
                Swal.fire({
                    icon: 'success',
                    title: 'Done',
                    text: 'Thank you for the rating!',
                });
            }).catch(err => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });

      }

    return (
        <div className="container">
            <div className="row">
                <div className="col-md-4">
                    <div className="row">
                        <div className="col-12 bg-white p-0 px-3 py-3 mb-5">
                            <div className="d-flex flex-column align-items-center">
                                <img src={require('../images/job-avatar.png')} />
                                <p className="fw-bold h4 mt-3">jobOffer</p>
                                
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
                                    <h4>Rate us</h4>
                                </div>
                                <ReactStars
                                count={5}
                                onChange={ratingChanged}
                                size={24}
                                color2={'#ffd700'} 
                                />
                                <button onClick={()=>addGrade()} style={{marginTop: "1%"}} type="button" class="btn btn-outline-primary" disabled={ratingDisabled}>Confirm</button>
                                <br></br>
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