use std::net::SocketAddr;

use axum::{routing::{get, post}, Router};
use tower_http::trace::TraceLayer;
use tracing_subscriber::EnvFilter;

mod http;
mod core;
mod services;
mod middleware;
{{- if .WithGRPC }}
mod grpc;
{{- end }}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Tracing
    let filter = EnvFilter::try_from_default_env().unwrap_or_else(|_| EnvFilter::new("info"));
    tracing_subscriber::fmt().with_env_filter(filter).init();

    // Build router
    let app = http::routes::create_router()
        .layer(TraceLayer::new_for_http());

    // HTTP address
    let port: u16 = std::env::var("PORT")
        .ok()
        .and_then(|s| s.parse().ok())
        .unwrap_or(8080);
    let http_addr: SocketAddr = ([0, 0, 0, 0], port).into();

    let http_task = async move {
        let listener = tokio::net::TcpListener::bind(http_addr).await?;
        tracing::info!("HTTP listening on {}", http_addr);
        axum::serve(listener, app).await?;
        Ok::<(), Box<dyn std::error::Error>>(())
    };

    {{- if .WithGRPC }}
    use tonic::transport::Server;
    // gRPC address
    let grpc_port: u16 = std::env::var("GRPC_PORT")
        .ok()
        .and_then(|s| s.parse().ok())
        .unwrap_or(50051);
    let grpc_addr: SocketAddr = ([0, 0, 0, 0], grpc_port).into();

    let grpc_task = async move {
        let svc = grpc::service::{{.DomainTitle}}GrpcService::default();
        let server = grpc::service::pb::{{.DomainTitle}}ServiceServer::new(svc);
        tracing::info!("gRPC listening on {}", grpc_addr);
        Server::builder()
            .add_service(server)
            .serve(grpc_addr)
            .await?;
        Ok::<(), Box<dyn std::error::Error>>(())
    };

    tokio::try_join!(http_task, grpc_task)?;
    {{- else }}
    http_task.await?;
    {{- end }}

    Ok(())
}
