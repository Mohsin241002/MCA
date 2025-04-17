// Auth state management
document.addEventListener('DOMContentLoaded', function() {
    checkAuthState();

    // Logout button event listener
    const logoutBtn = document.getElementById('logoutBtn');
    if (logoutBtn) {
        logoutBtn.addEventListener('click', function(e) {
            e.preventDefault();
            logout();
        });
    }
});

// Check authentication state
function checkAuthState() {
    const token = localStorage.getItem('token');
    
    if (token) {
        // User is logged in
        document.getElementById('loginNav')?.classList.add('d-none');
        document.getElementById('registerNav')?.classList.add('d-none');
        document.getElementById('bookingsNav')?.classList.remove('d-none');
        document.getElementById('logoutNav')?.classList.remove('d-none');
    } else {
        // User is logged out
        document.getElementById('loginNav')?.classList.remove('d-none');
        document.getElementById('registerNav')?.classList.remove('d-none');
        document.getElementById('bookingsNav')?.classList.add('d-none');
        document.getElementById('logoutNav')?.classList.add('d-none');
    }
}

// Logout user
function logout() {
    localStorage.removeItem('token');
    checkAuthState();
    window.location.href = '/';
}

// API call helper with authentication
async function fetchAuth(url, options = {}) {
    const token = localStorage.getItem('token');
    
    // Set default options
    options.headers = options.headers || {};
    
    // Add authentication header if token exists
    if (token) {
        options.headers['Authorization'] = `Bearer ${token}`;
    }
    
    // Add default content type if not specified
    if (!options.headers['Content-Type'] && options.method !== 'GET') {
        options.headers['Content-Type'] = 'application/json';
    }
    
    try {
        const response = await fetch(url, options);
        
        // Handle unauthorized responses (invalid/expired token)
        if (response.status === 401) {
            // Clear token and update UI
            localStorage.removeItem('token');
            checkAuthState();
            
            // Redirect to login if on a protected page
            if (window.location.pathname.includes('/bookings')) {
                window.location.href = '/login';
            }
        }
        
        return response;
    } catch (error) {
        console.error('Fetch error:', error);
        throw error;
    }
} 