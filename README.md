# Simple Bank App build with Golang

This app is follow tutorial from <https://www.blog.duomly.com/golang-course-with-building-a-fintech-banking-app-lesson-1-start-the-project/>

---

## Run (on VSCode)

`go mod vendor`

`go run main.go`

---

## Test Endpoint

tested using _httpie_

### Register

```sh
http -v 127.0.0.1:8888/register username=5CharUsername email=valid@email.format password=5CharMinPassword
```

### Login

```sh
http -v 127.0.0.1:8888/login username=5CharUsername password=5CharMinPassword
```
