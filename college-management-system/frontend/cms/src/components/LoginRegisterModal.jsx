// src/components/LoginRegisterModal.jsx
import { useState } from 'react';
import '../styles/loginRegisterModal.css';

const LoginRegisterModal = ({ onClose }) => {
  const [isLogin, setIsLogin] = useState(true);
  const [role, setRole] = useState('Student');
  const [course, setCourse] = useState('');
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    name: '',
  });

  const courses = {
    Student: ['B.Tech', 'M.Tech', 'MBA'],
    Faculty: ['CSE', 'ECE', 'Mechanical', 'Management'],
    Admin: ['--'], // Placeholder
  };

  const toggleMode = () => {
    setIsLogin(!isLogin);
    setFormData({ email: '', password: '', name: '' });
    setRole('Student');
    setCourse('');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const endpoint = isLogin ? '/api/login' : '/api/register-temp';

    const payload = {
      ...formData,
      role,
      course,
    };

    try {
      const res = await fetch(endpoint, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });

      const data = await res.json();
      alert(data.message);
    } catch (err) {
      alert('Something went wrong. Please try again.');
    }
  };

  return (
    <div className="modal-backdrop">
      <div className="auth-modal">
        <button className="close-btn" onClick={onClose}>×</button>
        <h2>{isLogin ? 'Login' : 'Register'}</h2>

        <form onSubmit={handleSubmit}>
          {!isLogin && (
            <input
              type="text"
              placeholder="Full Name"
              value={formData.name}
              required
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
            />
          )}

          <input
            type="email"
            placeholder="Email"
            value={formData.email}
            required
            onChange={(e) => setFormData({ ...formData, email: e.target.value })}
          />

          <input
            type="password"
            placeholder="Password"
            value={formData.password}
            required
            onChange={(e) => setFormData({ ...formData, password: e.target.value })}
          />

          <select value={role} onChange={(e) => setRole(e.target.value)}>
            <option>Student</option>
            <option>Faculty</option>
            <option>Admin</option>
          </select>

          <select value={course} onChange={(e) => setCourse(e.target.value)} required>
            <option value="">-- Select Course --</option>
            {courses[role].map((c, i) => (
              <option key={i} value={c}>{c}</option>
            ))}
          </select>

          <button type="submit">{isLogin ? 'Login' : 'Register'}</button>
        </form>

        <p className="toggle-auth">
          {isLogin ? 'Don’t have an account?' : 'Already registered?'}{' '}
          <span onClick={toggleMode}>
            {isLogin ? 'Register Here' : 'Login Here'}
          </span>
        </p>
      </div>
    </div>
  );
};

export default LoginRegisterModal;
