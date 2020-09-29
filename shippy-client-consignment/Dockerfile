FROM golang:alpine as builder

RUN apk update && apk upgrade && \
       apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=auto

COPY ./shippy-service-consignment/proto/consignment ./shippy-service-consignment/proto/consignment 
COPY ./shippy-client-consignment ./shippy-client-consignment

RUN ls -l
WORKDIR /app/shippy-service-consignment/proto/consignment
RUN go mod download
WORKDIR /app/shippy-client-consignment
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-cli-consignment


#Run Container
FROM alpine:latest 

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD  ./shippy-client-consignment/consignment.json /app/consignment.json
COPY --from=builder /app/shippy-client-consignment/shippy-cli-consignment .

CMD ["./shippy-cli-consignment"]

