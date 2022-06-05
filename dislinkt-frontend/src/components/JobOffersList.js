import "../css/userCard.css";
import JobOfferCard from "./JobOfferCard";

const JobOffersList = (props) => {


    return (
        <section className="our-webcoderskull padding-lg">
            <ul className="row">
                {props.jobOffers.length == 0 && <h3 style={{ textAlign: "center" }}>No job offers found.</h3>}
                {
                    (props.jobOffers).map((jobOffer, index) => {
                        return (
                            <JobOfferCard key={index}
                                jobOffer={jobOffer}
                            />
                        );
                    })}
            </ul>
        </section>
    );
}

export default JobOffersList;