// Function to search games
function searchGames() {
    // Get the input value
    let input = document.getElementById('searchInput').value.toLowerCase();
    
    // Get the game gallery and all game thumbnails
    let gamesGallery = document.getElementById('gamesGallery');
    let games = gamesGallery.getElementsByClassName('game-thumbnail');
    
    // Loop through all game thumbnails and hide those that don't match the search query
    for (let i = 0; i < games.length; i++) {
      let gameTitle = games[i].getElementsByTagName('p')[0].textContent.toLowerCase();
      
      if (gameTitle.indexOf(input) > -1) {
        games[i].style.display = '';
      } else {
        games[i].style.display = 'none';
      }
    }
  }
  
  // Function to make each game clickable
  function setupGameClickListeners() {
    // Get the game thumbnails
    let gamesGallery = document.getElementById('gamesGallery');
    let games = gamesGallery.getElementsByClassName('game-thumbnail');
  
    // Add click event listener to each game thumbnail
    for (let i = 0; i < games.length; i++) {
      games[i].addEventListener('click', function() {
        // Navigate to the details page (add a unique identifier if needed)
        window.location.href = 'details.html';
      });
    }
  }
  
  // Initialize the click listeners
  setupGameClickListeners();
  