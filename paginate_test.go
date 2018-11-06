package ripper

import (
	"testing"
)

func Test_paginate(t *testing.T) {
	type args struct {
		p         *PaginateParams
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
		{"nextPageToken should be empty if there is no next page", args{p: &PaginateParams{Page: 1, PageSize: 15}, totalSize: 15}, 0, 15, "", false},
		{"Page should be larger than zero", args{p: &PaginateParams{Page: -1, PageSize: 15}, totalSize: 15}, 0, 0, "", true},
		{"PageSize should be larger than zero", args{p: &PaginateParams{Page: 1, PageSize: -15}, totalSize: 15}, 0, 0, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOffset, gotLimit, gotNextPageToken, err := paginate(tt.args.p, tt.args.totalSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("paginate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOffset != tt.wantOffset {
				t.Errorf("paginate() gotOffset = %v, want %v", gotOffset, tt.wantOffset)
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("paginate() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
			if gotNextPageToken != tt.wantNextPageToken {
				t.Errorf("paginate() gotNextPageToken = %v, want %v", gotNextPageToken, tt.wantNextPageToken)
			}
		})
	}
}

func Test_hasNextPage(t *testing.T) {
	type args struct {
		page      int
		pageSize  int
		totalSize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"flag shoule be true if there is next page", args{page: 1, pageSize: 15, totalSize: 30}, true},
		{"flag shoule be false if there is no next page", args{page: 2, pageSize: 15, totalSize: 15}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasNextPage(tt.args.page, tt.args.pageSize, tt.args.totalSize); got != tt.want {
				t.Errorf("hasNextPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
