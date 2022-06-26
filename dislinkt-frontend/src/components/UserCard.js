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
                props.buttonClickChanger();
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
                //console.log(res)
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: res.data.response,
                });
                props.buttonClickChanger();
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
                props.buttonClickChanger();
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
                props.buttonClickChanger();
            });
    }

    const unblock = async (e) => {
        e.preventDefault();

        axios.delete(axios.defaults.baseURL + 'followers/block/' + localStorage.getItem('user_id') + '/' + props.user.id)
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
                props.buttonClickChanger();
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
                {
                    props.blocked && <button onClick={(e) => unblock(e)} type="button" className="btn btn-outline-primary">Unblock</button>
                }
            </div>
        </li>
    );
}

export default UserCard;