<img alt="Golang" src="https://img.shields.io/badge/Golang-276DC3.svg?logo=go&logoColor=white"> [![Go](https://github.com/joshuaray/go-proxy/actions/workflows/go-build.yml/badge.svg)](https://github.com/joshuaray/go-proxy/actions/workflows/go-build.yml)

# go-proxy
Proxy server for Go

## Request Example
```javascript
const http = require('http');

function request() {
  const options = {
    hostname: 'localhost',
    port: 33333,
    path: 'http://localhost',
    method: 'GET',
    headers: {
      'X-Token': 'Authentication token',
      'X-Url': 'Proxy URL'
    }
  };

  const req = http.request(options, function (r) {
    res.writeHead(r.statusCode, r.headers);
    r.pipe(res, {
      end: true
    });
  });
}
```

## Headers
* X-Token: AES-256 encoded value in base64 using scrypt derivation function
* X-Url: Proxy URL, must be of a whitelisted domain
