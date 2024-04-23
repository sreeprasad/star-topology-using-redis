## Star Topology using Redis

## How to run

Start the Redis by running docker

```shell
docker-compose up -- build
```

Start the subscriber

```shell
cd subscriber
go build
./subscriber
```

Start the publishier

```shell
cd publisher
go build
./publisher
```

## Publisher

Fetch top 10 posts from `r/space/top`
and extracts the title from those posts
and publishes to Redis channel `reddit_space`

## Subscriber

As soon as publisher publishes the post,
subscriber running will subscribe to the redis
channel `reddit_space` and get the post title
and prints

As Redis does not store the data, if any subscriber is not
not subscribed to redis channel `reddit_space`, then it willnto receive that
message. This means older messages `will not` be printed.

## Screenshots

![star topology](screenshots/star-topology.jpg)
