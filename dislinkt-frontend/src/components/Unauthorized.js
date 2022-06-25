import { useNavigate } from 'react-router-dom';
import { useEffect } from "react";
import Swal from 'sweetalert2';

const Unauthorized = () => {
    const history = useNavigate();

    useEffect(() => {
        Swal.fire({
            position: 'top-end',
            icon: 'error',
            title: 'Unauthorized access!',
            showConfirmButton: false,
            timer: 2000
          })
        history('/');
      }, [])

    return (
        <>
        </>
    );
}

export default Unauthorized;