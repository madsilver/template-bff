# template-bff

The project's architecture is based on [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

```sh
curl --location 'http://localhost:8080/query' \
--header 'x-roles: admin' \
--header 'Content-Type: application/json' \
--data '{"query":"query { health{ message, status_code }}"}'
```
