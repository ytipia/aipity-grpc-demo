#!/bin/bash

# 2023.12.04 by aipity

# variables set
DOMAIN="localhost"
CERT_DIR="certs"
COUNTRY="CN"
STATE="SHANGHAI"
CITY="SHANGHAI"
ORG_NAME="AIPITY"
EMAIL="admin@aipity.com"
KEY_SIZE="4096"
DAYS="3650"
PASS="aipity"
PASSPHRASEFILE="$CERT_DIR/passphrase.txt"

# if DOMAIN is not empty, use it
if [ -n "$1" ]; then
    DOMAIN="$1"
fi

echo "DOMAIN is: $DOMAIN"

mkdir certs

echo "################ 1. create PASSPHRASEFILE ###############"
# create PASSPHRASEFILE
touch $PASSPHRASEFILE
echo $PASS > $PASSPHRASEFILE

# create CA certs
echo "################ 2. Create CA certs ########################"

openssl genrsa -aes256 -out $CERT_DIR/ca.key  --passout file:$CERT_DIR/passphrase.txt $KEY_SIZE
openssl req -new -x509 -sha256 -days $DAYS -key $CERT_DIR/ca.key -out $CERT_DIR/ca.crt -subj "/C=$COUNTRY/ST=$STATE/L=$CITY/O=$ORG_NAME/CN=$DOMAIN/emailAddress=$EMAIL" --passin file:$CERT_DIR/passphrase.txt

# create server certs
echo "################ 3. Create Server side certs #######################"

openssl genrsa -out $CERT_DIR/server.key $KEY_SIZE
openssl req -new -sha256 -key $CERT_DIR/server.key -out $CERT_DIR/server.csr -subj "/C=$COUNTRY/ST=$STATE/L=$CITY/O=$ORG_NAME/CN=$DOMAIN/emailAddress=$EMAIL" --passin file:$CERT_DIR/passphrase.txt
openssl x509 -req -days $DAYS -sha256 -in $CERT_DIR/server.csr -CA $CERT_DIR/ca.crt -CAkey $CERT_DIR/ca.key -set_serial 1 -out $CERT_DIR/server.crt -extfile <(printf "subjectAltName=DNS:$DOMAIN") --passin file:$CERT_DIR/passphrase.txt
# display server cert content
echo "the server cert content is:"
echo ""
openssl x509 -noout -text -in $CERT_DIR/server.crt


# Create client certs
echo "################ 4. Create Client side certs #######################"

openssl genrsa -out $CERT_DIR/client.key $KEY_SIZE
openssl req -new -key $CERT_DIR/client.key -out $CERT_DIR/client.csr -subj "/C=$COUNTRY/ST=$STATE/L=$CITY/O=$ORG_NAME/CN=$DOMAIN/emailAddress=$EMAIL" --passin file:$CERT_DIR/passphrase.txt
openssl x509 -req -days $DAYS -sha256 -in $CERT_DIR/client.csr -CA $CERT_DIR/ca.crt -CAkey $CERT_DIR/ca.key -set_serial 2 -out $CERT_DIR/client.crt -extfile <(printf "subjectAltName=DNS:$DOMAIN") --passin file:$CERT_DIR/passphrase.txt
# display client cert content
echo "the client cert content is:"
echo ""
openssl x509 -noout -text -in $CERT_DIR/client.crt

