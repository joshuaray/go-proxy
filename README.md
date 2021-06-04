# go-proxy
Proxy server for Go

## Request Example
```javascript
const http = require('http');

function request() {
  console.log('serve: ' + req.url);

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
* X-Token: AES-256 encoded value in base64
* X-Url: Proxy URL, must be of a whitelisted domain
