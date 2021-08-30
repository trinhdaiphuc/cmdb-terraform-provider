package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CMDB_API_VERSION"); v == "" {
		t.Fatal("CMDB_API_VERSION must be set for acceptance tests")
	}
	if v := os.Getenv("CMDB_HOST"); v == "" {
		t.Fatal("CMDB_HOST must be set for acceptance tests")
	}
}

func TestMainProvider(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"cmdb_config": func() (tfprotov6.ProviderServer, error) {
				// newProvider is your function that returns a
				// tfsdk.Provider implementation
				return tfsdk.NewProtocol6Server(New()), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCmdbConfigBasic("db.host", "localhost"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmdbConfigrExists("cmdb_config.new"),
				),
			},
		},
	})
}
