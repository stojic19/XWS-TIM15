import { useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Navbar from './components/Navbar';
import Login from './components/Login';
import Registration from './components/Registration';
import AdminNavbar from './components/AdminNavbar';
import RegularNavbar from './components/RegularNavbar';

import axios from 'axios';

function App() {

  useEffect(() => {
    document.title = "Agent"
  }, [])

  axios.defaults.baseURL = "https://localhost:44316/";
  //axios.defaults.headers.post['Content-Type'] = 'application/json';
  //axios.defaults.headers.post['Accept'] = 'application/json';
  //axios.defaults.headers.post['Access-Control-Allow-Origin'] = 'http://localhost:3000';

  axios.interceptors.request.use(function (config) {
    const token = localStorage.getItem('token');
    config.headers.token = token ? `${token}` : '';
    return config;
  });

  return (
    <div>
      <Router>
        <Routes>
          <Route key='/' exact path="/" element={[<Navbar key='/' />, <Login key='/1' />]} />
          <Route key='/registration' exact path="/registration" element={[<Navbar key='/registration' />, <Registration key='/registration1' />]} />
          <Route key='/companyRegistrationRequests' exact path="/companyRegistrationRequests" element={[<AdminNavbar key='/companyRegistrationRequests' />]} />
          <Route key='/home' exact path="/home" element={[<RegularNavbar key='/home' />]} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
