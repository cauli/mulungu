package tree

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	Convey("Given an initial tree with a valid id", t, func() {
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

			Convey("Then we should not have any error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the value of the resulting tree must resemble the initial tree", func() {
				So(*resultingTree, ShouldResemble, initialTree)
			})
		})
	})
}

func TestFindNode(t *testing.T) {
	Convey("Given an initial Tree with any valid name", t, func() {
		chart, _ := Create("normal")
		var rootNode *Node

		Convey("When I try to access it's Root Node directly", func() {
			var rootErr error
			rootNode, rootErr = chart.GetRoot()

			Convey("Then we should not have any error", func() {
				So(rootErr, ShouldEqual, nil)
			})

			Convey("When I ask to find the Root Node using the FindNode method", func() {
				foundNode, err := chart.FindNode((*rootNode).Id, nil)

				Convey("Then we should not have any error", func() {
					So(err, ShouldEqual, nil)
				})

				Convey("And the found node should equal the root node", func() {
					So(rootNode, ShouldEqual, foundNode)
				})
			})
		})
	})
}

func TestInsertNode(t *testing.T) {
	Convey("Given an initial Tree with any valid name", t, func() {
		chart, _ := Create("normal")

		Convey("When I insert a new node to the root node", func() {
			rootNode, _ := chart.GetRoot()
			insertedNode := &Node{Id: "2"}

			err := chart.InsertNode(insertedNode, rootNode)

			Convey("Then we should not have any error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the new node must be findable", func() {
				foundNode, _ := chart.FindNode(insertedNode.Id, nil)
				So(*insertedNode, ShouldResemble, *foundNode)
			})
		})
	})

	Convey("Given an initial Tree with any valid name", t, func() {
		chart, _ := Create("normal")

		Convey("When I insert many nodes", func() {
			rootNode, _ := chart.GetRoot()
			aNode := &Node{Id: "a"}
			bNode := &Node{Id: "b"}
			cNode := &Node{Id: "c"}

			aErr := chart.InsertNode(aNode, rootNode)
			bErr := chart.InsertNode(bNode, rootNode)
			cErr := chart.InsertNode(cNode, aNode)

			Convey("Then we should not have an error inserting `a`", func() {
				So(aErr, ShouldEqual, nil)
			})

			Convey("Then we should not have an error inserting `b`", func() {
				So(bErr, ShouldEqual, nil)
			})

			Convey("Then we should not have an error inserting `c`", func() {
				So(cErr, ShouldEqual, nil)
			})

			Convey("Then the deepest node must be findable by ID", func() {
				foundNode, errFindC := chart.FindNode(cNode.Id, nil)

				Convey("and then should not have an error finding `c`", func() {
					So(errFindC, ShouldEqual, nil)
				})

				Convey("and then the foundNode should not be nil", func() {
					So(foundNode, ShouldNotEqual, nil)
				})

				spew.Dump(chart.ToJSON())

				Convey("and then the foundNode.Id should be `c`", func() {
					So((*foundNode).Id, ShouldEqual, "c")
				})
			})
		})
	})
}
