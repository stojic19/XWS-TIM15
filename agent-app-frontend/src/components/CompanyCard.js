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
        axios.delete(axios.defaults.baseURL + 'api/Companies/' + props.company.id, {headers: {}})
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                window.location.reload()
            });
    }

    const getFullName = () => {
        return props.company.owner.personalInfo.firstName + ' (' + props.company.owner.personalInfo.middleName + ') ' + props.company.owner.personalInfo.lastName
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
                    !props.admin && <h6><a href={'/companyProfile/' + props.company.id}>{props.company.companyInfo.name}</a></h6>
                }
                <img style={{height: "45%", width: "40%"}} src={require("../images/company-avatar.jpg")} className="img-responsive" alt=""></img>
                <br></br>
                <br></br>
                <h3>{getFullName()}</h3>
                <a>@{props.company.owner.username}</a>
                <br></br>
                <a>{props.company.owner.personalInfo.email}</a>
                <br></br>
                <br></br>
                <h3>Company info:</h3>
                <p>{props.company.companyInfo.address}</p>
                <p>{props.company.companyInfo.email}</p>
                <p>{props.company.companyInfo.phoneNumber}</p>
                {
                    props.admin &&
                    <>
                        <button onClick={(e) => accept(e)} type="button" className="btn btn-outline-primary m-1">Accept</button>
                        <button onClick={(e) => decline(e)} type="button" className="btn btn-outline-primary m-1">Decline</button>
                    </>
                }
            </div>
        </li>
    );
}

export default CompanyCard;