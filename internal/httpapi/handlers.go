package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pymax-hashes/internal/storage"
)

func (s *Router) GetVersions(w http.ResponseWriter, r *http.Request) {
	versions, err := s.storage.GetVersions()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"ok": false}`)
		return
	}

	jsonData, err := json.MarshalIndent(versions, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"ok": false}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control:", "public, max-age=600")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (s *Router) AddVersion(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("Authorization")

	if authKey != s.config.AuthKey {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"ok":false}`)
		return
	}

	var incoming storage.VersionData

	if err := json.NewDecoder(r.Body).Decode(&incoming); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"ok":false}`)
		return
	}

	versions, err := s.storage.GetVersions()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"ok": false}`)
		return
	}

	versionName := r.PathValue("version")
	versions[versionName] = incoming

	err = s.storage.SaveVersions(versions)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"ok": false}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"ok": true}`)
}
