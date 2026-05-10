# Go HTTP API Contract IO

Automatically handle strict request decoding, transparent validation with JSON-tag reflection, and standardized JSON responses.

The library generates a standardized error response body designed for consistent API communication:

```json
{
  "error": {
    "message": "VALIDATION_FAILED",
    "details": {}
  }
}
```
