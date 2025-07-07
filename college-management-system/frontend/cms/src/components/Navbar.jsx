// src/components/Navbar.jsx
import logo from '../assets/logo.png';
import '../styles/navbar.css';

const Navbar = () => {
/*
  useEffect(() => {
  const token = localStorage.getItem('token');
  if (!token) window.location.href = '/login';
}, []);
*/

  return (
    <header>
      {/* Top bar with logo and name */}
      <div className="top-bar">
        <img src={logo} alt="College Logo" className="college-logo" />
        <h1 className="college-name">ABC Institute of Technology</h1>
      </div>

      {/* Navigation bar */}
      <nav className="nav-bar">
        <ul className="nav-links">
          <li><a href="/">Home</a></li>
          <li><a href="/academics">Academics</a></li>
          <li><a href="/departments">Departments</a></li>
          <li><a href="/admissions">Admissions</a></li>
          <li><a href="/facilities">Facilities</a></li>
          <li><a href="/placement">Placement</a></li>
          <li><a href="/contact">Contact</a></li>
          <li> <button onClick={() => {
                       localStorage.removeItem('token');
                       window.location.href = '/';
                  }}>Logout</button>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Navbar;
