# Simple Rate Limit

### Introduction
The project used to implement simple rate limit project.

Used techonologies:
- Use `echo` as web service framework to handle http request.
- Use `redis` as cache database to calculate request limit.
- Use `viper` as config component.
- Use `logrus` as log component.
- Use `dig` as dependency injection component.
- Use `dockertest` and `httpexpect` as test utility for integration testing.


### Rule
- Server can only accept 60 requests per minute each IP.
- It will response the text of N if N <= 60.
- It will response the text "Error" if N > 60.

### Demo
The project is hosted in [Heroku](https://young-wave-60838.herokuapp.com/) and uses `Heroku Redis` as database.
It's easy to send request by `curl -X POST https://young-wave-60838.herokuapp.com/post` to test the function if works.

### Test (For Mac)
Required: Docker Desktop, Go.
Execute `go test -run=. -v ./tests` to see the test result.

The detail are listed below:
- Use `dockertest` to run docker `redis` container as database.
- Start running http server for integration testing.
- Use `httpexpect` to check responses if correct.
    - First, call `/post` 60 times, it must be responsed the text of requested count with http status `200`.
    - Then, call `/post` once, it must be responsed the text "Error" with http status `429`.
- Stop running http server and purge `redis` container.