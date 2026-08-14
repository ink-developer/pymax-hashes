package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type Storage struct {
	path string
	mu   sync.RWMutex
}

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

func New(path string) *Storage {
	return &Storage{path: path}
}

func (s *Storage) GetVersions() (Versions, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	content, err := os.ReadFile(s.path)

	if err != nil {
		return nil, err
	}

	var versions Versions

	err = json.Unmarshal(content, &versions)

	if err != nil {
		return nil, err
	}

	return versions, nil
}

func (s *Storage) SaveVersions(versions Versions) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(versions, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(s.path)

	tmp, err := os.CreateTemp(dir, ".versions-*.tmp")
	if err != nil {
		return err
	}

	tmpName := tmp.Name()

	defer func() {
		tmp.Close()
		os.Remove(tmpName)
	}()

	if _, err := tmp.Write(data); err != nil {
		return err
	}

	if err := tmp.Sync(); err != nil {
		return err
	}

	if err := tmp.Close(); err != nil {
		return err
	}

	return os.Rename(tmpName, s.path)
}
