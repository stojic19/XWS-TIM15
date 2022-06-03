import "../css/userCard.css"

const UserCard = (user) => {
    return(
        <li className="col-12 col-md-4 col-lg-3">
                    <div className="cnt-block equal-hight" style={{height: "360px"}}>
                        <h6>Public</h6>
                        <figure><img src={require("../images/user-avatar.png")} class="img-responsive" alt=""></img></figure>
                        <h3><a href="">Michael Brown</a></h3>
                        <p>Freelance Web Developer</p>
                        <button type="button" class="btn btn-outline-primary">Follow</button>
                    </div>
        </li>
    );
}

export default UserCard;