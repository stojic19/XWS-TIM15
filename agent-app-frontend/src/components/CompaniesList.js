import "../css/userCard.css"
import CompanyCard from "./CompanyCard";

const CompaniesList = (props) => {

    return (
        <>
            {props.companies.length == 0 && props.admin && <h3 style={{ textAlign: "center" }}>No requests found.</h3>}
            {props.companies.length == 0 && !props.admin && <h3 style={{ textAlign: "center" }}>No companies found.</h3>}
            <section className="our-webcoderskull padding-m" style={{ maxWidth: '95%' }}>
                <ul className="row">
                    {
                        (props.companies).map((company, index) => {
                            return (
                                <CompanyCard key={index} company={company} admin={props.admin} />
                            );
                        })}
                </ul>
            </section>
        </>
    );
}

export default CompaniesList;