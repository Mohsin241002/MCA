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
    
    const registerForm = document.getElementById('register-form');
    const errorMessage = document.getElementById('error-message');
    
    registerForm.addEventListener('submit', function(e) {
        e.preventDefault();
        
        const fullName = document.getElementById('fullName').value;
        const email = document.getElementById('email').value;
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;
        const confirmPassword = document.getElementById('confirmPassword').value;
        const phoneNumber = document.getElementById('phoneNumber').value;
        
        // Validation
        if (password !== confirmPassword) {
            errorMessage.textContent = 'Passwords do not match';
            return;
        }
        
        // First create user in Firebase Auth
        firebase.auth().createUserWithEmailAndPassword(email, password)
            .then((userCredential) => {
                // Signed in 
                const user = userCredential.user;
                
                // Update profile with the display name
                user.updateProfile({
                    displayName: fullName
                }).then(() => {
                    console.log("Profile updated successfully");
                }).catch((error) => {
                    console.error("Error updating profile:", error);
                });
                
                // Now create user record in our backend
                return fetch('/api/users/register', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        id: user.uid,
                        email: email,
                        username: username,
                        full_name: fullName,
                        phone_number: phoneNumber,
                        role: 'user'
                    })
                });
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to complete registration');
                }
                return response.json();
            })
            .then(data => {
                console.log("Registration successful:", data);
                // Redirect to login page
                window.location.href = '/login?registered=true';
            })
            .catch((error) => {
                console.error("Registration error:", error);
                let errorMsg = 'Registration failed. Please try again.';
                
                if (error.code) {
                    switch(error.code) {
                        case 'auth/email-already-in-use':
                            errorMsg = 'Email is already in use';
                            break;
                        case 'auth/invalid-email':
                            errorMsg = 'Invalid email address';
                            break;
                        case 'auth/weak-password':
                            errorMsg = 'Password is too weak';
                            break;
                    }
                }
                
                errorMessage.textContent = errorMsg;
            });
    });
}); 