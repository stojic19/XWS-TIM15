import "../css/userCard.css"

const JobOfferCard = (props) => {

    return (
        <li className="col-12 col-md-4 col-lg-3">
            {props.jobOffer &&
                <div className="cnt-block equal-hight" style={{ maxHeight: "100%" }}>
                    <h6>
                        {
                            props.jobOffer.isActive ? 'Active' : 'Closed'
                        }
                    </h6>
                    <figure><img src={require("../images/user-avatar.png")} className="img-responsive" alt=""></img></figure>
                    <h3>{props.jobOffer.name}</h3>
                    <p>{props.jobOffer.position}</p>
                    <p>Description : {props.jobOffer.description}</p>
                    <p>Requirements : {props.jobOffer.requirements}</p>
                </div>
            }
        </li>
    );
}

export default JobOfferCard;