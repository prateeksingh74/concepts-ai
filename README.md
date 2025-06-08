# Concept Explainer

A small API using Go and Python that explains computer science concepts with a textbook definition and a real-life analogy. Has a brief hard-coded dictionary and a default response for concepts not included in the dictionary.

## Folder Structure

```
concept-explainer/
├── go-server/
│   └── main.go             # Go server with /explain endpoint
├── ai-engine/
│   └── explainer.py        # Python script that returns definition + analogy
```

## Setup

1. **Install Go** (https://golang.org/dl/)
2. **Install Python 3** (https://www.python.org/downloads/)

## Running the Server

```
cd go-server
# Start the Go server
go run main.go
```

## Example API Call

```
curl -X POST http://localhost:8080/explain \
  -H "Content-Type: application/json" \
  -d '{"concept": "hashmap"}'
```

## Example Response

```
{
  "definition": "A hashmap is a data structure used to store key-value pairs for fast lookup.",
  "analogy": "Think of it like a phone book — you search for a name (key) to get the number (value)."
}
```

## How it Works

- The Go server accepts POST requests at `/explain` with a JSON body: `{ "concept": "recursion" }`
- It calls the Python script: `python3 ../ai-engine/explainer.py recursion`
- The Python script returns a JSON object with a definition and analogy
- The Go server relays this JSON back to the client
