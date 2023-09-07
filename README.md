
```
microservice_go
├─ .env
├─ add_user_microservice
│  ├─ add_user.exe
│  ├─ app.env
│  ├─ bin
│  ├─ Dockerfile
│  ├─ gapi
│  │  ├─ rpc_add
│  │  └─ server.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main
│  ├─ main.go
│  ├─ pb
│  │  ├─ rpc_add.pb.go
│  │  ├─ rpc_transfert.pb.go
│  │  ├─ service_add.pb.go
│  │  ├─ service_add.pb.gw.go
│  │  ├─ service_add_grpc.pb.go
│  │  └─ service_transfert_grpc.pb.go
│  ├─ proto
│  │  ├─ google
│  │  │  └─ api
│  │  │     ├─ annotations.proto
│  │  │     ├─ field_behavior.proto
│  │  │     ├─ http.proto
│  │  │     └─ httpbody.proto
│  │  ├─ rpc_add.proto
│  │  └─ service_add.proto
│  ├─ tools
│  │  └─ tools.go
│  └─ utils
│     ├─ add.go
│     └─ config.go
├─ admin_microservice
│  ├─ app.env
│  ├─ Dockerfile
│  ├─ go.mod
│  ├─ go.sum
│  ├─ index.html
│  ├─ main.go
│  └─ utils
│     └─ config.go
├─ Capture d’écran 2023-07-19 205703.jpg
├─ conf
│  └─ app.conf
├─ data
│  ├─ access
│  ├─ authelia
│  │  └─ config
│  │     ├─ configuration.yml
│  │     ├─ db.sqlite3
│  │     ├─ notification.txt
│  │     └─ users_database.yml
│  ├─ custom_ssl
│  │  ├─ npm-10
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  ├─ npm-11
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  ├─ npm-12
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  ├─ npm-13
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  ├─ npm-8
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  └─ npm-9
│  │     ├─ fullchain.pem
│  │     └─ privkey.pem
│  ├─ database.sqlite
│  ├─ databases
│  │  ├─ neo4j
│  │  │  ├─ database_lock
│  │  │  ├─ id-buffer.tmp.0
│  │  │  ├─ neostore
│  │  │  ├─ neostore.counts.db
│  │  │  ├─ neostore.indexstats.db
│  │  │  ├─ neostore.labeltokenstore.db
│  │  │  ├─ neostore.labeltokenstore.db.id
│  │  │  ├─ neostore.labeltokenstore.db.names
│  │  │  ├─ neostore.labeltokenstore.db.names.id
│  │  │  ├─ neostore.nodestore.db
│  │  │  ├─ neostore.nodestore.db.id
│  │  │  ├─ neostore.nodestore.db.labels
│  │  │  ├─ neostore.nodestore.db.labels.id
│  │  │  ├─ neostore.propertystore.db
│  │  │  ├─ neostore.propertystore.db.arrays
│  │  │  ├─ neostore.propertystore.db.arrays.id
│  │  │  ├─ neostore.propertystore.db.id
│  │  │  ├─ neostore.propertystore.db.index
│  │  │  ├─ neostore.propertystore.db.index.id
│  │  │  ├─ neostore.propertystore.db.index.keys
│  │  │  ├─ neostore.propertystore.db.index.keys.id
│  │  │  ├─ neostore.propertystore.db.strings
│  │  │  ├─ neostore.propertystore.db.strings.id
│  │  │  ├─ neostore.relationshipgroupstore.db
│  │  │  ├─ neostore.relationshipgroupstore.db.id
│  │  │  ├─ neostore.relationshipgroupstore.degrees.db
│  │  │  ├─ neostore.relationshipstore.db
│  │  │  ├─ neostore.relationshipstore.db.id
│  │  │  ├─ neostore.relationshiptypestore.db
│  │  │  ├─ neostore.relationshiptypestore.db.id
│  │  │  ├─ neostore.relationshiptypestore.db.names
│  │  │  ├─ neostore.relationshiptypestore.db.names.id
│  │  │  ├─ neostore.schemastore.db
│  │  │  ├─ neostore.schemastore.db.id
│  │  │  └─ schema
│  │  │     └─ index
│  │  │        └─ token-lookup-1.0
│  │  │           ├─ 1
│  │  │           │  └─ index-1
│  │  │           └─ 2
│  │  │              └─ index-2
│  │  ├─ store_lock
│  │  └─ system
│  │     ├─ database_lock
│  │     ├─ id-buffer.tmp.0
│  │     ├─ neostore
│  │     ├─ neostore.counts.db
│  │     ├─ neostore.indexstats.db
│  │     ├─ neostore.labeltokenstore.db
│  │     ├─ neostore.labeltokenstore.db.id
│  │     ├─ neostore.labeltokenstore.db.names
│  │     ├─ neostore.labeltokenstore.db.names.id
│  │     ├─ neostore.nodestore.db
│  │     ├─ neostore.nodestore.db.id
│  │     ├─ neostore.nodestore.db.labels
│  │     ├─ neostore.nodestore.db.labels.id
│  │     ├─ neostore.propertystore.db
│  │     ├─ neostore.propertystore.db.arrays
│  │     ├─ neostore.propertystore.db.arrays.id
│  │     ├─ neostore.propertystore.db.id
│  │     ├─ neostore.propertystore.db.index
│  │     ├─ neostore.propertystore.db.index.id
│  │     ├─ neostore.propertystore.db.index.keys
│  │     ├─ neostore.propertystore.db.index.keys.id
│  │     ├─ neostore.propertystore.db.strings
│  │     ├─ neostore.propertystore.db.strings.id
│  │     ├─ neostore.relationshipgroupstore.db
│  │     ├─ neostore.relationshipgroupstore.db.id
│  │     ├─ neostore.relationshipgroupstore.degrees.db
│  │     ├─ neostore.relationshipstore.db
│  │     ├─ neostore.relationshipstore.db.id
│  │     ├─ neostore.relationshiptypestore.db
│  │     ├─ neostore.relationshiptypestore.db.id
│  │     ├─ neostore.relationshiptypestore.db.names
│  │     ├─ neostore.relationshiptypestore.db.names.id
│  │     ├─ neostore.schemastore.db
│  │     ├─ neostore.schemastore.db.id
│  │     └─ schema
│  │        └─ index
│  │           ├─ range-1.0
│  │           │  ├─ 3
│  │           │  │  └─ index-3
│  │           │  ├─ 4
│  │           │  │  └─ index-4
│  │           │  ├─ 7
│  │           │  │  └─ index-7
│  │           │  └─ 8
│  │           │     └─ index-8
│  │           └─ token-lookup-1.0
│  │              ├─ 1
│  │              │  └─ index-1
│  │              └─ 2
│  │                 └─ index-2
│  ├─ dbms
│  │  └─ auth.ini
│  ├─ keys.json
│  ├─ letsencrypt-acme-challenge
│  │  └─ .well-known
│  │     └─ acme-challenge
│  │        └─ test-challenge
│  ├─ logs
│  │  ├─ fallback_access.log
│  │  ├─ fallback_error.log
│  │  ├─ letsencrypt-requests_access.log
│  │  ├─ letsencrypt-requests_error.log
│  │  ├─ proxy-host-11_access.log
│  │  ├─ proxy-host-11_error.log
│  │  ├─ proxy-host-1_access.log
│  │  ├─ proxy-host-1_error.log
│  │  ├─ proxy-host-2_access.log
│  │  ├─ proxy-host-2_error.log
│  │  ├─ proxy-host-6_access.log
│  │  ├─ proxy-host-6_error.log
│  │  ├─ proxy-host-7_access.log
│  │  ├─ proxy-host-7_error.log
│  │  ├─ proxy-host-8_access.log
│  │  ├─ proxy-host-8_error.log
│  │  ├─ proxy-host-9_access.log
│  │  └─ proxy-host-9_error.log
│  ├─ nginx
│  │  ├─ dead_host
│  │  ├─ default_host
│  │  ├─ default_www
│  │  ├─ proxy_host
│  │  │  ├─ 11.conf
│  │  │  ├─ 6.conf
│  │  │  ├─ 7.conf
│  │  │  ├─ 8.conf
│  │  │  └─ 9.conf
│  │  ├─ redirection_host
│  │  ├─ snippets
│  │  │  ├─ authelia-authrequest.conf
│  │  │  ├─ authelia-location.conf
│  │  │  └─ proxy.conf
│  │  ├─ stream
│  │  └─ temp
│  ├─ server_id
│  └─ transactions
│     ├─ neo4j
│     │  ├─ checkpoint.0
│     │  └─ neostore.transaction.db.0
│     └─ system
│        ├─ checkpoint.0
│        └─ neostore.transaction.db.0
├─ delete_microservice
│  ├─ app.env
│  ├─ bin
│  ├─ Dockerfile
│  ├─ gapi
│  │  ├─ rpc_delete.go
│  │  └─ server.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main.go
│  ├─ pb
│  │  ├─ rpc_delete.pb.go
│  │  ├─ service_delete.pb.go
│  │  ├─ service_delete.pb.gw.go
│  │  └─ service_delete_grpc.pb.go
│  ├─ proto
│  │  ├─ google
│  │  │  └─ api
│  │  │     ├─ annotations.proto
│  │  │     ├─ field_behavior.proto
│  │  │     ├─ http.proto
│  │  │     └─ httpbody.proto
│  │  ├─ rpc_delete.proto
│  │  └─ service_delete.proto
│  ├─ tools
│  │  └─ tools.go
│  └─ utils
│     ├─ config.go
│     └─ delete.go
├─ dnsmasq.conf
├─ docker-compose.yml
├─ Dockerfile
├─ letsencrypt
│  ├─ accounts
│  │  └─ acme-v02.api.letsencrypt.org
│  │     └─ directory
│  │        └─ b476b2138f6c5ce3c835c724abf3e4a3
│  │           ├─ meta.json
│  │           ├─ private_key.json
│  │           └─ regr.json
│  └─ renewal
├─ mlmcli
│  ├─ cmd
│  │  ├─ addUser.go
│  │  ├─ deleteUser.go
│  │  ├─ root.go
│  │  └─ transferUser.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ LICENSE
│  ├─ main.go
│  ├─ mlmcli
│  ├─ mlmcli.exe
│  └─ utils
│     ├─ request.go
│     ├─ response.go
│     └─ utils.go
├─ Testing
│  ├─ delete.go
│  └─ go.mod
└─ transfer_user_microservice
   ├─ app.env
   ├─ bin
   ├─ Dockerfile
   ├─ gapi
   │  ├─ rpc_transfert.go
   │  └─ server.go
   ├─ go.mod
   ├─ go.sum
   ├─ main.go
   ├─ pb
   │  ├─ rpc_transfert.pb.go
   │  ├─ service_transfert.pb.go
   │  ├─ service_transfert.pb.gw.go
   │  ├─ service_transfert_grpc.pb.go
   │  └─ user.pb.go
   ├─ proto
   │  ├─ google
   │  │  └─ api
   │  │     ├─ annotations.proto
   │  │     ├─ field_behavior.proto
   │  │     ├─ http.proto
   │  │     └─ httpbody.proto
   │  ├─ rpc_transfert.proto
   │  ├─ service_transfert.proto
   │  └─ user.proto
   ├─ tools
   │  └─ tools.go
   └─ utils
      ├─ config.go
      └─ transport.go

```
```
microservice_go
├─ .env
├─ add_user_microservice
│  ├─ add_user.exe
│  ├─ app.env
│  ├─ bin
│  ├─ Dockerfile
│  ├─ gapi
│  │  ├─ rpc_add.go
│  │  └─ server.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main
│  ├─ main.go
│  ├─ pb
│  │  ├─ rpc_add.pb.go
│  │  ├─ rpc_transfert.pb.go
│  │  ├─ service_add.pb.go
│  │  ├─ service_add.pb.gw.go
│  │  ├─ service_add_grpc.pb.go
│  │  └─ service_transfert_grpc.pb.go
│  ├─ proto
│  │  ├─ google
│  │  │  └─ api
│  │  │     ├─ annotations.proto
│  │  │     ├─ field_behavior.proto
│  │  │     ├─ http.proto
│  │  │     └─ httpbody.proto
│  │  ├─ rpc_add.proto
│  │  └─ service_add.proto
│  ├─ tools
│  │  └─ tools.go
│  └─ utils
│     ├─ add.go
│     └─ config.go
├─ admin_microservice
│  ├─ app.env
│  ├─ Dockerfile
│  ├─ go.mod
│  ├─ go.sum
│  ├─ index.html
│  ├─ main.go
│  └─ utils
│     └─ config.go
├─ conf
│  └─ app.conf
├─ data
│  ├─ access
│  ├─ authelia
│  │  └─ config
│  │     ├─ configuration.yml
│  │     ├─ db.sqlite3
│  │     ├─ notification.txt
│  │     └─ users_database.yml
│  ├─ custom_ssl
│  │  ├─ npm-10
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  ├─ npm-11
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  ├─ npm-12
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  ├─ npm-13
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  ├─ npm-8
│  │  │  ├─ fullchain.pem
│  │  │  └─ privkey.pem
│  │  └─ npm-9
│  │     ├─ fullchain.pem
│  │     └─ privkey.pem
│  ├─ database.sqlite
│  ├─ databases
│  │  ├─ neo4j
│  │  │  ├─ database_lock
│  │  │  ├─ id-buffer.tmp.0
│  │  │  ├─ neostore
│  │  │  ├─ neostore.counts.db
│  │  │  ├─ neostore.indexstats.db
│  │  │  ├─ neostore.labeltokenstore.db
│  │  │  ├─ neostore.labeltokenstore.db.id
│  │  │  ├─ neostore.labeltokenstore.db.names
│  │  │  ├─ neostore.labeltokenstore.db.names.id
│  │  │  ├─ neostore.nodestore.db
│  │  │  ├─ neostore.nodestore.db.id
│  │  │  ├─ neostore.nodestore.db.labels
│  │  │  ├─ neostore.nodestore.db.labels.id
│  │  │  ├─ neostore.propertystore.db
│  │  │  ├─ neostore.propertystore.db.arrays
│  │  │  ├─ neostore.propertystore.db.arrays.id
│  │  │  ├─ neostore.propertystore.db.id
│  │  │  ├─ neostore.propertystore.db.index
│  │  │  ├─ neostore.propertystore.db.index.id
│  │  │  ├─ neostore.propertystore.db.index.keys
│  │  │  ├─ neostore.propertystore.db.index.keys.id
│  │  │  ├─ neostore.propertystore.db.strings
│  │  │  ├─ neostore.propertystore.db.strings.id
│  │  │  ├─ neostore.relationshipgroupstore.db
│  │  │  ├─ neostore.relationshipgroupstore.db.id
│  │  │  ├─ neostore.relationshipgroupstore.degrees.db
│  │  │  ├─ neostore.relationshipstore.db
│  │  │  ├─ neostore.relationshipstore.db.id
│  │  │  ├─ neostore.relationshiptypestore.db
│  │  │  ├─ neostore.relationshiptypestore.db.id
│  │  │  ├─ neostore.relationshiptypestore.db.names
│  │  │  ├─ neostore.relationshiptypestore.db.names.id
│  │  │  ├─ neostore.schemastore.db
│  │  │  ├─ neostore.schemastore.db.id
│  │  │  └─ schema
│  │  │     └─ index
│  │  │        └─ token-lookup-1.0
│  │  │           ├─ 1
│  │  │           │  └─ index-1
│  │  │           └─ 2
│  │  │              └─ index-2
│  │  ├─ store_lock
│  │  └─ system
│  │     ├─ database_lock
│  │     ├─ id-buffer.tmp.0
│  │     ├─ neostore
│  │     ├─ neostore.counts.db
│  │     ├─ neostore.indexstats.db
│  │     ├─ neostore.labeltokenstore.db
│  │     ├─ neostore.labeltokenstore.db.id
│  │     ├─ neostore.labeltokenstore.db.names
│  │     ├─ neostore.labeltokenstore.db.names.id
│  │     ├─ neostore.nodestore.db
│  │     ├─ neostore.nodestore.db.id
│  │     ├─ neostore.nodestore.db.labels
│  │     ├─ neostore.nodestore.db.labels.id
│  │     ├─ neostore.propertystore.db
│  │     ├─ neostore.propertystore.db.arrays
│  │     ├─ neostore.propertystore.db.arrays.id
│  │     ├─ neostore.propertystore.db.id
│  │     ├─ neostore.propertystore.db.index
│  │     ├─ neostore.propertystore.db.index.id
│  │     ├─ neostore.propertystore.db.index.keys
│  │     ├─ neostore.propertystore.db.index.keys.id
│  │     ├─ neostore.propertystore.db.strings
│  │     ├─ neostore.propertystore.db.strings.id
│  │     ├─ neostore.relationshipgroupstore.db
│  │     ├─ neostore.relationshipgroupstore.db.id
│  │     ├─ neostore.relationshipgroupstore.degrees.db
│  │     ├─ neostore.relationshipstore.db
│  │     ├─ neostore.relationshipstore.db.id
│  │     ├─ neostore.relationshiptypestore.db
│  │     ├─ neostore.relationshiptypestore.db.id
│  │     ├─ neostore.relationshiptypestore.db.names
│  │     ├─ neostore.relationshiptypestore.db.names.id
│  │     ├─ neostore.schemastore.db
│  │     ├─ neostore.schemastore.db.id
│  │     └─ schema
│  │        └─ index
│  │           ├─ range-1.0
│  │           │  ├─ 3
│  │           │  │  └─ index-3
│  │           │  ├─ 4
│  │           │  │  └─ index-4
│  │           │  ├─ 7
│  │           │  │  └─ index-7
│  │           │  └─ 8
│  │           │     └─ index-8
│  │           └─ token-lookup-1.0
│  │              ├─ 1
│  │              │  └─ index-1
│  │              └─ 2
│  │                 └─ index-2
│  ├─ dbms
│  │  └─ auth.ini
│  ├─ keys.json
│  ├─ letsencrypt-acme-challenge
│  │  └─ .well-known
│  │     └─ acme-challenge
│  │        └─ test-challenge
│  ├─ logs
│  │  ├─ fallback_access.log
│  │  ├─ fallback_error.log
│  │  ├─ letsencrypt-requests_access.log
│  │  ├─ letsencrypt-requests_error.log
│  │  ├─ proxy-host-11_access.log
│  │  ├─ proxy-host-11_error.log
│  │  ├─ proxy-host-1_access.log
│  │  ├─ proxy-host-1_error.log
│  │  ├─ proxy-host-2_access.log
│  │  ├─ proxy-host-2_error.log
│  │  ├─ proxy-host-6_access.log
│  │  ├─ proxy-host-6_error.log
│  │  ├─ proxy-host-7_access.log
│  │  ├─ proxy-host-7_error.log
│  │  ├─ proxy-host-8_access.log
│  │  ├─ proxy-host-8_error.log
│  │  ├─ proxy-host-9_access.log
│  │  └─ proxy-host-9_error.log
│  ├─ nginx
│  │  ├─ dead_host
│  │  ├─ default_host
│  │  ├─ default_www
│  │  ├─ proxy_host
│  │  │  ├─ 11.conf
│  │  │  ├─ 6.conf
│  │  │  ├─ 7.conf
│  │  │  ├─ 8.conf
│  │  │  └─ 9.conf
│  │  ├─ redirection_host
│  │  ├─ snippets
│  │  │  ├─ authelia-authrequest.conf
│  │  │  ├─ authelia-location.conf
│  │  │  └─ proxy.conf
│  │  ├─ stream
│  │  └─ temp
│  ├─ server_id
│  └─ transactions
│     ├─ neo4j
│     │  ├─ checkpoint.0
│     │  └─ neostore.transaction.db.0
│     └─ system
│        ├─ checkpoint.0
│        └─ neostore.transaction.db.0
├─ delete_microservice
│  ├─ app.env
│  ├─ bin
│  ├─ Dockerfile
│  ├─ gapi
│  │  ├─ rpc_delete.go
│  │  └─ server.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main.go
│  ├─ pb
│  │  ├─ rpc_delete.pb.go
│  │  ├─ service_delete.pb.go
│  │  ├─ service_delete.pb.gw.go
│  │  └─ service_delete_grpc.pb.go
│  ├─ proto
│  │  ├─ google
│  │  │  └─ api
│  │  │     ├─ annotations.proto
│  │  │     ├─ field_behavior.proto
│  │  │     ├─ http.proto
│  │  │     └─ httpbody.proto
│  │  ├─ rpc_delete.proto
│  │  └─ service_delete.proto
│  ├─ tools
│  │  └─ tools.go
│  └─ utils
│     ├─ config.go
│     └─ delete.go
├─ docker-compose.yml
├─ Dockerfile
├─ letsencrypt
│  ├─ accounts
│  │  └─ acme-v02.api.letsencrypt.org
│  │     └─ directory
│  │        └─ b476b2138f6c5ce3c835c724abf3e4a3
│  │           ├─ meta.json
│  │           ├─ private_key.json
│  │           └─ regr.json
│  └─ renewal
├─ mlmcli
│  ├─ cmd
│  │  ├─ addUser.go
│  │  ├─ deleteUser.go
│  │  ├─ root.go
│  │  └─ transferUser.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ LICENSE
│  ├─ main.go
│  ├─ mlmcli
│  ├─ mlmcli.exe
│  └─ utils
│     ├─ request.go
│     ├─ response.go
│     └─ utils.go
├─ README.md
├─ transfer_user_microservice
│  ├─ app.env
│  ├─ bin
│  ├─ Dockerfile
│  ├─ gapi
│  │  ├─ rpc_transfert.go
│  │  └─ server.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main.go
│  ├─ pb
│  │  ├─ rpc_transfert.pb.go
│  │  ├─ service_transfert.pb.go
│  │  ├─ service_transfert.pb.gw.go
│  │  ├─ service_transfert_grpc.pb.go
│  │  └─ user.pb.go
│  ├─ proto
│  │  ├─ google
│  │  │  └─ api
│  │  │     ├─ annotations.proto
│  │  │     ├─ field_behavior.proto
│  │  │     ├─ http.proto
│  │  │     └─ httpbody.proto
│  │  ├─ rpc_transfert.proto
│  │  ├─ service_transfert.proto
│  │  └─ user.proto
│  ├─ tools
│  │  └─ tools.go
│  └─ utils
│     ├─ config.go
│     └─ transport.go
└─ vizialisation
   ├─ app.js
   ├─ index.html
   ├─ package-lock.json
   ├─ package.json
   ├─ records.json
   └─ style.css

```