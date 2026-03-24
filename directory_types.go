package alanube

// DirectoryEntry represents an active company in the electronic document directory.
type DirectoryEntry struct {
	Nombre        string `json:"nombre"`
	RNC           string `json:"rnc"`
	URLRecepcion  string `json:"urlRecepcion"`
	URLAceptacion string `json:"urlAceptacion"`
	URLOpcional   string `json:"urlOpcional"`
}

// DirectoryQueryOptions represents query parameters for directory lookups.
type DirectoryQueryOptions struct {
	RNC string // 9-11 chars
}
