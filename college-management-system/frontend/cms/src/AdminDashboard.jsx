import { useEffect, useState } from 'react';
import './admindashboard.css';
import Footer from './components/Footer';
import Navbar from './components/Navbar';

const AdminPanel = () => {
  const [users, setUsers] = useState([]);

  const fetchUsers = () =>
    fetch('/api/admin/temp-users').then(res => res.json()).then(setUsers);

  const handleApprove = (user) =>
    fetch('/api/admin/approve', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(user),
    }).then(fetchUsers);

  const handleReject = (email) =>
    fetch('/api/admin/reject', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    }).then(fetchUsers);

  useEffect(() => {
    fetchUsers();
  }, []);

  return (
    <div> 
      <Navbar/>
      
    <div className="admin-panel">
      <h2>Pending Registrations</h2>
      <table>
        <thead><tr><th>Name</th><th>Email</th><th>Role</th><th>Course</th><th>Actions</th></tr></thead>
        <tbody>
          {users.map((u, i) => (
            <tr key={i}>
              <td>{u.name}</td>
              <td>{u.email}</td>
              <td>{u.role}</td>
              <td>{u.course}</td>
              <td>
                <button onClick={() => handleApprove(u)}>Approve</button>
                <button onClick={() => handleReject(u.email)}>Reject</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>

     <Footer/>
    </div>
  );
};

export default AdminPanel;
