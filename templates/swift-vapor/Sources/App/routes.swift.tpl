import Vapor

public func routes(_ app: Application) throws {
    // Health check
    app.get("health") { req async throws -> String in
        return "OK"
    }

    // API v1 routes
    let api = app.grouped("api", "v1")

    // Domain routes: /api/v1/{{.DomainLower}}
    let {{.DomainLower}}Service = {{.DomainTitle}}Service()
    let {{.DomainLower}}Controller = {{.DomainTitle}}Controller(service: {{.DomainLower}}Service)
    try {{.DomainLower}}Controller.boot(routes: api)
}
