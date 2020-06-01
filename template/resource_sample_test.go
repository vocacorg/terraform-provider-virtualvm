package template

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccBitbucketSample_basic(t *testing.T) {
	var repo Sample

	testUser := os.Getenv("BITBUCKET_USERNAME")
	testAccBitbucketSampleConfig := fmt.Sprintf(`
		resource "bitbucket_Sample" "test_repo" {
			owner = "%s"
			name = "test-repo-for-Sample-test"
		}
	`, testUser)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBitbucketSampleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBitbucketSampleConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBitbucketSampleExists("bitbucket_Sample.test_repo", &repo),
				),
			},
		},
	})
}

func TestAccBitbucketSample_camelcase(t *testing.T) {
	var repo Sample

	testUser := os.Getenv("BITBUCKET_USERNAME")
	testAccBitbucketSampleConfig := fmt.Sprintf(`
		resource "bitbucket_Sample" "test_repo" {
			owner = "%s"
			name = "TestRepoForSampleTest"
			slug = "test-repo-for-Sample-test"
		}
	`, testUser)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBitbucketSampleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBitbucketSampleConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBitbucketSampleExists("bitbucket_Sample.test_repo", &repo),
				),
			},
		},
	})
}

func testAccCheckBitbucketSampleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)
	rs, ok := s.RootModule().Resources["bitbucket_Sample.test_repo"]
	if !ok {
		return fmt.Errorf("Not found %s", "bitbucket_Sample.test_repo")
	}

	response, _ := client.Get(fmt.Sprintf("2.0/repositories/%s/%s", rs.Primary.Attributes["owner"], rs.Primary.Attributes["name"]))

	if response.StatusCode != 404 {
		return fmt.Errorf("Sample still exists")
	}

	return nil
}

func testAccCheckBitbucketSampleExists(n string, Sample *Sample) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Sample ID is set")
		}
		return nil
	}
}
