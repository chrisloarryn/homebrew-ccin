use axum::{routing::{get, post}, Router};

use crate::http::handlers::{{.DomainLower}}_handler;

pub fn create_router() -> Router {
    Router::new()
        .route("/health", get(|| async { "ok" }))
        .route("/api/{{.DomainLower}}", get({{.DomainLower}}_handler::list))
        .route("/api/{{.DomainLower}}", post({{.DomainLower}}_handler::create))
}
