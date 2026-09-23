# Go Caching Proxy

A simple HTTP caching proxy built with **Go, Gin, Redis, and CLI flags**. It forwards requests to an origin server and caches responses in Redis to serve subsequent requests faster.

## Features

* HTTP caching with Redis
* Cache HIT/MISS detection using `X-Cache`
* Configurable origin server
* Configurable server port
* CLI-based configuration
* Clear Redis cache from the CLI
* Automatic forwarding of requests to the origin server

## Usage

### Start the proxy

```bash
./caching-proxy --port 3000 --origin http://dummyjson.com/products
```

The proxy will be available at:

```text
http://localhost:3000/products
```

### Clear the cache

```bash
./caching-proxy --clear-cache
```

This flushes the Redis database and exits.

### Example

First request:

```text
X-Cache: MISS
```

The proxy fetches the response from the origin and stores it in Redis.

Subsequent requests:

```text
X-Cache: HIT
```

The response is served from Redis.


https://roadmap.sh/projects/caching-server