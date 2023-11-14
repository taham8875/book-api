# Book API with Golang

This is a simple Book API implemented in Golang.

## Introduction

This is a simple Book API implemented in Golang. It is a RESTful API that allows you to perform CRUD operations on a book resource.

This project is for the sake of learning.

## Features

- Create a book
- Get all books
- Get a single book
- Update a book
- Delete a book

## Getting Started

To get started with this project, you need to have Golang installed on your local machine. You can download and install Golang from [here](https://golang.org/dl/).

```bash
# Clone the repository
git clone https://github.com/taham8875/book-api

# Change into the project directory
cd book-api

# Install dependencies from the go.mod file
go mod download

# Run the application
go run main.go
```

You may use a REST client like [Postman](https://www.postman.com/) to test the API endpoints.

## Demo


https://github.com/taham8875/book-api/assets/92264237/29a979c9-6b38-4ab4-87c3-8ac455c28bd9




## API Endpoints

| Endpoint | Method | Description |
| --- | --- | --- |
| `/api/books` | `GET` | Get all books |
| `/api/books/{id}` | `GET` | Get a single book |
| `/api/books` | `POST` | Create a book |
| `/api/books/{id}` | `PUT` | Update a book |
| `/api/books/{id}` | `DELETE` | Delete a book |
