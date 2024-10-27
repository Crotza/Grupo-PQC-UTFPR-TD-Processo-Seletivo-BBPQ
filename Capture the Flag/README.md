# Flag Encrypted

This project contains a Python script that processes an encrypted image using various image processing techniques. The script is designed to run inside a Docker container.

## Prerequisites

- Docker must be installed on your system. You can download and install Docker from [here](https://www.docker.com/products/docker-desktop).

## Instructions

### 1. Verify Docker Installation

First, verify that Docker is installed correctly by checking the Docker version:

```sh
docker --version
```

You should see output similar to:

```sh
Docker version 27.2.0, build 3ab4256
```

### 2. Build the Docker Image

Navigate to the directory containing the `Dockerfile`, `requirements.txt`, `flag-encrypted.py`, and png files. Then, build the Docker image using the following command:

```sh
docker build -t flag-encrypted .
```

This command will create a Docker image named flag-encrypted.

### 3. Run the Docker Container

After building the Docker image, run the Docker container using the following command:

```sh
docker run --rm flag-encrypted
```

This command will execute the `flag-encrypted.py` script inside the Docker container. The output will be displayed in the terminal.


### Summary
- **Verify Docker Installation**: Use `docker --version` to check if Docker is installed.
- **Build Docker Image**: Use `docker build -t flag-encrypted .` to build the Docker image.
- **Run Docker Container**: Use `docker run --rm flag-encrypted` to run the container and execute the script.