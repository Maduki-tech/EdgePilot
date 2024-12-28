# EdgePilot

## Featurelist

- [ ] Reverse Proxy: Forward requests to different backend
services based on the URL path.

- [ ] Caching: Cache responses for frequently requested resources.

- [ ] Rate Limiting: Prevent abuse by limiting the number of requests per client.

- [ ] Content Negotiation: Serve responses in JSON or XML
depending on the client's preference.

- [ ] Error Handling: Standardize error responses for all backend services.

- [ ] HTTP/2 Support: Enable HTTP/2 for faster communication.

- [ ] Monitoring: Log request metrics and visualize traffic patterns.

- [ ] Authentication: Implement basic token-based authentication for API access.

## Techstack

- Go
- Docker
- `net/http`
- `golang.org/x/net/http2`
- `go-cache`
- `rate`

## Implementation Plan

### Phase 1: Reverse Proxy

Route requests to mock services based on URL paths.
Implement basic request forwarding.

### Phase 2: Caching

Add caching for responses based on the requested resource.
Use Cache-Control and ETag headers to make it HTTP-compliant.

### Phase 3: Rate Limiting

Use the rate library to track and throttle clients.
Implement rate limiting per client IP using a token bucket algorithm.

### Phase 4: Content Negotiation

Serve responses in JSON or XML based on the Accept header.

### Phase 5: Error Handling

Standardize error messages across all routes.
Implement retry headers for specific error cases.

### Phase 6: HTTP/2 Support

Enable HTTP/2 for the API gateway.

### Phase 7: Monitoring

Log request and response times.
Track and visualize the most hit endpoints using Grafana or a simple dashboard.

### Phase 8: Authentication

Add token-based authentication for access control.
Validate tokens before forwarding requests to backend services
