import { useState, useEffect } from "react";
import Swal from "sweetalert2";
import { useNavigate } from 'react-router-dom';
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

import axios from 'axios';

const EditProfile = () => {
    const [user, setUser] = useState();
    const [loading, setLoading] = useState(true);
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [repeatPassword, setRepeatPassword] = useState("");
    const [telephoneNumber, setTelephoneNumber] = useState("");
    const [gender, setGender] = useState("");
    const [name, setName] = useState("");
    const [dateOfBirth, setDateOfBirth] = useState(new Date());
    const [biography, setBiography] = useState("");
    const [isPrivate, setIsPrivate] = useState(false);
    const [isPending, setIsPending] = useState(false);
    const history = useNavigate();

    useEffect(() => {
        const fetchUser = async () => {
            let id = localStorage.getItem('user_id');
            setLoading(true);
            axios.get(axios.defaults.baseURL + 'getUserForEdit/' + id)
                .then(res => {
                    let user = res.data.user;
                    setEmail(user.email);
                    setUsername(user.username);
                    setPassword(user.password);
                    setRepeatPassword(user.password);
                    setTelephoneNumber(user.telephoneNo);
                    setGender(user.gender);
                    setName(user.name);
                    setDateOfBirth(new Date(user.dateOfBirth));
                    setBiography(user.biography);
                    setIsPrivate(user.isPrivate);
                    setUser(user);
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
    }, []);

    const Validate = () => {
        if (email === "" || username === "" || password === "" || repeatPassword === "" || telephoneNumber === ""
            || gender === "" || name === "" || dateOfBirth === "" || biography === "") {
            Swal.fire({
                icon: 'warning',
                title: 'Oops...',
                text: 'All inputs must be filled!',
            });
            return false;
        }
        if (password !== repeatPassword) {
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: 'Passwords must be same!',
            });
            return false;
        }
        return true;
    }
    const FormatDate = (date) => {
        var month = date.getUTCMonth() + 1; //months from 1-12.
        var day = date.getUTCDate() + 1;
        var year = date.getUTCFullYear();
        return day + '/' + month + '/' + year;
    }
    const onSubmit = async (e) => {
        e.preventDefault();
        if (!Validate())
            return;

        setIsPending(true);
        const update = {
            "username": user.username,
            "newUsername": username,
            "password": password,
            "email": email,
            "name": name,
            "telephoneNo": telephoneNumber,
            "gender": gender,
            "biography": biography,
            "isPrivate": isPrivate === "false" ? false : true,
            "dateOfBirth": FormatDate(dateOfBirth),
        };

        axios.put(axios.defaults.baseURL + 'users', update)
            .then(res => {
                if (res.data.error) {
                    setIsPending(false);
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: res.data.error,
                    });
                } else {
                    setIsPending(false);
                    Swal.fire({
                        icon: 'success',
                        title: 'Success!',
                        text: res.data.response,
                    });
                    //TODO: ADD REDIRECT TO PROFILE PAGE
                }
            });
    }

    return (
        <div className="container rounded bg-white mt-5 mb-5">
            {loading && <h3>Loading...</h3>}
            {!loading && <div className="row">
                <div className="col-md-3 border-right">
                    <div className="d-flex flex-column align-items-center text-center p-3 py-5"><img className="rounded-circle mt-5" width="150px" src="https://st3.depositphotos.com/15648834/17930/v/600/depositphotos_179308454-stock-illustration-unknown-person-silhouette-glasses-profile.jpg" />
                        <span className="font-weight-bold">{user.name}</span>
                        <span className="text-black-50">{user.email}</span>
                        <span> </span></div>
                </div>
                <div className="col-md-5 border-right">
                    <div className="p-3 py-5">
                        <div className="d-flex justify-content-between align-items-center mb-3">
                            <h4 className="text-right">Profile Settings</h4>
                        </div>
                        <div className="row mt-2">
                            <div className="col-md-6">
                                <label className="labels">Username</label>
                                <input type="text" className="form-control" placeholder="Enter username" value={username} onChange={(e) => setUsername(e.target.value)} />
                            </div>
                            <div className="col-md-6">
                                <label className="labels">Email</label>
                                <input type="text" className="form-control" value={email} placeholder="Enter email" onChange={(e) => setEmail(e.target.value)} />
                            </div>
                        </div>
                        <div className="row mt-3">
                            <div className="col-md-12">
                                <label className="labels">Name</label>
                                <input type="text" className="form-control" placeholder="Enter name" value={name} onChange={(e) => setName(e.target.value)} />
                            </div>
                            <div className="col-md-12">
                                <label className="labels">Telephone number</label>
                                <input type="text" className="form-control" placeholder="Enter telephone number" value={telephoneNumber} onChange={(e) => setTelephoneNumber(e.target.value)} />
                            </div>
                            <div className="col-md-12">
                                <label className="labels">Gender</label>
                                <select id="InputGender"
                                    name="gender"
                                    className="form-control"
                                    value={gender}
                                    onChange={(e) => setGender(e.target.value)}
                                >
                                    <option value="" disabled>Choose gender</option>
                                    <option value="Female">Female</option>
                                    <option value="Male">Male</option>
                                </select>
                            </div>
                            <div className="col-md-12"><label className="labels">Date of birth</label>
                                <DatePicker dateFormat="dd/MM/yyyy" selected={dateOfBirth} onChange={(date) => setDateOfBirth(date)} className="form-control" />
                            </div>
                            <div className="col-md-12"><label className="labels">Biography</label>
                                <textarea type="text" className="form-control" placeholder="Enter biography" value={biography} onChange={(e) => setBiography(e.target.value)} id="InputBiography" /></div>
                            <div className="col-md-12"><label className="labels">Private profile</label>
                                <input value={isPrivate} onChange={(e) => setIsPrivate(!isPrivate)} type="checkbox" className="form-check-input ml-2" id="CheckIsPrivate" />
                            </div>
                            <div className="col-md-12"><label className="labels">Password</label>
                                <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" className="form-control" id="InputPassword" /></div>
                            <div className="col-md-12"><label className="labels">Repeat password</label>
                                <input value={repeatPassword} onChange={(e) => setRepeatPassword(e.target.value)} type="password" className="form-control" id="InputRepeatPassword" />
                            </div>
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

export default EditProfile;