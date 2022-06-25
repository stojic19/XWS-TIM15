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
import CreatePost from './components/CreatePost';
import CreateJobOffer from './components/CreateJobOffer';
import Chat from './components/Chat';

import axios from 'axios';
import { v4 as uuidv4 } from 'uuid';

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
  useEffect(() => {
    document.title = "Dislinkt"
  }, [])
  
  return (
    <div className="App">
      <Router>
      <Routes>
          <Route key='/' exact path="/" element={[<Navbar key={uuidv4()}/>, <PublicProfileSearch key={uuidv4()} displayFollowButtons={false}/>]}/>
          <Route key='/login' exact path="/login" element={[<Navbar key='/login1'/>,<Login key='/login2'/>]}/>
          <Route key='/registration' exact path="/registration" element={[<Navbar key='/registration1'/>,<Registration key='/registration2'/>]}/>
          <Route key='/userPosts' exact path="/userPosts/:id" element={[<Navbar key={uuidv4()}/>,<PostsByUserId key={uuidv4()}/>]}/>
          <Route key='/home' exact path="/home" element={[<UserNavbar key={uuidv4()}/>, <UserHome key={uuidv4()}/>]}/>
          <Route key='/publicProfiles' exact path="/publicProfiles" element={[<UserNavbar key={uuidv4()}/>, <PublicProfileSearch displayFollowButtons={true} key={uuidv4()}/>]}/>
          <Route key='/allProfiles' exact path="/allProfiles" element={[<UserNavbar key={uuidv4()}/>, <AllProfiles displayFollowButtons={true} key={uuidv4()}/>]}/>
          <Route key='/jobOffers' exact path='/jobOffers' element={[<UserNavbar key={uuidv4()}/>,<AllJobOffers key={uuidv4()}/>]}/>
          <Route key='/profile' exact path="/profile/:id" element={[<UserNavbar key={uuidv4()}/>, <UserProfile key={uuidv4()}/>]}/>
          <Route key='/publicProfile' exact path="/publicProfile/:id" element={[<Navbar key={uuidv4()}/>, <UserProfile key={uuidv4()}/>]}/>
          <Route key='/users' exact path="/users" element={[<Navbar key={uuidv4()}/>,<Users key={uuidv4()}/>]}/>
          <Route key='/editProfile' exact path="/editProfile" element={[<UserNavbar key={uuidv4()}/>,<EditProfile key={uuidv4()}/>]}/>
          <Route key='/editWorkExperience' exact path="/editWorkExperience" element={[<UserNavbar key={uuidv4()}/>,<EditWorkExperience key={uuidv4()}/>]}/>
          <Route key='/editEducation' exact path="/editEducation" element={[<UserNavbar key={uuidv4()}/>,<EditEducation key={uuidv4()}/>]}/>
          <Route key='/editSkillsAndInterests' exact path="/editSkillsAndInterests" element={[<UserNavbar key='/editSkillsAndInterests1'/>,<EditSkillsAndInterests key='/editSkillsAndInterests2'/>]}/>
          <Route key='/personalProfile' exact path='/personalProfile' element={[<UserNavbar key='/personalProfile1'/>,<PersonalProfile key='/personalProfile2'/>]}/>
          <Route key='/followRequests' exact path='/followRequests' element={[<UserNavbar key='/followRequests1'/>,<FollowRequestsList key='/followRequests2'/>]}/>   
          <Route key='/createPost' exact path='/createPost' element={[<UserNavbar key='/createPost1'/>,<CreatePost key='/createPost2'/>]}/>      
          <Route key='/createJobOffer' exact path='/createJobOffer' element={[<UserNavbar key='/createJobOffer1'/>,<CreateJobOffer key='/createJobOffer2'/>]}/>    
          <Route key='/chat' exact path='/chat/:id' element={[<UserNavbar key='/chat1'/>,<Chat key='/chat2'/>]}/> 
          </Routes>
        </Router>
    </div>
  );
}

export default App;
