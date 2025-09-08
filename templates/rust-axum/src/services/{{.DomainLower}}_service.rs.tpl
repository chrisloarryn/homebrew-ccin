use crate::core::{{.DomainLower}}::{{.DomainTitle}};

pub struct {{.DomainTitle}}Service;

impl {{.DomainTitle}}Service {
    pub fn list() -> Vec<{{.DomainTitle}}> {
        vec![{{.DomainTitle}}::new(1, format!("sample-{{.DomainLower}}"))]
    }

    pub fn create(name: String) -> {{.DomainTitle}} {
        {{.DomainTitle}}::new(1, name)
    }
}
