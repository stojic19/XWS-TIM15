import Profile from "./Profile";
import { useParams } from "react-router";
import { useState, useEffect } from "react";
import axios from "axios";

const CompanyProfile = () =>{

    const { id } = useParams();
    const [company, setCompany] = useState({})

        

    useEffect(() => {
        const getCompanyById = async () => {
            axios.get(axios.defaults.baseURL + 'api/Companies/' + id)
                    .then(res => {
                        setCompany(res.data)
                    }).catch(err => {
                        console.log(err);
                    });
        };
        getCompanyById();
    }, [])
        
    return(
        <>{ company && <Profile company={company}></Profile>}</>
    )
}
export default CompanyProfile;