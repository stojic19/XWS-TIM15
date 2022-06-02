import { useState } from "react";
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import Swal from "sweetalert2";

const Registration = () => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [telephoneNumber, setTelephoneNumber] = useState("");
    const [gender, setGender] = useState("");
    const [name, setName] = useState("");
    const [dateOfBirth, setDateOfBirth] = useState("");
    const [biography, setBiography] = useState("");
    const [isPrivate, setIsPrivate] = useState(false);
    const [isPending, setIsPending] = useState(false);
    const history = useNavigate();

    const Validate = () => {
        if (email === "" || username === "" || password === "" || telephoneNumber === ""
            || gender === "" || name === "" || dateOfBirth === "" || biography === "") {
            Swal.fire({
                icon: 'warning',
                title: 'Oops...',
                text: 'All inputs must be filled!',
            });
            return false;
        }
        return true;
    }
    const FormatDate = (date) => {
        var list = date.split('-');
        return list[2] + '/' + list[1] + '/' + list[0];
    }
    const onSubmit = async (e) => {
        e.preventDefault();
        if (!Validate())
            return;
        setIsPending(true);
        const registration = {
            "username": username,
            "password": password,
            "email": email,
            "name": name,
            "telephoneNo": telephoneNumber,
            "gender": gender,
            "biography": biography,
            "isPrivate": isPrivate === "false" ? false : true,
            "dateOfBirth": FormatDate(dateOfBirth),
        };
        console.log(registration)
        axios.post(axios.defaults.baseURL + 'users', registration)
        .then(res => {
            if(res.data.response.includes("Added user")){
                setIsPending(false);
                history('/');
            }else{
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
        <div className="m-5">
            <form style={{ maxWidth: "50%", alignContent: "center", alignItems: "center", margin: "auto" }}>
                <div className="mb-3">
                    <label className="form-label">Email address</label>
                    <input value={email} onChange={(e) => setEmail(e.target.value)} type="email" className="form-control" id="InputEmail" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Username</label>
                    <input value={username} onChange={(e) => setUsername(e.target.value)} type="text" className="form-control" id="InputUsername" />
                    <div id="usernameHelp" className="form-text">Username must be unique.</div>
                </div>
                <div className="mb-3">
                    <label className="form-label">Password</label>
                    <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" className="form-control" id="InputPassword" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Name</label>
                    <input value={name} onChange={(e) => setName(e.target.value)} type="text" className="form-control" id="InputName" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Telephone number</label>
                    <input value={telephoneNumber} onChange={(e) => setTelephoneNumber(e.target.value)} type="text" className="form-control" id="InputTelNu" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Gender</label>
                    <select id="InputGender"
                        name="gender"
                        className="form-control"
                        value={gender}
                        onChange={(e) => setGender(e.target.value)}
                    >
                        <option value="" disabled>Choose gender</option>
                        <option value="Female">Female</option>
                        <option value="Male">Male</option>
                    </select>
                </div>
                <div className="mb-3">
                    <label className="form-label">Date of birth</label>
                    <input selected={dateOfBirth} onChange={(e) => setDateOfBirth(e.target.value)} type="date" className="form-control" id="InputDateOfBirth" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Biography</label>
                    <textarea value={biography} onChange={(e) => setBiography(e.target.value)} type="text" className="form-control" id="InputBiography" />
                </div>
                <div className="mb-3 form-check">
                    <input value={isPrivate} onChange={(e) => setIsPrivate(!isPrivate)} type="checkbox" className="form-check-input" id="CheckIsPrivate" />
                    <label className="form-check-label">Private profile</label>
                </div>
                <div>
                    {!isPending && <button onClick={(e) => onSubmit(e)} type="submit" className="btn btn-primary">Submit</button>}
                    {isPending && <label>Registration...</label>}
                </div>
            </form>
        </div>
    );
}

export default Registration;