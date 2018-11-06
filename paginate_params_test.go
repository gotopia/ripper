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
		name    string
		args    args
		wantP   *PaginateParams
		wantErr bool
	}{
		{"default page should be 1", args{pageSize: 15}, &PaginateParams{Page: 1, PageSize: 15}, false},
		{"default pageSize should be 15", args{page: 1}, &PaginateParams{Page: 1, PageSize: 15}, false},
		{"page should be larger than 0", args{page: -1, pageSize: 15}, nil, true},
		{"pageSize should be larger than 0", args{page: 1, pageSize: -1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := NewPaginateParams(tt.args.page, tt.args.parent, tt.args.pageSize, tt.args.orderBy, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPaginateParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("NewPaginateParams() = %v, want %v", gotP, tt.wantP)
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
				p, _ := NewPaginateParams(1, "parents/1", 15, "update_time", "name=@string")
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
