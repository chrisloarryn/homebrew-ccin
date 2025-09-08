// swift-tools-version:6.1
import PackageDescription

let package = Package(
    name: "{{.ProjectName}}",
    platforms: [
        .macOS(.v13)
    ],
    products: [
        .library(name: "App", targets: ["App"]),
        .executable(name: "Run", targets: ["Run"]) 
    ],
    dependencies: [
        .package(url: "https://github.com/vapor/vapor.git", from: "4.87.0"),
        {{- if .WithGRPC}}
        .package(url: "https://github.com/grpc/grpc-swift.git", from: "1.20.2"),
        {{- end}}
    ],
    targets: [
        .target(
            name: "App",
            dependencies: [
                .product(name: "Vapor", package: "vapor")
                {{- if .WithGRPC}},
                .product(name: "GRPC", package: "grpc-swift")
                {{- end}}
            ],
            path: "Sources/App"
        ),
        .executableTarget(
            name: "Run",
            dependencies: ["App"],
            path: "Sources/Run"
        )
    ]
)
