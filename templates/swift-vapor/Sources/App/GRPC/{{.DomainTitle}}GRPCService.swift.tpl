#if canImport(GRPC)
import GRPC
import NIO
import Foundation

// This is a placeholder for your generated gRPC service implementation.
// After generating Swift code from Proto/{{.DomainLower}}.proto using protoc with grpc-swift,
// replace this file or extend the generated provider to hook into your domain Service.
final class {{.DomainTitle}}GRPCService /*: Generated{{.DomainTitle}}ServiceProvider*/ {
    // let service: {{.DomainTitle}}Service
    // init(service: {{.DomainTitle}}Service) { self.service = service }
}
#endif
