// src/components/NoticeSidebar.jsx
import { useEffect, useState } from 'react';
import '../styles/noticeSidebar.css';

const NoticeSidebar = () => {
  const [notices, setNotices] = useState([]);

  useEffect(() => {
    const fetchNotices = async () => {
      try {
        const res = await fetch('/api/notices');
        const data = await res.json();
        setNotices(data);
      } catch {
        setNotices([{ title: "Unable to load notices" }]);
      }
    };
    fetchNotices();
  }, []);

  return (
    <div className="notice-sidebar">
      <h3>📢 Notices</h3>
      <ul>
        {notices.map((n, i) => (
          <li key={i}>{n.title}</li>
        ))}
      </ul>
    </div>
  );
};

export default NoticeSidebar;
