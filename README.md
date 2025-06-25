# Notification Feature

goal of this repository is to practice design notification feature for multi-tenant microservice architecture

- [ ] request api that create notification across tenant
- [ ] notification on user side can be read or unread
- [ ] basic crud tenant
- [ ] basic crud user

## Stack

- Docker
- Golang
- Kafka
- MongoDB

## Flow

```mermaid
flowchart TD

    client["`Client`"]

    gateway["`App service
    (Golang)`"]

    kafka["`Message Queuing Service
    (Apache Kafka)`"]

    worker["`Worker service
    (Golang)`"]

    db["`Database
    (MongoDB)`"]

    client --> gateway
    subgraph container [Docker Container]
        gateway --> kafka --> worker --> db
        gateway --> db
    end

```
