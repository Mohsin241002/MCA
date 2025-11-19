import { useState, useEffect } from 'react'
import { useSearchParams } from 'react-router-dom'
import { Search, SlidersHorizontal } from 'lucide-react'
import PropertyCard from '../components/PropertyCard'

const SearchResults = () => {
  const [searchParams] = useSearchParams()
  const [searchQuery, setSearchQuery] = useState(searchParams.get('q') || '')
  const [showFilters, setShowFilters] = useState(false)
  const [filters, setFilters] = useState({
    type: searchParams.get('type') || '',
    location: searchParams.get('location') || '',
    feature: searchParams.get('feature') || '',
    minGuests: '',
    maxGuests: '',
    minPrice: '',
    maxPrice: ''
  })

  // Mock search results - in a real app, this would come from an API
  const allProperties = [
    {
      id: 1,
      name: "Thicket Priory",
      location: "North Yorkshire",
      guests: 70,
      bedrooms: 35,
      bathrooms: 35,
      weekendPrice: 22000,
      weekPrice: 48000,
      image: "https://images.unsplash.com/photo-1564013799919-ab600027ffc6?w=800&h=600&fit=crop",
      description: "Nestled in the North Yorkshire countryside, however only 20 minutes drive from the center of York!"
    },
    {
      id: 2,
      name: "Balinakill Country House",
      location: "Argyll, Scotland",
      guests: 30,
      bedrooms: 13,
      bathrooms: 12,
      weekendPrice: 3450,
      weekPrice: 5460,
      image: "https://images.unsplash.com/photo-1505142468610-359e7d316be0?w=800&h=600&fit=crop",
      description: "A welcoming Victorian country house, on West Scotland's Kintyre Peninsula"
    },
    {
      id: 3,
      name: "Ringshall Grange",
      location: "Suffolk",
      guests: 23,
      bedrooms: 8,
      bathrooms: 3,
      weekendPrice: 3500,
      weekPrice: 4500,
      image: "https://images.unsplash.com/photo-1512917774080-9991f1c4c750?w=800&h=600&fit=crop",
      description: "A beautiful, private moated manor house, in the depths of the Suffolk countryside."
    },
    {
      id: 4,
      name: "River Wye Lodge",
      location: "Gloucestershire",
      guests: 26,
      bedrooms: 12,
      bathrooms: 11,
      weekendPrice: 3250,
      weekPrice: 4500,
      image: "https://images.unsplash.com/photo-1600596542815-ffad4c1539a9?w=800&h=600&fit=crop",
      description: "Set on the banks of the River Wye, River Wye Lodge is perfect for groups of up to 26 guests."
    },
    {
      id: 5,
      name: "Boxted Hall",
      location: "Suffolk",
      guests: 40,
      bedrooms: 18,
      bathrooms: 15,
      weekendPrice: 8000,
      weekPrice: 12000,
      image: "https://images.unsplash.com/photo-1600047509807-ba8f99d2cdde?w=800&h=600&fit=crop",
      description: "A magnificent Georgian manor house with stunning gardens and lake views."
    },
    {
      id: 6,
      name: "Cowdray House",
      location: "West Sussex",
      guests: 50,
      bedrooms: 22,
      bathrooms: 20,
      weekendPrice: 15000,
      weekPrice: 25000,
      image: "https://images.unsplash.com/photo-1600585154340-be6161a56a0c?w=800&h=600&fit=crop",
      description: "Historic Tudor mansion with polo grounds and championship golf course."
    }
  ]

  const [filteredProperties, setFilteredProperties] = useState(allProperties)

  useEffect(() => {
    let filtered = allProperties

    // Apply search query filter
    if (searchQuery) {
      filtered = filtered.filter(property =>
        property.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        property.location.toLowerCase().includes(searchQuery.toLowerCase()) ||
        property.description.toLowerCase().includes(searchQuery.toLowerCase())
      )
    }

    // Apply other filters
    if (filters.minGuests) {
      filtered = filtered.filter(property => property.guests >= parseInt(filters.minGuests))
    }
    if (filters.maxGuests) {
      filtered = filtered.filter(property => property.guests <= parseInt(filters.maxGuests))
    }
    if (filters.minPrice) {
      filtered = filtered.filter(property => property.weekendPrice >= parseInt(filters.minPrice))
    }
    if (filters.maxPrice) {
      filtered = filtered.filter(property => property.weekendPrice <= parseInt(filters.maxPrice))
    }

    setFilteredProperties(filtered)
  }, [searchQuery, filters])

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    // Search logic is handled by useEffect
  }

  const handleFilterChange = (key: string, value: string) => {
    setFilters(prev => ({ ...prev, [key]: value }))
  }

  const clearFilters = () => {
    setFilters({
      type: '',
      location: '',
      feature: '',
      minGuests: '',
      maxGuests: '',
      minPrice: '',
      maxPrice: ''
    })
  }

  return (
    <div className="search-results">
      <div className="container">
        {/* Search Header */}
        <div className="search-header">
          <form className="search-form-results" onSubmit={handleSearch}>
            <div className="search-input-group">
              <Search className="search-icon" />
              <input
                type="text"
                placeholder="Search by location, property name, or features..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
              />
              <button type="submit" className="search-btn">Search</button>
            </div>
          </form>

          <button 
            className="filters-toggle"
            onClick={() => setShowFilters(!showFilters)}
          >
            <SlidersHorizontal size={20} />
            Filters
          </button>
        </div>

        {/* Filters Panel */}
        {showFilters && (
          <div className="filters-panel">
            <div className="filters-grid">
              <div className="filter-group">
                <label>House Type</label>
                <select 
                  value={filters.type}
                  onChange={(e) => handleFilterChange('type', e.target.value)}
                >
                  <option value="">All Types</option>
                  <option value="holiday">Holiday Homes</option>
                  <option value="party">Party Houses</option>
                  <option value="wedding">Wedding Houses</option>
                  <option value="mansion">Mansions</option>
                </select>
              </div>

              <div className="filter-group">
                <label>Location</label>
                <select 
                  value={filters.location}
                  onChange={(e) => handleFilterChange('location', e.target.value)}
                >
                  <option value="">All Locations</option>
                  <option value="devon-cornwall">Devon & Cornwall</option>
                  <option value="london">London & Home Counties</option>
                  <option value="scotland">Scotland</option>
                  <option value="wales">Wales</option>
                  <option value="suffolk">Suffolk</option>
                  <option value="norfolk">Norfolk</option>
                </select>
              </div>

              <div className="filter-group">
                <label>Min Guests</label>
                <input
                  type="number"
                  placeholder="Min guests"
                  value={filters.minGuests}
                  onChange={(e) => handleFilterChange('minGuests', e.target.value)}
                />
              </div>

              <div className="filter-group">
                <label>Max Guests</label>
                <input
                  type="number"
                  placeholder="Max guests"
                  value={filters.maxGuests}
                  onChange={(e) => handleFilterChange('maxGuests', e.target.value)}
                />
              </div>

              <div className="filter-group">
                <label>Min Price (Weekend)</label>
                <input
                  type="number"
                  placeholder="Min price"
                  value={filters.minPrice}
                  onChange={(e) => handleFilterChange('minPrice', e.target.value)}
                />
              </div>

              <div className="filter-group">
                <label>Max Price (Weekend)</label>
                <input
                  type="number"
                  placeholder="Max price"
                  value={filters.maxPrice}
                  onChange={(e) => handleFilterChange('maxPrice', e.target.value)}
                />
              </div>
            </div>

            <div className="filters-actions">
              <button onClick={clearFilters} className="clear-filters-btn">
                Clear All Filters
              </button>
            </div>
          </div>
        )}

        {/* Results */}
        <div className="results-section">
          <div className="results-header">
            <h1>Search Results</h1>
            <p>{filteredProperties.length} properties found</p>
          </div>

          {filteredProperties.length > 0 ? (
            <div className="properties-grid">
              {filteredProperties.map(property => (
                <PropertyCard key={property.id} property={property} />
              ))}
            </div>
          ) : (
            <div className="no-results">
              <h3>No properties found</h3>
              <p>Try adjusting your search criteria or filters to find more properties.</p>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default SearchResults
