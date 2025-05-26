# Go RabbitMQ Playground

A simple Go module to help me learn and experiment with message brokers using [RabbitMQ](https://www.rabbitmq.com/).

> This is a learning project — not production-ready code.

---

## Goals

- Understand basic message broker concepts (queues, exchanges, routing, etc.)
- Learn how to use RabbitMQ with Go
- Experiment with publishing and consuming messages

---

## Requirements

- [Go](https://golang.org/dl/) 1.20+
- [RabbitMQ](https://www.rabbitmq.com/download.html) (can be run locally via Docker)

---

# How to Test

1. run consumer app first
2. after consumer started, run producer app

this will make consumer "asking" for message, after consumer asked, we run producer to serve the message