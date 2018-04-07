# JWT-Decode

[![Build Status](https://travis-ci.org/snamber/jwt-decode.svg?branch=master)](https://travis-ci.org/snamber/jwt-decode)

JWT-Decode is a json web token decoder

# Usage

Get the tool via

```sh
go get github.com/snamber/jwt-decode
```

Use either by piping a token into `jwt-decode`

```sh
echo eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMwOTg2MTMsIk5hbWUiOiJ0ZXN0IiwiS2luZCI6MH0.moBp8MqKCi | jwt-decode
```

or by specifying the token as command line parameter

```sh
./jwt-decode -token eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMwOTg2MTMsIk5hbWUiOiJ0ZXN0IiwiS2luZCI6MH0.moBp8MqKCi
```

both will print the following

```json
{
  "alg": "RS256",
  "typ": "JWT"
}
{
  "Kind": 0,
  "Name": "test",
  "exp": 1.523098613e+09
}
```
