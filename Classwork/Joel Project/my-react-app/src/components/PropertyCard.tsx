import { Link } from 'react-router-dom'
import { Users, Bed, Bath, MapPin } from 'lucide-react'

interface Property {
  id: number
  name: string
  location: string
  guests: number
  bedrooms: number
  bathrooms: number
  weekendPrice: number
  weekPrice: number
  image: string
  description: string
}

interface PropertyCardProps {
  property: Property
}

const PropertyCard = ({ property }: PropertyCardProps) => {
  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('en-GB', {
      style: 'currency',
      currency: 'GBP',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(price)
  }

  return (
    <div className="property-card">
      <Link to={`/property/${property.id}`} className="property-link">
        <div className="property-image">
          <img src={property.image} alt={property.name} />
          <div className="property-location">
            <MapPin size={14} />
            <span>{property.location}</span>
          </div>
        </div>
        
        <div className="property-content">
          <h3 className="property-name">{property.name}</h3>
          
          <div className="property-stats">
            <div className="stat">
              <Users size={16} />
              <span>{property.guests} Guests</span>
            </div>
            <div className="stat">
              <Bed size={16} />
              <span>{property.bedrooms} Bedrooms</span>
            </div>
            <div className="stat">
              <Bath size={16} />
              <span>{property.bathrooms} Bathrooms</span>
            </div>
          </div>
          
          <p className="property-description">{property.description}</p>
          
          <div className="property-pricing">
            <div className="price-item">
              <span className="price-label">Weekends From</span>
              <span className="price-value">{formatPrice(property.weekendPrice)}</span>
            </div>
            <div className="price-item">
              <span className="price-label">Weeks From</span>
              <span className="price-value">{formatPrice(property.weekPrice)}</span>
            </div>
          </div>
        </div>
      </Link>
    </div>
  )
}

export default PropertyCard
