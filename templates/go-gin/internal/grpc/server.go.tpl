{{- if .WithGRPC}}
package grpc

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"{{.ProjectName}}/internal/services"
	pb "{{.ProjectName}}/proto"

	"google.golang.org/grpc"
)

// Server implements the gRPC server
type Server struct {
	pb.Unimplemented{{.DomainTitle}}ServiceServer
	{{.DomainLower}}Service *services.{{.DomainTitle}}Service
}

// StartServer starts the gRPC server
func StartServer(port string, db *sql.DB) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", port, err)
	}

	s := grpc.NewServer()
	
	{{.DomainLower}}Service := services.New{{.DomainTitle}}Service(db)
	server := &Server{
		{{.DomainLower}}Service: {{.DomainLower}}Service,
	}
	
	pb.Register{{.DomainTitle}}ServiceServer(s, server)

	log.Printf("gRPC server listening on port %s", port)
	return s.Serve(lis)
}
{{- else}}
package grpc

import (
	"database/sql"
)

// StartServer is a no-op when gRPC is not enabled
func StartServer(port string, db *sql.DB) error {
	// No-op
	return nil
}
{{- end}}
