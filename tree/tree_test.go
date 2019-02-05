package tree

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	Convey("Given an initial tree with a valid id", t, func() {
		name := "normal"
		initialTree := Tree{
			Id: name,
			Root: &Node{
				ID: "root",
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
				foundNode, err := chart.FindNode((*rootNode).ID, nil)

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

func TestAttachNode(t *testing.T) {
	Convey("Given an initial Tree with any valid name", t, func() {
		chart, _ := Create("normal")

		Convey("When I insert a new node to the root node", func() {
			rootNode, _ := chart.GetRoot()
			insertedNode := &Node{ID: "2"}

			err := chart.AttachNode(insertedNode, rootNode)

			Convey("Then we should not have any error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the new node must be findable", func() {
				foundNode, _ := chart.FindNode(insertedNode.ID, nil)
				So(*insertedNode, ShouldResemble, *foundNode)
			})
		})
	})

	Convey("Given an initial Tree with any valid name", t, func() {
		chart, _ := Create("normal")

		Convey("When I insert many nodes", func() {
			rootNode, _ := chart.GetRoot()
			aNode := Node{ID: "a"}
			bNode := Node{ID: "b"}
			cNode := Node{ID: "c"}
			dNode := Node{ID: "d"}

			aErr := chart.AttachNode(&aNode, rootNode)
			bErr := chart.AttachNode(&bNode, rootNode)
			cErr := chart.AttachNode(&cNode, &aNode)
			dErr := chart.AttachNode(&dNode, rootNode)

			Convey("Then we should not have an error inserting `a`", func() {
				So(aErr, ShouldEqual, nil)
			})

			Convey("Then we should not have an error inserting `b`", func() {
				So(bErr, ShouldEqual, nil)
			})

			Convey("Then we should not have an error inserting `c`", func() {
				So(cErr, ShouldEqual, nil)
			})

			Convey("Then we should not have an error inserting `d`", func() {
				So(dErr, ShouldEqual, nil)
			})

			Convey("Then the deepest node must be findable by ID", func() {
				foundNode, errFindC := chart.FindNode(cNode.ID, nil)

				Convey("and should not have an error finding `c`", func() {
					So(errFindC, ShouldEqual, nil)
				})

				Convey("and the foundNode should not be nil", func() {
					So(foundNode, ShouldNotEqual, nil)
				})

				Convey("and foundNode.Id should be `c`", func() {
					So((*foundNode).ID, ShouldEqual, "c")
				})
			})
		})
	})
}

func TestDetachNode(t *testing.T) {
	Convey("Given an initial Tree with any valid name", t, func() {
		chart, _ := Create("normal")

		Convey("When I insert nodes to create a complex tree", func() {
			rootNode, _ := chart.GetRoot()
			aNode := Node{ID: "a"}
			bNode := Node{ID: "b"}
			cNode := Node{ID: "c"}
			dNode := Node{ID: "d"}
			eNode := Node{ID: "e"}
			fNode := Node{ID: "f"}

			chart.AttachNode(&aNode, rootNode)
			chart.AttachNode(&bNode, rootNode)
			chart.AttachNode(&cNode, &aNode)
			chart.AttachNode(&dNode, rootNode)
			chart.AttachNode(&eNode, &cNode)
			chart.AttachNode(&fNode, &cNode)

			Convey("Then when I get detach the node `c`", func() {
				detachedNode, detachErr := chart.DetachNode(&cNode)

				Convey("I should not have an error", func() {
					So(detachErr, ShouldEqual, nil)
				})

				response, descErr := rootNode.GetDescendants()

				Convey("The count of direct subordinates should equal 3", func() {
					So(response.Subordinates.Count.Direct, ShouldEqual, 3)
				})

				Convey("I should not have an error getting the descendants of the detached tree", func() {
					So(descErr, ShouldEqual, nil)
				})

				Convey("The `hierarchy` of the tree should resemble expected JSON", func() {
					expectedHierarchyJSON := []byte(`[{
							"id": "a",
							"height": 1,
							"parentId": "root",
							"rootId": "root"
						},
						{
							"id": "b",
							"height": 1,
							"parentId": "root",
							"rootId": "root"
						},
						{
							"id": "d",
							"height": 1,
							"parentId": "root",
							"rootId": "root"
						}
					]`)
					expectedHierarchy := []*Node{}
					json.Unmarshal(expectedHierarchyJSON, &expectedHierarchy)

					resultingHierarchyJSON, _ := json.Marshal(response.Subordinates.Hierarchy)
					resultingHierarchy := []*Node{}
					resultingErr := json.Unmarshal([]byte(resultingHierarchyJSON), &resultingHierarchy)

					Convey("I should not have an unmarshall error for the resulting structure", func() {
						So(resultingErr, ShouldEqual, nil)
					})

					Convey("It resembles the expected JSON", func() {
						So(response.Subordinates.Hierarchy, ShouldResemble, expectedHierarchy)
					})
				})

				Convey("And when I attach the detached node to the root", func() {
					chart.AttachNode(detachedNode, chart.Root)

					Convey("The count of direct subordinates should equal 4", func() {
						response, _ := rootNode.GetDescendants()
						So(response.Subordinates.Count.Direct, ShouldEqual, 4)
					})
				})
			})
		})
	})

	Convey("Given an initial simplistic Tree", t, func() {
		chart, _ := Create("normal")
		rootNode, _ := chart.GetRoot()

		Convey("When I get detach the root node", func() {
			_, detachErr := chart.DetachNode(rootNode)

			Convey("I should have an error, because you should not be able to detach a root node", func() {
				So(detachErr, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a linear tree", t, func() {
		chart, _ := Create("complex")
		rootNode, _ := chart.GetRoot()
		aNode := Node{ID: "a"}
		bNode := Node{ID: "b"}
		cNode := Node{ID: "c"}

		chart.AttachNode(&aNode, rootNode)
		chart.AttachNode(&bNode, &aNode)
		chart.AttachNode(&cNode, &bNode)

		Convey("When I detach a higher node (`a`)", func() {
			detachedA, _ := chart.DetachNode(&aNode)

			Convey("And try to attach it to one of its descendants (`c`)", func() {
				attachErr := chart.AttachNode(detachedA, &cNode)

				Convey("Then we should have an error avoiding this edge case", func() {
					So(attachErr, ShouldNotEqual, nil)
				})
			})
		})
	})
}

func TestMoveNode(t *testing.T) {
	Convey("Given a Tree with some nodes", t, func() {
		chart, _ := Create("normal")

		rootNode, _ := chart.GetRoot()
		aNode := Node{ID: "a"}
		bNode := Node{ID: "b"}
		cNode := Node{ID: "c"}
		dNode := Node{ID: "d"}
		fNode := Node{ID: "d"}

		chart.AttachNode(&aNode, rootNode)
		chart.AttachNode(&bNode, rootNode)
		chart.AttachNode(&cNode, &aNode)
		chart.AttachNode(&dNode, rootNode)
		chart.AttachNode(&fNode, rootNode)

		Convey("When the initial chart is ready", func() {
			Convey("The count of direct subordinates of root should equal 4", func() {
				response, _ := rootNode.GetDescendants()
				So(response.Subordinates.Count.Direct, ShouldEqual, 4)
			})

			Convey("The count of direct subordinates of `d` should equal 0", func() {
				response, _ := dNode.GetDescendants()
				So(response.Subordinates.Count.Direct, ShouldEqual, 0)
			})
		})

		Convey("When I move node `a` to new parent `d`", func() {
			moveErr := chart.MoveNode(&aNode, &dNode)

			Convey("I should not have an error", func() {
				So(moveErr, ShouldEqual, nil)
			})

			Convey("The count of direct subordinates of root should equal 3", func() {
				response, _ := rootNode.GetDescendants()
				So(response.Subordinates.Count.Direct, ShouldEqual, 3)
			})

			Convey("The count of direct subordinates of `d` should equal 1", func() {
				response, _ := dNode.GetDescendants()
				So(response.Subordinates.Count.Direct, ShouldEqual, 1)
			})

			Convey("The total count subordinates of `d` should equal 2", func() {
				response, _ := dNode.GetDescendants()
				So(response.Subordinates.Count.Total, ShouldEqual, 2)
			})
		})
	})
}
