# GoHire

## Description

This is a simple project to demonstrate the use of the Go programming language to create a simple REST API.

## Installation

To install the project, you need to have Go and Docker installed on your machine. You can find the installation instructions for Go [here](https://golang.org/doc/install) and for Docker [here](https://docs.docker.com/get-docker/).

After installing Go, you can clone the project using the following command:

```bash
git clone
```

## Usage

To run the project, first you need to setup the database. Copy the `.env.example` file to `.env` and set the environment variables to the desired values. Then, you can start the database using the following command:

```bash
docker-compose up -d
```

After starting the database, you can run the project using the following command:

```bash
go run main.go
```

By default, the project will run the project with Swagger UI enabled. You can access the Swagger UI by navigating to `http://localhost:8080/swagger/index.html`.

Alternatively, you can use the following `make` commands to interact with the project:

- `make run`: Run the project
- `make run-with-docs`: Generate the Swagger documentation and run the project
- `make test`: Run the tests
- `make build`: Build the project into a binary file
- `make docs`: Generate the Swagger documentation
- `make clean`: Clean the project by removing the binary file and the generated documentation

## Project Structure

The project is structured as follows:

- `config`: Contains the configuration files for the project
- `docs`: Contains the generated Swagger documentation
- `handler`: Contains the request handlers for the project
- `router`: Contains the router for the project
- `schema`: Contains the schema definitions for the project
