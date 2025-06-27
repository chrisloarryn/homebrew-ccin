module {{.ProjectName}}

go 1.24.4

require (
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/lib/pq v1.10.9
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
	{{- if .WithGRPC}}
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
	{{- end}}
	{{- if .GCPProject}}
	cloud.google.com/go/monitoring v1.19.0
	{{- end}}
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
