import { Link } from "react-router-dom";
import { useState } from "react";
import { useNavigate  } from 'react-router-dom';
import axios from 'axios';

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [isPending, setIsPending] = useState(false);
    const [message, setMessage] = useState("");
    const history = useNavigate();

    const onSubmit = async (e) => {
        e.preventDefault();
        const login = { username: email, password: password };
        setIsPending(true);
        const res = await axios.post(axios.defaults.baseURL + 'login', login);
        if (res.status === 200) {
            setIsPending(false);
            localStorage.setItem('token', res.data.access_token);
            //localStorage.setItem('auth_name', res.data.name);
            history.push('/index');
        } else {
            setIsPending(false);
            setMessage("Invalid credentials.");
        }
    }

    return (
        <div className="container max-width:80% max-height:80% align-content: center display: flex align-items: center mt-5">
            <form>
                <div className="mb-3">
                    <label className="form-label">Email address</label>
                    <input value={email} onChange={(e) => setEmail(e.target.value)} required type="email" className="form-control" id="InputEmail" aria-describedby="emailHelp" />
                    <div id="emailHelp" className="form-text">We'll never share your email with anyone else.</div>
                </div>
                <div className="mb-3">
                    <label className="form-label">Password</label>
                    <input value={password} onChange={(e) => setPassword(e.target.value)} required type="password" className="form-control" id="InputPassword" />
                </div>
                {!isPending && <span className="right">
                    <Link to="#" onClick={(e) => onSubmit(e)} type="submit" className="btn btn-primary">Login</Link>
                </span>}
                {isPending && <label>Logging...</label>}
                <div className="mb-3">
                    {message}
                </div>
            </form>
        </div>
    );
}

export default Login;