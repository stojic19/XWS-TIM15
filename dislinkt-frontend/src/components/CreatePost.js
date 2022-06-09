import { useEffect, useState } from 'react';
import '../css/createPost.css'
import Swal from 'sweetalert2';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { projectStorage, projectFirestore, timestamp } from '../firebase/config';
import { v4 as uuidv4 } from 'uuid';

const CreatePost = () => {
    const history = useNavigate()
    const [title, setTitle] = useState('')
    const [contentText, setContentText] = useState('')
    const [images, setImages] = useState('');

    const saveImage = (file) => {
        // references
        const imageId = uuidv4();
        const storageRef = projectStorage.ref(imageId);
        const collectionRef = projectFirestore.collection('images');

        storageRef.put(file).on('state_changed', (snap) => {
            let percentage = (snap.bytesTransferred / snap.totalBytes) * 100;
            //setProgress(percentage);
        }, (err) => {
            //setError(err);
        }, async () => {
            const url = await storageRef.getDownloadURL();
            const createdAt = timestamp();
            await collectionRef.add({ url, createdAt });
            var imagesTemp = []
            imagesTemp = imagesTemp.concat(imageId)
            setImages(imagesTemp)
            //setUrl(url);
        });
    }

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

    const addPost = async (e) => {
        e.preventDefault()
        console.log(images);
        if (!Validate())
            return;

        let config = {
            newPost: {
                title: title,
                content: {
                    text: contentText,
                    links: [],
                    images: images
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
        axios.post(axios.defaults.baseURL + 'posts/posts', config, { headers })
            .then(res => {
                console.log(config)
                console.log(res.data)
                history('/home')
            }).catch(err => {
                console.log(err);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    }

    const [file, setFile] = useState(null);
    const [error, setError] = useState(null);

    const types = ['image/png', 'image/jpeg'];

    const handleChange = (e) => {
        let selected = e.target.files[0];

        if (selected && types.includes(selected.type)) {
            setFile(selected);
            saveImage(selected);
            setError('');
        } else {
            setFile(null);
            setError('Please select an image file (png or jpg)');
        }
    };

    return (
        <div className="container">
            <div className="row">
                <div className="col-md-12">
                    <div className="well well-sm">
                        <form className="form-horizontal">
                            <legend className="text-left header">New post</legend>

                            <div className="form-group">
                                <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                <div className="col-md-8">
                                    <input type="text" onChange={(e) => setTitle(e.target.value)} placeholder="Title" className="form-control" />
                                </div>
                            </div>

                            <div className="form-group">
                                <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                <div className="col-md-8">
                                    <textarea className="form-control" onChange={(e) => setContentText(e.target.value)} placeholder="Write..." rows="7"></textarea>
                                </div>
                            </div>
                            <div className="form-group">
                                <div className="col-md-8">
                                    <input type="file" className="form-control" onChange={handleChange} />
                                    {error && <p className="error">{error}</p>}
                                </div>
                            </div>
                            <div className="form-group">
                                <div className="col-md-12 text-left">
                                    <br></br>
                                    <button onClick={(e) => addPost(e)} className="btn btn-primary btn-lg">Done</button>
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