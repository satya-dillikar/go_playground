https://youngkin.github.io/post/gohttpsclientserver/

 brew install certstrap

 git clone https://github.com/youngkin/gohttps


generate cert
--------------
certstrap init --common-name "ExampleCA"
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Created out/ExampleCA.key
Created out/ExampleCA.crt
Created out/ExampleCA.crl

certstrap request-cert --domain "localhost"
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Created out/localhost.key
Created out/localhost.csr

certstrap request-cert --domain "client"
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Created out/client.key
Created out/client.csr

certstrap sign localhost --CA ExampleCA
Created out/localhost.crt from out/localhost.csr signed by out/ExampleCA.key

certstrap sign client --CA ExampleCA
Created out/client.crt from out/client.csr signed by out/ExampleCA.key

simple server
-------------
go run gohttps/simpleserver/server.go -srvcert out/localhost.crt -srvkey out/localhost.key -host localhost

go run gohttps/simpleserver/server.go -srvcert out/localhost.crt -srvkey out/localhost.key -host localhost -port 8888

client run
----------
#curl request with no client certificate validation.
curl -vi -d "world" --cacert out/ExampleCA.crt https://localhost:8888

#curl request with full client certificate validation.
curl -vi -d "world" --cacert out/ExampleCA.crt  --cert out/client.crt --key out/client.key https://localhost:8888

go run gohttps/client/client.go -clientcert out/client.crt -clientkey out/client.key -cacert out/ExampleCA.crt -port 8888

go run gohttps/client/client.go -cacert out/ExampleCA.crt -port 8888

Reach APP client 
----------
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