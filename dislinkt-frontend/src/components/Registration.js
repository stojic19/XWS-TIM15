import { useState } from "react";
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import Swal from "sweetalert2";
import DatePicker from "react-datepicker";

const Registration = () => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [telephoneNumber, setTelephoneNumber] = useState("");
    const [gender, setGender] = useState("");
    const [name, setName] = useState("");
    const [dateOfBirth, setDateOfBirth] = useState("");
    const [biography, setBiography] = useState("");
    const [isPrivate, setIsPrivate] = useState("");
    const [isPending, setIsPending] = useState(false);
    const history = useNavigate();

    const Validate = () => {
        if (email === "" || username === "" || password === "" || telephoneNumber === ""
            || gender === "" || name === "" || dateOfBirth === "" || biography === "" || isPrivate === "") {
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
        var month = date.getUTCMonth() + 1; //months from 1-12.
        var day = date.getUTCDate() + 1;
        var year = date.getUTCFullYear();
        return day + '/' + month + '/' + year;
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
        //console.log(registration)
        axios.post(axios.defaults.baseURL + 'users', registration)
            .then(res => {
                if (res.data.response.includes("Id:")) {
                    let id = res.data.response.split(':')
                    const notificationSettings = {
                        "userId": id[1],
                        "followerIdsForPosts": [],
                        "followerIdsForMessages": [],
                        "getNotificationsForMyPosts": true,
                    };
                    axios.post(axios.defaults.baseURL + 'notificationSettings', { notificationSettings: notificationSettings })
                        .then(res => {
                            console.log(res);
                            setIsPending(false);
                            history('/');
                        });
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
                    <DatePicker dateFormat="dd/MM/yyyy" selected={dateOfBirth} onChange={(date) => setDateOfBirth(date)} className="form-control" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Biography</label>
                    <textarea value={biography} onChange={(e) => setBiography(e.target.value)} type="text" className="form-control" id="InputBiography" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Profile visibility</label>
                    <select id="InputVisibility"
                        name="visibility"
                        className="form-control"
                        value={isPrivate}
                        onChange={(e) => setIsPrivate(e.target.value)}
                    >
                        <option value="" disabled>Choose visibility</option>
                        <option value="true">Private</option>
                        <option value="false">Public</option>
                    </select>
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