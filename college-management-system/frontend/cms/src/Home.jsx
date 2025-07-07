import { useState } from 'react';
import AboutCollege from './components/AboutCollege';
import Carousel from './components/Carousel';
import Footer from './components/Footer';
import LoginRegisterModal from './components/LoginRegisterModal';
import Navbar from './components/Navbar';
import NoticeSidebar from './components/NoticeSidebar';
import PlacementAchievements from './components/PlacementAchievements';
import VirtualTour from './components/VirtualTour';

function Home() {     
 // const [count, setCount] = useState(0);

  const [showModal,setShowModal]=useState(false);

  return (
    <>
      <Navbar />
      <Carousel/>

      <button onClick={() => setShowModal(true)}>Login/Register</button>
      {showModal && <LoginRegisterModal onClose={() => setShowModal(false)} />}

      <NoticeSidebar/>
      <AboutCollege/>
      <PlacementAchievements/>
      <VirtualTour/>
      <Footer/>
    </>
  )
}

export default Home
