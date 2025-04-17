// Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyBJzM8YzR-jcBsdWgLD7zsngamDBFQt2y8",
  authDomain: "cinema-34494.firebaseapp.com",
  projectId: "cinema-34494",
  storageBucket: "cinema-34494.firebasestorage.app",
  messagingSenderId: "571118173278",
  appId: "1:571118173278:web:75c461367bc17d5a4988cf",
  measurementId: "G-G59KGFNR9E"
};

// Initialize Firebase
firebase.initializeApp(firebaseConfig);

document.addEventListener('DOMContentLoaded', function() {
    // Redirect if already logged in
    const token = localStorage.getItem('token');
    if (token) {
        window.location.href = '/';
        return;
    }
    
    const loginForm = document.getElementById('login-form');
    const errorMessage = document.getElementById('error-message');
    
    loginForm.addEventListener('submit', function(e) {
        e.preventDefault();
        
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        
        // Firebase Authentication
        firebase.auth().signInWithEmailAndPassword(email, password)
            .then((userCredential) => {
                // Signed in
                const user = userCredential.user;
                console.log("User logged in:", user);
                
                // Get user details from the backend
                fetch(`/api/users/${user.uid}`)
                    .then(response => {
                        if (response.ok) {
                            // Redirect to homepage or dashboard
                            window.location.href = '/';
                        } else {
                            throw new Error('Failed to fetch user data');
                        }
                    })
                    .catch(error => {
                        console.error("Error fetching user data:", error);
                        errorMessage.textContent = 'Failed to complete login. Please try again.';
                    });
            })
            .catch((error) => {
                console.error("Login error:", error);
                let errorMsg = 'Login failed. Please check your credentials.';
                
                switch(error.code) {
                    case 'auth/user-not-found':
                    case 'auth/wrong-password':
                        errorMsg = 'Invalid email or password';
                        break;
                    case 'auth/too-many-requests':
                        errorMsg = 'Too many unsuccessful login attempts. Please try again later.';
                        break;
                    case 'auth/invalid-email':
                        errorMsg = 'Invalid email address';
                        break;
                }
                
                errorMessage.textContent = errorMsg;
            });
    });
    
    // Handle password reset link
    const resetPasswordLink = document.getElementById('reset-password');
    if (resetPasswordLink) {
        resetPasswordLink.addEventListener('click', function(e) {
            e.preventDefault();
            
            const email = document.getElementById('email').value;
            if (!email) {
                errorMessage.textContent = 'Please enter your email address';
                return;
            }
            
            firebase.auth().sendPasswordResetEmail(email)
                .then(() => {
                    alert('Password reset email sent! Check your inbox.');
                })
                .catch((error) => {
                    console.error("Password reset error:", error);
                    errorMessage.textContent = 'Failed to send password reset email. Please try again.';
                });
        });
    }
}); 