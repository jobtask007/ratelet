FROM golang:1.24.3-bookworm AS build

WORKDIR /app

COPY . .

ENV GOTOOLCHAIN=local

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ratelet cmd/main.go


FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=build /app/ratelet /
CMD ["/ratelet"]