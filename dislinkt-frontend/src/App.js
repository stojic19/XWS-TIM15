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

import axios from 'axios';

axios.defaults.baseURL = "http://localhost:8000/";
//axios.defaults.headers.post['Content-Type'] = 'application/json';
//axios.defaults.headers.post['Accept'] = 'application/json';
//axios.defaults.headers.post['Access-Control-Allow-Origin'] = 'http://localhost:3000';

/*axios.interceptors.request.use(function(config){
  const token = localStorage.getItem('auth_token');
  config.headers.Authorization = token ? `Bearer ${token}` : '';
  return config; 
});*/

function App() {
  return (
    <div className="App">
      <Router>
      <Routes>
          <Route exact path="/" element={[<Navbar />, <Home/>]}/>
          <Route exact path="/login" element={[<Navbar />,<Login />]}/>
          <Route exact path="/registration" element={[<Navbar />,<Registration />]}/>
          <Route exact path="/userPosts/:id" element={[<Navbar />,<PostsByUserId />]}/>
          <Route exact path="/home" element={[<UserNavbar />]}/>
          <Route exact path="/profiles" element={[<UserNavbar />]}/>
          <Route exact path="/jobOffers" element={[<UserNavbar />]}/>
          <Route exact path="/profile/:id" element={[<UserNavbar />, <UserProfile/>]}/>
          <Route exact path="/users" element={[<Navbar />,<Users/>]}/>
          <Route exact path="/editProfile" element={[<UserNavbar />,<EditProfile/>]}/>
          <Route exact path="/editWorkExperience" element={[<UserNavbar />,<EditWorkExperience/>]}/>
          <Route exact path='/personalProfile' element={[<UserNavbar />,<PersonalProfile/>]}/>
          </Routes>
        </Router>
    </div>
  );
}

export default App;
