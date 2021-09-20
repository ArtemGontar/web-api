## Get project

```
go get github.com/ArtemGontar/web-api
```

## Contract

### Get Request

```
curl http://localhost:8080/albums
```

### Get Response
``` json
[
  {
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
  },
  {
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
  },
  {
    "id": "3",
    "title": "Sarah Vaughan and Clifford Brown",
    "artist": "Sarah Vaughan",
    "price": 39.99
  }
]
```
### GetById Request

```
curl http://localhost:8080/albums/2
```

### GetById Response
``` json
{
  "id": "2",
  "title": "Jeru",
  "artist": "Gerry Mulligan",
  "price": 17.99
}
```
### Post Request
```
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
### Post Response

HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Wed, 02 Jun 2021 00:34:12 GMT
Content-Length: 116

``` json
{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
```

### Delete Request

```
curl http://localhost:8080/albums/2 \
  --header "Content-Type: application/json" \
  --request "DELETE" \
```

### Delete Response
``` json
{
  "id": "2",
  "title": "Jeru",
  "artist": "Gerry Mulligan",
  "price": 17.99
}
```
