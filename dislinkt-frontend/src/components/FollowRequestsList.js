import { useState, useEffect } from "react";
import Swal from "sweetalert2";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

import axios from 'axios';

const FollowRequestsList = () => {
    const [loading, setLoading] = useState(true);
    const [email, setEmail] = useState("");;
    const [name, setName] = useState("");
    const [isPending, setIsPending] = useState(false);
    const [followRequests, setFollowRequests] = useState([]);

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
        const fetchFollowRequests = async () => {
            let id = localStorage.getItem('user_id');
            setLoading(true);
            axios.get(axios.defaults.baseURL + 'followers/followRequests/' + id)
                .then(res => {
                    let followRequests = res.data.followRequests;
                    setFollowRequests(followRequests);
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
        fetchFollowRequests();
    }, []);

    const accept = async (e, id) => {
        e.preventDefault();
        setIsPending(true);

        const update = {
            "followerId": localStorage.getItem('user_id'),
            "followedId": id
        };
        axios.put(axios.defaults.baseURL + 'followers/followRequest', update)
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

    const remove = async (e, id) => {
        e.preventDefault();
        setIsPending(true);

        const update = {
            "followerId": localStorage.getItem('user_id'),
            "followedId": id
        };
        axios.put(axios.defaults.baseURL + 'followers/followRequest', update)
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

    const getUsername = (id) => {
        return id;
    }

    const getName = (id) => {
        return id;
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
                            <h4 className="text-right">Follow requests</h4>
                        </div>
                        {followRequests.length == 0 && <p className="text-right">No follow requests for display.</p>}
                        {followRequests && followRequests.map((followRequest, index) => {
                            return (
                                <div key={index}>
                                    <div className="row mt-2">
                                        <div className="col-md-12">
                                            <label className="labels">Username</label>
                                            <input type="text" className="form-control" value={getUsername(followRequest.id)} readOnly />
                                        </div>
                                    </div>
                                    <div className="row mt-2">
                                        <div className="col-md-12">
                                            <label className="labels">Name</label>
                                            <input type="text" className="form-control" value={getName(followRequest.id)} readOnly />
                                        </div>
                                    </div>
                                    <div className="row mt-2">
                                        <div className="col-md-12">
                                            <label className="labels">Date of request</label>
                                            <DatePicker dateFormat="dd/MM/yyyy" selected={getDateFormatForDisplay(followRequest.time)} className="form-control" readOnly />
                                        </div>
                                    </div>
                                    <div className="row mt-2">
                                        <div className="col-md-6  text-center">
                                            <button onClick={(e) => accept(e, followRequest.id)} className="btn btn-primary button" type="button">Accept</button>
                                        </div>
                                        <div className="col-md-6  text-center">
                                            <button onClick={(e) => remove(e, followRequest.id)} className="btn btn-primary button" type="button">Delete</button>
                                        </div>
                                    </div>
                                    <hr />
                                </div>
                            )
                        })
                        }
                    </div>
                </div>
            </div>}
        </div>
    );
}

export default FollowRequestsList;