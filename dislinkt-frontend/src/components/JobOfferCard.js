import "../css/userCard.css"
import { useState, useEffect } from "react";
import Swal from "sweetalert2";

import axios from 'axios';

const JobOfferCard = (props) => {

    const subscribe = async (e) => {
        e.preventDefault();

        const update = {
            "id": localStorage.getItem('user_id'),
            "jobOfferId": props.jobOffer.id
        };
        axios.put(axios.defaults.baseURL + 'job_offers/subscribe', update)
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                window.location.reload()
            });
    }
    const unsubscribe = async (e) => {
        e.preventDefault();

        const update = {
            "id": localStorage.getItem('user_id'),
            "jobOfferId": props.jobOffer.id
        };
        axios.put(axios.defaults.baseURL + 'job_offers/unsubscribe', update)
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                window.location.reload()
            });
    }

    return (
        <li className="col-12 col-md-4 col-lg-3">
            {props.jobOffer &&
                <div className="cnt-block equal-hight" style={{ maxHeight: "100%" }}>
                    <h6>
                        {
                            props.jobOffer.isActive ? 'Active' : 'Closed'
                        }
                    </h6>
                    <figure><img src={require("../images/user-avatar.png")} className="img-responsive" alt=""></img></figure>
                    <h3>{props.jobOffer.name}</h3>
                    <p>{props.jobOffer.position}</p>
                    <p>Description : {props.jobOffer.description}</p>
                    <p>Requirements : {props.jobOffer.requirements}</p>
                    {
                        !props.jobOffer.subscribers.some(u => u.id == localStorage.getItem('user_id'))
                        && <button onClick={(e) => subscribe(e)} type="button" className="btn btn-outline-primary">Subscribe</button>
                    }
                    {
                        props.jobOffer.subscribers.some(u => u.id == localStorage.getItem('user_id'))
                        && <button onClick={(e) => unsubscribe(e)} type="button" className="btn btn-outline-primary">Unsubscribe</button>
                    }
                </div>
            }
        </li>
    );
}

export default JobOfferCard;