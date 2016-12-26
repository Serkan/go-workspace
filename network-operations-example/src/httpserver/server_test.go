package httpserver

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	abs, _ := filepath.Abs("../../testdata/http_get.txt")
	f, _ := os.Open(abs)
	req := parse(f)
	if req == nil {
		t.Fatal("Resource is nil")
	}
	if req.requestedResource != "index.html" {
		t.Fatalf("Requested resource parsed incorrectly expected 'index.html', actual %s", req.requestedResource)
	}
}

func TestServer(t *testing.T) {
	srv := NewServerBuilder().AddResource("index.html", "/Users/serkan/Desktop/index.html").Build("localhost", 8080)
	srv.Start()
}
