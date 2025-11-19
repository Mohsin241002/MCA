import { useParams } from 'react-router-dom'
import { Users, Bed, Bath, MapPin, Wifi, Car, Utensils, Waves, Calendar, Phone, Mail } from 'lucide-react'

const PropertyDetail = () => {
  const { id } = useParams()

  // Mock property data - in a real app, this would come from an API
  const property = {
    id: parseInt(id || '1'),
    name: "Thicket Priory",
    location: "North Yorkshire",
    guests: 70,
    bedrooms: 35,
    bathrooms: 35,
    weekendPrice: 22000,
    weekPrice: 48000,
    images: [
      "https://images.unsplash.com/photo-1564013799919-ab600027ffc6?w=1200&h=800&fit=crop",
      "https://images.unsplash.com/photo-1600585154340-be6161a56a0c?w=1200&h=800&fit=crop",
      "https://images.unsplash.com/photo-1600047509807-ba8f99d2cdde?w=1200&h=800&fit=crop",
      "https://images.unsplash.com/photo-1512917774080-9991f1c4c750?w=1200&h=800&fit=crop"
    ],
    description: "Thicket Priory is nestled in the North Yorkshire countryside, however only 20 minutes drive from the center of York! The perfect celebration get-away! This magnificent Georgian manor house offers unparalleled luxury and space for large groups seeking an unforgettable experience.",
    features: [
      "Hot Tub",
      "Indoor Pool", 
      "Tennis Court",
      "Wedding Licence",
      "Pet Friendly",
      "Parking for 30 cars",
      "Professional Kitchen",
      "Grand Ballroom",
      "Beautiful Gardens",
      "Historic Architecture"
    ],
    amenities: [
      { icon: Wifi, label: "High-speed WiFi" },
      { icon: Car, label: "Parking Available" },
      { icon: Utensils, label: "Full Kitchen" },
      { icon: Waves, label: "Swimming Pool" }
    ]
  }

  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('en-GB', {
      style: 'currency',
      currency: 'GBP',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(price)
  }

  return (
    <div className="property-detail">
      <div className="container">
        {/* Property Header */}
        <div className="property-header">
          <div className="property-title">
            <h1>{property.name}</h1>
            <div className="property-location-detail">
              <MapPin size={20} />
              <span>{property.location}</span>
            </div>
          </div>
          
          <div className="property-stats-detail">
            <div className="stat-large">
              <Users size={24} />
              <div>
                <span className="stat-number">{property.guests}</span>
                <span className="stat-label">Guests</span>
              </div>
            </div>
            <div className="stat-large">
              <Bed size={24} />
              <div>
                <span className="stat-number">{property.bedrooms}</span>
                <span className="stat-label">Bedrooms</span>
              </div>
            </div>
            <div className="stat-large">
              <Bath size={24} />
              <div>
                <span className="stat-number">{property.bathrooms}</span>
                <span className="stat-label">Bathrooms</span>
              </div>
            </div>
          </div>
        </div>

        {/* Image Gallery */}
        <div className="image-gallery">
          <div className="main-image">
            <img src={property.images[0]} alt={property.name} />
          </div>
          <div className="thumbnail-grid">
            {property.images.slice(1).map((image, index) => (
              <div key={index} className="thumbnail">
                <img src={image} alt={`${property.name} ${index + 2}`} />
              </div>
            ))}
          </div>
        </div>

        <div className="property-content-grid">
          {/* Property Description */}
          <div className="property-main-content">
            <section className="property-description">
              <h2>About this property</h2>
              <p>{property.description}</p>
            </section>

            {/* Features */}
            <section className="property-features">
              <h2>Features & Amenities</h2>
              <div className="features-grid">
                {property.features.map((feature, index) => (
                  <div key={index} className="feature-item">
                    <span>{feature}</span>
                  </div>
                ))}
              </div>
            </section>

            {/* Amenities */}
            <section className="property-amenities">
              <h2>What this place offers</h2>
              <div className="amenities-grid">
                {property.amenities.map((amenity, index) => (
                  <div key={index} className="amenity-item">
                    <amenity.icon size={20} />
                    <span>{amenity.label}</span>
                  </div>
                ))}
              </div>
            </section>
          </div>

          {/* Booking Sidebar */}
          <div className="booking-sidebar">
            <div className="booking-card">
              <div className="pricing-info">
                <div className="price-item">
                  <span className="price-label">Weekends From</span>
                  <span className="price-value">{formatPrice(property.weekendPrice)}</span>
                </div>
                <div className="price-item">
                  <span className="price-label">Weeks From</span>
                  <span className="price-value">{formatPrice(property.weekPrice)}</span>
                </div>
              </div>

              <div className="booking-form">
                <h3>Check Availability</h3>
                <form>
                  <div className="form-group">
                    <label>Check-in</label>
                    <input type="date" />
                  </div>
                  <div className="form-group">
                    <label>Check-out</label>
                    <input type="date" />
                  </div>
                  <div className="form-group">
                    <label>Guests</label>
                    <select>
                      <option>Select number of guests</option>
                      <option>1-10 guests</option>
                      <option>11-20 guests</option>
                      <option>21-30 guests</option>
                      <option>31-50 guests</option>
                      <option>50+ guests</option>
                    </select>
                  </div>
                  <button type="submit" className="enquiry-btn">
                    <Calendar size={20} />
                    Send Enquiry
                  </button>
                </form>
              </div>

              <div className="contact-info">
                <h4>Contact Owner</h4>
                <div className="contact-methods">
                  <a href="tel:+441234567890" className="contact-method">
                    <Phone size={16} />
                    <span>Call Now</span>
                  </a>
                  <a href="mailto:owner@example.com" className="contact-method">
                    <Mail size={16} />
                    <span>Send Email</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default PropertyDetail
