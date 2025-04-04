# okport

**okport** is a super lightweight HTTP server written in Go that listens on multiple ports and responds with `200 OK` to all incoming requests.  
It is designed specifically for infrastructure validation during the initial phase of container orchestration systems like Kubernetes, Nomad, or Docker Swarm.

---

## ✨ Features

- ✅ Listens on multiple ports simultaneously
- ✅ Always returns `200 OK` regardless of path or method
- ✅ Zero configuration, no external dependencies
- ✅ Secure by design — minimal and transparent functionality
- ✅ Ideal for load balancer health checks and orchestration readiness probes

---

## 🚀 Usage

### Run with Go

```bash
go run main.go 8080 9090 10080
```

### Build and Run

```bash
go build -o okport main.go
./okport 8080 9090 10080
```

### Docker

```bash
docker build -t okport .
docker run -p 8080:8080 -p 9090:9090 -p 10080:10080 okport 8080 9090 10080
```

---

## 🔧 Use Cases

- Initial deployment of container orchestrators (Kubernetes, Nomad, etc.)
- Multi-port health checks for L4/L7 load balancers
- Connectivity validation for firewalls, security groups, or ingress controllers
- Dummy service endpoints for DNS round-robin or service mesh testing

---

## 🔒 Security

This project follows the principle of least privilege:

- No input is reflected in the response — safe from XSS and injection
- No authentication, state, or cookies
- Only responds with `200 OK` and a plain text body (`OK\n`)
- Not recommended for public exposure without reverse proxy or TLS termination
- Suitable for **internal networking and automation environments**

---

## 🛠 Build Requirements

- Go 1.18 or later
- No third-party libraries

---

## 📄 License

MIT License

---

## 🙌 Contributing

This project is intentionally kept simple.  
Bug reports, suggestions, or questions are welcome via issues or pull requests.
