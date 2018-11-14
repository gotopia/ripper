package ripper

import (
	"testing"
)

func TestPaginate(t *testing.T) {
	p := NewPaginateParams(1, "", 15, "", "")
	totalSize := 15
	limit, offset, nextPageToken, _ := paginate(p, totalSize)
	type args struct {
		p         *PaginateParams
		pageToken string
		totalSize int
	}
	tests := []struct {
		name              string
		args              args
		wantOffset        int
		wantLimit         int
		wantNextPageToken string
		wantErr           bool
	}{
		{"pageToken could be empty string", args{p: p, pageToken: "", totalSize: 15}, limit, offset, nextPageToken, false},
		{"there should be an error if pageToken is invalid", args{p: p, pageToken: "invalid pageToken", totalSize: 15}, 0, 0, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOffset, gotLimit, gotNextPageToken, err := Paginate(tt.args.p, tt.args.pageToken, tt.args.totalSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Paginate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOffset != tt.wantOffset {
				t.Errorf("Paginate() gotOffset = %v, want %v", gotOffset, tt.wantOffset)
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("Paginate() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
			if gotNextPageToken != tt.wantNextPageToken {
				t.Errorf("Paginate() gotNextPageToken = %v, want %v", gotNextPageToken, tt.wantNextPageToken)
			}
		})
	}
}

func TestSetDefaultPageSize(t *testing.T) {
	type args struct {
		pageSize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"defaultPageSize should be larger than zero", args{pageSize: 25}, true},
		{"defaultPageSize should not be smaller than or equal to zero", args{pageSize: -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDefaultPageSize(tt.args.pageSize); got != tt.want {
				t.Errorf("SetDefaultPageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetMaxPageSize(t *testing.T) {
	type args struct {
		pageSize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"maxPageSize should be larger than or equal to defaultPageSize", args{pageSize: defaultPageSize + 1}, true},
		{"maxPageSize should not be smaller than defaultPageSize", args{pageSize: defaultPageSize - 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetMaxPageSize(tt.args.pageSize); got != tt.want {
				t.Errorf("SetMaxPageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
