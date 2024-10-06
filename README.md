# ⚙️ Caesar Cipher Decoder

Welcome to the **Caesar Cipher Decoder**! This project is designed to help you decode Caesar ciphers while allowing me to practice my Go (Golang) skills. If you're interested in cryptography or just want to have some fun decoding messages, you've come to the right place!

## Table of Contents

- [Technologies Used](#technologies-used)
- [Run Locally](#run-locally)
- [Docker](#docker)
  - [Prerequisites](#prerequisites)
  - [Building the Docker Image](#building-the-docker-image)
  - [Running the Application](#running-the-application)
- [Screenshot](#screenshot)
- [Support](#support)

## Technologies Used

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) 
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## Run Locally

To run the application locally, follow these steps:

1. **Clone the project:**
   ```bash
   git clone https://github.com/Xanoor/Caesar-cipher-decoder
   ```

2. **Navigate to the project directory:**
   ```bash
   cd Caesar-cipher-decoder
   ```

3. **Create the executable:**
   ```bash
   go mod init main
   go build
   ```

4. **Run the executable:**
   - Use \`./main\` (or \`main.exe\` on Windows) in the terminal to run the application.

## Docker

This project includes a Docker setup, making it easy to build and run the application in a containerized environment.

### Prerequisites

Ensure you have the following installed on your machine:

- **Docker**: Download and install Docker from [Docker's official website](https://www.docker.com/get-started).

### Building the Docker Image

To build the Docker image for the Caesar Cipher Decoder, follow these steps:

1. **Clone the repository to your local machine:**
   ```bash
   git clone https://github.com/Xanoor/Caesar-cipher-decoder
   cd Caesar-cipher-decoder
   ```

2. **Build the Docker image:**
   ```bash
   docker build -t caesar-cipher-decoder .
   ```

3. **Running the Application**

    After building the image, run the application in a Docker container with the following command:

    ```bash
    docker run -it caesar-cipher-decoder
    ```

    This command starts the container and runs the Caesar Cipher Decoder application interactively.

## Screenshot

![Demo Screen](https://github.com/Xanoor/Caesar-cipher-decoder/blob/main/example.png)

## Support

If you have any questions or need support, feel free to reach out to me on Discord: **xanoor1**.

---

Have fun decoding!
