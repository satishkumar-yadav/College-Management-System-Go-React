import logo from '../assets/logo.png';
import '../styles/footer.css';

function Footer(){

    return (
        <footer style={{textAlign:'center', padding:'20px', backgroundColor: '#da222', color: 'black'}}>
             <div className="bottom-bar">
                    <img src={logo} alt="College Logo" className="college-logo" />
                   <p className="college-name">&copy; {new Date().getFullYear()} ABC Institute of Technology, Patna</p>
                  </div>
           
        </footer>   
    );
}

export default Footer;