const mysql = require('mysql2');

// Create a connection to the MySQL database
const connection = mysql.createConnection({
    host: 'localhost',
    user: 'root', // Replace with your MySQL username
    password: 'admin', // Replace with your MySQL password
    database: 'nodejs_mysql_example'
});

// Connect to MySQL
connection.connect((err) => {
    if (err) {
        console.error('Error connecting to MySQL:', err.stack);
        return;
    }
    console.log('Connected to MySQL as ID', connection.threadId);
});

// Insert a new user
const addUser = (name, email, age) => {
    const query = 'INSERT INTO users (name, email, age) VALUES (?, ?, ?)';
    connection.execute(query, [name, email, age], (err, results) => {
        if (err) {
            console.error('Error adding user:', err.stack);
            return;
        }
        console.log('User added:', results.insertId);
    });
};

// Fetch all users
const getUsers = () => {
    const query = 'SELECT * FROM users';
    connection.query(query, (err, results) => {
        if (err) {
            console.error('Error fetching users:', err.stack);
            return;
        }
        console.log('Users:', results);
    });
};

// Update a user
const updateUser = (id, name, email, age) => {
    const query = 'UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?';
    connection.execute(query, [name, email, age, id], (err, results) => {
        if (err) {
            console.error('Error updating user:', err.stack);
            return;
        }
        console.log('User updated:', results.affectedRows);
    });
};

// Delete a user
const deleteUser = (id) => {
    const query = 'DELETE FROM users WHERE id = ?';
    connection.execute(query, [id], (err, results) => {
        if (err) {
            console.error('Error deleting user:', err.stack);
            return;
        }
        console.log('User deleted:', results.affectedRows);
    });
};

// Example usage
addUser('John Doe', 'john.doe@example.com', 30);
getUsers();
updateUser(1, 'Jane Doe', 'jane.doe@example.com', 25);
deleteUser(1);

// Close the connection when you're done
connection.end();
