package main

import (
	"reflect"
	"testing"
)

func Test_findDeret(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			args:    args{arr: []int{7, 1, 2, 9, 7, 2, 1}},
			want:    []int{1, 2},
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{arr: []int{7, 1, 2, 3, 3, 2, 1}},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "test3",
			args:    args{arr: []int{2, 1, 5, 3, 3, 8, 1}},
			want:    []int{},
			wantErr: true,
		},
		{
			name:    "test4",
			args:    args{arr: []int{}},
			want:    []int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findDeret(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("findDeret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDeret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDeretReverse(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			args:    args{arr: []int{7, 1, 2, 9, 7, 2, 1}},
			want:    []int{1, 2},
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{arr: []int{7, 1, 2, 3, 3, 2, 1}},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "test3",
			args:    args{arr: []int{2, 1, 5, 3, 3, 8, 1, 5}},
			want:    []int{},
			wantErr: true,
		},
		{
			name:    "test4",
			args:    args{arr: []int{}},
			want:    []int{},
			wantErr: true,
		},
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findDeretReverse(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("findDeretReverse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDeretReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkHighest(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			args:    args{arr: []int{7, 1, 2, 9, 7, 2, 1}},
			want:    2,
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{arr: []int{7, 1, 2, 3, 3, 2, 1}},
			want:    3,
			wantErr: false,
		},
		{
			name:    "test3",
			args:    args{arr: []int{7, 1, 8, 3, 9, 2, 1}},
			want:    0,
			wantErr: true,
		},
		{
			name:    "test4",
			args:    args{arr: []int{7, 1, 2, 3, 8, 3, 9, 2, 1}},
			want:    0,
			wantErr: true,
		},
		{
			name:    "test5",
			args:    args{arr: []int{}},
			want:    0,
			wantErr: true,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkHighest(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkHighest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}
