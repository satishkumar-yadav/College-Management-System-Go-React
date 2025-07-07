
import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import AdminPanel from './AdminDashboard';
import Dashboard from './Dashboard';
import Home from './Home';

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/admin" element={<AdminPanel />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </Router>
  );
}