import Vapor

final class {{.DomainTitle}}Service {
    private var storage: [UUID: {{.DomainTitle}}] = [:]

    init() {}

    func list(req: Request) async throws -> [{{.DomainTitle}}] {
        Array(storage.values)
    }

    func get(id: UUID, req: Request) async throws -> {{.DomainTitle}} {
        guard let item = storage[id] else {
            throw Abort(.notFound, reason: "{{.DomainTitle}} not found")
        }
        return item
    }

    func create(_ dto: {{.DomainTitle}}, req: Request) async throws -> {{.DomainTitle}} {
        var entity = dto
        let id = dto.id ?? UUID()
        entity.id = id
        entity.createdAt = Date()
        storage[id] = entity
        return entity
    }

    func update(id: UUID, _ dto: {{.DomainTitle}}, req: Request) async throws -> {{.DomainTitle}} {
        guard var existing = storage[id] else {
            throw Abort(.notFound, reason: "{{.DomainTitle}} not found")
        }
        existing.name = dto.name
        existing.updatedAt = Date()
        storage[id] = existing
        return existing
    }

    func delete(id: UUID, req: Request) async throws {
        guard storage.removeValue(forKey: id) != nil else {
            throw Abort(.notFound, reason: "{{.DomainTitle}} not found")
        }
    }
}
