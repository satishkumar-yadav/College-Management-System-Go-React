import { useEffect, useState } from 'react';
import API from '../api/api';

export default function AdminPanel() {
  const [pendingUsers, setPendingUsers] = useState([]);

  const fetchPending = async () => {
    try {
      const res = await API.get('/admin/temp-users');
      setPendingUsers(res.data);
    } catch {
      alert("Unauthorized or failed to fetch users.");
    }
  };

  const approveUser = async (user) => {
    try {
      await API.post('/admin/approve', user);
      alert(`${user.name} approved.`);
      fetchPending(); // refresh list
    } catch {
      alert("Approval failed.");
    }
  };

  const rejectUser = async (email) => {
    try {
      await API.post('/admin/reject', { email });
      alert("User rejected.");
      fetchPending();
    } catch {
      alert("Rejection failed.");
    }
  };

  useEffect(() => {
    fetchPending();
  }, []);

  return (
    <div className="admin-panel">
      <h2>Admin Panel – Pending Approvals</h2>
      {pendingUsers.length === 0 ? (
        <p>No pending registrations.</p>
      ) : (
        <table border="1" cellPadding="10">
          <thead>
            <tr>
              <th>Name</th><th>Email</th><th>Role</th><th>Course</th><th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {pendingUsers.map((user, i) => (
              <tr key={i}>
                <td>{user.name}</td>
                <td>{user.email}</td>
                <td>{user.role}</td>
                <td>{user.course}</td>
                <td>
                  <button onClick={() => approveUser(user)}>Approve</button>
                  <button onClick={() => rejectUser(user.email)}>Reject</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
