import JobOfferPost from "./JobOfferPost";


const JobOfferPostList = (props) => {

    return (
        <div className="container align-content: center display: flex align-items: center mt-5">
            {props.offers &&
                (props.offers).map((offer, index) => {
                    return (

                        <JobOfferPost key={index} offer={offer} />

                    );
                })}
            {props.offers.length==0 && <h3 style={{textAlign: 'center'}}>No offer to show.</h3>}
        </div>
    );
}

export default JobOfferPostList;