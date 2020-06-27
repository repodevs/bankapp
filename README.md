# Simple Bank App build with Golang

This app is follow tutorial from <https://www.blog.duomly.com/golang-course-with-building-a-fintech-banking-app-lesson-1-start-the-project/>

---

## Run (on VSCode)

`go mod vendor`

`go run main.go`

---

## Test Endpoint

tested using _httpie_

### `POST /register` for register a user

```sh
http -v 127.0.0.1:8888/register username=5CharUsername email=valid@email.format password=5CharMinPassword
```

```sh
POST /register HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 71
Content-Type: application/json
Host: 127.0.0.1:8888
User-Agent: HTTPie/1.0.2

{
    "email": "reva@reva.com",
    "password": "revaxz",
    "username": "revania"
}

HTTP/1.1 200 OK
Content-Length: 277
Content-Type: text/plain; charset=utf-8
Date: Sat, 27 Jun 2020 12:01:39 GMT

{
    "data": {
        "Accounts": [
            {
                "Balance": 0,
                "ID": 10,
                "Name": "revania's account"
            }
        ],
        "Email": "reva@reva.com",
        "ID": 10,
        "Username": "revania"
    },
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE1OTMyNjI4OTksInVzZXJfaWQiOjEwfQ.Ok_M3Q8affLhTRS4jOZuCAPE1KHbrcFqAjf1Wg6dxWY",
    "message": "ok"
}
```

### `POST /login` for login a user

```sh
http -v 127.0.0.1:8888/login username=5CharUsername password=5CharMinPassword
```

```sh
HTTP/1.1 200 OK
Content-Length: 273
Content-Type: text/plain; charset=utf-8
Date: Sat, 27 Jun 2020 12:00:23 GMT

{
    "data": {
        "Accounts": [
            {
                "Balance": 9900,
                "ID": 7,
                "Name": "Anton's account"
            }
        ],
        "Email": "Anton@bank.app",
        "ID": 7,
        "Username": "Anton"
    },
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE1OTMyNjI4MjMsInVzZXJfaWQiOjd9.wwJ5tIWqFlQoeNcTv6ExE-S3ntINDaWcQ5waU989Xf8",
    "message": "ok"
}
```

### `GET /user/{id}` for get detail of a user

```sh
export JWT_AUTH_TOKEN=JWTTokenOfUser
http -v --auth-type=jwt 127.0.0.1:8888/user/7
```

```sh
GET /user/7 HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE1OTMyNDI2MzYsInVzZXJfaWQiOjd9.aEugwOF9CLPafnr8Djbc7K_LAbxFOKvUv_B9Uccpyks
Connection: keep-alive
Host: 127.0.0.1:8888
User-Agent: HTTPie/1.0.2


HTTP/1.1 200 OK
Content-Length: 139
Content-Type: text/plain; charset=utf-8
Date: Sat, 27 Jun 2020 11:53:05 GMT

{
    "data": {
        "Accounts": [
            {
                "Balance": 9950,
                "ID": 7,
                "Name": "Anton's account"
            }
        ],
        "Email": "Anton@bank.app",
        "ID": 7,
        "Username": "Anton"
    },
    "message": "ok"
}
```

### `POST /transaction` for create transaction between users

```sh
export JWT_AUTH_TOKEN=JWTTokenOfUser
http -v --auth-type=jwt  127.0.0.1:8888/transaction userid:=7 from:=7 to:=9 amount:=50
```

```sh
POST /transaction HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE1OTMyNDI2MzYsInVzZXJfaWQiOjd9.aEugwOF9CLPafnr8Djbc7K_LAbxFOKvUv_B9Uccpyks
Connection: keep-alive
Content-Length: 47
Content-Type: application/json
Host: 127.0.0.1:8888
User-Agent: HTTPie/1.0.2

{
    "amount": 50,
    "from": 7,
    "to": 9,
    "userid": 7
}

HTTP/1.1 200 OK
Content-Length: 73
Content-Type: text/plain; charset=utf-8
Date: Sat, 27 Jun 2020 11:56:05 GMT

{
    "data": {
        "Balance": 9900,
        "ID": 7,
        "Name": "Anton's account"
    },
    "message": "ok"
}
```
