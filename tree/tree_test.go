package tree

import (
	"testing"

	"github.com/go-test/deep"
)

func TestCreate(t *testing.T) {
	type args struct {
		treeID string
	}

	tests := []struct {
		name    string
		args    args
		want    *Tree
		wantErr bool
	}{
		{
			"Create tree with normal name",
			args{
				"normal",
			},
			&Tree{
				Id: "normal",
				Root: &Node{
					Id: "1",
					Data: MetaData{
						Name:  "#1",
						Title: "Founder",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.treeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("tree.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
