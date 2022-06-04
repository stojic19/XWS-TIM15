const UserNavbar = () => {
    return (
        <div className="m-5 p-1">
        <ul className="nav nav-pills nav-fill ">
            <li className="nav-item" key={1}>
                <a style={{textDecoration: "none"}} href="/"><h1>Dislinkt</h1></a>
            </li>
            <li className="nav-item" key={2}>
                <a className="nav-link active" aria-current="page" href="/home">Home</a>
            </li>
            <li className="nav-item" key={3}>
                <a className="nav-link" href="/publicProfiles">Profiles</a>
            </li>
            <li className="nav-item" key={4}>
                <a className="nav-link" href="/jobOffers">Job offers</a>
            </li>
            <li className="nav-item" key={5}>
                <a className="nav-link" href="/personalProfile">Profile</a>
            </li>
            <li className="nav-item" key={6}>
                <a className="nav-link" href="/followRequests">Follow requests</a>
            </li>
            <li className="nav-item" key={7}>
                <a className="nav-link disabled" href="/" aria-disabled="true">Disabled</a>
            </li>
        </ul>
        <hr></hr>
    </div>
    );
}

export default UserNavbar;