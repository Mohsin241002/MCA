document.addEventListener("DOMContentLoaded", () => {
    // Smooth scrolling for internal links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
      anchor.addEventListener("click", function(e) {
        e.preventDefault();
        document.querySelector(this.getAttribute("href")).scrollIntoView({
          behavior: "smooth"
        });
      });
    });
  
    // Search functionality
    const searchInput = document.getElementById('search');
    if (searchInput) {
      searchInput.addEventListener('input', () => {
        const filter = searchInput.value.toLowerCase();
        document.querySelectorAll('.game-thumbnail').forEach(game => {
          const text = game.textContent.toLowerCase();
          game.style.display = text.includes(filter) ? "" : "none";
        });
      });
    }
  
    // Fade-in animations
    const faders = document.querySelectorAll('.fade-in');
    const appearOptions = {
      threshold: 0,
      rootMargin: "0px 0px -100px 0px"
    };
  
    const appearOnScroll = new IntersectionObserver((entries, appearOnScroll) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add('appear');
          appearOnScroll.unobserve(entry.target);
        }
      });
    }, appearOptions);
  
    faders.forEach(fader => {
      appearOnScroll.observe(fader);
    });
  });
  

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
