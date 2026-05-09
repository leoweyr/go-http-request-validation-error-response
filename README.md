# Go Http Request Validation Error Response

Automatically transform go-playground/validator errors into structured, integration-ready JSON.

The library generates a standardized error response body designed for consistent API communication:

```json
{
  "error": {
    "message": "VALIDATION_FAILED",
    "details": {}
  }
}
```
