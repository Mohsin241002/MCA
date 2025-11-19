import { Link } from 'react-router-dom'
import { Home, Mail, Phone, MapPin } from 'lucide-react'

const Footer = () => {
  return (
    <footer className="footer">
      <div className="container">
        <div className="footer-content">
          <div className="footer-section">
            <Link to="/" className="footer-logo">
              <Home className="logo-icon" />
              <span>Landed Houses</span>
            </Link>
            <p>Your insider guide to hiring Britain's best big houses.</p>
            <div className="contact-info">
              <div className="contact-item">
                <Mail size={16} />
                <span>info@landedhouses.com</span>
              </div>
              <div className="contact-item">
                <Phone size={16} />
                <span>+44 20 7123 4567</span>
              </div>
              <div className="contact-item">
                <MapPin size={16} />
                <span>London, United Kingdom</span>
              </div>
            </div>
          </div>

          <div className="footer-section">
            <h4>House Types</h4>
            <ul>
              <li><Link to="/search?type=holiday">Large Holiday Homes</Link></li>
              <li><Link to="/search?type=party">Party Houses</Link></li>
              <li><Link to="/search?type=wedding">Wedding Houses</Link></li>
              <li><Link to="/search?type=mansion">Mansions</Link></li>
            </ul>
          </div>

          <div className="footer-section">
            <h4>Popular Locations</h4>
            <ul>
              <li><Link to="/search?location=devon-cornwall">Devon & Cornwall</Link></li>
              <li><Link to="/search?location=london">London & Home Counties</Link></li>
              <li><Link to="/search?location=scotland">Scotland</Link></li>
              <li><Link to="/search?location=wales">Wales</Link></li>
            </ul>
          </div>

          <div className="footer-section">
            <h4>Company</h4>
            <ul>
              <li><Link to="/about">About Us</Link></li>
              <li><Link to="/list-property">List Your Property</Link></li>
              <li><Link to="/contact">Contact</Link></li>
              <li><Link to="/help">Help & Support</Link></li>
            </ul>
          </div>
        </div>

        <div className="footer-bottom">
          <p>&copy; 2024 Landed Houses. All rights reserved.</p>
          <div className="footer-links">
            <Link to="/privacy">Privacy Policy</Link>
            <Link to="/terms">Terms of Service</Link>
            <Link to="/cookies">Cookie Policy</Link>
          </div>
        </div>
      </div>
    </footer>
  )
}

export default Footer
