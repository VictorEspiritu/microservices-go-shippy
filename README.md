### Microservices in Go: shippy

#### Consignment Microservices
- Workdir: /
- `docker build -t shippy-service-consignment -f shippy-service-consignment/Dockerfile  `
- `docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment `

#### Vessel Microservices
- Workdir: /shippy-service-vessel
- `docker build -t shippy-service-vessel .`
- `docker run -p 50052:50052 -e MICRO_SERVER_ADDRESS=:50052 shippy-service-consignment `

#### Client Consignment Microservices
- Workdir /
- `docker build -t shippy-cli-consignment  -f shippy-client-consignment/Dockerfile . `
- `docker run -p 50053:50053 -e MICRO_SERVER_ADDRESS=:50053 shippy-cli-consignment`

