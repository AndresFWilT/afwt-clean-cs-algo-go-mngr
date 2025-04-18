# afwt-cs-api-clean-go-mngr

A Go microservice using the Gin web framework and Clean Architecture principles to expose classic computer science algorithms and data structures as RESTful APIs. This project serves as a personal learning tool to relearn and master fundamental CS concepts like sorting, trees, graphs, and dynamic programming while practicing idiomatic Go and scalable API design.

---

## 📚 Table of Contents

- [✨ Features](#-features)
- [📦 Installation](#-installation)
- [🚀 Quick Start](#-quick-start)
- [📁 Project Structure](#-project-structure)
- [🧩 Architecture Overview](#-architecture-overview)
- [📈 Roadmap](#-roadmap)
- [🧠 Algorithms Included](#-algorithms-included)
- [🔧 Future Ideas](#-future-ideas)
- [📄 Contributors](#contributors)

---

## ✨ Features

- ✅ RESTful API endpoints for classic CS algorithms
- ✅ Modular, idiomatic Go code
- ✅ Follows Clean Architecture (hexagonal) principles
- ✅ Great for learning + testing CS fundamentals
- ✅ Easily extendable with new algorithms or data structures

---

## 📦 Installation

```bash
go get github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr
```

---

## 🚀 Quick Start

```go
package main

import "github.com/AndresFWilT/afwt-clean-cs-algo-go-mngr/internal/config"

func main() {
	serverConfigured, port := config.InitializeServer()
	err := serverConfigured.Run(":" + port)
	if err != nil {
		return
	}
}
```

Environment variable:
```
API_PORT=8080
API_BASE_PATH=/api/v1
```

Then make a request:
```bash
curl -X POST http://localhost:8080/api/v1/sort/bubble -H "Content-Type: application/json" -d '{"array": [5, 2, 4, 3, 1]}'
```

---

## 📁 Project Structure

```
.
├── cmd/
│   └── server/               # Entrypoint (optional)
├── internal/
│   ├── config/               # Server + route configuration
│   ├── adapter/
│   │   ├── route/sort/       # Route definition for /sort endpoints
│   │   └── handler/sort/     # HTTP handlers for sorting algorithms
│   ├── usecase/sort/         # Algorithm implementations (BubbleSort)
│   ├── domain/
│   │   ├── model/sort/       # Request/response models
│   │   └── port/sort/        # Sorter interface
├── go.mod
└── main.go
```

---

## 🧩 Architecture Overview

This microservice follows Clean Architecture:

- **Domain Layer (Core Logic)**
    - `port.Sorter`: interface for any sorting algorithm
    - `model.Array`: input struct used across layers

- **Use Case Layer**
    - `BubbleSort`: implementation of bubble sort (adapts to `Sorter`)

- **Adapter Layer**
    - `BubbleSortHandler`: receives requests, delegates to service
    - `BubbleRoute`: defines Gin route group for Bubble Sort

- **Config Layer**
    - Initializes server and wires routes with their use cases

---

## 📈 Roadmap

- [x] Sorting: Bubble Sort
- [ ] Add Merge Sort, Quick Sort
- [ ] Searching: Binary Search
- [ ] Trees: Binary Tree, AVL, Red-Black
- [ ] Graphs: BFS, DFS, Dijkstra
- [ ] Dynamic Programming: Fibonacci, LCS

---

## 🧠 Algorithms Included

- POST `/api/v1/sort/bubble` → sorts array using Bubble Sort

Request:
```json
{
  "array": [5, 3, 2, 4, 1]
}
```

Response:
```json
{
  "data": [1, 2, 3, 4, 5]
}
```

---

## 🔧 Future Ideas

- [ ] Swagger/OpenAPI docs
- [ ] WASM UI or CLI tool to test endpoints
- [ ] Benchmarking suite for each algorithm
- [ ] Add runtime complexity annotations

---

## 📄 Contributors

- Andrés Felipe Wilches Torres (AndresFWilT) [Contact me!](mailto:andresfwilchestdev@gmail.com)
