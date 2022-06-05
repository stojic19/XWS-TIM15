import { useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Navbar from './components/Navbar';

function App() {

  useEffect(() => {
    document.title = "Agent"
  }, [])

  return (
    <div>
      <Router>
        <Routes>
        <Route key='/' exact path="/" element={[<Navbar key='/' />]}/>
        </Routes>
      </Router>
    </div>
  );
}

export default App;
