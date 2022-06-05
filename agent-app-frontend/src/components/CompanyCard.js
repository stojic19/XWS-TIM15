import "../css/userCard.css"
import { useState, useEffect } from "react";
import Swal from "sweetalert2";

import axios from 'axios';

const CompanyCard = (props) => {

    const accept = (e) => {
        e.preventDefault();

        axios.put(axios.defaults.baseURL + 'api/Companies/' + props.company.id + '/Register')
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                window.location.reload()
            }).catch((err) => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oopss...',
                    text: err.data.response,
                });
            });
    }

    const decline = (e) => {
        e.preventDefault();
        return;
        /*const update = {
            "followerId": localStorage.getItem('user_id'),
            "followedId": props.user.id
        };
        axios.delete(axios.defaults.baseURL + 'followers/followRequest', {headers: {}, data: update})
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                window.location.reload()
            });*/
    }

    const getFullName = () => {
        return props.company.owner.personalInfo.firstName + ' ' + props.company.owner.personalInfo.middleName + ' ' + props.company.owner.personalInfo.lastName
    }
    const getUsernameAndEmail = () => {
        return props.company.owner.username + ' ' + props.company.owner.personalInfo.email;
    }
    return (
        <li className="col-12 col-md-4 col-lg-3 m-2">
            <div className="cnt-block equal-hight" style={{ maxHeight: "120%", maxWidth: "90%" }}>
                {
                    props.admin && <h6>{props.company.companyInfo.name}</h6>
                }
                {
                    !props.admin &&<h6><a href={'/jobOffers/' + props.company.id}>{props.company.companyInfo.name}</a></h6>
                }
                <figure><img src={require("../images/user-avatar.png")} className="img-responsive" alt=""></img></figure>
                <h3>{getFullName()}</h3>
                <h3>{getUsernameAndEmail()}</h3>
                <p>{props.company.companyInfo.address}</p>
                <p>{props.company.companyInfo.email}</p>
                <p>{props.company.companyInfo.phoneNumber}</p>
                <p>Description: {props.company.companyInfo.description}</p>
                <p>Culture: {props.company.companyInfo.culture}</p>
                {
                    props.admin &&
                    <>
                        <button onClick={(e) => accept(e)} type="button" className="btn btn-outline-primary m-1">Accept</button>
                        {/*<button onClick={(e) => decline(e)} type="button" className="btn btn-outline-primary m-1">Decline</button>*/}
                    </>
                }
            </div>
        </li>
    );
}

export default CompanyCard;