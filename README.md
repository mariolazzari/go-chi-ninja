# Go Chi Ninja

YouTube [playlist](https://www.youtube.com/watch?v=wpnN3RIRSxs&t=103s)

## Getting started

### goenv

Docs [page](https://pkg.go.dev/github.com/drewgonzales360/goenv)

```sh
brew install goenv
echo 'eval "$(goenv init -)"' >> ~/.zshrc
```

### Project creation

```sh
go mod init github.com/mariolazzari/go-chi-ninja
```

### Chi

Github [repository](https://github.com/go-chi/chi)

### Go get

```sh
go get github.com/go-chi/chi/v5
```

### Go tidy

```sh
go mod tidy
```

### Redis

Go Redix [client](https://github.com/redis/go-redis)

```sh
go get github.com/redis/go-redis/v9
docker run -p 6379:6379 redis:latest 
```
