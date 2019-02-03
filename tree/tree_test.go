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

func TestFindNode(t *testing.T) {
	Convey("Given an initial Tree with any name", t, func() {
		chart, _ := Create("normal")
		var rootNode *Node

		Convey("Then I should be able to directly find it's Root Node", func() {
			rootNode, _ = chart.GetRoot()

			Convey("When I ask to find the Root Node using the FindNode method", func() {
				foundNode, err := chart.FindNode((*rootNode).Id, nil)

				Convey("We should not have any error", func() {
					So(err, ShouldEqual, nil)
				})

				Convey("And the found node should equal the root node", func() {
					So(rootNode, ShouldEqual, foundNode)
				})
			})
		})

	})
}
