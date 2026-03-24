package alanube

// ServiceStatus represents the status of a single DGII service.
type ServiceStatus struct {
	Service string `json:"servicio"`
	Status  string `json:"estatus"`
}

// DGIIStatusResponse is the response when checking DGII status without params.
// The response wraps services in {"response": [[...]]}.
type DGIIStatusResponse struct {
	Response [][]ServiceStatus `json:"response"`
}

// EnvironmentStatus represents the status of a specific environment.
type EnvironmentStatus struct {
	Status string `json:"estado"`
}

// MaintenanceWindow represents a scheduled maintenance window.
type MaintenanceWindow struct {
	Environment string   `json:"ambiente"`
	StartTime   string   `json:"horaInicio"`
	EndTime     string   `json:"horaFin"`
	Days        []string `json:"dias"`
}

// MaintenanceResponse is the response for maintenance schedule queries.
type MaintenanceResponse struct {
	Windows []MaintenanceWindow `json:"ventanaMantenimientos"`
}

// DGIIEnvironment represents available DGII environments.
type DGIIEnvironment string

const (
	DGIIPreCertification DGIIEnvironment = "1"
	DGIIProduction       DGIIEnvironment = "2"
	DGIICertification    DGIIEnvironment = "3"
)

// DGIIQueryOptions represents query parameters for DGII status endpoints.
type DGIIQueryOptions struct {
	Environment DGIIEnvironment // "1", "2", "3" - takes priority over Maintenance
	Maintenance bool            // true to query maintenance windows
}
