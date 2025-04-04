# ---------- Stage 1: Build ----------
    FROM golang:1.24-alpine AS build

    WORKDIR /app
    COPY . .
    RUN go build -o okport main.go
    
    # ---------- Stage 2: Minimal runtime ----------
    FROM alpine:latest
    
    WORKDIR /root/
    
    # Copy binary only (no source code)
    COPY --from=build /app/okport .
    
    # Expose no ports explicitly (up to user via -p)
    ENTRYPOINT ["./okport"]