import "../css/userCard.css"

const UserCard = (props) => {
    return(
        <li className="col-12 col-md-4 col-lg-3">
                    <div className="cnt-block equal-hight" style={{height: "360px"}}>
                        <h6>
                            {
                                props.user.isPrivate ? 'Private' : 'Public'
                            }
                        </h6>
                        <figure><img src={require("../images/user-avatar.png")} class="img-responsive" alt=""></img></figure>
                        <h3><a href={'/profile/' + props.user.id}>{props.user.name}</a></h3>
                        {/* <p>Freelance Web Developer</p> */}
                        <button type="button" class="btn btn-outline-primary">Follow</button>
                    </div>
        </li>
    );
}

export default UserCard;