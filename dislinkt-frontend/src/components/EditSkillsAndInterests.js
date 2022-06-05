import { useState, useEffect } from "react";
import Swal from "sweetalert2";
import { useNavigate } from 'react-router-dom';
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

import axios from 'axios';

const EditSkillsAndInterests = () => {
    const [loading, setLoading] = useState(true);
    const [email, setEmail] = useState("");;
    const [name, setName] = useState("");
    const [isPending, setIsPending] = useState(false);
    const [skills, setSkills] = useState([]);
    const [interests, setInterests] = useState([]);
    const [skill, setSkill] = useState('');
    const [interest, setInterest] = useState('');

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
        const fetchSkills = async () => {
            let id = localStorage.getItem('user_id');
            setLoading(true);
            axios.get(axios.defaults.baseURL + 'skills/' + id)
                .then(res => {
                    let skills = res.data.skills;
                    setSkills(skills);
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
        const fetchInterests = async () => {
            let id = localStorage.getItem('user_id');
            setLoading(true);
            axios.get(axios.defaults.baseURL + 'interests/' + id)
                .then(res => {
                    let interests = res.data.interests;
                    setInterests(interests);
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
        fetchSkills();
        fetchInterests();
    }, []);

    const ValidateSkill = () => {
        if (skill === "") {
            Swal.fire({
                icon: 'warning',
                title: 'Oops...',
                text: 'Skill input must be filled!',
            });
            return false;
        }
        return true;
    }

    const ValidateInterest = () => {
        if (interest === "") {
            Swal.fire({
                icon: 'warning',
                title: 'Oops...',
                text: 'Interest input must be filled!',
            });
            return false;
        }
        return true;
    }

    const onSubmitSkill = async (e) => {
        e.preventDefault();
        if (!ValidateSkill())
            return;

        setIsPending(true);

        let skillsUpdate = []

        skills.map((skillUpdate) => {
            skillsUpdate = skillsUpdate.concat(skillUpdate);
        })
        skillsUpdate = skillsUpdate.concat(skill);

        const update = {
            "userId": localStorage.getItem('user_id'),
            "skills": skillsUpdate
        };
        axios.put(axios.defaults.baseURL + 'skills', update)
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

    const removeSkill = async (e, index) => {
        e.preventDefault();
        setIsPending(true);

        let skillsUpdate = []

        skills.map((skillUpdate, checkIndex) => {
            if (checkIndex != index)
                skillsUpdate = skillsUpdate.concat(skillUpdate);
        })

        const update = {
            "userId": localStorage.getItem('user_id'),
            "skills": skillsUpdate
        };
        axios.put(axios.defaults.baseURL + 'skills', update)
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

    const onSubmitInterest = async (e) => {
        e.preventDefault();
        if (!ValidateInterest())
            return;

        setIsPending(true);

        let interestsUpdate = []

        interests.map((interestUpdate) => {
            interestsUpdate = interestsUpdate.concat(interestUpdate);
        })
        interestsUpdate = interestsUpdate.concat(interest);

        const update = {
            "userId": localStorage.getItem('user_id'),
            "interests": interestsUpdate
        };
        axios.put(axios.defaults.baseURL + 'interests', update)
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

    const removeInterest = async (e, index) => {
        e.preventDefault();
        setIsPending(true);

        let interestsUpdate = []

        interests.map((interestUpdate, checkIndex) => {
            if (checkIndex != index)
            interestsUpdate = interestsUpdate.concat(interestUpdate);
        })

        const update = {
            "userId": localStorage.getItem('user_id'),
            "interests": interestsUpdate
        };
        axios.put(axios.defaults.baseURL + 'interests', update)
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
                            <h4 className="text-right">Skills Settings</h4>
                        </div>
                        {skills && skills.map((skill, index) => {
                            return (
                                <div key={index}>
                                    <div className="row mt-2">
                                        <div className="col-md-6">
                                            <input type="text" className="form-control" placeholder="Enter skill" value={skill} readOnly />
                                        </div>
                                        <div className="col-md-6 align-items-center text-right">
                                            <button onClick={(e) => removeSkill(e, index)} className="btn btn-primary button" type="button">Remove</button>
                                        </div>
                                    </div>
                                    <hr />
                                </div>
                            )
                        })
                        }
                        <div>
                            <div className="row mt-2">
                                <div className="col-md-6">
                                    <input type="text" className="form-control" placeholder="Enter skill" value={skill} onChange={(e) => setSkill(e.target.value)} />
                                </div>
                                <div className="col-md-6 align-items-center text-right">
                                    {isPending && <label>Update in progress...</label>}
                                    {!isPending && <button onClick={(e) => onSubmitSkill(e)} className="btn btn-primary profile-button" type="button">Save</button>}
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="p-3 py-5">
                        <div className="d-flex justify-content-between align-items-center mb-3">
                            <h4 className="text-right">Interests Settings</h4>
                        </div>
                        {interests && interests.map((inter, index) => {
                            return (
                                <div key={index+'interests'}>
                                    <div className="row mt-2">
                                        <div className="col-md-6">
                                            <input type="text" className="form-control" placeholder="Enter skill" value={inter} readOnly />
                                        </div>
                                        <div className="col-md-6 align-items-center text-right">
                                            <button onClick={(e) => removeInterest(e, index)} className="btn btn-primary button" type="button">Remove</button>
                                        </div>
                                    </div>
                                    <hr />
                                </div>
                            )
                        })
                        }
                        <div>
                            <div className="row mt-2">
                                <div className="col-md-6">
                                    <input type="text" className="form-control" placeholder="Enter skill" value={interest} onChange={(e) => setInterest(e.target.value)} />
                                </div>
                                <div className="col-md-6 align-items-center text-right">
                                    {isPending && <label>Update in progress...</label>}
                                    {!isPending && <button onClick={(e) => onSubmitInterest(e)} className="btn btn-primary profile-button" type="button">Save</button>}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                
            </div>}
        </div>
    );
}

export default EditSkillsAndInterests;