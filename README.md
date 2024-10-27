# Number Search Service

A high-performance REST service for efficient number search operations within a sorted dataset. The service loads a predefined sorted list of numbers (0 to 1000000) and provides fast lookup capabilities with configurable approximation tolerance. The service returns the position (index) of the requested number in the dataset.

## Features

- Fast binary search implementation
- Configurable approximation tolerance (10% by default)
- REST API endpoint for number position lookups
- Docker support
- Comprehensive logging with configurable levels
- Environment-based configuration
- Unit test coverage

## Prerequisites

- Go 1.23 or higher
- Docker (optional)
- Make

## Installation

1. Clone the repository:
```bash
git clone https://github.com/KRR19/number-search.git
cd number-search
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

The service can be configured using environment variables:

| Variable    | Description                    | Default Value  |
|------------|--------------------------------|----------------|
| PORT       | Server port                    | :8080         |
| LOG_LEVEL  | Logging level (info/debug/error)| info          |
| VARIATION  | Approximation tolerance (%)     | 10            |
| FILE_PATH  | Path to input numbers file     | ./input.txt   |

## Running the Service

### Using Make

```bash
# Run the service
make run

# Build the binary
make build

# Run tests
make test

# Format code
make fmt

# Run linter
make lint
```

### Using Docker

```bash
# Build and run with Docker
make docker-up

# Or build and run separately
make docker-build
make docker-run
```

## API Endpoints

### GET /v1/search/{number}

Search for a number's position in the sorted list using binary search algorithm.

#### Parameters
- `number` (path parameter): The number to search for

#### Response

```json
{
    "position": 3,
    "error": null
}
```

If the exact number is not found, the service will search for the closest number within the specified variation threshold (default 10%). For example, if searching for 1150 and the dataset contains 1100 and 1200, the service will return the position of either 1100 or 1200 if they fall within the variation threshold.

### GET /v2/search/{number}

Optimized version of the search endpoint specifically designed for a sorted sequence of numbers from 0 to 1000000 with a step of 100 (0, 100, 200, ..., 1000000). This endpoint provides better performance but works only with this specific number pattern.

For example:
- GET /v2/search/100 returns position 1
- GET /v2/search/200 returns position 2
- GET /v2/search/1000000 returns position 10000

If a number doesn't match the exact step pattern (100), the endpoint will attempt to find the closest matching position within the variation threshold.

#### Status Codes for both endpoints
- 200: Success
- 400: Invalid number format
- 404: Number not found (no matching numbers within variation threshold)
- 500: Internal server error

## Error Handling

The service implements approximate matching within a 10% variation by default. When searching for a number:
1. If exact match is found - returns its position
2. If no exact match - searches for numbers within the variation threshold and returns the position of the closest match
3. If no numbers found within threshold - returns a 404 error
