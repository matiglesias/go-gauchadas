# GAUCHADAS

This is a project created by Mati Iglesias and Agus Garcia.

## Requirements
* Docker
* Go 1.14 or newer

## Getting Started
This project needs to have a MongoDB instance running. Run `make db` to automatically retrieve a Docker image for it, create an instance and run it. To stop it (and also delete it), run `make db-stop`.

Once you have the database running, you can do `make` or `make run` to automatically build and run the project. Use `make clean` to delete the executable.
