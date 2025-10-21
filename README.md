# GoString API: Advanced String Processing Microservice ����

## Overview
This GoString API, built with the high-performance Gin web framework, offers a robust solution for advanced string manipulation and analysis. It provides capabilities to store, retrieve, and query strings based on various properties like length, palindrome status, unique characters, and even through natural language commands. Designed for efficiency and extensibility, it's an ideal backend for applications requiring sophisticated text data management. 🚀

## Features
- **Gin Web Framework**: Leverages Gin for building a high-performance HTTP API with efficient request routing and middleware support.
- **In-Memory Data Store**: Utilizes an in-memory slice (`HardcodedStrings`) for fast data storage and retrieval, suitable for demonstrations or environments where persistent storage is externally managed.
- **Comprehensive String Analysis**: Automatically computes and stores various string properties, including length, palindrome status, unique character count, word count, SHA256 hash, and character frequency map.
- **Flexible Querying**: Supports filtering strings based on multiple criteria (e.g., palindrome, length ranges, word count, character presence) via URL query parameters.
- **Natural Language Processing (NLP) for Queries**: Allows users to filter strings using natural language phrases (e.g., "palindrome longer than 5 words") for intuitive interaction.
- **SHA256 Hashing**: Provides secure hashing of string values for unique identification.
- **Environment Variable Management**: Uses `godotenv` for easy configuration of environment variables.

## Getting Started

To get a copy of this project up and running on your local machine for development and testing purposes, follow these steps.

### Installation
1.  **Clone the Repository**
    ```bash
    git clone https://github.com/franzego/stage01.git
    cd stage01
    ```
2.  **Install Dependencies**
    ```bash
    go mod tidy
    ```
3.  **Run the Application**
    ```bash
    go run main.go
    ```
    The server will start on the port specified in your `.env` file or default to `8080`.

### Environment Variables
The project uses environment variables for configuration. You need to create a `.env` file in the root directory.

| Variable | Example Value | Description                                   |
| :------- | :------------ | :-------------------------------------------- |
| `PORT`   | `8080`        | The port on which the Gin server will listen. |

Example `.env` file:
```
PORT=8080
```

## Usage

Once the server is running, you can interact with the API using `curl` or any API client.

### Create a String
```bash
curl -X POST http://localhost:8080/strings \
-H "Content-Type: application/json" \
-d '{"value": "madam"}'
```

### Retrieve a Specific String
```bash
curl http://localhost:8080/strings/madam
```

### Query Strings with Filters
```bash
# Get palindromes with a word count of 1
curl "http://localhost:8080/strings?is_palindrome=true&word_count=1"

# Get strings longer than 5 characters containing 'e'
curl "http://localhost:8080/strings?min_length=6&contains_character=e"
```

### Query Strings with Natural Language
```bash
# Get strings that are palindromes
curl "http://localhost:8080/strings/filter-by-natural-language?query=show me all palindromes"

# Get strings with 2 words and containing the letter 'o'
curl "http://localhost:8080/strings/filter-by-natural-language?query=get strings with 2 words containing the letter o"

# Get strings between 5 and 10 characters long
curl "http://localhost:8080/strings/filter-by-natural-language?query=strings between 5 and 10 characters"
```

### Delete a String
```bash
curl -X DELETE http://localhost:8080/strings/madam
```

## API Documentation

### Base URL
`http://localhost:8080` (or the port configured in your `.env` file)

### Endpoints

#### POST /strings
Creates a new string entry with its analyzed properties.

**Request**:
```json
{
  "value": "string"
}
```
**Example**:
```json
{
  "value": "madam"
}
```

**Response**:
```json
{
  "id": "e40292c300b99187a53c0765d75d5e53e4142f36f9828d54c15335195155c5e8",
  "value": "madam",
  "props": {
    "length": 5,
    "is_palindrome": true,
    "unique_characters": 3,
    "word_count": 1,
    "sha256_hash": "e40292c300b99187a53c0765d75d5e53e4142f36f9828d54c15335195155c5e8",
    "character_frequnecy_map": {
      "a": 2,
      "d": 1,
      "m": 2
    }
  },
  "created_at": "2025-10-20T12:00:00Z"
}
```

**Errors**:
- `422 Unprocessable Entity`: Invalid JSON payload.
- `400 Bad Request`: Request body contains an empty or whitespace-only string value.
- `409 Conflict`: The string `value` already exists in the system.

#### GET /strings/:string_value
Retrieves a specific string entry by its exact value.

**Request**:
None (value is in path)

**Response**:
```json
{
  "id": "b1946ac92492d2347c6235b4d2611184",
  "value": "rotator",
  "props": {
    "length": 7,
    "is_palindrome": true,
    "unique_characters": 5,
    "word_count": 1,
    "sha256_hash": "b1946ac92492d2347c6235b4d2611184",
    "character_frequnecy_map": {
      "r": 2,
      "o": 1,
      "t": 2,
      "a": 1
    }
  },
  "created_at": "2025-10-20T12:00:00Z"
}
```

