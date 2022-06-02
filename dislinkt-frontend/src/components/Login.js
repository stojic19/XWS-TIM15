import { Link } from "react-router-dom";
import { useState } from "react";
import { useNavigate } from 'react-router-dom';
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
            "username": username,
            "password": password
        };
        setIsPending(true);
        axios.post(axios.defaults.baseURL + 'login', login).then(res => {
            if (res.data.response.includes("Added user")) {
                setIsPending(false);
                //localStorage.setItem('token', res.data.access_token);
                //localStorage.setItem('auth_name', res.data.name);
                history('/');
            } else {
                setIsPending(false);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: res.data.response,
                });
            }
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