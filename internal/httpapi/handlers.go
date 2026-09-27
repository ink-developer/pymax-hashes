package httpapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pymax-hashes/internal/storage"
	"slices"
	"strconv"
	"strings"
)

func versionParts(v string) ([3]int, error) {
	parts := strings.Split(v, ".")

	var out [3]int

	if len(parts) != 3 {
		return out, fmt.Errorf("invalid version %q", v)
	}

	for i := 0; i < 3; i++ {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			return out, fmt.Errorf("invalid version %q: %w", v, err)
		}

		out[i] = n
	}

	return out, nil
}

func compareVersions(a, b string) (int, error) {
	av, err := versionParts(a)
	if err != nil {
		return 0, err
	}

	bv, err := versionParts(b)
	if err != nil {
		return 0, err
	}

	for i := 0; i < 3; i++ {
		if av[i] < bv[i] {
			return -1, nil
		}
		if av[i] > bv[i] {
			return 1, nil
		}
	}

	return 0, nil
}

func sendError(w http.ResponseWriter, status int, err string) {
	data := struct {
		Error string `json:"error"`
	}{
		Error: err,
	}

	sendJson(w, data, status)
}

func sendJson(w http.ResponseWriter, data any, status int) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		sendError(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

func getSortedKeys(versions storage.Versions) ([]string, error) {
	keys := make([]string, 0, len(versions))
	for k := range versions {
		keys = append(keys, k)
	}

	var sortErr error

	slices.SortFunc(keys, func(a, b string) int {
		cmp, err := compareVersions(a, b)
		if err != nil {
			sortErr = err
			return 0
		}

		return cmp
	})

	return keys, sortErr
}

func (s *Router) GetLatest(w http.ResponseWriter, r *http.Request) {
	versions, err := s.storage.GetVersions()
	if err != nil {
		sendError(w, http.StatusInternalServerError, "internal error")
		return
	}

	keys, sortErr := getSortedKeys(versions)

	if sortErr != nil {
		sendError(w, http.StatusInternalServerError, "internal error")
		return
	}

	if len(keys) == 0 {
		sendError(w, http.StatusNotFound, "no versions")
		return
	}

	key := keys[len(keys)-1]
	data := struct {
		Version string              `json:"version"`
		Data    storage.VersionData `json:"data"`
	}{
		Version: key,
		Data:    versions[key],
	}

	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", s.config.LatestCacheTime))
	sendJson(w, data, http.StatusOK)
}

func (s *Router) GetVersions(w http.ResponseWriter, r *http.Request) {
	versions, err := s.storage.GetVersions()
	if err != nil {
		sendError(w, http.StatusInternalServerError, "internal error")
		return
	}

	keys, sortErr := getSortedKeys(versions)

	if sortErr != nil {
		sendError(w, http.StatusInternalServerError, "internal error")
		return
	}

	since := r.URL.Query().Get("since")

	if since != "" {
		if _, err := versionParts(since); err != nil {
			sendError(w, http.StatusBadRequest, err.Error())
			return
		}

		index := len(keys)

		for i, key := range keys {
			cmp, err := compareVersions(key, since)
			if err != nil {
				sendError(w, http.StatusInternalServerError, "internal error")
				return
			}

			if cmp >= 0 {
				index = i
				break
			}
		}

		keys = keys[index:]
	}

	outVersions := make(storage.Versions, len(keys))

	for _, k := range keys {
		outVersions[k] = versions[k]
	}

	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", s.config.GlobalCacheTime))
	sendJson(w, outVersions, http.StatusOK)
}

func (s *Router) AddVersion(w http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("Authorization")

	if authKey != s.config.AuthKey {
		sendError(w, http.StatusUnauthorized, "invalid auth")
		return
	}

	var incoming storage.VersionData

	if err := json.NewDecoder(r.Body).Decode(&incoming); err != nil {
		sendError(w, http.StatusBadRequest, "invalid body")
		return
	}

	versions, err := s.storage.GetVersions()

	if err != nil {
		log.Printf("failed to save versions: %v", err)
		sendError(w, http.StatusInternalServerError, "internal error")
		return
	}

	versionName := r.PathValue("version")
	versions[versionName] = incoming

	err = s.storage.SaveVersions(versions)

	if err != nil {
		sendError(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Router) NotFound(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("frontend/404.html")
	if err != nil {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write(data)
}
