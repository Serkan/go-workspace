package httpserver

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func TestWriteHeader(t *testing.T) {
	writer := &StringWriter{}
	writeHeader(writer, "Content-Type", "text/html")
	str := writer.String()
	exp, _ := regexp.Compile("[A-z]+\\-??[A-z]+\\:.*")
	if !exp.Match([]byte(str)) {
		t.Fatalf("Written header does not match with HTTP header regex, header: %s", str)
	}
}

func TestWriteLine(t *testing.T) {
	writer := &StringWriter{}
	writeLine(writer, "status line")
	str := writer.String()
	// make sure it ends with end-line char "\r\n"
	if !strings.HasSuffix(str, "\r\n") {
		t.Fatal("It does not end with end-line char")
	}
}

func TestLookupWebResourceWildCard(t *testing.T) {
	abs, _ := filepath.Abs("../../server_pages/505.html")
	srv := NewServerBuilder().AddResource(".*", abs).Build("localhost", 8080)
	wr := srv.lookupWebResource("anything")
	if wr.location != abs || wr.contentType != "text/html" {
		t.Fatal("Assertion failed")
	}
}

func TestLookupWebResourceExact(t *testing.T) {
	abs, _ := filepath.Abs("../../server_pages/505.html")
	srv := NewServerBuilder().AddResource("505.html", abs).Build("localhost", 8080)
	wr := srv.lookupWebResource("505.html")
	if wr.location != abs || wr.contentType != "text/html" {
		t.Fatal("Assertion failed")
	}
}

func TestServe(t *testing.T) {
	/*abs, _ := filepath.Abs("../../server_pages/505.html")
	exp, _ := regexp.Compile(".*")
	wr := &WebResource{
		pattern:     exp,
		location:    abs,
		contentType: "text/html",
	}
	writerActual := &StringWriter{}
	wr.serve(writerActual)
	actual := writerActual.String()
	f, _ := os.Open(abs)
	defer f.Close()
	writerExpected := &StringWriter{}
	io.Copy(writerExpected, f)
	expected := writerExpected.String()
	if actual != expected {
		t.Fatal("Served content does not match ")
	}
	*/
	// TODO should parse HTTP response and match headers and content and date (approximatley)
}
