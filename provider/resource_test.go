package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccCheckCmdbConfigBasic(name, value string) string {
	return fmt.Sprintf(`
	resource "cmdb_config" "new" {
		config = {
			name = "%v"
			value = "%v"
	  	}
	}
	`, name, value)
}

func testAccCheckCmdbConfigrExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No config id set")
		}

		return nil
	}
}
