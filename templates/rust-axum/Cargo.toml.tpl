[package]
name = "{{.ProjectName}}"
version = "0.1.0"
edition = "2021"

[dependencies]
tokio = { version = "1", features = ["rt-multi-thread", "macros"] }
axum = "0.7"
hyper = { version = "1", features = ["full"] }
tower = "0.4"
tower-http = { version = "0.5", features = ["trace", "cors"] }
tracing = "0.1"
tracing-subscriber = { version = "0.3", features = ["fmt", "env-filter"] }
serde = { version = "1", features = ["derive"] }
serde_json = "1"
thiserror = "1"
{{- if .WithGRPC }}
tonic = { version = "0.12", features = ["transport"] }
prost = "0.13"
prost-types = "0.13"
{{- end }}

[build-dependencies]
{{- if .WithGRPC }}
tonic-build = "0.12"
{{- end }}
