#
<!--
NOTE FOR AI:
When generating new changelog entries, DO NOT use the 'Unreleased' section.
Instead, create a new section with the version and date as specified by the user (e.g., 'Bump patch', 'Bump minor', 'Bump major').
Use the format: '## [version] - YYYY-MM-DD'.
Always follow the user's instruction for versioning and date.
-->
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.10] - 2025-09-07
### Added
- CHANGELOG.md file following Keep a Changelog best practices.

### Changed
- Updated NestJS template to use Node.js 24.7.0 (alpine3.22) and added `.nvmrc` for version consistency.
- Updated all NestJS, Go Gin, and Go Fiber template dependencies to their latest stable versions.
- Updated Dockerfiles in all templates to match the latest runtime versions.
- Updated EXAMPLES.md to reflect the new Node.js version and dependency changes.

### Fixed
- Fixed version mismatches between documentation and template files.

### Removed
- Deprecated references to older Node.js and library versions in documentation and templates.
