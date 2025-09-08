import Foundation
import Vapor

struct {{.DomainTitle}}: Content, Equatable, Codable, Identifiable {
    var id: UUID?
    var name: String
    var createdAt: Date?
    var updatedAt: Date?

    init(id: UUID? = nil, name: String, createdAt: Date? = nil, updatedAt: Date? = nil) {
        self.id = id
        self.name = name
        self.createdAt = createdAt
        self.updatedAt = updatedAt
    }
}
