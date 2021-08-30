terraform {
  required_providers {
    cmdb = {
      version = "0.4"
      source = "zalopay.com.vn/top/cmdb"
    }
  }
}

provider "cmdb" {
  api_version = "v1"
  host = "http://localhost:8080"
}

resource "cmdb_config" "new" {
  config = {
    name = "db.host"
    value = "localhost"
  }
}

output "new_config" {
  value = cmdb_config.new
}

data "cmdb_config" "history" {
  name = "db.host"
}

output "db_host" {
  value = data.cmdb_config.history
}
