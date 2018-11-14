package ripper

import (
	"reflect"
	"testing"

	"github.com/gotopia/ripper/proto"
)

func TestNewPaginateParams(t *testing.T) {
	type args struct {
		page     int
		parent   string
		pageSize int
		orderBy  string
		filter   string
	}
	tests := []struct {
		name string
		args args
		want *PaginateParams
	}{
		{"default page should be 1", args{pageSize: 15}, &PaginateParams{Page: 1, PageSize: 15}},
		{"default pageSize should be 15", args{page: 1}, &PaginateParams{Page: 1, PageSize: 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaginateParams(tt.args.page, tt.args.parent, tt.args.pageSize, tt.args.orderBy, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaginateParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItop(t *testing.T) {
	type args struct {
		i IPaginateParams
	}
	tests := []struct {
		name  string
		args  args
		wantP *PaginateParams
	}{
		{
			"pb struct should be converted to PaginateParams",
			args{
				&proto.PaginateParams{Parent: "parents/1", PageSize: 15, PageToken: "page_token", OrderBy: "update_time", Filter: "name=@string"},
			},
			func() *PaginateParams {
				p := NewPaginateParams(1, "parents/1", 15, "update_time", "name=@string")
				return p
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := Itop(tt.args.i); !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("Itop() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
