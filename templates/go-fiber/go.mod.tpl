module {{.ProjectName}}

go 1.25.1

require (
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/lib/pq v1.10.9
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
	{{- if .WithGRPC}}
	google.golang.org/grpc v1.66.0
	google.golang.org/protobuf v1.34.2
	{{- end}}
	{{- if .GCPProject}}
	cloud.google.com/go/monitoring v1.20.3
	{{- end}}
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.55.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
