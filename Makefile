build:
	go build -o bin/cmdb-provider main.go

copy:
	mv bin/cmdb-provider ~/.terraform.d/plugins/zalopay.vn/top/cmdb/0.1/darwin_amd64
