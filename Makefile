build:
	go build -o bin/terraform-provider-cmdb main.go

copy:
	mv bin/terraform-provider-cmdb ~/.terraform.d/plugins/zalopay.vn/top/cmdb/0.4/darwin_amd64
