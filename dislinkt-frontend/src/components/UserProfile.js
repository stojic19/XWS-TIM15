import { useEffect, useState } from "react";
import Profile from "./Profile";
import axios from "axios";
import { useParams } from "react-router-dom";

const UserProfile = () => {
    const { id } = useParams();
    const [user, setUser] = useState(null)
    const [education, setEducation] = useState([])
    const [experience, setExperience] = useState([])
    const [skills, setSkills] = useState([])
    const [interests, setInterests] = useState([])
    const [personalProfile, setPersonalProfile] = useState(false);
    const [posts, setPosts] = useState([])

    useEffect(() => {
        const getUserById = async () => {
            axios.get(axios.defaults.baseURL + 'users/' + id)
                    .then(res => {
                        console.log(res.data)
                        setUser(res.data.user)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const getWorkExperience = async () => {
            axios.get(axios.defaults.baseURL + 'workExperience/' + id)
                    .then(res => {
                        setExperience(res.data.workExperience)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const getSkills = async () => {
            axios.get(axios.defaults.baseURL + 'education/' + id)
                    .then(res => {
                        setEducation(res.data.education)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const getEducation = async () => {
            axios.get(axios.defaults.baseURL + 'skills/' + id)
                    .then(res => {
                        setSkills(res.data.skills)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const getInterests = async () => {
            axios.get(axios.defaults.baseURL + 'interests/' + id)
                    .then(res => {
                        setInterests(res.data.interests)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const isPersonalProfile = () => {
            setPersonalProfile(id===localStorage.getItem('user_id'));
        }

        const getPosts = async () => {
            axios.get(axios.defaults.baseURL + 'posts/postsFromUser/' + id)
                .then(res => {
                    let posts = Array.from(res.data.posts)
                    setPosts(posts);
                }).catch(err => {
                    console.log(err);
                });
        };

        getUserById();
        getWorkExperience();
        getSkills();
        getEducation();
        getInterests();
        isPersonalProfile();
        getPosts();
    }, [])

    return(
        <>{user && <Profile user={user} 
                experience={experience} 
                skills={skills} 
                education={education} 
                interests={interests}
                personalProfile={personalProfile}
                posts={posts}></Profile>}</>
    );
}

export default UserProfile;