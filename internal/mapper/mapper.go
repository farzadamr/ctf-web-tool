package mapper

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Endpoint represents a discovered endpoint with risk score

type Endpoint struct {
	Path  string
	Score int
	Hints []string
}

// MapTarget fetches a URL and returns ranked endpoints
func MapTarget(target string) ([]Endpoint, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	body := string(data)

	paths := extractPaths(body)
	endpoints := scorePaths(paths)

	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].Score > endpoints[j].Score
	})

	// normalize relative paths
	for i := range endpoints {
		if strings.HasPrefix(endpoints[i].Path, "/") {
			endpoints[i].Path = u.Scheme + "://" + u.Host + endpoints[i].Path
		}
	}

	return endpoints, nil
}

// extractPaths finds endpoints from HTML & JS
func extractPaths(body string) []string {
	set := map[string]bool{}

	re := regexp.MustCompile(`(?i)(/api/[a-z0-9_/-]+|/[a-z0-9_-]+)`) // simple heuristic
	matches := re.FindAllString(body, -1)

	for _, m := range matches {
		if len(m) > 1 && !strings.Contains(m, ".") {
			set[m] = true
		}
	}

	paths := make([]string, 0, len(set))
	for p := range set {
		paths = append(paths, p)
	}
	return paths
}

// scorePaths assigns risk scores based on keywords
func scorePaths(paths []string) []Endpoint {
	var endpoints []Endpoint

	high := []string{"admin", "debug", "internal", "upload", "config", "manage"}
	medium := []string{"profile", "user", "account", "api"}

	for _, p := range paths {
		score := 1
		hints := []string{}

		for _, h := range high {
			if strings.Contains(p, h) {
				score += 5
				hints = append(hints, "HIGH value keyword: "+h)
			}
		}

		for _, m := range medium {
			if strings.Contains(p, m) {
				score += 2
				hints = append(hints, "Possible IDOR / auth check")
			}
		}

		endpoints = append(endpoints, Endpoint{
			Path:  p,
			Score: score,
			Hints: hints,
		})
	}

	return endpoints
}
func MapTargetString(target string) string {
	endpoints, err := MapTarget(target)
	if err != nil {
		return err.Error()
	}

	if len(endpoints) == 0 {
		return "No endpoints found"
	}

	var out strings.Builder

	for _, e := range endpoints {
		level := "LOW"
		if e.Score >= 7 {
			level = "HIGH"
		} else if e.Score >= 4 {
			level = "MEDIUM"
		}

		out.WriteString(
			"[" + level + "] " + e.Path +
				" (score=" + strconv.Itoa(e.Score) + ")\n",
		)

		for _, h := range e.Hints {
			out.WriteString("  - " + h + "\n")
		}
	}

	return out.String()
}
