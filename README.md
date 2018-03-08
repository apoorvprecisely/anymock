# anymock
Go Binary to mock any GET/POST APIs on a given port

### Install :

`go get -u github.com/apoorvprecisely/anymock`

`go install github.com/apoorvprecisely/anymock`

### Usage : 

```
anymock --config /path/of/config/file

```

### Sample Config File :

```
{
  "object": {
    "apiData": [
      {
        "url": "/abc/xyz",
        "response": "Success"
      }
    ],
    "port": "1234"
  }
}
```
