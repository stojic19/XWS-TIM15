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
    const [firstName, setFirstName] = useState("");
    const [middleName, setMiddleName] = useState("");
    const [lastName, setLastName] = useState("");
    const [dateOfBirth, setDateOfBirth] = useState("");
    const [isPending, setIsPending] = useState(false);
    const history = useNavigate();

    const Validate = () => {
        if (email === "" || username === "" || password === "" || telephoneNumber === ""
            || gender === "" || firstName === "" || dateOfBirth === "" || middleName === "" || lastName === "") {
            Swal.fire({
                icon: 'warning',
                title: 'Oops...',
                text: 'All inputs must be filled!',
            });
            return false;
        }
        return true;
    }
    /*const FormatDate = (date) => {
        var month = date.getUTCMonth() + 1; //months from 1-12.
        var day = date.getUTCDate() + 1;
        var year = date.getUTCFullYear();
        return day + '/' + month + '/' + year;
    }*/
    const onSubmit = async (e) => {
        e.preventDefault();
        if (!Validate())
            return;
        setIsPending(true);
        const personalInfo = {
            "FirstName": firstName,
            "MiddleName": middleName,
            "LastName": lastName,
            "PhoneNumber": telephoneNumber,
            "Gender": gender,
            "Email": email,
            "BirthDate": dateOfBirth,
        }
        const registration = {
            "Username": username,
            "Password": password,
            "PersonalInfo": personalInfo
        };
        console.log(registration)
        axios.post(axios.defaults.baseURL + 'api/Users', registration)
            .then(() => {
                setIsPending(false);
                history('/');
            }).catch((err) => {
                console.log(err);
                setIsPending(false);
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.response.data,
                });
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
                    <label className="form-label">First name</label>
                    <input value={firstName} onChange={(e) => setFirstName(e.target.value)} type="text" className="form-control" id="InputFName" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Middle name</label>
                    <input value={middleName} onChange={(e) => setMiddleName(e.target.value)} type="text" className="form-control" id="InputMName" />
                </div>
                <div className="mb-3">
                    <label className="form-label">Last name</label>
                    <input value={lastName} onChange={(e) => setLastName(e.target.value)} type="text" className="form-control" id="InputLName" />
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
                        <option value="Other">Other</option>
                    </select>
                </div>
                <div className="mb-3">
                    <label className="form-label">Date of birth</label>
                    <DatePicker dateFormat="dd/MM/yyyy" selected={dateOfBirth} onChange={(date) => setDateOfBirth(date)} className="form-control" />
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