package hltb

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func readTestFile(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("test_files", name))
	if err != nil {
		t.Fatalf("read test file %q: %v", name, err)
	}
	return data
}

func newTestServer(t *testing.T, singleJson []byte, arrayJson []byte) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/hltb/55663":
			_, _ = w.Write(singleJson)
			return
		case r.Method == http.MethodGet && r.URL.Path == "/hltb/55663/refresh":
			_, _ = w.Write(singleJson)
			return
		case r.Method == http.MethodGet && r.URL.Path == "/steam/991270":
			_, _ = w.Write(singleJson)
			return
		case r.Method == http.MethodGet && r.URL.Path == "/gog/1702886318":
			_, _ = w.Write(singleJson)
			return
		case r.Method == http.MethodPost && r.URL.Path == "/hltb/search":
			_, _ = w.Write(arrayJson)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error":"not found"}`))
			return
		}
	}))
}

func Test_GetByHltbId(t *testing.T) {
	singleJson := readTestFile(t, "game_single.json")
	arrayJson := readTestFile(t, "game_array.json")
	server := newTestServer(t, singleJson, arrayJson)
	defer server.Close()

	c := New(WithBaseUrl(server.URL))

	game, err := c.GetByHltbId(context.Background(), 55663)
	if err != nil {
		t.Fatalf("GetByHltbId error: %v", err)
	}
	if game == nil {
		t.Fatalf("GetByHltbId returned nil game")
	}
	if game.Id != 1 {
		t.Fatalf("unexpected game id: got %d want %d", game.Id, 1)
	}
}

func Test_RefreshByHltbId(t *testing.T) {
	singleJson := readTestFile(t, "game_single.json")
	arrayJson := readTestFile(t, "game_array.json")
	server := newTestServer(t, singleJson, arrayJson)
	defer server.Close()

	c := New(WithBaseUrl(server.URL))

	game, err := c.RefreshByHltbId(context.Background(), 55663)
	if err != nil {
		t.Fatalf("RefreshByHltbId error: %v", err)
	}
	if game == nil {
		t.Fatalf("RefreshByHltbId returned nil game")
	}
	if game.Id != 1 {
		t.Fatalf("unexpected game id: got %d want %d", game.Id, 1)
	}
}

func Test_SearchByGameTitle(t *testing.T) {
	singleJson := readTestFile(t, "game_single.json")
	arrayJson := readTestFile(t, "game_array.json")
	server := newTestServer(t, singleJson, arrayJson)
	defer server.Close()

	c := New(WithBaseUrl(server.URL))

	games, err := c.SearchByGameTitle(context.Background(), "Trails of cold Steel III", nil)
	if err != nil {
		t.Fatalf("SearchByGameTitle error: %v", err)
	}
	if len(games) != 1 {
		t.Fatalf("unexpected games count: got %d want %d", len(games), 1)
	}
	if games[0].Id != 1 {
		t.Fatalf("unexpected game id: got %d want %d", games[0].Id, 1)
	}
}

func Test_GetBySteamAppId(t *testing.T) {
	singleJson := readTestFile(t, "game_single.json")
	arrayJson := readTestFile(t, "game_array.json")
	server := newTestServer(t, singleJson, arrayJson)
	defer server.Close()

	c := New(WithBaseUrl(server.URL))

	game, err := c.GetBySteamAppId(context.Background(), 991270)
	if err != nil {
		t.Fatalf("GetBySteamAppId error: %v", err)
	}
	if game == nil {
		t.Fatalf("GetBySteamAppId returned nil game")
	}
	if game.Id != 1 {
		t.Fatalf("unexpected game id: got %d want %d", game.Id, 1)
	}
}

func Test_GetByGogAppId(t *testing.T) {
	singleJson := readTestFile(t, "game_single.json")
	arrayJson := readTestFile(t, "game_array.json")
	server := newTestServer(t, singleJson, arrayJson)
	defer server.Close()

	c := New(WithBaseUrl(server.URL))

	game, err := c.GetByGogAppId(context.Background(), 1702886318)
	if err != nil {
		t.Fatalf("GetByGogAppId error: %v", err)
	}
	if game == nil {
		t.Fatalf("GetByGogAppId returned nil game")
	}
	if game.Id != 1 {
		t.Fatalf("unexpected game id: got %d want %d", game.Id, 1)
	}
}
