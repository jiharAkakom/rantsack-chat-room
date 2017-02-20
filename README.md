# Rantsack Chat Room

Simple chat room functionality using `Golang` and `Docker`.

Rantsack is a play on the words Rant (speak or shout at length in a wild, impassioned way) and Ransack (go hurriedly through (a place) stealing things and causing damage).

## Running

Install `docker` and `docker-compose`.

Launch with `docker-compose up -d`. Stop with `docker-compose stop`.

Chat room is now accessible on port `8080` (localhost:8080 / IP:8080).

## To Do

- Set up Redis store https://github.com/go-redis/redis.
- Put messages in store.
- Store user connections.
- Identify messages by connection.
- Show users present in room.
- Sub chat rooms.
- Single sign on.

