import { useState } from "react";
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';
import Swal from "sweetalert2";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [isPending, setIsPending] = useState(false);
    const history = useNavigate();

    const Validate = () => {
        if (username === "" || password === "") {
            Swal.fire({
                icon: 'warning',
                title: 'Oops...',
                text: 'All inputs must be filled!',
            });
            return false;
        }
        return true;
    }

    const onSubmit = async (e) => {
        e.preventDefault();
        if (!Validate())
            return;
        const login = {
            "Username": username,
            "Password": password
        };
        setIsPending(true);
        axios.post(axios.defaults.baseURL + 'api/LogIn', login).then(res => {
            setIsPending(false);
            localStorage.setItem('token', res.data.token);
            if(res.data.role==='Regular'){
                // check if owner and set localstorage variable
                localStorage.setItem('user_id', res.data.id);
                axios.get(axios.defaults.baseURL + 'api/Companies/User/' + res.data.id, login)
                .then(res => {
                    localStorage.setItem('user_id_owner', res.data.length!=0);
                    history('/home');
                })
                .catch(()=>{
                    localStorage.setItem('user_id_owner', false);
                    history('/home');
                })
            }
            else
                history('/companyRegistrationRequests');

        }).catch(err => {
            setIsPending(false);
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: err.response.data,
            });
        });
    }

    return (
        <div className="container align-content: center display: flex align-items: center mt-5">
            <form style={{ maxWidth: "50%", margin: "auto" }}>
                <div className="mb-3">
                    <label className="form-label">Username</label>
                    <input value={username} onChange={(e) => setUsername(e.target.value)} required type="text" className="form-control" id="InputUsername" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Password</label>
                    <input value={password} onChange={(e) => setPassword(e.target.value)} required type="password" className="form-control" id="InputPassword" />
                </div>
                <div className="mb-3">
                    {!isPending && <span className="right">
                        <Link to="#" onClick={(e) => onSubmit(e)} type="submit" className="btn btn-primary">Login</Link>
                    </span>}
                    {isPending && <label>Logging...</label>}
                </div>
            </form>
        </div>
    );
}

export default Login;