## BOOKIFY A BOOK MANAGEMENT SYSTEM- BACKEND

# Bookify is a simple web application built with the Go programming language and the Gin web framework. It provides a comprehensive backend solution for managing a collection of books. 

**It supports the following CRUD (Create, Read, Update, Delete) operations:**

```GET /books:```             Get a list of all books.
```GET /books/:title:```      Get a book by title.
```POST /books:```            Create a new book.
```PUT /books/:title:```      Update a book.
```DELETE /books/:title:```   Delete a book.

* The web application uses the Fetch API to make requests to the API.
* The API runs on port 8080.

# Dockerfile for Backend

**Key Highlights:**

* Utilizes a multi-stage build approach 1.Building the App 2.Create the production Image
* Optimizes the build process by leveraging Go modules for dependency management.
* Builds a statically-linked binary with CGO disabled for enhanced portability and security.
* Creates a minimalistic production image by starting from the scratch base image.
* Copies the built application binary into the production image.


This Dockerfile demonstrates best practices for containerizing a Golang backend application, ensuring a streamlined and efficient deployment process.