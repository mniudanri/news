A Go (golang) simple news app APIs project

Available endpoints:

> Post news `\news`.
>
> Get news `\news`

## Prerequisites

1.  go custom environtment [github.com/joho/godotenv](https://github.com/joho/godotenv)

```shell
go get -u github.com/joho/godotenv
```

2.  Postgres library [github.com/lib/pq](https://github.com/lib/pq)

```shell
go get -u github.com/lib/pq
```

3.  Elasticsearch (version 7.x) as search engine
    ES installation can be found [here](https://www.elastic.co/downloads/elasticsearch)

4.  Postgres as DB

## How to implement

1. Install prerequisites
2. Change `sample.env` to `.env`
3. Configure `.env` file

## Example Uses of Endpoint
1. Example json body `POST/news`
```shell
curl --location --request POST 'http://localhost:3000/news' \
--header 'Content-Type: application/json' \
--data-raw '{
	"author": "Author",
	"body": "This is my first content"
}'
```
2. Example get 
```shell
curl --location --request GET 'http://localhost:3000/news?limit=10&page=1'
```

## Testing

Testing go using `go test`

```shell
go test ./...
```

## Contributing

Contributions are most welcome!

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
