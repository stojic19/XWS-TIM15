import "../css/userCard.css"
import UserCard from "./UserCard";
import { useState } from "react";

const UserList = (props) => {
    const[follows, setFollows] = useState(props.follows);
    const[followRequests, setFollowRequests] = useState(props.followRequests);
    const[displayFollowButtons, setDisplayFollowButtons] = useState(props.displayFollowButtons);

    const userFollowsUser = (id) => {
        let isFollowing = false;
        if(follows)
        follows.map((follow) =>{
            if(follow.id == id){
                isFollowing = true;
            }
        })
        return isFollowing;
    }

    const userSentFollowRequest = (id) => {
        let sentFollowRequest = false;
        if(followRequests)
        followRequests.map((follow) =>{
            if(follow.id == id){
                sentFollowRequest = true;
            }
        })
        return sentFollowRequest;
    }

    const DisplayFollowButtons = (id) => {
        if(props.displayFollowButtons){
            let userId = localStorage.getItem('user_id');
            if(id==userId)
                return false;
            return true;
        }
        return false;
    }

    return(
        <section className="our-webcoderskull padding-lg">
            <ul className="row">
            {props.users.length == 0 && <h3 style={{textAlign: "center"}}>No users found.</h3>}
            {
            (props.users).map((user, index) => {
                return (
                    <UserCard key={index} user={user} displayFollowButtons={DisplayFollowButtons(user.id)} isFollowing={userFollowsUser(user.id)} sentFollowRequest={userSentFollowRequest(user.id)}/>
                );
            })}
            </ul>
        </section>
    );
}

export default UserList;