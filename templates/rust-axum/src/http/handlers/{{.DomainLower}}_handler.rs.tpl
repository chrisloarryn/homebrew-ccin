use axum::{Json, http::StatusCode};
use serde::{Serialize, Deserialize};

use crate::services::{{.DomainLower}}_service::{{.DomainTitle}}Service;
use crate::core::{{.DomainLower}}::{{.DomainTitle}};

#[derive(Deserialize)]
pub struct Create{{.DomainTitle}}Request {
    pub name: String,
}

#[derive(Serialize)]
pub struct ApiResponse<T> {
    pub data: T,
}

pub async fn list() -> Json<ApiResponse<Vec<{{.DomainTitle}}>>> {
    let items = {{.DomainTitle}}Service::list();
    Json(ApiResponse { data: items })
}

pub async fn create(Json(req): Json<Create{{.DomainTitle}}Request>) -> (StatusCode, Json<ApiResponse<{{.DomainTitle}}>>) {
    let item = {{.DomainTitle}}Service::create(req.name);
    (StatusCode::CREATED, Json(ApiResponse { data: item }))
}
