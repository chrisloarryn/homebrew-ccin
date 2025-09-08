import Vapor

struct MetricsMiddleware: AsyncMiddleware {
    func respond(to request: Request, chainingTo next: AsyncResponder) async throws -> Response {
        let start = Date()
        let response = try await next.respond(to: request)
        let duration = Date().timeIntervalSince(start)
        request.logger.info("[Metrics] \(request.method.string) \(request.url.path) -> \(response.status.code) in \(String(format: "%.3f", duration))s")
        return response
    }
}
