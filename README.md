# Fetch Rewards Receipt Processor Challenge

This repository contains my submission for Fetch Rewards' [Receipt Processor Challenge](https://github.com/fetch-rewards/receipt-processor-challenge) as part of my application for the Backend Software Engineer role.

---

## Prerequisites

Ensure the following tools are installed on your system before running the application:

- Docker
- Docker Compose (v2 or later)

---

## Running the Application

The solution is containerized and can be easily run using Docker Compose. Follow these steps:

1. Start the application by running:

   ```bash
   docker compose up -d
   ```

   The REST API is exposed on port `8080` in the Docker container, and this is mapped to port `8080` on your host machine. If you wish to use a different port, update the mapping in the `docker-compose.yml` file.

2. Stop the application when you're done by running:

   ```bash
   docker compose down
   ```
