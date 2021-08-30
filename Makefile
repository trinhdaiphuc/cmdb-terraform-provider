build:
	go build -o bin/terraform-provider-cmdb main.go

copy:
	mv bin/terraform-provider-cmdb ~/.terraform.d/plugins/trinhdaiphuc/cmdb/0.4/darwin_amd64
