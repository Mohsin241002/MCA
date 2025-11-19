import { useState } from 'react'
import { Menu, User } from 'lucide-react'

const Header = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false)

  return (
    <header className="absolute top-0 left-0 right-0 z-50 bg-transparent">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Left side - Hamburger menu */}
          <div className="flex items-center">
            <button
              className="p-2 text-gray-800 hover:text-gray-600 transition-colors"
              onClick={() => setIsMenuOpen(!isMenuOpen)}
            >
              <Menu className="h-6 w-6" />
            </button>
            <span className="ml-2 text-gray-800 font-medium">Search houses</span>
          </div>

          {/* Center - Logo */}
          <div className="flex items-center">
            <div className="flex items-center">
              <svg className="h-8 w-8 text-orange-500 mr-2" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
              </svg>
              <span className="text-xl font-semibold text-gray-800">landed houses</span>
            </div>
          </div>

          {/* Right side - User actions */}
          <div className="flex items-center space-x-4">
            <button className="p-2 text-gray-800 hover:text-gray-600 transition-colors">
              <User className="h-5 w-5" />
            </button>
            <span className="text-gray-800 font-medium">Log in</span>
            <button className="bg-gray-900 text-white px-4 py-2 rounded-md hover:bg-gray-800 transition-colors font-medium">
              List your house
            </button>
          </div>
        </div>

        {/* Mobile menu */}
        {isMenuOpen && (
          <div className="md:hidden">
            <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3 bg-white/90 backdrop-blur-sm rounded-md mt-2">
              <a href="#" className="block px-3 py-2 text-base font-medium text-gray-700 hover:text-orange-500">
                Search houses
              </a>
              <a href="#" className="block px-3 py-2 text-base font-medium text-gray-700 hover:text-orange-500">
                About
              </a>
              <a href="#" className="block px-3 py-2 text-base font-medium text-gray-700 hover:text-orange-500">
                Help
              </a>
            </div>
          </div>
        )}
      </div>
    </header>
  )
}

export default Header
