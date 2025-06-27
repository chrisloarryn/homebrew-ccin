PORT={{.Port}}
{{- if .WithGRPC}}
GRPC_PORT=50051
{{- end}}
DATABASE_URL=postgres://localhost/{{.ProjectName}}_dev?sslmode=disable
{{- if .GCPProject}}
GCP_PROJECT={{.GCPProject}}
{{- end}}
