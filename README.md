# COC Proxy

## Overview

COC proxy is a simple API written in Go that acts as a proxy for making requests to the Clash of Clans API.

It is especially useful for making requests to the Clash of Clans API from a constanly changing IP address, such as a Google spreadsheet app script.

## Features

- Exposes an endpoint to make requests to the Clash of Clans API.
- Handles requests for player data by playerId and url parameters.

## requirements

- Go 1.16 or higher installed on your system.
- Docker (optional, for running the application in a containerized environment).

## Installation

1. Clone the repository:

```bash
git clone git@github.com:ZiplEix/COC_proxy.git
```

2. Change to the project directory:

```bash
cd COC_proxy
```

3. Create a `.env` file in the root of the project and add the following environment variables:

```bash
COC_TOKEN=<your_clash_of_clans_api_token>
```

## Usage

### Running the application locally

1. Make sure you have set up your environment variables. Refer to the ``.env.example`` file for required environment variables.

2. Build and run the application:

```bash
go run main.go
```

3. The application will start on port **8080** by default.

4. You can access the API at **<http://localhost:8080/>**.

### Running the application with Docker

1. Make sure you have set up your environment variables. Refer to the ``.env.example`` file for required environment variables.

2. Build the Docker image:

```bash
docker build -t coc-proxy .
```

3. Run the Docker container:

```bash
docker run -p 8080:8080 coc-proxy
```

4. The API server will start running inside a Docker container on port **8080**.

5. You can access the API at **<http://localhost:8080/>**.

## API Documentation

### Endpoint 1

- `/`

#### Parameters

- `url`: URL for making requests to the Clash of Clans API.

#### Example

```bash
curl -X GET "http://localhost:8080/?url=https://api.clashofclans.com/v1/players/%23P0LYJC8C"
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find any bugs or have suggestions for improvements.
