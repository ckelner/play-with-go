package bdd_cli

import (
	"log"
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// A simple demo example
func TestSpec(t *testing.T) {
	Convey("Given that a user has created an environment with the name 'kelner-test'", t, func() {
		// TODO: Perform setup to create environment; this was a simple hack to start
		// Will need to capture the environment name and id

		Convey("When the user lists their environments", func() {
			output, err := exec.Command("sh", "-c", "bx schematics environment list").Output()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("bx schematics environment list: %v", string(output))

			Convey("Then the environment 'kelner-test' should be in the output", func() {
				So(string(output), ShouldContainSubstring, "kelner-test   d047b72086d5ecc899d87b4103157a62")
			})
		})
	})
}
