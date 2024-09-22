import React, { useState } from 'react';

// Sample data for gaming cards
const gamesData = [
  {
    id: 1,
    title: 'Cyberpunk 2077',
    features: ['Open World', 'RPG', 'First-Person Shooter'],
    image: 'cyber.jpg', // Replace with actual image URL
  },
  {
    id: 2,
    title: 'The Witcher 3',
    features: ['Fantasy', 'Story Driven', 'Open World'],
    image: 'witcher.jpg', // Replace with actual image URL
  },
  {
    id: 3,
    title: 'Elden Ring',
    features: ['Dark Fantasy', 'Exploration', 'Challenging Combat'],
    image: 'elden.jpg', // Replace with actual image URL
  },
];

// GameCard component to display individual game information
const GameCard = ({ game }) => {
  const [liked, setLiked] = useState(false); // Track the like status of each game

  const toggleLike = () => {
    setLiked(!liked); // Toggle the like state when the button is clicked
  };

  return (
    <div className="max-w-xs rounded overflow-hidden shadow-lg my-4 bg-gray-800 text-white p-4">
      {/* Updated image styling to prevent cropping */}
      <img className="w-full h-48 object-contain" src={game.image} alt={game.title} />
      <div className="px-6 py-4">
        <div className="font-bold text-xl mb-2">{game.title}</div>
        <ul>
          {game.features.map((feature, index) => (
            <li key={index}>- {feature}</li>
          ))}
        </ul>
      </div>
      <div className="px-6 py-4 flex justify-end">
        {/* Like button with dynamic scaling and color */}
        <button
          className={`text-2xl ${liked ? 'text-red-500' : 'text-gray-400'} transition-transform transform ${
            liked ? 'scale-125' : 'scale-100'
          }`}
          onClick={toggleLike}
        >
          ♥
        </button>
      </div>
    </div>
  );
};

// Main App component
const App = () => {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-4xl font-bold text-center text-gray-100 mb-8">Gaming Cards</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
        {/* Render each game card */}
        {gamesData.map((game) => (
          <GameCard key={game.id} game={game} />
        ))}
      </div>
    </div>
  );
};

export default App;
