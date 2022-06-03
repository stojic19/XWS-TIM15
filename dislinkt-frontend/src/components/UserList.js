import "../css/userCard.css"
import UserCard from "./UserCard";

const UserList = (props) => {
    return(
        <section className="our-webcoderskull padding-lg">
            <ul className="row">
            {
            (props.users).map((user, index) => {
                return (
                    <UserCard key={index} user={user} />
                );
            })}
                {/* <li className="col-12 col-md-4 col-lg-3">
                    <div className="cnt-block equal-hight" style={{height: "360px"}}>
                        <h6>Public</h6>
                        <figure><img src={require("../images/user-avatar.png")} class="img-responsive" alt=""></img></figure>
                        <h3><a href="">Michael Brown</a></h3>
                        <p>Freelance Web Developer</p>
                        <button type="button" class="btn btn-outline-primary">Follow</button>
                    </div>
                </li>
                <li className="col-12 col-md-4 col-lg-3">
                    <div className="cnt-block equal-hight" style={{height: "360px"}}>
                        <h6>Public</h6>
                        <figure><img src={require("../images/user-avatar.png")} class="img-responsive" alt=""></img></figure>
                        <h3><a href="">Michael Brown</a></h3>
                        <p>Freelance Web Developer</p>
                        <button type="button" class="btn btn-outline-primary">Follow</button>
                    </div>
                 </li> */}
            </ul>
        </section>
    );
}

export default UserList;