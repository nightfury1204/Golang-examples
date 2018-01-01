#Secure hello server

```
go run main.go --help

#for generating certificate
go run main.go initcert

#run https server and it doesn't require client certificate
go run main.go serve --mutual=false

#run https server and require client certificate
go run main.go serve --mutual=false

```

####Client
```apple js
#run https client in client directory
go run main.go

```