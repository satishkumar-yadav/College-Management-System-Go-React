import { useEffect, useState } from 'react';
import Footer from './components/Footer';
import Navbar from './components/Navbar';
import './dashboard.css';

const Dashboard = () => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // Fetch user from token
    const token = localStorage.getItem('token');
    fetch('/api/dashboard', {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then(res => res.json())
      .then(data => setUser(data));
  }, []);

  if (!user) return <p>Loading...</p>;

  return (
    <div> 
      <Navbar/>
      
    <div className="dashboard">
      <h2>Welcome, {user.name}</h2>
      <p>Role: {user.role}</p>
      <p>Course: {user.course}</p>
      {/* Add more features: assignments, attendance, results etc. */}
    </div>

      <Footer/>
    </div>
  );
};

export default Dashboard;


/*

import React, { useEffect, useState } from 'react';
import API from '../api/api';

export default function Dashboard() {
  const [info, setInfo] = useState(null);

  useEffect(() => {
    API.get('/dashboard')
      .then(res => setInfo(res.data))
      .catch(() => alert('Unauthorized or expired token!'));
  }, []);

  return info ? (
    <div>
      <h2>Welcome to Dashboard</h2>
      <p><b>Email:</b> {info.email}</p>
      <p><b>Role:</b> {info.role}</p>
      <p><b>Course:</b> {info.course}</p>
    </div>
  ) : (
    <p>Loading...</p>
  );
}


*/