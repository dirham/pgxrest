# PgXRest
Is package (or server) that will turn your postgres database into RESTful API without efforts. This project is highly inspired by [PostgREST](https://postgrest.org/).

## Todo
- [ ] Support more query/filters, like field select and other filters
- [ ] Support relationship on query
- [ ] Json response directly from postgres
- [ ] Add Authentication and authorization on endpoint layer
- [ ] Add Example
- [ ] Complete Unit Tests
---
### Used Packages
1. **PostgreSQL driver [PGX](https://github.com/jackc/pgx)**<br/>
`pgx` is a pure Go driver and toolkit for PostgreSQL. is a low-level, high performance interface that exposes PostgreSQL-specific features such as LISTEN / NOTIFY and COPY

2. HTTP Router [Go-Chi](https://github.com/jackc/pgx) <br/>
`chi` is a lightweight, idiomatic and composable router for building Go HTTP services. It's especially good at helping you write large REST API services that are kept maintainable as your project grows and changes. It also support `net/http` package that in my opinion is best choice to go.

### License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
