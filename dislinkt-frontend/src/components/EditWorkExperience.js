import { useState, useEffect } from "react";
import Swal from "sweetalert2";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

import axios from 'axios';

const EditWorkExperience = () => {
    const [loading, setLoading] = useState(true);
    const [email, setEmail] = useState("");;
    const [name, setName] = useState("");
    const [isPending, setIsPending] = useState(false);
    const [workExperience, setWorkExperience] = useState([]);
    const [startDate, setStartDate] = useState(new Date());
    const [endDate, setEndDate] = useState(new Date());
    const [jobTitle, setJobTitle] = useState('');
    const [companyName, setCompanyName] = useState('');

    useEffect(() => {
        const fetchUser = async () => {
            let id = localStorage.getItem('user_id');
            setLoading(true);
            axios.get(axios.defaults.baseURL + 'getUserForEdit/' + id)
                .then(res => {
                    let user = res.data.user;;
                    setEmail(user.email);
                    setName(user.name);
                    setLoading(false);
                }).catch(err => {
                    console.log(err);
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: err.data,
                    });
                });
        };
        const fetchWorkExperience = async () => {
            let id = localStorage.getItem('user_id');
            setLoading(true);
            axios.get(axios.defaults.baseURL + 'workExperience/' + id)
                .then(res => {
                    let workExperience = res.data.workExperience;
                    setWorkExperience(workExperience);
                    setLoading(false);
                }).catch(err => {
                    console.log(err);
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: err.data,
                    });
                });
        };
        fetchUser();
        fetchWorkExperience();
    }, []);

    const Validate = () => {
        if (companyName === "" || jobTitle === "" || startDate === "" || endDate === "") {
            Swal.fire({
                icon: 'warning',
                title: 'Oops...',
                text: 'All inputs must be filled!',
            });
            return false;
        }
        if (startDate >= endDate) {
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: 'Start date must be before end date!',
            });
            return false;
        }
        return true;
    }
    const FormatDate = (date) => {
        var month = date.getUTCMonth() + 1; //months from 1-12.
        var day = date.getUTCDate();
        var year = date.getUTCFullYear();
        return day + '/' + month + '/' + year;
    }
    const onSubmit = async (e) => {
        e.preventDefault();
        if (!Validate())
            return;

        setIsPending(true);

        let workExperiences = []

        workExperience.map((workExperience) => {
            workExperiences = workExperiences.concat({
                "jobTitle": workExperience.jobTitle,
                "companyName": workExperience.companyName,
                "startDate": FormatDate(new Date(workExperience.startDate)),
                "endDate": FormatDate(new Date(workExperience.endDate))
            })
        })
        workExperiences = workExperiences.concat({
            "jobTitle": jobTitle,
            "companyName": companyName,
            "startDate": FormatDate(startDate),
            "endDate": FormatDate(endDate)
        })

        const update = {
            "userId": localStorage.getItem('user_id'),
            "workExperiences": workExperiences
        };
        axios.put(axios.defaults.baseURL + 'workExperience', update)
            .then(res => {
                setIsPending(false);
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                window.location.reload()
            });
    }

    const remove = async (e, index) => {
        e.preventDefault();
        setIsPending(true);

        let workExperiences = []

        workExperience.map((workExperience, checkIndex) => {
            if (checkIndex != index)
                workExperiences = workExperiences.concat({
                    "jobTitle": workExperience.jobTitle,
                    "companyName": workExperience.companyName,
                    "startDate": FormatDate(new Date(workExperience.startDate)),
                    "endDate": FormatDate(new Date(workExperience.endDate))
                })
        })

        const update = {
            "userId": localStorage.getItem('user_id'),
            "workExperiences": workExperiences
        };
        axios.put(axios.defaults.baseURL + 'workExperience', update)
            .then(res => {
                setIsPending(false);
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                window.location.reload()
            });
    }

    const getDateFormatForDisplay = (date) => {
        return new Date(date);
    }
    return (
        <div className="container rounded bg-white mt-5 mb-5">
            {loading && <h3>Loading...</h3>}
            {!loading && <div className="row">
                <div className="col-md-3 border-right">
                    <div className="d-flex flex-column align-items-center text-center p-3 py-5"><img className="rounded-circle mt-5" width="150px" src="https://st3.depositphotos.com/15648834/17930/v/600/depositphotos_179308454-stock-illustration-unknown-person-silhouette-glasses-profile.jpg" />
                        <span className="font-weight-bold">{name}</span>
                        <span className="text-black-50">{email}</span>
                        <span> </span></div>
                </div>
                <div className="col-md-5 border-right">
                    <div className="p-3 py-5">
                        <div className="d-flex justify-content-between align-items-center mb-3">
                            <h4 className="text-right">Work Experience Settings</h4>
                        </div>
                        {workExperience && workExperience.map((ex, index) => {
                            return (
                                <div key={index}>
                                    <div className="row mt-2">
                                        <div className="col-md-6">
                                            <label className="labels">Company name</label>
                                            <input type="text" className="form-control" placeholder="Enter company name" value={ex.companyName} readOnly />
                                        </div>
                                        <div className="col-md-6">
                                            <label className="labels">Job title</label>
                                            <input type="text" className="form-control" value={ex.jobTitle} placeholder="Enter job title" readOnly />
                                        </div>
                                    </div>
                                    <div className="row mt-2">
                                        <div className="col-md-6">
                                            <label className="labels">Start date</label>
                                            <DatePicker dateFormat="dd/MM/yyyy" selected={getDateFormatForDisplay(ex.startDate)} className="form-control" readOnly />
                                        </div>
                                        <div className="col-md-6">
                                            <label className="labels">End date</label>
                                            <DatePicker dateFormat="dd/MM/yyyy" selected={getDateFormatForDisplay(ex.endDate)} className="form-control" readOnly />
                                        </div>
                                    </div>
                                    <div className="align-items-center m-2 text-center">
                                        <button onClick={(e) => remove(e, index)} className="btn btn-primary button" type="button">Remove</button>
                                    </div>
                                    <hr />
                                </div>
                            )
                        })
                        }
                        <div>
                            <div className="row mt-2">
                                <div className="col-md-6">
                                    <label className="labels">Company name</label>
                                    <input type="text" className="form-control" placeholder="Enter company name" value={companyName} onChange={(e) => setCompanyName(e.target.value)} />
                                </div>
                                <div className="col-md-6">
                                    <label className="labels">Job title</label>
                                    <input type="text" className="form-control" value={jobTitle} placeholder="Enter job title" onChange={(e) => setJobTitle(e.target.value)} />
                                </div>
                            </div>
                            <div className="row mt-2">
                                <div className="col-md-6">
                                    <label className="labels">Start date</label>
                                    <DatePicker dateFormat="dd/MM/yyyy" selected={startDate} className="form-control" onChange={(date) => setStartDate(date)} />
                                </div>
                                <div className="col-md-6">
                                    <label className="labels">End date</label>
                                    <DatePicker dateFormat="dd/MM/yyyy" selected={endDate} className="form-control" onChange={(date) => setEndDate(date)} />
                                </div>
                            </div>
                            <hr />
                        </div>
                        <div className="mt-5 text-center">
                            {isPending && <label>Update in progress...</label>}
                            {!isPending && <button onClick={(e) => onSubmit(e)} className="btn btn-primary profile-button" type="button">Save Profile</button>}
                        </div>
                    </div>
                </div>
            </div>}
        </div>
    );
}

export default EditWorkExperience;