document.addEventListener('DOMContentLoaded', function() {
    // Check if user is logged in
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = '/login';
        return;
    }
    
    // Get show ID from URL
    const urlParams = new URLSearchParams(window.location.search);
    const showId = urlParams.get('id');
    
    if (!showId) {
        window.location.href = '/movies';
        return;
    }
    
    // Initialize booking
    loadShowDetails(showId);
    loadSeats(showId);
    
    // Start polling for seat updates
    startSeatPolling(showId);
    
    // Book button click handler
    document.getElementById('bookButton').addEventListener('click', function() {
        bookSeat(showId);
    });
});

// Selected seat state
let selectedSeatId = null;
let polling = null;

// Load show details
async function loadShowDetails(showId) {
    const showInfoContainer = document.getElementById('showInfo');
    
    try {
        const response = await fetchAuth(`/api/shows/${showId}`);
        
        if (!response.ok) {
            throw new Error('Failed to fetch show details');
        }
        
        const show = await response.json();
        
        // Get movie details
        const movieResponse = await fetch(`/api/movies/${show.movie_id}`);
        if (!movieResponse.ok) {
            throw new Error('Failed to fetch movie details');
        }
        
        const movie = await movieResponse.json();
        
        // Format dates
        const showDate = new Date(show.show_time);
        const endTime = new Date(show.end_time);
        
        showInfoContainer.innerHTML = `
            <div class="col-md-12">
                <div class="card mb-4">
                    <div class="card-body">
                        <div class="row">
                            <div class="col-md-3">
                                <img src="${movie.poster_url || '/static/img/no-poster.jpg'}" class="img-fluid rounded" alt="${movie.title}">
                            </div>
                            <div class="col-md-9">
                                <h2>${movie.title}</h2>
                                <p class="text-muted">${movie.genre} | ${movie.duration} min</p>
                                <hr>
                                <div class="row">
                                    <div class="col-md-6">
                                        <p><strong>Date:</strong> ${showDate.toLocaleDateString()}</p>
                                        <p><strong>Time:</strong> ${showDate.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})} - ${endTime.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})}</p>
                                    </div>
                                    <div class="col-md-6">
                                        <p><strong>Hall:</strong> ${show.hall_number}</p>
                                        <p><strong>Price:</strong> $${show.price.toFixed(2)}</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        `;
        
        // Update booking summary
        updateBookingSummary(show);
        
    } catch (error) {
        console.error('Error loading show details:', error);
        showInfoContainer.innerHTML = `
            <div class="col-12">
                <div class="alert alert-danger">
                    Error loading show details. Please try again later.
                </div>
            </div>
        `;
    }
}

// Load seats for a show
async function loadSeats(showId) {
    const seatMapContainer = document.getElementById('seatMap');
    
    try {
        const response = await fetchAuth(`/api/shows/${showId}/seats`);
        
        if (!response.ok) {
            throw new Error('Failed to fetch seats');
        }
        
        const seats = await response.json();
        
        // Organize seats into rows (A1, A2, ... B1, B2, ...)
        const seatsByRow = {};
        
        seats.forEach(seat => {
            const rowLetter = seat.seat_number.charAt(0);
            if (!seatsByRow[rowLetter]) {
                seatsByRow[rowLetter] = [];
            }
            seatsByRow[rowLetter].push(seat);
        });
        
        // Sort rows alphabetically
        const sortedRows = Object.keys(seatsByRow).sort();
        
        // Generate seat map HTML
        let seatMapHTML = '';
        
        sortedRows.forEach(rowLetter => {
            // Sort seats in this row by number
            const rowSeats = seatsByRow[rowLetter].sort((a, b) => {
                const numA = parseInt(a.seat_number.substring(1));
                const numB = parseInt(b.seat_number.substring(1));
                return numA - numB;
            });
            
            seatMapHTML += `
                <div class="seat-row">
                    <div class="row-label">${rowLetter}</div>
                    ${rowSeats.map(seat => `
                        <div class="seat ${seat.is_booked ? 'booked' : 'available'}" 
                             data-seat-id="${seat.id}" 
                             data-seat-number="${seat.seat_number}">
                            ${seat.seat_number.substring(1)}
                        </div>
                    `).join('')}
                </div>
            `;
        });
        
        seatMapContainer.innerHTML = seatMapHTML;
        
        // Add event listeners to seats
        document.querySelectorAll('.seat.available').forEach(seatElement => {
            seatElement.addEventListener('click', function() {
                selectSeat(this.dataset.seatId, this.dataset.seatNumber);
            });
        });
        
    } catch (error) {
        console.error('Error loading seats:', error);
        seatMapContainer.innerHTML = `
            <div class="alert alert-danger">
                Error loading seats. Please try again later.
            </div>
        `;
    }
}

