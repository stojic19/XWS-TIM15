import "../css/userCard.css"
import { useState, useEffect } from "react";
import Swal from "sweetalert2";

import axios from 'axios';

const UserCard = (props) => {

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
                window.location.reload()
            });
    }

    const unfollow = async (e) => {
        e.preventDefault();

        const update = {
            "followerId": localStorage.getItem('user_id'),
            "followedId": props.user.id
        };
        axios.delete(axios.defaults.baseURL + 'followers/follow', {headers: {}, data: update})
            .then(res => {
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
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
                window.location.reload()
            });
    }

    const cancelFollowRequest = async (e) => {
        e.preventDefault();

        const update = {
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
            });
    }
    return (
        <li className="col-12 col-md-4 col-lg-3">
            <div className="cnt-block equal-hight" style={{ height: "360px" }}>
                <h6>
                    {
                        props.user.isPrivate ? 'Private' : 'Public'
                    }
                </h6>
                <figure><img src={require("../images/user-avatar.png")} className="img-responsive" alt=""></img></figure>
                {localStorage.getItem('user_id') && <h3><a href={'/profile/' + props.user.id}>{props.user.name}</a></h3>}
                {!localStorage.getItem('user_id') && <h3><a href={'/publicProfile/' + props.user.id}>{props.user.name}</a></h3>}
                {/* <p>Freelance Web Developer</p> */}
                {
                    !props.user.isPrivate && props.displayFollowButtons && !props.isFollowing 
                    && <button onClick={(e) => follow(e)} type="button" className="btn btn-outline-primary">Follow</button>
                }
                {
                    props.isFollowing && props.displayFollowButtons 
                    && <button onClick={(e) => unfollow(e)} type="button" className="btn btn-outline-primary">Unfollow</button>
                }
                {
                    props.user.isPrivate && props.displayFollowButtons && !props.isFollowing && !props.sentFollowRequest 
                    && <button onClick={(e) => followRequest(e)} type="button" className="btn btn-outline-primary">Follow request</button>
                }
                {
                    props.user.isPrivate && props.displayFollowButtons && !props.isFollowing && props.sentFollowRequest 
                    && <button onClick={(e) => cancelFollowRequest(e)} type="button" className="btn btn-outline-primary">Cancel follow request</button>
                }
            </div>
        </li>
    );
}

export default UserCard;