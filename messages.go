package yaramsg

// struct to handle matches
type Match struct {
	Rule      string
	Namespace string
	Tags      []string
}

// Upload response

type UploadResponse struct {
	Message string
	Result  bool
}

// Remove response

type RemoveResponse struct {
	Message string
	Result  bool
}

// List Response

type ListResponse struct {
	Files []string
}

// Scan Respnse

type ScanResponse struct {
	Matches []Match
}
