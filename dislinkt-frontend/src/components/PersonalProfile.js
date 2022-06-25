import { useEffect, useState } from "react";
import Profile from "./Profile";
import axios from "axios";

const PersonalProfile = () => {
    const [user, setUser] = useState([])
    const [education, setEducation] = useState([])
    const [experience, setExperience] = useState([])
    const [skills, setSkills] = useState([])
    const [interests, setInterests] = useState([])
    const [posts, setPosts] = useState([]);

    useEffect(() => {
        const getUserById = async () => {
            axios.get(axios.defaults.baseURL + 'users/' + localStorage.getItem('user_id'))
                    .then(res => {
                        setUser(res.data.user)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const getWorkExperience = async () => {
            axios.get(axios.defaults.baseURL + 'workExperience/' + localStorage.getItem('user_id'))
                    .then(res => {
                        setExperience(res.data.workExperience)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const getSkills = async () => {
            axios.get(axios.defaults.baseURL + 'education/' + localStorage.getItem('user_id'))
                    .then(res => {
                        setEducation(res.data.education)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const getEducation = async () => {
            axios.get(axios.defaults.baseURL + 'skills/' + localStorage.getItem('user_id'))
                    .then(res => {
                        setSkills(res.data.skills)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const getInterests = async () => {
            axios.get(axios.defaults.baseURL + 'interests/' + localStorage.getItem('user_id'))
                    .then(res => {
                        setInterests(res.data.interests)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        const fetchPosts = async () => {
            axios.get(axios.defaults.baseURL + 'posts/postsFromUser/' + localStorage.getItem('user_id'))
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
        fetchPosts();
    }, [])

    return(
        <Profile user={user} 
                experience={experience} 
                skills={skills} 
                education={education} 
                interests={interests}
                personalProfile={true}
                posts={posts}
                ></Profile>
    );
}

export default PersonalProfile;