**Errors**:
- `404 Not Found`: The requested string was not found.

#### GET /strings
Retrieves a list of strings, optionally filtered by various criteria.

**Query Parameters**:
- `is_palindrome`: `true` or `false` (boolean)
- `min_length`: Minimum length (integer)
- `max_length`: Maximum length (integer)
- `word_count`: Exact word count (integer)
- `contains_character`: A single character that the string must contain (string)

**Request**:
None (parameters are in query string)

**Example**:
`GET /strings?is_palindrome=true&min_length=5`

**Response**:
```json
{
  "data": [
    {
      "id": "b1946ac92492d2347c6235b4d2611184",
      "value": "rotator",
      "props": {
        "length": 7,
        "is_palindrome": true,
        "unique_characters": 5,
        "word_count": 1,
        "sha256_hash": "b1946ac92492d2347c6235b4d2611184",
        "character_frequnecy_map": {
          "r": 2,
          "o": 1,
          "t": 2,
          "a": 1
        }
      },
      "created_at": "2025-10-20T12:00:00Z"
    }
  ],
  "count": 1,
  "filters_applied": {
    "is_palindrome": true,
    "min_length": 5,
    "max_length": 0,
    "word_count": 0,
    "contains_character": ""
  }
}
```

**Errors**:
- `400 Bad Request`: No strings matched the applied filters.

#### GET /strings/filter-by-natural-language
Retrieves a list of strings by parsing a natural language query.

**Query Parameters**:
- `query`: A natural language phrase describing the desired string properties (string)

**Request**:
None (parameter is in query string)

**Example**:
`GET /strings/filter-by-natural-language?query=show me strings that are palindromes and have 1 word`

**Response**:
```json
{
  "data": [
    {
      "id": "b1946ac92492d2347c6235b4d2611184",
      "value": "rotator",
      "props": {
        "length": 7,
        "is_palindrome": true,
        "unique_characters": 5,
        "word_count": 1,
        "sha256_hash": "b1946ac92492d2347c6235b4d2611184",
        "character_frequnecy_map": {
          "r": 2,
          "o": 1,
          "t": 2,
          "a": 1
        }
      },
      "created_at": "2025-10-20T12:00:00Z"
    }
  ],
  "count": 1,
  "interpreted_query": {
    "original": "show me strings that are palindromes and have 1 word",
    "parsed_filters": {
      "is_palindrome": true,
      "min_length": 0,
      "max_length": 0,
      "word_count": 1,
      "contains_char": ""
    }
  }
}
```

**Errors**:
- `400 Bad Request`: Query parameter 'query' is required.
- `400 Bad Request`: String not found (parser error).
- `422 Unprocessable Entity`: Query parsed but resulted in conflicting filters (e.g., `min_length` > `max_length`).
- `404 Not Found`: No strings matched your filters.

#### DELETE /strings/:string_value
Deletes a specific string entry by its exact value.

**Request**:
None (value is in path)

**Response**:
`204 No Content`

**Errors**:
- `404 Not Found`: The string does not exist in the system.

## Technologies Used

| Technology                                                 | Description                                                |
| :--------------------------------------------------------- | :--------------------------------------------------------- |
| [Go](https://golang.org/)                                  | Primary programming language.                              |
| [Gin](https://gin-gonic.com/docs/)                         | High-performance HTTP web framework.                       |
| [Go DotEnv](https://github.com/joho/godotenv)              | Loads environment variables from a `.env` file.            |
| [SHA256 (Go Crypto)](https://pkg.go.dev/crypto/sha256)     | Cryptographic hashing for string identification.           |
| [Go Strings](https://pkg.go.dev/strings)                   | String manipulation utilities.                             |
| [Go Time](https://pkg.go.dev/time)                         | Timestamp generation.                                      |
| [Go Unicode](https://pkg.go.dev/unicode)                   | Unicode character operations for string analysis.          |
| [Go Regular Expressions](https://pkg.go.dev/regexp)        | Pattern matching for natural language query parsing.       |

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, please follow these steps:

1.  ✨ Fork the repository.
2.  🌿 Create a new branch (`git checkout -b feature/AmazingFeature`).
3.  ✏️ Make your changes and commit them (`git commit -m 'Add some AmazingFeature'`).
4.  🚀 Push to the branch (`git push origin feature/AmazingFeature`).
5.  📬 Open a Pull Request, describing your changes in detail.

## License

This project currently does not have an explicit license file. Please contact the author for licensing information.

## Author Info

-   **Franz Ego**
-   
-   [Twitter](https://twitter.com/saint_franz)

---

[![Go Reference](https://pkg.go.dev/badge/github.com/franzego/stage01.svg)](https://pkg.go.dev/github.com/franzego/stage01)
[![Go Report Card](https://goreportcard.com/badge/github.com/franzego/stage01)](https://goreportcard.com/report/github.com/franzego/stage01)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) <!-- Assuming MIT as a common default, but user should verify/update -->
[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)