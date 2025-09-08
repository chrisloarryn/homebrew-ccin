import Vapor

public func configure(_ app: Application) throws {
    // Server configuration
    app.http.server.configuration.port = {{.Port}}
    app.http.server.configuration.hostname = "0.0.0.0"

    // Middlewares
    app.middleware.use(FileMiddleware(publicDirectory: app.directory.publicDirectory))
    app.middleware.use(MetricsMiddleware())

    // Routes
    try routes(app)

    {{- if .WithGRPC}}
    // gRPC scaffold (requires generating Swift gRPC code from Proto/{{.DomainLower}}.proto)
    try startGRPCServer(app)
    {{- end}}
}

#if canImport(GRPC)
import GRPC

// Placeholder gRPC bootstrap. Add your generated service providers here.
func startGRPCServer(_ app: Application) throws {
    app.logger.info("gRPC scaffold enabled. TODO: generate service stubs from Proto/{{.DomainLower}}.proto and bind the server.")
    // Example (after generating providers):
    // let group = app.eventLoopGroup
    // let server = Server.insecure(group: group)
    // let provider = {{.DomainTitle}}ServiceProvider() // from generated code
    // _ = try server.withServiceProviders([provider]).bind(host: "0.0.0.0", port: 50051).wait()
}
#else
func startGRPCServer(_ app: Application) throws { /* gRPC not available */ }
#endif
