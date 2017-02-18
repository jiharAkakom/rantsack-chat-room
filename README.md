# Rantsack Chat Room

Simple chat room functionality using Golang.

Rantsack is a play on the words Rant (speak or shout at length in a wild, impassioned way) and Ransack (go hurriedly through (a place) stealing things and causing damage).

## Running

Install `go`.

Run with `go run MAIN.go`.

Access the chat room by going to `localhost:8080` in a browser.

*More detailed [setup documentation](docs/SETUP.md) is available on setting up the chat service.*

## To Do

- Push change to clients instead of making clients poll.
- Set up Redis store.
- Put messages in store.
- Store user connections.
- Identify messages by connection.
- Show users present in room.
- Dockerize.
- Sub chat rooms.
- Single sign on.

