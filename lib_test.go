package protobufs

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"testing"
	"time"
)

func TestCreateTimestamps(t *testing.T) {
	type args struct {
		input CreateTimestampsInput
	}

	layout := "2006-01-02T15:04:05.000Z"
	val := "2014-11-12T11:45:26.371Z"

	ts, _ := time.Parse(layout, val)

	tests := []struct {
		name string
		args args
		want *Timestamps
	}{
		{
			name: "Happy path",
			args: args{
				input: CreateTimestampsInput{
					CreatedAt: ts,
				},
			},
			want: &Timestamps{
				CreatedAt: timestamppb.New(ts),
				UpdatedAt: timestamppb.New(ts),
				DeletedAt: nil,
			},
		},
		{
			name: "Happy path with deleted",
			args: args{
				input: CreateTimestampsInput{
					CreatedAt: ts,
					DeletedAt: ts,
				},
			},
			want: &Timestamps{
				CreatedAt: timestamppb.New(ts),
				UpdatedAt: timestamppb.New(ts),
				DeletedAt: timestamppb.New(ts),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTimestamps(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTimestamps() = %v, want %v", got, tt.want)
			}
		})
	}
}
