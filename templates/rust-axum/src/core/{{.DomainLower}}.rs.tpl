use serde::{Serialize, Deserialize};

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct {{.DomainTitle}} {
    pub id: u64,
    pub name: String,
}

impl {{.DomainTitle}} {
    pub fn new(id: u64, name: String) -> Self {
        Self { id, name }
    }
}
