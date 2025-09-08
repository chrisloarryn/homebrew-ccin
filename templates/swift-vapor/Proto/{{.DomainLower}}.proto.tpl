syntax = "proto3";
package {{.ProjectName}}.{{.DomainLower}};

// Basic example proto for {{.DomainTitle}} service
service {{.DomainTitle}}Service {
  rpc List({{.DomainTitle}}ListRequest) returns ({{.DomainTitle}}ListResponse);
  rpc Get({{.DomainTitle}}GetRequest) returns ({{.DomainTitle}}Item);
  rpc Create({{.DomainTitle}}Item) returns ({{.DomainTitle}}Item);
  rpc Update({{.DomainTitle}}Item) returns ({{.DomainTitle}}Item);
  rpc Delete({{.DomainTitle}}DeleteRequest) returns ({{.DomainTitle}}DeleteResponse);
}

message {{.DomainTitle}}Item {
  string id = 1;
  string name = 2;
}

message {{.DomainTitle}}ListRequest {}
message {{.DomainTitle}}ListResponse { repeated {{.DomainTitle}}Item items = 1; }
message {{.DomainTitle}}GetRequest { string id = 1; }
message {{.DomainTitle}}DeleteRequest { string id = 1; }
message {{.DomainTitle}}DeleteResponse { bool success = 1; }
