terraform {
  required_providers {
    cmdb = {
      version = "0.1"
      source = "zalopay.vn/top/cmdb"
    }
  }
}

provider "cmdb" {
  api_version = "v1"
  hostname = "localhost"
  protocol = "http"
  port = 8080
}

resource "cmdb" "new" {
  config {
    name = "db.host"
    value = "localhost"
  }
}

output "new_config" {
  value = cmdb.new
}


data "cmdb" "db_host" {
  id = "db.host"
}

output "db_host" {
  value = data.cmdb.db_host
}
