package nexus

import (
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"
)

func testAccDataSourceSecuritySAML() (*security.SAML, error) {
	dat, err := ioutil.ReadFile("../examples/saml-testconfig.xml")
	if err != nil {
		return nil, err
	}

	return &security.SAML{
		// https://samltest.id/saml/idp
		IdpMetadata:                string(dat),
		EntityId:                   "http://example.test/service/rest/v1/security/saml/metadata",
		ValidateAssertionSignature: false,
		ValidateResponseSignature:  true,
		UsernameAttribute:          "username2",
		FirstNameAttribute:         "firstName",
		LastNameAttribute:          "lastName",
		EmailAttribute:             "email",
		GroupsAttribute:            "groups",
	}, nil
}

func TestAccDataSourceSecuritySaml(t *testing.T) {
	if getEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}

	saml, err := testAccDataSourceSecuritySAML()
	assert.Nil(t, err)
	resName := "data.nexus_security_saml.acceptance"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecuritySAMLConfig(*saml),
				Check:  nil,
			},
			{
				Config: testAccResourceSecuritySAMLConfig(*saml) + testAccDataSourceSecuritySAMLConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "idp_metadata", saml.IdpMetadata),
					resource.TestCheckResourceAttr(resName, "entity_id", saml.EntityId),
					resource.TestCheckResourceAttr(resName, "validate_response_signature", strconv.FormatBool(saml.ValidateResponseSignature)),
					resource.TestCheckResourceAttr(resName, "validate_assertion_signature", strconv.FormatBool(saml.ValidateAssertionSignature)),
					resource.TestCheckResourceAttr(resName, "username_attribute", saml.UsernameAttribute),
					resource.TestCheckResourceAttr(resName, "first_name_attribute", saml.FirstNameAttribute),
					resource.TestCheckResourceAttr(resName, "last_name_attribute", saml.LastNameAttribute),
					resource.TestCheckResourceAttr(resName, "email_attribute", saml.EmailAttribute),
					resource.TestCheckResourceAttr(resName, "groups_attribute", saml.GroupsAttribute),
				),
			},
			{
				ResourceName:      resName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataSourceSecuritySAMLConfig() string {
	return `
data "nexus_security_saml" "acceptance" {
}
`
}
