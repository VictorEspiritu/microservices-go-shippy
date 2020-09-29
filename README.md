### Microservices in Go: shippy

#### Server Microservices
- Workdir: /shippy-service-consignment
- `docker build -t shippy-service-consignment .  `
- `docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment `

#### Client Microservices
- Workdir /
- `docker build -t shippy-cli-consignment  -f shippy-client-consignment/Dockerfile . `
- `docker run shippy-cli-consignment `


