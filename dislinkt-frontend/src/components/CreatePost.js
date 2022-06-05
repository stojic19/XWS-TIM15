import { useState } from 'react';
import '../css/createPost.css'
import Swal from 'sweetalert2';
import axios from 'axios';

const CreatePost = () =>{
    const [title, setTitle] = useState('')
    const [contentText, setContentText] = useState('')

    const Validate = () => {
        if (title === "" || contentText === "") {
            Swal.fire({
                icon: 'warning',
                title: 'Oops...',
                text: 'All inputs must be filled!',
            });
            return false;
        }
        if (localStorage.getItem('user_id').length === 0) {
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: 'You are not logged in!',
            });
            return false;
        }
        return true;
    }

    const addPost = async (e)=>{
        e.preventDefault()
        if (!Validate())
            return;

        let config={
            newPost: {
                title: title,
                content: {
                    text: contentText,
                    links: [],
                    images: []
                },
                owner: {
                    id: localStorage.getItem('user_id')
                }
            }
        } 
        const headers = {
            'token': localStorage.getItem('token')
        }
        console.log(config)
        axios.post(axios.defaults.baseURL + 'posts/posts', config, {headers})
                .then(res => {
                    console.log(config)
                    console.log(res.data)
                }).catch(err => {
                    console.log(err);
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: err.data,
                    });
                });
    } 

    return(
        <div className="container">
            <div className="row">
                <div className="col-md-12">
                    <div className="well well-sm">
                        <form className="form-horizontal">
                                <legend className="text-left header">New post</legend>

                                <div className="form-group">
                                    <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                    <div className="col-md-8">
                                        <input type="text" onChange={(e)=>setTitle(e.target.value)} placeholder="Title" className="form-control"/>
                                    </div>
                                </div>

                                <div className="form-group">
                                    <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                    <div className="col-md-8">
                                        <textarea className="form-control"onChange={(e)=>setContentText(e.target.value)} placeholder="Write..." rows="7"></textarea>
                                    </div>
                                </div>
                                <div className="form-group">
                                    <div className="col-md-8">          
                                        <input type="file" className="form-control" />
                                    </div>
                                </div>
                                <div className="form-group">
                                    <div className="col-md-12 text-left">
                                        <br></br>
                                        <button onClick={(e)=>addPost(e)} className="btn btn-primary btn-lg">Done</button>
                                    </div>
                                </div>
                                
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default CreatePost;