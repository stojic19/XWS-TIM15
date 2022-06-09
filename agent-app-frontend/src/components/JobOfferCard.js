import "../css/userCard.css"
import Swal from "sweetalert2";
import axios from 'axios';

const JobOfferCard = (props) => {

    const avtivate = () => {
        console.log(props)
        const headers = {
            'token': localStorage.getItem('token')
        }
        var data = {
            "id": props.jobOffer.id,
            "companyId": props.jobOffer.companyId
          }
          console.log(data)
        axios.put(axios.defaults.baseURL + 'api/Companies/JobOffer/Activate' , data, { headers })
            .then(res => {
                console.log(res)
                Swal.fire({
                    icon: 'success',
                    title: 'Success!',
                    text: 'Successfully activated!',
                });
                window.location.reload();
            }).catch(err => {
                console.log(err)
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: err.data,
                });
            });
    };

    return (
        <li className="col-12 col-md-4 col-lg-3">
            {props.jobOffer &&
                <div className="cnt-block equal-hight" style={{ maxHeight: "100%", maxWidth: "90%" }}>
                    <h6>
                        {
                            props.jobOffer.isActive ? 'Active' :
                                (props.myJobOffers ?
                                    <button onClick={() => avtivate()} className="btn btn-primary me-2">Activate</button>
                                    : 'Closed')
                        }
                    </h6>
                    <figure><img src={require("../images/user-avatar.png")} className="img-responsive" alt=""></img></figure>
                    <h3>{props.jobOffer.position}</h3>
                    <p>Description : {props.jobOffer.description}</p>
                    <p>Requirements : {props.jobOffer.requirements}</p>
                </div>
            }
        </li>
    );
}

export default JobOfferCard;