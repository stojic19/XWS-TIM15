import "../css/userCard.css";
import JobOfferCard from "./JobOfferCard";

const JobOffersList = (props) => {


    return (
        <>
            {!props.hideTitle && props.jobOffers.length == 0 && <h3 style={{ textAlign: "center" }}>No job offers found.</h3>}
            <section className="our-webcoderskull padding-lg p-1" style={{ maxWidth: "90%" }}>
                <ul className="row">
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
        </>
    );
}

export default JobOffersList;