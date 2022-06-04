import "../css/userCard.css"
import UserCard from "./UserCard";

const UserList = (props) => {
    return(
        <section className="our-webcoderskull padding-lg">
            <ul className="row">
            {props.users.length == 0 && <h3 style={{textAlign: "center"}}>No users found.</h3>}
            {
            (props.users).map((user, index) => {
                return (
                    <UserCard key={index} user={user} />
                );
            })}
            </ul>
        </section>
    );
}

export default UserList;