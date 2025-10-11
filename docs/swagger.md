# API Documentation

## Swagger/OpenAPI Documentation

This project includes comprehensive API documentation using OpenAPI 3.1.0 specification.

### Viewing the Documentation

#### Online Swagger Editor
1. Open [Swagger Editor](https://editor.swagger.io/)
2. Copy the content from [`docs/openapi.yaml`](openapi.yaml)
3. Paste it into the editor to view interactive API documentation

#### Local Swagger UI
You can set up Swagger UI locally to view the documentation:

1. **Using Docker** (Recommended):
   ```bash
   docker run -p 80:8080 -e SWAGGER_JSON=/openapi.yaml -v $(pwd)/docs/openapi.yaml:/openapi.yaml swaggerapi/swagger-ui
   ```

2. **Using Node.js**:
   ```bash
   npm install -g swagger-ui-dist
   swagger-ui-serve -f docs/openapi.yaml -p 8080
   ```

3. **Using Python**:
   ```bash
   pip install flask-swagger-ui
   # Create a simple Flask app to serve the documentation
   ```

### API Endpoints Summary

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/shorten` | Create a shortened URL from a long URL |
| GET | `/{short_code}` | Retrieve original URL by short code |

### Response Format

All API responses follow a consistent format:

```json
{
  "success": boolean,
  "message": "string",
  "data": object,
  "timestamp": "ISO 8601 datetime",
  "request_id": "string"
}
```

### Error Handling

The API returns appropriate HTTP status codes and error messages:

- `200` - Success
- `201` - Created (for URL shortening)
- `400` - Bad Request (invalid input)
- `404` - Not Found (URL not found)
- `500` - Internal Server Error

### Authentication

Currently, the API does not require authentication. The OpenAPI spec includes JWT Bearer authentication setup for future implementation.

### Rate Limiting

Rate limiting is not currently implemented but can be added using middleware.

### Examples

#### Shorten URL Example
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://www.example.com/very/long/url"
  }'
```

#### Get URL Example
```bash
curl -X GET http://localhost:8080/KHlY5UTAnS
```

### Integration with Code

The OpenAPI specification can be used to:

1. **Generate client SDKs** using tools like OpenAPI Generator
2. **Automated testing** with tools like Dredd or Postman
3. **API mocking** during development
4. **Documentation generation** for various formats

### Code Generation

You can generate Go client code using:

```bash
# Install OpenAPI Generator
npm install @openapitools/openapi-generator-cli

# Generate Go client
openapi-generator-cli generate -i docs/openapi.yaml -g go -o client/go
```

### Validation

The OpenAPI specification can be used for request/response validation:

```bash
# Validate the specification
swagger-codegen validate -i docs/openapi.yaml

# Use spectral for advanced linting
npm install -g @stoplight/spectral-cli
spectral lint docs/openapi.yaml