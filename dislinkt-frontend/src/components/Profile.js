import '../css/userProfile.css'
import { useNavigate } from 'react-router-dom';
import PostList from './PostList';
import Swal from "sweetalert2";
import { useState } from "react";

import axios from 'axios';

const Profile = (props) => {
    const [relationship, setRelationship] = useState(props.relationship);
    const history = useNavigate();

    const followRequests = () => {
        history('/followRequests');
    }

    const editPersonalInfo = () => {
        history('/editProfile');
    }

    const blocked = () => {
        history('/blocked');
    }

    const getRelationship = () => {
        if (relationship === 'NO RELATIONSHIP')
            return true;
        return false;
    }

    const follow = async (e) => {
        e.preventDefault();

        const update = {
            "followerId": localStorage.getItem('user_id'),
            "followedId": props.user.id
        };

        axios.put(axios.defaults.baseURL + 'followers/follow', update)
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                setRelationship('FOLLOWING');
                window.location.reload()
            });
    }

    const unfollow = async (e) => {
        e.preventDefault();

        const update = {
            "followerId": localStorage.getItem('user_id'),
            "followedId": props.user.id
        };
        axios.delete(axios.defaults.baseURL + 'followers/follow', { headers: {}, data: update })
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                setRelationship('NO RELATIONSHIP');
                window.location.reload()
            });
    }

    const followRequest = async (e) => {
        e.preventDefault();

        const update = {
            "followerId": localStorage.getItem('user_id'),
            "followedId": props.user.id
        };
        axios.put(axios.defaults.baseURL + 'followers/follow', update)
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                setRelationship('FOLLOW REQUEST');
                window.location.reload()
            });
    }

    const cancelFollowRequest = async (e) => {
        e.preventDefault();

        const update = {
            "followerId": localStorage.getItem('user_id'),
            "followedId": props.user.id
        };
        axios.delete(axios.defaults.baseURL + 'followers/followRequest', { headers: {}, data: update })
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                setRelationship('NO RELATIONSHIP');
                window.location.reload()
            });
    }

    const block = async (e) => {
        e.preventDefault();

        const update = {
            "subjectId": localStorage.getItem('user_id'),
            "objectId": props.user.id
        };
        axios.put(axios.defaults.baseURL + 'followers/block', update)
            .then(res => {
                console.log(res)
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                    position: 'top-end',
                    showConfirmButton: false,
                    timer: 2000
                });
                history('/home');
            });
    }

    const showApiKey = async (e) => {
        e.preventDefault();

        Swal.fire({
            icon: 'info',
            title: 'ApiKey',
            text: props.user.apikey,
            position: 'center',
            showConfirmButton: true,
            timer: 5000
        });
    }

    return (
        <div className="container">
            <div className="row">
                <div className="col-md-5">
                    <div className="row">
                        <div className="col-12 bg-white p-0 px-3 py-3 mb-5">
                            <div className="d-flex flex-column align-items-center">
                                <h6>
                                    {
                                        props.user.isPrivate ? 'Private' : 'Public'
                                    }
                                </h6>
                                <img src={require('../images/user-avatar.png')} />
                                <p className="fw-bold h4 mt-3">{props.user.name}</p>
                                {/*<p className="text-muted">Software developer</p>*/} {/* sortirati listu experience pa podesiti poslednji posao */}
                                <p className="text-muted mb-3">@{props.user.username}</p>
                                {
                                    !props.hiddenButtons && !props.personalProfile && !props.hiddenContent &&
                                    <div className="d-flex ">
                                        {getRelationship(props.relationship)
                                            && <button onClick={(e) => follow(e)} className="btn btn-primary follow me-2">Follow</button>
                                        }
                                        {!getRelationship(props.relationship) && <>
                                            <button onClick={(e) => unfollow(e)} className="btn btn-primary follow me-2">Unfollow</button>
                                            <button onClick={() => history("/chat/" + props.user.id)} className="btn btn-outline-primary message">Message</button>
                                        </>}
                                    </div>
                                }
                                {
                                    !props.hiddenButtons && !props.personalProfile && props.hiddenContent &&
                                    <div className="d-flex ">
                                        {getRelationship(props.relationship)
                                            && <button onClick={(e) => followRequest(e)} className="btn btn-primary follow me-2">Follow request</button>
                                        }
                                        {!getRelationship(props.relationship)
                                            && <button onClick={(e) => cancelFollowRequest(e)} className="btn btn-primary follow me-2">Cancel follow request</button>
                                        }
                                    </div>
                                }
                                {
                                    props.personalProfile &&
                                    <>
                                        <div className="d-flex ">
                                            <button onClick={(e) => editPersonalInfo(e)} className="btn btn-primary follow me-2">Edit personal info</button>
                                        </div>
                                        {
                                            props.user.isPrivate && <div className="d-flex m-2">
                                                <button onClick={(e) => followRequests(e)} className="btn btn-primary follow me-2">Follow requests</button>
                                            </div>
                                        }
                                        <div className="d-flex m-2">
                                            <button onClick={(e) => blocked(e)} className="btn btn-primary follow me-2">View blocked</button>
                                        </div>
                                    </>
                                }
                                {
                                    !props.personalProfile &&
                                    <>
                                        <div className="d-flex m-2">
                                            <button onClick={(e) => block(e)} className="btn btn-primary follow me-2">Block</button>
                                        </div>
                                    </>
                                }
                                {
                                    props.personalProfile &&
                                    <>
                                        <div className="d-flex m-2">
                                            <button onClick={(e) => showApiKey(e)} className="btn btn-primary follow me-2">Show ApiKey</button>
                                        </div>
                                    </>
                                }
                            </div>
                        </div>
                        {
                            !props.hiddenContent &&
                            <div className="col-12 bg-white px-3 pb-2 ">
                                <h6 className="d-flex align-items-center mb-3 fw-bold py-3 justify-content-center">
                                    <i className="text-info me-2">Skills</i>
                                    {props.personalProfile && <button onClick={() => history("/editSkillsAndInterests")} className="btn btn-primary">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-pencil-square" viewBox="0 0 16 16">
                                            <path d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z" />
                                            <path fillRule="evenodd" d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z" />
                                        </svg>Edit
                                    </button>}
                                </h6>
                                {props.skills.length === 0 && <p style={{ textAlign: "center" }}>No skills to show.</p>}
                                {
                                    props.skills.map((skill, index) => {
                                        return (
                                            <p style={{ textAlign: "center" }} key={index}>{skill}</p>

                                        )
                                    })
                                }
                            </div>
                        }
                        {!props.hiddenContent &&
                            <div className="col-12 bg-white px-3 pb-2 ">
                                <h6 className="d-flex align-items-center mb-3 fw-bold py-3 justify-content-center"><i
                                    className="text-info me-2">Interests</i>
                                    {props.personalProfile && <button onClick={() => history("/editSkillsAndInterests")} className="btn btn-primary">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-pencil-square" viewBox="0 0 16 16">
                                            <path d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z" />
                                            <path fillRule="evenodd" d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z" />
                                        </svg>Edit
                                    </button>
                                    }
                                </h6>
                                {props.interests.length === 0 && <p style={{ textAlign: "center" }}>No interests to show.</p>}
                                {
                                    props.interests.map((interest, index) => {
                                        return (
                                            <p style={{ textAlign: "center" }} key={index}>{interest}</p>

                                        )
                                    })
                                }
                            </div>
                        }
                    </div>
                </div>
                {!props.hiddenContent &&
                    <div className="col-md-7 ps-md-4">
                        <div className="row">
                            <div className="col-12 bg-white p-0 px-2 pb-3 mb-3">
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <h4>About</h4>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Email</p>
                                    <p className="py-2 text-muted">{props.user.email}</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Mobile</p>
                                    <p className="py-2 text-muted">{props.user.telephoneNo}</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Date of birth</p>
                                    <p className="py-2 text-muted">{props.user.dateOfBirth}</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Gender</p>
                                    <p className="py-2 text-muted">{props.user.gender}</p>
                                </div>
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <p className="py-2">Biography</p>
                                    <p className="py-2 text-muted">{props.user.biography}</p>
                                </div>
                            </div>
                            <br></br>
                            <div className="col-12 bg-white p-0 px-2 pb-3 mb-3">
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <h4>Work experience</h4>
                                    {props.personalProfile && <button onClick={() => history("/editWorkExperience")} className="btn btn-primary m-2">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-pencil-square" viewBox="0 0 16 16">
                                            <path d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z" />
                                            <path fillRule="evenodd" d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z" />
                                        </svg>Edit
                                    </button>
                                    }
                                </div>
                                {props.experience.length === 0 && <p style={{ textAlign: "center" }}>No work experience to show.</p>}
                                {
                                    (props.experience).map((ex, index) => {
                                        return (
                                            <div key={index} className="d-flex align-items-center justify-content-between border-bottom">
                                                <p className="py-2">{ex.startDate} - {ex.endDate}</p>
                                                <p className="py-2">{ex.companyName} </p>
                                                <p className="py-2 text-muted">{ex.jobTitle}</p>
                                            </div>
                                        )
                                    })
                                }
                            </div>
                            <br></br>
                            <div className="col-12 bg-white p-0 px-2 mb-3 pb-3">
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <h4>Education</h4>
                                    {props.personalProfile && <button onClick={() => history("/editEducation")} className="btn btn-primary m-2">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-pencil-square" viewBox="0 0 16 16">
                                            <path d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z" />
                                            <path fillRule="evenodd" d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z" />
                                        </svg>Edit
                                    </button>
                                    }
                                </div>
                                {props.education.length === 0 && <p style={{ textAlign: "center" }}>No education to show.</p>}
                                {
                                    props.education.map((edu, index) => {
                                        return (
                                            <div key={index} className="d-flex align-items-center justify-content-between border-bottom">
                                                <p className="py-2">{edu.startDate} - {edu.endDate}</p>
                                                <p className="py-2 text-muted">{edu.institutionName}</p>
                                            </div>
                                        );
                                    })
                                }
                            </div>

                            <div className="col-12 bg-white p-0 px-2 mb-3 pb-3">
                                <div className="d-flex align-items-center justify-content-between border-bottom">
                                    <h4>Posts</h4>
                                </div>
                                {
                                    props.user.isPrivate && !props.personalProfile ?
                                        <p style={{ textAlign: "center" }}>No posts to show.</p>
                                        :
                                        (
                                            !props.posts || props.posts.length === 0 ?
                                                <p style={{ textAlign: "center" }}>No posts to show.</p>
                                                : <PostList posts={props.posts} />
                                        )
                                }
                            </div>

                        </div>
                    </div>
                }
            </div>
        </div>
    );
}

export default Profile;