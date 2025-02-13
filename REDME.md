# Spy Cat Agency

## Prerequisites

Ensure you have the following installed on your system:
- Go
- Bash
- Docker
- Basic GNU utilities (e.g., `make`, `curl`)

## Setup

Follow these steps to set up and start the application:

1. **Clone the repository**:

    ```bash
    git clone https://github.com/Oleksandr-Stetsyshyn/spy-cat-agency.git
    cd spy-cat-agency
    ```

2. **Start the database and the application**:

    ```bash
    docker-compose up -d
    go run main.go
    ```

## Testing the Application

Once the application is running, you can test it by sending requests to the endpoints. For example, to get a list of cats:

```bash
curl http://localhost:8080/cats