package tree

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateHeight(t *testing.T) {
	Convey("Given a tree with a deep linear structure", t, func() {
		chart := New("line")
		rootNode, _ := chart.GetRoot()
		aNode := Node{ID: "a"}
		bNode := Node{ID: "b"}
		cNode := Node{ID: "c"}
		dNode := Node{ID: "d"}
		eNode := Node{ID: "e"}
		fNode := Node{ID: "f"}

		chart.AttachNode(&aNode, rootNode)
		chart.AttachNode(&bNode, &aNode)
		chart.AttachNode(&cNode, &bNode)
		chart.AttachNode(&dNode, &cNode)
		chart.AttachNode(&eNode, &dNode)
		chart.AttachNode(&fNode, &eNode)

		Convey("When the initial chart is ready", func() {
			Convey("The height of `f` should be 6", func() {
				So(fNode.Height, ShouldEqual, 6)
			})

			Convey("The height of `d` should be 4", func() {
				So(dNode.Height, ShouldEqual, 4)
			})

			Convey("The height of `b` should be 2", func() {
				So(bNode.Height, ShouldEqual, 2)
			})

			Convey("The height of `root` should be 0", func() {
				So(rootNode.Height, ShouldEqual, 0)
			})
		})

		Convey("When I move `d` to `root`", func() {
			chart.MoveNode(&dNode, rootNode)

			Convey("The height of `f` should now be 3", func() {
				So(fNode.Height, ShouldEqual, 3)
			})

			Convey("The height of `d` should now be 1", func() {
				So(dNode.Height, ShouldEqual, 1)
			})

			Convey("The height of `b` should still be 2", func() {
				So(bNode.Height, ShouldEqual, 2)
			})

			Convey("The height of `root` should be 0", func() {
				So(rootNode.Height, ShouldEqual, 0)
			})
		})
	})
}

func TestGetDescendants(t *testing.T) {
	Convey("Given an initial Tree with any valid name", t, func() {
		chart := New("normal")

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

			Convey("Then when I get the descendants of the root node", func() {
				response := rootNode.GetDescendants()

				Convey("The count of direct subordinates should equal 3", func() {
					So(response.Subordinates.Count.Direct, ShouldEqual, 3)
				})

				Convey("The total count of subordinates should equal 6", func() {
					So(response.Subordinates.Count.Total, ShouldEqual, 6)
				})

				Convey("The `hierarchy` should resemble expected JSON", func() {
					expectedHierarchyJSON := []byte(`[{
							"id": "a",
							"height": 1,
							"parentId": "root",
							"rootId": "root",
							"children": [{
								"id": "c",
								"height": 2,
								"parentId": "a",
								"rootId": "root",
								"children": [{
									"id": "e",
									"height": 3,
									"parentId": "c",
									"rootId": "root"
								}, {
									"id": "f",
									"height": 3,
									"parentId": "c",
									"rootId": "root"
								}]
							}]
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
			})
		})
	})
}
