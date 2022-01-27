#!/bin/bash

# generate ca.key 
openssl genrsa -out ca.key 4096
#output : ca.key
# generate certificate
openssl req -new -x509 -key ca.key -sha256 -subj "/C=SE/ST=HL/O=Example, INC." -days 365 -out ca.cert
#output ca.cert
# generate the server key
openssl genrsa -out server.key 4096
#output server.key
# Generate the csr
openssl req -new -key server.key -out server.csr -config certificate.conf
# output server.csr
openssl x509 -req -in server.csr -CA ca.cert -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256 -extfile certificate.conf -extensions req_ext
#output server.crt
openssl x509 -in server.crt -out server.pem
#output server.pem
#server.crt or server.pem are same files