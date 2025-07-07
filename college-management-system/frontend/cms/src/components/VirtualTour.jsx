// src/components/VirtualTour.jsx
import '../styles/virtualTour.css';

const VirtualTour = () => (
  <section className="virtual-tour">
    <h2>🎥 College Virtual Tour</h2>
    <div className="video-container">
      <iframe
        src="https://www.youtube.com/embed/VIDEO_ID"
        title="Campus Tour"
        allowFullScreen
      ></iframe>
    </div>
  </section>
);

export default VirtualTour;
