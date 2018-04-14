#!/bin/bash
#^
echo "if you get an error go install openssl"
echo "its not my fault if thats hard for you"
unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac
echo ${machine}

openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
case "${machine}" in
	Linux) openssl req -new -x509 -sha256 -key server.key -out server.crt -days 9000  -subj "/C=US/ST=CA/O=Acme, Inc./CN=example.com";;
	Mac) openssl req -new -x509 -sha256 -key server.key -out server.crt -days 9000  -subj "/C=US/ST=CA/O=Acme, Inc./CN=example.com";;
	MinGw) openssl req -new -x509 -sha256 -key server.key -out server.crt -days 9000  -subj "//C=US\ST=CA\O=Acme, Inc.\CN=example.com";;
	*) echo "no."; exit 1
esac
