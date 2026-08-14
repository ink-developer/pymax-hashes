package httpapi

type Versions map[string]VersionData

type VersionData struct {
	SignatureScheme       string       `json:"signature_scheme"`
	CertificateCount      int          `json:"certificate_count"`
	CertificateMetaSHA256 string       `json:"certificate_meta_sha256"`
	CertificateSHA256     []string     `json:"certificate_sha256"`
	DexMetaSHA256         string       `json:"dex_meta_sha256"`
	SoMetaSHA256Arm64V8A  string       `json:"so_meta_sha256_arm64_v8a"`
	SoMetaSHA256          SoMetaSHA256 `json:"so_meta_sha256"`
	BuildNumber           int          `json:"build_number"`
}

type SoMetaSHA256 struct {
	Arm64V8A   string `json:"arm64-v8a"`
	ArmeabiV7A string `json:"armeabi-v7a"`
	X86        string `json:"x86"`
	X8664      string `json:"x86_64"`
}
