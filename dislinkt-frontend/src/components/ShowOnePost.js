import PostCard from "./PostCard";
import {useParams} from "react-router-dom";
import { useState, useEffect } from "react";
import axios from "axios";

const ShowOnePost = () => {

    const { id } = useParams();
    const [post, setPost] = useState();
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchPost();
    }, [])

    const fetchPost = async () =>{
        setLoading(true);
        axios.get(axios.defaults.baseURL + 'posts/posts/' + id)
            .then(res => {
                setPost(res.data.post)
                //console.log(res.data.post)
                setLoading(false);
            });

    }

    return(
        <>
            {!loading && post && <PostCard post={post}></PostCard>}
        </>
    );
}

export default ShowOnePost;