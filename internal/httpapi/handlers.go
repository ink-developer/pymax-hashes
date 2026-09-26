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

func (s *Router) GetVersions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	versions, err := s.storage.GetVersions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"ok": false}`)
		return
	}

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

	if sortErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"ok": false, "error": %q}`, sortErr.Error())
		return
	}

	since := r.URL.Query().Get("since")

	if since != "" {
		if _, err := versionParts(since); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"ok": false, "error": %q}`, err.Error())
			return
		}

		index := len(keys)

		for i, key := range keys {
			cmp, err := compareVersions(key, since)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"ok": false, "error": %q}`, err.Error())
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

	jsonData, err := json.MarshalIndent(outVersions, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"ok": false}`)
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=600")
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
		log.Printf("failed to save versions: %v", err)
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
