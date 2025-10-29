# balance-api
Balance API used for check inquiry balance and withdraw balance from user

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Project](#running-the-project)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Features
- Check user balance
- Withdraw balance
- Transaction history

## Prerequisites
- Go (1.16 or later)
- MySQL Server
- Git (optional, for cloning the repository)

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/balance-api.git
   cd balance-api
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Create the MySQL database:**
   - Open MySQL command line or a GUI tool and run:
     ```sql
     CREATE DATABASE IF NOT EXISTS balance_db;
     ```

4. **Run the schema SQL script:**
   - Save your schema in a file named `schema.sql` and run:
     ```bash
     mysql -u root -p balance_db < schema.sql
     ```

## Configuration

1. **Create a configuration file:**
   - Create a file named `config.go` in the `config` directory with the following structure:
   ```go
   package config

   type Config struct {
       DBName     string
       DBUser     string
       DBPassword string
       DBHost     string
       DBPort     string
   }

   func NewConfig() *Config {
       return &Config{
           DBName:     "balance_db",
           DBUser:     "root", // your MySQL username
           DBPassword: "your-password-here",
           DBHost:     "localhost",
           DBPort:     "3306",
       }
   }
   ```

## Running the Project

1. **Run the application:**
   ```go 
   run cmd/api/main.go
   ```

3. **Access the API:**
   - The API will be running on `http://localhost:8080`.

## API Endpoints

- **GET /balance**: Retrieve the balance for a user.
- **POST /balance**: Withdraw balance from a user.