// Select a seat
function selectSeat(seatId, seatNumber) {
    // Clear previous selection
    document.querySelectorAll('.seat.selected').forEach(seat => {
        seat.classList.remove('selected');
        seat.classList.add('available');
    });
    
    // Update selected seat
    selectedSeatId = seatId;
    
    // Update UI
    if (seatId) {
        const seatElement = document.querySelector(`.seat[data-seat-id="${seatId}"]`);
        seatElement.classList.remove('available');
        seatElement.classList.add('selected');
        
        // Enable book button
        document.getElementById('bookButton').disabled = false;
        
        // Update booking summary
        const summaryElement = document.getElementById('bookingSummary');
        summaryElement.innerHTML += `
            <p><strong>Selected Seat:</strong> ${seatNumber}</p>
        `;
    } else {
        // Disable book button
        document.getElementById('bookButton').disabled = true;
    }
}

// Update booking summary
function updateBookingSummary(show) {
    const summaryElement = document.getElementById('bookingSummary');
    
    const showDate = new Date(show.show_time);
    
    summaryElement.innerHTML = `
        <h5>Booking Details</h5>
        <p><strong>Show Date:</strong> ${showDate.toLocaleDateString()}</p>
        <p><strong>Show Time:</strong> ${showDate.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})}</p>
        <p><strong>Ticket Price:</strong> $${show.price.toFixed(2)}</p>
        <hr>
    `;
}

// Book a seat
async function bookSeat(showId) {
    if (!selectedSeatId) return;
    
    const bookButton = document.getElementById('bookButton');
    const errorElement = document.getElementById('bookingError');
    
    // Update UI
    bookButton.disabled = true;
    bookButton.innerHTML = '<span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span> Booking...';
    errorElement.classList.add('d-none');
    
    try {
        const response = await fetchAuth('/api/bookings', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                show_id: parseInt(showId),
                seat_id: parseInt(selectedSeatId)
            })
        });
        
        const data = await response.json();
        
        if (!response.ok) {
            // Show error message
            errorElement.textContent = data.error || 'Failed to book seat';
            errorElement.classList.remove('d-none');
            bookButton.disabled = false;
            bookButton.textContent = 'Book Now';
            return;
        }
        
        // Booking successful, redirect to bookings page
        window.location.href = '/bookings';
        
    } catch (error) {
        console.error('Booking error:', error);
        errorElement.textContent = 'An error occurred. Please try again.';
        errorElement.classList.remove('d-none');
        bookButton.disabled = false;
        bookButton.textContent = 'Book Now';
    }
}

// Start polling for seat updates
function startSeatPolling(showId) {
    // Stop any existing polling
    if (polling) clearInterval(polling);
    
    // Poll every 5 seconds
    polling = setInterval(() => {
        updateSeats(showId);
    }, 5000);
    
    // Stop polling when user leaves the page
    window.addEventListener('beforeunload', () => {
        if (polling) clearInterval(polling);
    });
}

// Update seats without reloading the entire seat map
async function updateSeats(showId) {
    try {
        const response = await fetch(`/api/shows/${showId}/seats`);
        
        if (!response.ok) {
            throw new Error('Failed to fetch seat updates');
        }
        
        const seats = await response.json();
        
        // Update seat status
        seats.forEach(seat => {
            const seatElement = document.querySelector(`.seat[data-seat-id="${seat.id}"]`);
            
            if (seatElement) {
                // If seat is now booked, update UI
                if (seat.is_booked && !seatElement.classList.contains('booked')) {
                    seatElement.classList.remove('available', 'selected');
                    seatElement.classList.add('booked');
                    
                    // Remove click event
                    seatElement.onclick = null;
                    
                    // If this was the selected seat, clear selection
                    if (selectedSeatId === seat.id.toString()) {
                        selectedSeatId = null;
                        document.getElementById('bookButton').disabled = true;
                    }
                }
            }
        });
        
    } catch (error) {
        console.error('Error updating seats:', error);
    }
} 