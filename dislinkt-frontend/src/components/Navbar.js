const Navbar = () => {
    return (
        <div className="m-5 p-1">
        <ul className="nav nav-pills nav-fill ">
            <li className="nav-item" key={1}>
                <a style={{textDecoration: "none"}} href="/"><h1>Dislinkt</h1></a>
            </li>
            <li className="nav-item" key={2}>
                <a className="nav-link active" aria-current="page" href="/">Home</a>
            </li>
            <li className="nav-item" key={3}>
                <a className="nav-link" href="/login">Login</a>
            </li>
            <li className="nav-item" key={4}>
                <a className="nav-link" href="/registration">Registration</a>
            </li>
            <li className="nav-item" key={5}>
                <a className="nav-link disabled" href="/" aria-disabled="true">Disabled</a>
            </li>
        </ul>
        <hr></hr>
    </div>
    );
}

export default Navbar;