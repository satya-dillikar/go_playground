# HTTPS TLS Server

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

* certstrap
```
 brew install certstrap
```
* clone repo
```
  git clone https://github.com/youngkin/gohttps
```

### Installing

* generate cert
```
certstrap init --common-name "ExampleCA"
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Created out/ExampleCA.key
Created out/ExampleCA.crt
Created out/ExampleCA.crl
```

```
certstrap request-cert --domain "localhost"
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Created out/localhost.key
Created out/localhost.csr
```

```
certstrap request-cert --domain "client"
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Created out/client.key
Created out/client.csr
```

```
certstrap sign localhost --CA ExampleCA
Created out/localhost.crt from out/localhost.csr signed by out/ExampleCA.key
```

```
certstrap sign client --CA ExampleCA
Created out/client.crt from out/client.csr signed by out/ExampleCA.key
```


* simple server
```
go run gohttps/simpleserver/server.go -srvcert out/localhost.crt -srvkey out/localhost.key -host localhost
```

```
go run gohttps/simpleserver/server.go -srvcert out/localhost.crt -srvkey out/localhost.key -host localhost -port 8888
```

* CLIENT run

* curl request with no client certificate validation.
```
curl -vi -d "world" --cacert out/ExampleCA.crt https://localhost:8888
```

* curl request with full client certificate validation.
```
curl -vi -d "world" --cacert out/ExampleCA.crt  --cert out/client.crt --key out/client.key https://localhost:8888

```


```
go run gohttps/client/client.go -clientcert out/client.crt -clientkey out/client.key -cacert out/ExampleCA.crt -port 8888
```

```
go run gohttps/client/client.go -cacert out/ExampleCA.crt -port 8888

```
### Executing program

* React APP client 
```
/Users/sdillikar/github/public/react_pub/my-app3-ts
mkdir -p .cert
cp /Users/sdillikar/github/projects/go_playground/https_tls_server/out/client.* .
package.json:
  "scripts": {
    "start": "HTTPS=true SSL_CRT_FILE=./.cert/client.crt SSL_KEY_FILE=./.cert/client.key react-scripts start",
    "build": "react-scripts build",
    "test": "react-scripts test",
    "eject": "react-scripts eject"
  },

    "start": "HTTPS=true SSL_CRT_FILE=./.cert/localhost.crt SSL_KEY_FILE=./.cert/localhost.key react-scripts start",

    "start": "HTTPS=true SSL_CRT_FILE=./.cert/server.crt SSL_KEY_FILE=./.cert/server.key react-scripts start",
```

## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```

## Authors

Contributors names and contact info

[@SatyaDillikar](https://twitter.com/SatyaDillikar)

## Version History

* 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release

## License

N/A

## Acknowledgments

Inspiration, code snippets, etc.
* [gohttpsclientserver](https://youngkin.github.io/post/gohttpsclientserver/)