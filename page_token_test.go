package ripper

import (
	"reflect"
	"testing"
)

var p, _ = NewPaginateParams(1, "", 15, "", "")
var pageToken, _ = encodeToPageToken(p)

func Test_validatePageToken(t *testing.T) {
	type args struct {
		p         *PaginateParams
		pageToken string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"pageToken should be corresponding to the PaginateParams", args{p, pageToken}, false},
		{"pageToken should not be corresponding to different PaginateParams", func() args {
			ap, _ := NewPaginateParams(1, "", 150, "name desc", "")
			return args{ap, pageToken}
		}(), true},
		{"arbitrary pageToken should be invalid", func() args {
			s := "arbitrary string"
			return args{p, s}
		}(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePageToken(tt.args.p, tt.args.pageToken); (err != nil) != tt.wantErr {
				t.Errorf("validatePageToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeToPageToken(t *testing.T) {
	type args struct {
		p *PaginateParams
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		{"encode and decode should be interconvertible", func() args {
			p, _ := decodePageToken(pageToken)
			return args{p}
		}(), pageToken, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := encodeToPageToken(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeToPageToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("encodeToPageToken() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_decodePageToken(t *testing.T) {
	type args struct {
		pageToken string
	}
	tests := []struct {
		name    string
		args    args
		wantP   *PaginateParams
		wantErr bool
	}{
		{"pageToken should not be empty", args{""}, nil, true},
		{"encode and decode should be interconvertible", func() args {
			pageToken, _ := encodeToPageToken(p)
			return args{pageToken}
		}(), p, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := decodePageToken(tt.args.pageToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodePageToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("decodePageToken() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
