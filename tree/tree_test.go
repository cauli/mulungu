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
				ID: "1",
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

func TestInsertNode(t *testing.T) {
	Convey("Given an initial Tree with any valid name", t, func() {
		chart, _ := Create("normal")

		Convey("When I insert a new node to the root node", func() {
			rootNode, _ := chart.GetRoot()
			insertedNode := &Node{ID: "2"}

			err := chart.InsertNode(insertedNode, rootNode)

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

			aErr := chart.InsertNode(&aNode, rootNode)
			bErr := chart.InsertNode(&bNode, rootNode)
			cErr := chart.InsertNode(&cNode, &aNode)
			dErr := chart.InsertNode(&dNode, rootNode)

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

func TestGetDescendants(t *testing.T) {
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

			chart.InsertNode(&aNode, rootNode)
			chart.InsertNode(&bNode, rootNode)
			chart.InsertNode(&cNode, &aNode)
			chart.InsertNode(&dNode, rootNode)
			chart.InsertNode(&eNode, &cNode)
			chart.InsertNode(&fNode, &cNode)

			Convey("Then when I get the descendants of the root node", func() {
				response, subErr := rootNode.GetDescendants()

				Convey("I should not have an error", func() {
					So(subErr, ShouldEqual, nil)
				})

				Convey("The count of direct subordinates should equal 3", func() {
					So(response.subordinates.count.direct, ShouldEqual, 3)
				})

				Convey("The total count of subordinates should equal 6", func() {
					So(response.subordinates.count.total, ShouldEqual, 6)
				})

				Convey("The `hierarchy` should resemble expected JSON", func() {
					expectedHierarchyJSON := []byte(`[{
							"id": "a",
							"parentId": "1",
							"children": [{
								"id": "c",
								"parentId": "a",
								"children": [{
									"id": "e",
									"parentId": "c"
								}, {
									"id": "f",
									"parentId": "c"
								}]
							}]
						},
						{
							"id": "b",
							"parentId": "1"
						},
						{
							"id": "d",
							"parentId": "1"
						}
					]`)
					expectedHierarchy := []*Node{}
					unmarshallError := json.Unmarshal(expectedHierarchyJSON, &expectedHierarchy)

					Convey("I should not have an unmarshall error", func() {
						So(unmarshallError, ShouldEqual, nil)
					})

					resultingHierarchyJSON, _ := json.Marshal(response.subordinates.hierarchy)
					resultingHierarchy := []*Node{}
					resultingErr := json.Unmarshal([]byte(resultingHierarchyJSON), &resultingHierarchy)

					Convey("I should not have an unmarshall error for the resulting structure", func() {
						So(resultingErr, ShouldEqual, nil)
					})

					Convey("It resembles the expected JSON", func() {
						So(response.subordinates.hierarchy, ShouldResemble, expectedHierarchy)
					})

				})
			})
		})
	})
}
