# Soccer Api
List of soccer players in a team.

endpoint :
| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `GET`    | `/teams`                                 | Retrieve all teams.                      |
| `GET`    | `/teams/{id}`                            | Retrieve team {id}.                      |
| `POST`   | `/teams`                                 | Create a new team.                       |
| `GET`    | `/players`                               | Retrieve all players.                    |
| `GET`    | `/players/{id}`                          | Retrieve player {id}.                    |
| `POST`   | `/players`                               | Create a new player.                     |

## Quick Start
### Install Dependencies
```bash
go get ./...
```

### Run The App
```bash
go run main.go
```
### Run The Tests
```bash
go test -v ./...
```
```bash
go test -cover ./...
```

### Postman Collection
[Download file](go-soccer.postman_collection.json)
