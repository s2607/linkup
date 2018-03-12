#!/bin/sh
echo "if you get an error go install openssl"
echo "its not my fault if thats hard for you"

openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 9000
