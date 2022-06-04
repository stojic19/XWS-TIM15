import { useParams } from "react-router-dom";
import PostList from "./PostList";
import { useState, useEffect } from "react";
import Swal from "sweetalert2";

import axios from 'axios';

const UserHome = () => {
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchPosts = async () => {
            setLoading(true);
            if(localStorage.getItem('user_id'))
            axios.get(axios.defaults.baseURL + 'posts/postsFollowed/' + localStorage.getItem('user_id'))
                .then(res => {
                    let posts = Array.from(res.data.posts)
                    setPosts(posts);
                    setLoading(false);
                }).catch(err => {
                    console.log(err);
                    setError(err);
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: err.data,
                    });
                });
        };
        fetchPosts();
    }, []);

    return (
        <div>
            {error.length==0 && loading && <h3 style={{textAlign: 'center'}}>Loading...</h3>}
            {!loading && posts && <PostList posts={posts} />}
            {error.length!=0 && <h3 style={{textAlign: 'center'}}>Error!</h3>}
        </div>
    );
}

export default UserHome;