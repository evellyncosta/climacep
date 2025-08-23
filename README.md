# ClimaCEP

ClimaCEP is a Go application that receives a Brazilian ZIP code (CEP), identifies the city, and returns the current weather temperature in Celsius, Fahrenheit, and Kelvin.

LINK https://climacep-2okhjstqda-uc.a.run.app/

## Features

- Validates Brazilian ZIP codes (CEP)
- Fetches location information using ViaCEP API
- Retrieves current weather information using WeatherAPI
- Converts temperatures between Celsius, Fahrenheit, and Kelvin
- Handles various error scenarios with appropriate HTTP status codes
- Containerized with Docker for easy deployment

## Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose (for containerization)
- WeatherAPI API key

## Environment Variables

Create a `.env` file in the root directory with the following content:

```
WEATHER_API_KEY=your_weatherapi_key_here
```

## Running Locally

1. Clone the repository
2. Set up the environment variables
3. Run the application:

```bash
go run cmd/api/main.go cmd/api/handlers.go
```

## Running with Docker

1. Build and run the application using Docker Compose:

```bash
docker-compose up --build
```

## API Endpoints

### GET /weather

Fetches weather information based on a Brazilian ZIP code (CEP).

**Query Parameters:**

- `cep` (required): The Brazilian ZIP code (CEP) to lookup

**Success Response (200 OK):**

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

**Error Responses:**

- **422 Unprocessable Entity**: Invalid ZIP code format
  ```json
  {
    "message": "invalid zipcode"
  }
  ```

- **404 Not Found**: ZIP code not found
  ```json
  {
    "message": "can not find zipcode"
  }
  ```

## Testing

Run the tests:

```bash
go test ./...
```

## Deployment to Google Cloud Run

1. Build the Docker image:

```bash
docker build -t gcr.io/your-project-id/climacep .
```

2. Push the image to Google Container Registry:

```bash
docker push gcr.io/your-project-id/climacep
```

3. Deploy to Cloud Run:

```bash
gcloud run deploy climacep \
  --image gcr.io/your-project-id/climacep \
  --platform managed \
  --allow-unauthenticated \
  --region us-central1 \
  --set-env-vars WEATHER_API_KEY=your_api_key
```