import { useState } from "react"
import Swal from "sweetalert2"
import axios from "axios"
import { useNavigate } from "react-router-dom"

const CreateCompanyRequest = () => {
    const history = useNavigate();
    const [address, setAddress] = useState('');
    const [description, setDescription] = useState('');
    const [email, setEmail] = useState('');
    const [name, setName] = useState('');
    const [culture, setCulture] = useState('');
    const [phoneNumber, setPhoneNumber] = useState('');

    const Validate = () => {
        if (address === "" || description === "" || email === "" ||
        name === "" || culture === "" || phoneNumber === "") {
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

    const addRequest = async (e) => {
        e.preventDefault()
        if (!Validate())
            return;

        let companyInfo = {
            name: name,
            address: address,
            email: email,
            phoneNumber: phoneNumber,
            description: description,
            culture: culture,
        }

        let data = {
            companyInfo : companyInfo
        }

        const headers = {
            'token': localStorage.getItem('token')
        }

        axios.post(axios.defaults.baseURL + 'api/Companies', data, { headers })
            .then(res => {
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

    return (
        <div className="container">
            <div className="row">
                <div className="col-md-12">
                    <div className="well well-sm">
                        <form className="form-horizontal">
                            <legend className="text-left header">Company registration request</legend>

                            <div className="form-group">
                                <div className="col-md-8">
                                    <input type="text" value={name} onChange={(e) => setName(e.target.value)} placeholder="Name" className="form-control" />
                                </div>
                            </div>

                            <div className="form-group">
                                <div className="col-md-8">
                                    <input type="text" value={address} onChange={(e) => setAddress(e.target.value)} placeholder="Address" className="form-control" />
                                </div>
                            </div>

                            <div className="form-group">
                                <div className="col-md-8">
                                    <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="Email" className="form-control" />
                                </div>
                            </div>

                            <div className="form-group">
                                <div className="col-md-8">
                                    <input type="text" value={phoneNumber} onChange={(e) => setPhoneNumber(e.target.value)} placeholder="Phone number" className="form-control" />
                                </div>
                            </div>

                            <div className="form-group">
                                <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                <div className="col-md-8">
                                    <textarea className="form-control" value={description} onChange={(e) => setDescription(e.target.value)} placeholder="Description" rows="7"></textarea>
                                </div>
                            </div>
                            <div className="form-group">
                                <span className="col-md-1 col-md-offset-2 text-center"><i className="fa fa-pencil-square-o bigicon"></i></span>
                                <div className="col-md-8">
                                    <textarea className="form-control" value={culture} onChange={(e) => setCulture(e.target.value)} placeholder="Culture" rows="7"></textarea>
                                </div>
                            </div>
                            <div className="form-group">
                                <div className="col-md-12 text-left">
                                    <br></br>
                                    <button onClick={(e) => addRequest(e)} className="btn btn-primary btn-lg">Done</button>
                                </div>
                            </div>

                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}
export default CreateCompanyRequest;