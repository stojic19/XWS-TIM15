const UserNavbar = () => {

    const logout = () => {
        localStorage.setItem('token', '');
        localStorage.setItem('user_id', '');
    }

    return (
        <div className="m-5 p-1">
        <ul className="nav nav-pills nav-fill ">
            <li className="nav-item" key={1}>
                <a style={{textDecoration: "none"}} href="/" onClick={(e) => logout(e)}><h1>Dislinkt</h1></a>
            </li>
            <li className="nav-item" key={2}>
                <a className="nav-link active" aria-current="page" href="/home">Home</a>
            </li>
            <li className="nav-item" key={33}>
                <a className="nav-link" href="/createPost">New post</a>
            </li>
            <li className="nav-item" key={44}>
                <a className="nav-link" href="/createJobOffer">New job offer</a>
            </li>
            <li className="nav-item" key={3}>
                <a className="nav-link" href="/publicProfiles">Explore</a>
            </li>
            <li className="nav-item" key={4}>
                <a className="nav-link" href="/allProfiles">Profiles</a>
            </li>
            <li className="nav-item" key={5}>
                <a className="nav-link" href="/jobOffers">Job offers</a>
            </li>
            <li className="nav-item" key={6}>
                <a className="nav-link" href="/personalProfile">Profile</a>
            </li>
            <li className="nav-item" key={8}>
                <a className="nav-link" href="/" onClick={(e) => logout(e)}>Logout</a>
            </li>
        </ul>
        <hr></hr>
    </div>
    );
}

export default UserNavbar;