import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { useEffect } from 'react';

import Login from './components/Login';
import Navbar from './components/Navbar';
import Registration from './components/Registration';
import PostsByUserId from './components/PostsByUserId';
import UserNavbar from './components/UserNavbar';
import Home from './components/Home';
import Users from './components/Users';
import EditProfile from './components/EditProfile';
import UserProfile from './components/UserProfile';
import PersonalProfile from './components/PersonalProfile';
import EditWorkExperience from './components/EditWorkExperience';
import PublicProfileSearch from './components/PublicProfileSearch';
import EditEducation from './components/EditEducation';
import EditSkillsAndInterests from './components/EditSkillsAndInterests';
import FollowRequestsList from './components/FollowRequestsList';
import UserHome from './components/UserHome';
import AllProfiles from './components/AllProfiles';
import AllJobOffers from './components/AllJobOffers';

import axios from 'axios';
import CreatePost from './components/CreatePost';
import CreateJobOffer from './components/CreateJobOffer';

useEffect(() => {
  document.title = "Dislinkt"
}, [])

axios.defaults.baseURL = "http://localhost:8000/";
//axios.defaults.headers.post['Content-Type'] = 'application/json';
//axios.defaults.headers.post['Accept'] = 'application/json';
//axios.defaults.headers.post['Access-Control-Allow-Origin'] = 'http://localhost:3000';

axios.interceptors.request.use(function(config){
  const token = localStorage.getItem('token');
  config.headers.token = token ? `${token}` : '';
  return config; 
});

function App() {
  return (
    <div className="App">
      <Router>
      <Routes>
          <Route key='/' exact path="/" element={[<Navbar />, <PublicProfileSearch displayFollowButtons={false}/>]}/>
          <Route key='/login' exact path="/login" element={[<Navbar />,<Login />]}/>
          <Route key='/registration' exact path="/registration" element={[<Navbar />,<Registration />]}/>
          <Route key='/userPosts' exact path="/userPosts/:id" element={[<Navbar />,<PostsByUserId />]}/>
          <Route key='/home' exact path="/home" element={[<UserNavbar />, <UserHome/>]}/>
          <Route key='/publicProfiles' exact path="/publicProfiles" element={[<UserNavbar />, <PublicProfileSearch displayFollowButtons={true}/>]}/>
          <Route key='/allProfiles' exact path="/allProfiles" element={[<UserNavbar />, <AllProfiles displayFollowButtons={true}/>]}/>
          <Route key='/jobOffers' exact path='/jobOffers' element={[<UserNavbar />,<AllJobOffers />]}/>
          <Route key='/profile' exact path="/profile/:id" element={[<UserNavbar />, <UserProfile/>]}/>
          <Route key='/publicProfile' exact path="/publicProfile/:id" element={[<Navbar />, <UserProfile/>]}/>
          <Route key='/users' exact path="/users" element={[<Navbar />,<Users/>]}/>
          <Route key='/editProfile' exact path="/editProfile" element={[<UserNavbar />,<EditProfile/>]}/>
          <Route key='/editWorkExperience' exact path="/editWorkExperience" element={[<UserNavbar />,<EditWorkExperience/>]}/>
          <Route key='/editEducation' exact path="/editEducation" element={[<UserNavbar />,<EditEducation/>]}/>
          <Route key='/editSkillsAndInterests' exact path="/editSkillsAndInterests" element={[<UserNavbar />,<EditSkillsAndInterests/>]}/>
          <Route key='/personalProfile' exact path='/personalProfile' element={[<UserNavbar />,<PersonalProfile/>]}/>
          <Route key='/followRequests' exact path='/followRequests' element={[<UserNavbar />,<FollowRequestsList/>]}/>   
          <Route key='/createPost' exact path='/createPost' element={[<UserNavbar />,<CreatePost/>]}/>      
          <Route key='/createJobOffer' exact path='/createJobOffer' element={[<UserNavbar />,<CreateJobOffer/>]}/>     
          </Routes>
        </Router>
    </div>
  );
}

export default App;
