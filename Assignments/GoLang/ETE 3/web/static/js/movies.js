document.addEventListener('DOMContentLoaded', function() {
    // Load movies on page load
    loadMovies();
    
    // Search form event handler
    const searchForm = document.getElementById('searchForm');
    searchForm.addEventListener('submit', function(e) {
        e.preventDefault();
        loadMovies();
    });
});

// Load movies with optional filters
async function loadMovies() {
    const moviesContainer = document.getElementById('moviesList');
    const titleQuery = document.getElementById('searchTitle').value;
    const genreFilter = document.getElementById('filterGenre').value;
    
    // Show loading spinner
    moviesContainer.innerHTML = `
        <div class="col-12 text-center py-5">
            <div class="spinner-border" role="status">
                <span class="visually-hidden">Loading...</span>
            </div>
        </div>
    `;
    
    try {
        // Build query parameters
        const params = new URLSearchParams();
        if (titleQuery) params.append('title', titleQuery);
        if (genreFilter) params.append('genre', genreFilter);
        
        const response = await fetch(`/api/movies?${params.toString()}`);
        
        if (!response.ok) {
            throw new Error('Failed to fetch movies');
        }
        
        const movies = await response.json();
        
        if (movies.length === 0) {
            moviesContainer.innerHTML = `
                <div class="col-12 text-center py-5">
                    <p>No movies found. Try different search criteria.</p>
                </div>
            `;
            return;
        }
        
        // Render movies
        renderMovies(movies);
        
    } catch (error) {
        console.error('Error loading movies:', error);
        moviesContainer.innerHTML = `
            <div class="col-12 text-center py-5">
                <p class="text-danger">Error loading movies. Please try again later.</p>
            </div>
        `;
    }
}

// Render movie cards
function renderMovies(movies) {
    const moviesContainer = document.getElementById('moviesList');
    
    moviesContainer.innerHTML = movies.map(movie => `
        <div class="col-md-4 col-lg-3 mb-4">
            <div class="card movie-card h-100">
                <img src="${movie.poster_url || '/static/img/no-poster.jpg'}" class="card-img-top movie-poster" alt="${movie.title}">
                <div class="card-body d-flex flex-column">
                    <h5 class="card-title">${movie.title}</h5>
                    <p class="card-text text-muted">${movie.genre} | ${movie.duration} min</p>
                    <p class="card-text small text-truncate">${movie.description}</p>
                    <a href="/movies/${movie.id}" class="btn btn-primary mt-auto">View Details</a>
                </div>
            </div>
        </div>
    `).join('');
} 