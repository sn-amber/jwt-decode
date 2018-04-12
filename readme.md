# JWT-Decode

[![Build Status](https://travis-ci.org/snamber/jwt-decode.svg?branch=master)](https://travis-ci.org/snamber/jwt-decode)

JWT-Decode is a json web token decoder

# Installation

Get the tool via

```sh
go get github.com/snamber/jwt-decode
```

# Usage

Use either by piping a token into `jwt-decode`

```sh
echo eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMwOTg2MTMsIk5hbWUiOiJ0ZXN0IiwiS2luZCI6MH0.moBp8MqKCi | jwt-decode
```

or by specifying the token as command line parameter

```sh
jwt-decode -token eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMwOTg2MTMsIk5hbWUiOiJ0ZXN0IiwiS2luZCI6MH0.moBp8MqKCi
```

both will print the following

![{
  "alg": "RS256",
  "typ": "JWT"
}
{
  "Kind": 0,
  "Name": "test",
  "exp": 1.523098613e+09
}](assets/without_datetime.png)

Use the flag `-datetime` to print the `exp` expiration date in datetime format:

![{
  "alg": "RS256",
  "typ": "JWT"
}
{
  "Kind": 0,
  "Name": "test",
  "exp": "2018-04-07T12:56:53+02:00"
}
](assets/with_datetime.png)
