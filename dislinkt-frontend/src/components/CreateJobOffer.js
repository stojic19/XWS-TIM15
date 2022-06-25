import { useState } from "react"
import Swal from "sweetalert2"
import axios from "axios"
import { useNavigate } from "react-router-dom"

const CreateJobOffer =  () =>{
    const history = useNavigate()
    const [position, setPosition] = useState('')
    const [description, setDescription] = useState('')
    const [requirements, setRequirements] = useState('')

    const Validate = () => {
        if (position === "" || description === ""|| requirements === "") {
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

    const addJobOffer = async (e)=>{
        e.preventDefault()
        if (!Validate())
            return;

         let newOffer = {
                position: position,
                description: description,
                requirements: requirements
            }
    
        const headers = {
            'token': localStorage.getItem('token')
        }

        axios.post(axios.defaults.baseURL + 'job_offers', newOffer, {headers})
                .then(res => {
                    //console.log(res.data)
                    history('/jobOffers')
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
                            <legend className="text-left header">New job offer</legend>

                            <div className="form-group">
                                <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                <div className="col-md-8">
                                    <input type="text" onChange={(e)=>setPosition(e.target.value)} placeholder="Position" className="form-control"/>
                                </div>
                            </div>

                            <div className="form-group">
                                <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                <div className="col-md-8">
                                    <textarea className="form-control"onChange={(e)=>setDescription(e.target.value)} placeholder="Description" rows="7"></textarea>
                                </div>
                            </div>
                            <div className="form-group">
                                <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                <div className="col-md-8">
                                    <textarea className="form-control"onChange={(e)=>setRequirements(e.target.value)} placeholder="Requirements" rows="7"></textarea>
                                </div>
                            </div>
                            <div className="form-group">
                                <div className="col-md-12 text-left">
                                    <br></br>
                                    <button onClick={(e)=>addJobOffer(e)}  className="btn btn-primary btn-lg">Done</button>
                                </div>
                            </div>
                            
                    </form>
                </div>
            </div>
        </div>
    </div>
    )
}
export default CreateJobOffer;