import {  useEffect, useState } from 'react';
import axios from 'axios';

const SelectBoxCompany = ({selectedValue,setSelectedValue,path,message,className}) => {
    const [objects, setObjects] = useState([]);

    useEffect(() => {    
        axios.get(axios.defaults.baseURL + path).then(res => {
            return res.data;
         }).then(data => {
            setObjects(data);
        }).catch(err => {
            console.log(err);
        })
    }, [])
    return (  
        <select 
            className={className}
            id="select"
            name="select"
            value={selectedValue}
            onChange={(e) => setSelectedValue(e.target.value)}
            >
                <option disabled value="">{message}</option>
            {objects && (objects.map(object => (
                <option key={object.id} value={object.id}>{object.companyInfo.name}</option>
            )))}
								
        </select>
    );
}
 
export default SelectBoxCompany;