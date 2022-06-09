import { useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Navbar from './components/Navbar';
import Login from './components/Login';
import Registration from './components/Registration';
import AdminNavbar from './components/AdminNavbar';
import RegularNavbar from './components/RegularNavbar';
import CreateCompanyRequest from './components/CreateCompanyRequest';
import Requests from './components/Requests';
import MyJobOffers from './components/MyJobOffers';
import JobOffersByCompanyId from './components/JobOffersByCompanyId';
import AllRegisteredCompanies from './components/AllRegisteredCompanies';
import CompanyProfile from './components/CompanyProfile';

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
          <Route key='/companyRegistrationRequests' exact path="/companyRegistrationRequests" element={[<AdminNavbar key='/companyRegistrationRequests' />,<Requests key='/companyRegistrationRequests1'/>]} />
          <Route key='/home' exact path="/home" element={[<RegularNavbar key='/home' />, <AllRegisteredCompanies key='/home1' />]} />
          <Route key='/registrationRequest' exact path="/registrationRequest" element={[<RegularNavbar key='/registrationRequest' />,<CreateCompanyRequest key='/registrationRequest1'/>]} />
          <Route key='/myJobOffers' exact path="/myJobOffers" element={[<RegularNavbar key='/myJobOffers' />,<MyJobOffers key='/myJobOffers1'/>]} />
          <Route key='/jobOffers' exact path="/jobOffers/:id" element={[<RegularNavbar key='/jobOffers' />,<JobOffersByCompanyId key='/jobOffers1'/>]} />
          <Route path="/companyProfile/:id" element={[<RegularNavbar/>,<CompanyProfile/>]} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
