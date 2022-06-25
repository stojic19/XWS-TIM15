import { useNavigate } from 'react-router-dom';
import { useEffect } from "react";
import Swal from 'sweetalert2';

const Unauthorized = () => {
    const history = useNavigate();

    useEffect(() => {
        if (localStorage.getItem('user_id') === '') {
            Swal.fire({
                position: 'top-end',
                icon: 'error',
                title: 'Unauthorized access!',
                showConfirmButton: false,
                timer: 2000
            })
            history('/');
        }
        else {
            window.location.reload();
        }
    }, [])

    return (
        <>
        </>
    );
}

export default Unauthorized;