# Transportation Management System

This Transportation Management System is designed to manage and track transportation events and client profiles. It features a React frontend and a Go backend, providing capabilities such as event logging, client management, and alert generation based on predefined criteria.

## Prerequisites

Before you begin, ensure you have the following installed on your system:
- [Go](https://golang.org/dl/) - Download & install Go.
- [Node.js](https://nodejs.org/en/download/) - Download & install Node.js and the npm package manager.
- [Git](https://git-scm.com/downloads) - Download & install Git.
- A modern web browser, such as Google Chrome or Firefox.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Cloning the Repository

To get started, clone the repository to your local machine:

git clone https://github.com/yourusername/transportation-management-system.git
cd transportation-management-system
## Setting Up the Backend
Navigate to the backend directory from the root of the project:

cd backend
### Install the required Go modules (if any):

go mod tidy
### Configuring the Environment
Ensure your Go application is configured to connect to your database and other services. This might involve setting environment variables or configuring a .env file in your backend directory.

## Starting the Server
To start the Go server, run:

go run main.go
This will start the backend server on http://localhost:8080.

## Setting Up the Frontend
Open a new terminal window and navigate to the frontend directory from the root of the project:

cd frontend
### Install the required npm packages:

npm install
## Starting the Frontend Application
To start the React application, run:

npm start
This will open the Transportation Management System in your default web browser at http://localhost:3000.

## Usage
After starting both the frontend and backend, you can use the app to:

Add new transportation events.
View and manage client profiles.
Check alerts related to transportation events and client budget constraints.