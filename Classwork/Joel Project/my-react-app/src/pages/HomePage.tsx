import { useState } from 'react'

const HomePage = () => {
  const [searchType, setSearchType] = useState('stays')
  const [location, setLocation] = useState('')
  const [checkIn, setCheckIn] = useState('')
  const [checkOut, setCheckOut] = useState('')
  const [guests, setGuests] = useState('')

  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="relative h-screen flex items-center justify-center">
        {/* Background Image */}
        <div 
          className="absolute inset-0 bg-cover bg-center bg-no-repeat"
          style={{
            backgroundImage: "url('https://images.unsplash.com/photo-1600596542815-ffad4c1539a9?ixlib=rb-4.0.3&auto=format&fit=crop&w=2000&q=80')"
          }}
        />
        
        {/* Content */}
        <div className="relative z-10 text-center max-w-4xl mx-auto px-4">
          <h1 className="text-5xl md:text-6xl font-bold text-gray-900 mb-4 leading-tight">
            Find dream country houses
            <br />
            to rent
          </h1>
          <p className="text-lg md:text-xl text-gray-700 mb-12">
            Your insider guide to hiring Britain's best big houses.
          </p>
          
          {/* Search Form */}
          <div className="bg-white rounded-2xl p-6 shadow-lg max-w-4xl mx-auto">
            {/* Toggle Buttons */}
            <div className="flex items-center mb-6">
              <div className="flex bg-gray-100 rounded-full p-1">
                <button
                  onClick={() => setSearchType('stays')}
                  className={`px-4 py-2 rounded-full font-medium transition-colors ${
                    searchType === 'stays'
                      ? 'bg-gray-900 text-white'
                      : 'text-gray-600 hover:text-gray-900'
                  }`}
                >
                  Stays
                </button>
                <button
                  onClick={() => setSearchType('weddings')}
                  className={`px-4 py-2 rounded-full font-medium transition-colors ${
                    searchType === 'weddings'
                      ? 'bg-gray-900 text-white'
                      : 'text-gray-600 hover:text-gray-900'
                  }`}
                >
                  Weddings
                </button>
              </div>
            </div>

            {/* Search Fields */}
            <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
              <div className="relative">
                <label className="block text-sm font-medium text-gray-700 mb-2">Where</label>
                <input
                  type="text"
                  placeholder="Select region(s)"
                  value={location}
                  onChange={(e) => setLocation(e.target.value)}
                  className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent text-gray-900 placeholder-gray-500"
                />
              </div>
              
              <div className="relative">
                <label className="block text-sm font-medium text-gray-700 mb-2">Check in</label>
                <input
                  type="text"
                  placeholder="Select date"
                  value={checkIn}
                  onChange={(e) => setCheckIn(e.target.value)}
                  className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent text-gray-900 placeholder-gray-500"
                />
              </div>
              
              <div className="relative">
                <label className="block text-sm font-medium text-gray-700 mb-2">Check out</label>
                <input
                  type="text"
                  placeholder="Select date"
                  value={checkOut}
                  onChange={(e) => setCheckOut(e.target.value)}
                  className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent text-gray-900 placeholder-gray-500"
                />
              </div>
              
              <div className="relative">
                <label className="block text-sm font-medium text-gray-700 mb-2">Who</label>
                <input
                  type="text"
                  placeholder="Add guests"
                  value={guests}
                  onChange={(e) => setGuests(e.target.value)}
                  className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent text-gray-900 placeholder-gray-500"
                />
              </div>
            </div>
            
            <div className="mt-6">
              <button className="w-full bg-orange-500 text-white py-4 px-8 rounded-lg font-semibold text-lg hover:bg-orange-600 transition-colors">
                Search
              </button>
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}

export default HomePage
