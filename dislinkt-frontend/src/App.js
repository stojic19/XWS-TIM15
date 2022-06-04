import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

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

import axios from 'axios';

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
          <Route exact path="/" element={[<Navbar />, <PublicProfileSearch displayFollowButtons={false}/>]}/>
          <Route exact path="/login" element={[<Navbar />,<Login />]}/>
          <Route exact path="/registration" element={[<Navbar />,<Registration />]}/>
          <Route exact path="/userPosts/:id" element={[<Navbar />,<PostsByUserId />]}/>
          <Route exact path="/home" element={[<UserNavbar />, <UserHome/>]}/>
          <Route exact path="/publicProfiles" element={[<UserNavbar />, <PublicProfileSearch displayFollowButtons={true}/>]}/>
          <Route exact path="/jobOffers" element={[<UserNavbar />]}/>
          <Route exact path="/profile/:id" element={[<UserNavbar />, <UserProfile/>]}/>
          <Route exact path="/users" element={[<Navbar />,<Users/>]}/>
          <Route exact path="/editProfile" element={[<UserNavbar />,<EditProfile/>]}/>
          <Route exact path="/editWorkExperience" element={[<UserNavbar />,<EditWorkExperience/>]}/>
          <Route exact path="/editEducation" element={[<UserNavbar />,<EditEducation/>]}/>
          <Route exact path="/editSkillsAndInterests" element={[<UserNavbar />,<EditSkillsAndInterests/>]}/>
          <Route exact path='/personalProfile' element={[<UserNavbar />,<PersonalProfile/>]}/>
          <Route exact path='/followRequests' element={[<UserNavbar />,<FollowRequestsList/>]}/>
          </Routes>
        </Router>
    </div>
  );
}

export default App;
