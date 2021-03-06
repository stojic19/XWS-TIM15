const RegularNavbar = () => {

    const logout = () => {
        localStorage.setItem('token', '');
        localStorage.setItem('user_id', '');
        localStorage.setItem('user_id_owner','');
    }

    return (
        <div className="m-4 p-1">
            <ul className="nav nav-pills nav-fill ">
                <li className="nav-item" key={1}>
                    <a style={{ textDecoration: "none" }} onClick={(e) => logout(e)} href="/"><h1>Agent</h1></a>
                </li>
                <li className="nav-item" key={2}>
                    <a className="nav-link" href="/home">Home</a>
                </li>
                <li className="nav-item" key={3}>
                    <a className="nav-link" href="/registrationRequest">Registration request</a>
                </li>
                {localStorage.getItem('user_id_owner') === 'true' &&
                    <>
                        <li className="nav-item" key={4}>
                            <a className="nav-link" href="/myCompanies">My Companies</a>
                        </li>
                        <li className="nav-item" key={5}>
                            <a className="nav-link" href="/createJobOffer">Create job offer</a>
                        </li>
                        <li className="nav-item" key={6}>
                            <a className="nav-link" href="/myJobOffers">My job offers</a>
                        </li>
                    </>
                }
                <li className="nav-item" key={7}>
                    <a className="nav-link" href="/" onClick={(e) => logout(e)}>Logout</a>
                </li>
            </ul>
            <hr></hr>
        </div>
    );
}

export default RegularNavbar;