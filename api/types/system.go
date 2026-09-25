// Shared request/response types for the gluster API.
// Used by both the route handlers and any client, so the two stay in sync
package types

// Response body for GET /api/system/health
type HealthResponse struct {
	Status string `json:"status"`
}
