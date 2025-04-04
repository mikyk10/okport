# ---------- Stage 1: Build ----------
FROM golang:1.24 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o okport -ldflags="-s -w" main.go

# ---------- Stage 2: Minimal runtime ----------
FROM scratch

# Copy binary only
COPY --from=build /app/okport /okport

ENTRYPOINT ["/okport"]