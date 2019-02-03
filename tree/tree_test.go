package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	Convey("Given an initial Tree with a valid ID", t, func() {
		name := "normal"
		initialTree := Tree{
			Id: name,
			Root: &Node{
				Id: "1",
				Data: MetaData{
					Name:  "#1",
					Title: "Founder",
				},
			},
		}

		Convey("When we create a new Tree with that same name", func() {
			resultingTree, err := Create(name)

			Convey("We should not have any error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("And the value of the resulting tree must resemble the initial tree", func() {
				So(*resultingTree, ShouldResemble, initialTree)
			})
		})
	})
}
