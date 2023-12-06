# grpc: from 0 to 0.5 in action(go version)

---

**This repo is a simple grpc practice example. It contains such features:**

1. With grpc server and grpc client codes.
2. With a shell script can generate all related certs automatically.
3. With mTLS encryption for communation between grpc server and grpc client

---

## Steps to run:
### 1. Grpc server
#### 1. Generate all related certs
```shell
cd aipity-grpc-server-demo

./generate-all-certs.sh localhost # localhost is optional, see script for details
```
#### 2. Create grpc skeleton
```shell
# the steps to install protobuf and related plugins on mac
brew install protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
other OS version please refer [protoc-installation](https://grpc.io/docs/protoc-installation/)
```shell
#  steps to create grpc skeleton
#
#
cd aipity-grpc-server-demo/proto

# create user grpc info based on user.proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    user/user.proto

# create group grpc info based on group.proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    group/group.proto
```
#### 3. Init grpc server project
```shell
go mod tidy
```

#### 4. Start grpc server
```shell
air # run 'go run cmd/aipity/main.go' if air not configured
```

### 2. Grpc client
#### 1. Copy certs generated from grpc server here
```shell
cd aipity-grpc-client-demo

cp -r ../aipity-grpc-server-demo/certs .
```
#### 2. Copy grpc skeleton here
```shell
cp -r ../aipity-grpc-server-demo/proto .
```
#### 3. Init grpc client project
```shell
go mod tidy
```
#### 4. Start grpc client
```shell
air # run 'go run cmd/aipity/main.go' if air not configured
```

## Output
#### 1. Grpc server output
```shell
2023/12/05 22:53:03 grpc server with mTLS enabled listening at [::]:50051
2023/12/05 22:53:07 user create in=name:"aipity" password:"123456" email:"whsasf@aipity.com" phone:"111111111" status:1 role:1 createTime:1701787987 theme:1 language:1
2023/12/05 22:53:07 group create in=id:1 name:"aipity"
```
#### 2. Grpc client output
```shell
2023/12/05 22:53:07 creatd user: name:"aipity"  password:"123456"  email:"whsasf@aipity.com"  phone:"111111111"  status:1  role:1  createTime:1701787987  theme:1  language:1 added successfully
2023/12/05 22:53:07 created group: id:1  name:"aipity" added successfully
```

## TODO
#### 1. JWT auth