package appgate

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccConditionBasic(t *testing.T) {
	resourceName := "appgate_condition.test_condition"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConditionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCondition(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConditionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "condition-test"),
					resource.TestCheckResourceAttr(resourceName, "expression", "return true;"),
					resource.TestCheckResourceAttr(resourceName, "notes", "Managed by terraform"),

					resource.TestCheckResourceAttr(resourceName, "remedy_methods.#", "0"),

					resource.TestCheckResourceAttr(resourceName, "repeat_schedules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "repeat_schedules.2107984292", "13:32"),
					resource.TestCheckResourceAttr(resourceName, "repeat_schedules.3334954558", "1h"),

					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.2876187004", "api-created"),
					resource.TestCheckResourceAttr(resourceName, "tags.535570215", "terraform"),
				),
			},
		},
	})
}

func testAccCheckCondition() string {
	return fmt.Sprintf(`
resource "appgate_condition" "test_condition" {
    name = "condition-test"
    tags = [
      "terraform",
      "api-created"
    ]

    expression = "return true;"

    repeat_schedules = [
      "1h",
      "13:32"
    ]

}
`)
}

func testAccCheckConditionExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		token := testAccProvider.Meta().(*Client).Token
		api := testAccProvider.Meta().(*Client).API.ConditionsApi

		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		_, _, err := api.ConditionsIdGet(context.Background(), rs.Primary.ID).Authorization(token).Execute()
		if err != nil {
			return fmt.Errorf("error fetching condition with resource %s. %s", resource, err)
		}
		return nil
	}
}

func testAccCheckConditionDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "appgate_condition" {
			continue
		}

		token := testAccProvider.Meta().(*Client).Token
		api := testAccProvider.Meta().(*Client).API.ConditionsApi

		_, _, err := api.ConditionsIdGet(context.Background(), rs.Primary.ID).Authorization(token).Execute()
		if err == nil {
			return fmt.Errorf("Condition still exists, %+v", err)
		}
	}
	return nil
}