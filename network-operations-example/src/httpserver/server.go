package httpserver

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// WebResource servable resource representation
type WebResource struct {
	pattern     *regexp.Regexp
	location    string
	contentType string
	errorFlag   bool
}

// ResourceServer anything can be written to io.Writer comes from http connection
type ResourceServer interface {
	Serve(writer *io.Writer) (err error)
}

// Request parsed and structured HTTP request
type Request struct {
	requestedResource string
}

// Server representation of server instance ready to run
type Server struct {
	ip              string
	port            int
	resourceMapping []*WebResource
}

// ServerBuilder builder for server instance
type ServerBuilder interface {
	AddResource(pattern string, location string) ServerBuilder
	Build(ip string, port int) *Server
}

// AddResource creates a pattern mapping to specific static file on the machine, location is the absolute path of static resource
func (srv *Server) AddResource(pattern string, location string) ServerBuilder {
	exp, _ := regexp.Compile(pattern)
	contentType := convertToContentType(location)
	wr := &WebResource{
		pattern:     exp,
		location:    location,
		contentType: contentType,
	}
	srv.resourceMapping = append(srv.resourceMapping, wr)
	return srv
}

// Build creates a server instance which is ready to be run
func (srv *Server) Build(ip string, port int) (httpserver *Server) {
	srv.ip = ip
	srv.port = port
	return srv
}

// Start starts the server based on ip:port configuration and resource mappings
func (srv *Server) Start() (err error) {
	ln, err := net.Listen("tcp", srv.ip+":"+strconv.Itoa(srv.port))
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go srv.handeConnection(conn)
	}
}

// NewServer creates a new server instance but wont start it
func NewServerBuilder() (srvBuilder ServerBuilder) {
	srv := &Server{}
	p405, _ := filepath.Abs("../../server_pages/405.html")
	p505, _ := filepath.Abs("../../server_pages/505.html")
	srv.AddResource("405_method_not_allowed.error", p405)
	srv.AddResource("505_http_version_not_supported.error", p505)
	return srv
}

func convertToContentType(location string) string {
	lastIndex := strings.LastIndex(location, ".")
	ext := location[lastIndex+1 : len(location)]
	switch ext {
	case "html":
		return "text/html"
	case "htm":
		return "text/html"
	case "js":
		return "application/javascript"
	default:
		return "unknown"
	}
}

func (srv *Server) handeConnection(con net.Conn) {
	req := parse(con)
	res := req.requestedResource
	webres := srv.lookupWebResource(res)
	webres.serve(con)
}

func (resource WebResource) serve(writer io.Writer) (err error) {
	if resource.errorFlag { // TODO remove from here, this is an server error case not HTTP 404 maybe HTTP 5XX
		writeLine(writer, "HTTP/1.1 404 NOT FOUND")
		writeLine(writer, "<h1>404 NOT FOUND</h1>")
		return
	}
	// open resource file
	f, err := os.Open(resource.location)
	defer f.Close()
	info, _ := f.Stat()
	// preapre response headers
	writeLine(writer, "HTTP/1.1 200 OK")
	writeHeader(writer, "Content-Type", resource.contentType)
	writeHeader(writer, "Accept-Ranges", "bytes")
	writeHeader(writer, "Content-Length", strconv.FormatInt(info.Size(), 10))
	writeHeader(writer, "Date", currentDateForHeader())
	writeLine(writer, "")
	// io.Copy to writer
	io.Copy(writer, f)
	return nil
}

func (srv *Server) lookupWebResource(resourcename string) (res *WebResource) {
	// lookup from resource map for pattern and return web resource
	for _, wr := range srv.resourceMapping {
		if wr.pattern.Match([]byte(resourcename)) {
			return wr
		}
	}
	abs, _ := filepath.Abs("../../404.html")
	return &WebResource{
		location:    abs,
		contentType: "text/html",
		errorFlag:   true,
	}
}

func parse(reader io.Reader) (req *Request) {
	bufreader := bufio.NewReader(reader)
	line, err := bufreader.ReadString('\n') // fetch only first line
	if err != nil {
		return nil
	}
	columns := strings.Split(line, " ")
	if columns[0] != "GET" { // check method
		return &Request{
			requestedResource: "405_method_not_allowed.error",
		}
	}
	if columns[2] != "HTTP/1.1\r\n" { // check version
		return &Request{
			requestedResource: "505_http_version_not_supported.error",
		}
	}
	uri, err := url.Parse(columns[1])
	if err != nil {
		return nil
	}
	return &Request{
		requestedResource: uri.EscapedPath()[1:],
	}
}

func writeLine(writer io.Writer, line string) {
	writer.Write([]byte(line))
	writer.Write([]byte("\r\n"))
}

func writeHeader(writer io.Writer, headername string, headervalue string) {
	writer.Write([]byte(headername))
	writer.Write([]byte(": "))
	writer.Write([]byte(headervalue))
	writer.Write([]byte("\r\n"))
}

func currentDateForHeader() string {
	t := time.Now()
	// Date: <day-name>, <day> <month> <year> <hour>:<minute>:<second> GMT
	return t.Weekday().String() + "," + strconv.Itoa(t.Day()) + " " + t.Month().String() + " " + strconv.Itoa(t.Year()) + " " +
		strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second()) + " " + "GMT"
}

// StringWriter an implementation of io.Writer, it hold its buffer in-memory
type StringWriter struct {
	buffer bytes.Buffer
}

func (strwriter *StringWriter) Write(b []byte) (n int, err error) {
	written, err := strwriter.buffer.WriteString(string(b))
	if err != nil {
		return 0, err
	}
	return written, nil
}

func (strwriter StringWriter) String() string {
	return strwriter.buffer.String()
}
