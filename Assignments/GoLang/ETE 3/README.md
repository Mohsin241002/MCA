# Cinema Ticket Booking System

A modern ticket booking system for cinema theaters built with Go, PostgreSQL, GORM, and a HTML/CSS/JavaScript frontend.

## Features

- User authentication and registration
- Browse movies and showtimes
- View seat availability
- Book seats with concurrency control
- View booking history
- JWT-based authentication
- Real-time seat availability updates

## Project Structure

```
cinema-booking-system/
├── cmd/
│   └── server/              # Application entry point
├── internal/
│   ├── models/              # Database models using GORM
│   ├── handlers/            # HTTP request handlers
│   ├── middleware/          # Authentication, logging, etc.
│   ├── database/            # Database connection and migrations
│   └── utils/               # Helper functions
├── api/
│   └── routes/              # API route definitions
├── web/
│   ├── static/              # CSS, JS, images
│   └── templates/           # HTML templates
├── migrations/              # SQL migration files
├── tests/                   # Unit and integration tests
├── .env                     # Environment variables
└── README.md                # Documentation
```

## Setup and Installation

### Prerequisites

- Go 1.18 or higher
- PostgreSQL 13 or higher

### Steps

1. Clone the repository:

```bash
git clone https://github.com/yourusername/cinema-booking-system.git
cd cinema-booking-system
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up the PostgreSQL database:

```sql
CREATE DATABASE cinema_booking;
```

4. Configure environment variables in `.env` file:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=cinema_booking
SERVER_PORT=8080
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRY=24h
```

5. Run the application:

```bash
go run cmd/server/main.go
```

6. Open your browser and navigate to `http://localhost:8080`

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login and get a JWT token

### Movies

- `GET /api/movies` - Get all movies
- `GET /api/movies/{id}` - Get a specific movie

### Shows

- `GET /api/shows` - Get all shows
- `GET /api/shows/{id}` - Get a specific show
- `GET /api/shows/{id}/seats` - Get seats for a specific show

### Bookings

- `GET /api/bookings` - Get user's bookings (requires authentication)
- `POST /api/bookings` - Create a new booking (requires authentication)

## Concurrency Control

The application implements concurrency control for seat booking using:

- Goroutines for handling concurrent requests
- Mutex locks to prevent race conditions during booking
- Database transactions to ensure data integrity

## Future Improvements

- Add admin panel for managing movies and shows
- Implement seat selection interface
- Add payment gateway integration
- Support for multiple cinema locations
- Email notifications for bookings 