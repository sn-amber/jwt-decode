# JWT-Decode

JWT-Decode is a json web token decoder

# Usage

Get the tool via

```
go get -i github.com/sn-amber/jwt-decode
```

Use either by piping a token into `jwt-decode`

```
echo ./jwt-decode -token eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMwOTg2MTMsIk5hbWUiOiJ0ZXN0IiwiS2luZCI6MH0.moBp8MqKCi | jwt-decode
```

or by specifying the token as command line parameter

```
./jwt-decode -token eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMwOTg2MTMsIk5hbWUiOiJ0ZXN0IiwiS2luZCI6MH0.moBp8MqKCi
```