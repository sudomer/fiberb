FROM golang:1.20.10-bullseye as builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

ENV DATABASE_URI=${DATABASE_URI}
ENV EXPOSE_PORT=${EXPOSE_PORT}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.Version=v0.0.1'" -o /build_amd64 .

FROM gcr.io/distroless/static-debian9

COPY --from=builder /build_amd64 .

EXPOSE 8080

CMD [ "/build_amd64" ]