# HTTP Proxy Server in Go

This project implements an HTTP proxy server in Go, designed to proxy requests to external services and return responses in JSON format.

## Getting started

1. Clone the repository:
   ```bash
   git clone https://github.com/allwsaa/http-proxy-task1.git
   cd http-proxy-task1
2. Build the project:
    ```bash
    make build
3. Run the server:
    ```bash
    make run


## ENDPOINTS

The HTTP proxy server exposes the following endpoints:

## Health Check

- **URL:** `/health`
- **Method:** `GET`
- **Description:** Verifies if the server is operational.
- **Response:** Returns `200 OK` if the server is running.

### Proxy Endpoint

- **URL:** `/proxy`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Description:** Proxies HTTP requests to specified URLs.
- **Request Format:** Requires JSON payload specifying HTTP method, target URL, and optional headers.
- **Response Format:** Returns a JSON object with details including request ID, status code, headers, and response body length.


## Swagger LINK: 
http://localhost:8080/swagger/index.html

## Render Deployment:
https://http-proxy-task1-ntvlbl.onrender.com



