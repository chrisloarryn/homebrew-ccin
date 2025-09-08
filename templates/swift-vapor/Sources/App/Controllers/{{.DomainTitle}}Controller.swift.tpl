import Vapor

struct {{.DomainTitle}}Controller {
    let service: {{.DomainTitle}}Service

    func boot(routes: RoutesBuilder) throws {
        let group = routes.grouped("{{.DomainLower}}")
        group.get(use: index)
        group.get(":id", use: get)
        group.post(use: create)
        group.put(":id", use: update)
        group.delete(":id", use: delete)
    }

    func index(req: Request) async throws -> [{{.DomainTitle}}] {
        try await service.list(req: req)
    }

    func get(req: Request) async throws -> {{.DomainTitle}} {
        guard let id = req.parameters.get("id", as: UUID.self) else {
            throw Abort(.badRequest, reason: "Invalid ID")
        }
        return try await service.get(id: id, req: req)
    }

    func create(req: Request) async throws -> {{.DomainTitle}} {
        let dto = try req.content.decode({{.DomainTitle}}.self)
        return try await service.create(dto, req: req)
    }

    func update(req: Request) async throws -> {{.DomainTitle}} {
        guard let id = req.parameters.get("id", as: UUID.self) else {
            throw Abort(.badRequest, reason: "Invalid ID")
        }
        let dto = try req.content.decode({{.DomainTitle}}.self)
        return try await service.update(id: id, dto, req: req)
    }

    func delete(req: Request) async throws -> HTTPStatus {
        guard let id = req.parameters.get("id", as: UUID.self) else {
            throw Abort(.badRequest, reason: "Invalid ID")
        }
        try await service.delete(id: id, req: req)
        return .noContent
    }
}
