import { useParams } from "react-router-dom";
import PostList from "./PostList";
import { useState, useEffect } from "react";
import Swal from "sweetalert2";

import axios from 'axios';

const PostsByUserId = () => {
    const { id } = useParams();
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchPosts = async () => {
            setLoading(true);
            axios.get(axios.defaults.baseURL + 'posts/postsFromUser/' + id)
                .then(res => {
                    let posts = Array.from(res.data.posts)
                    setPosts(posts);
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
        fetchPosts();
    }, []);

    return (
        <div>
            {loading && <h3>Loading...</h3>}
            {!loading && posts && <PostList posts={posts} />}
        </div>
    );
}

export default PostsByUserId;