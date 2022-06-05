const AdminNavbar = () => {

    const logout = () => {
        localStorage.setItem('token', '');
    }

    return (
        <div className="m-4 p-1">
        <ul className="nav nav-pills nav-fill ">
            <li className="nav-item" key={1}>
                <a style={{textDecoration: "none"}} onClick={(e) => logout(e)} href="/"><h1>Agent</h1></a>
            </li>
            <li className="nav-item" key={2}>
                <a className="nav-link" href="/" onClick={(e) => logout(e)}>Logout</a>
            </li>
        </ul>
        <hr></hr>
    </div>
    );
}

export default AdminNavbar;