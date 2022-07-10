import NotificationsList from "./NotificationsList";
import { Dropdown } from "react-bootstrap";
import { v4 as uuidv4 } from 'uuid';

const UserNavbar = () => {

    const logout = () => {
        localStorage.setItem('token', '');
        localStorage.setItem('user_id', '');
    }

    return (
        <div className="m-5 p-1">
            <ul className="nav nav-pills nav-fill ">
                <li className="nav-item" key={uuidv4()}>
                    <a style={{ textDecoration: "none" }} href="/" onClick={(e) => logout(e)}><h1>Dislinkt</h1></a>
                </li>
                <li className="nav-item" key={uuidv4()}>
                    <a className="nav-link active" aria-current="page" href="/home">Home</a>
                </li>
                <li className="nav-item" key={uuidv4()}>
                    <a className="nav-link" href="/createPost">New post</a>
                </li>
                <li className="nav-item" key={uuidv4()}>
                    <Dropdown>
                        <Dropdown.Toggle variant="light" id="dropdown-basic">
                            <a className="nav-link" style={{ display: 'inline' }}>Users</a>
                        </Dropdown.Toggle>

                        <Dropdown.Menu>
                            <Dropdown.Item href="/allProfiles">All</Dropdown.Item>
                            <Dropdown.Item href="/publicProfiles">Explore</Dropdown.Item>
                        </Dropdown.Menu>
                    </Dropdown>
                </li>
                <li className="nav-item" key={uuidv4()}>
                    <Dropdown>
                        <Dropdown.Toggle variant="light" id="dropdown-basic">
                            <a className="nav-link" style={{ display: 'inline' }}>Job offers</a>
                        </Dropdown.Toggle>

                        <Dropdown.Menu>
                            <Dropdown.Item href="/createJobOffer">New</Dropdown.Item>
                            <Dropdown.Item href="/jobOffers">All</Dropdown.Item>
                            <Dropdown.Item href="/recommendedJobOffers">Recommended</Dropdown.Item>
                        </Dropdown.Menu>
                    </Dropdown>
                </li>
                <li className="nav-item" key={uuidv4()}>
                    <a className="nav-link" href="/personalProfile">Profile</a>
                </li>
                <li>
                    <NotificationsList></NotificationsList>
                </li>
                <li className="nav-item" key={uuidv4()}>
                    <a className="nav-link" href="/" onClick={(e) => logout(e)}>Logout</a>
                </li>
            </ul>
            <hr></hr>
        </div>
    );
}

export default UserNavbar;