package file

import (
	"testing"
)

func TestExistDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: ".から始まり存在する",
			args: args{path: "./conv"},
			want: false,
		},
		{
			name: ".から始まり存在しない",
			args: args{"./kadai1"},
			want: false,
		},
		{
			name: "..を使って存在する",
			args: args{path: "../"},
			want: true,
		},
		{
			name: "..を使って存在しない",
			args: args{path: "../kadai"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExistDir(tt.args.path); got != tt.want {
				t.Errorf("ExistDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetImgFiles1(t *testing.T) {
	type args struct {
		path     string
		beforeEx string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "階層が１",
			args: args{
				path:     "../dummyImg/png",
				beforeEx: "png",
			},
			wantErr: false,
		},
		{
			name: "階層が２",
			args: args{
				path:     "../dummyImg/",
				beforeEx: "png",
			},
			wantErr: false,
		},
		{
			name: "最初から存在しないパス",
			args: args{
				path:     "../path",
				beforeEx: "png",
			},
			wantErr: true,
		},
		{
			name: "途中までは存在するパス",
			args: args{
				path:     "../conv/path",
				beforeEx: "png",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetImgFiles(tt.args.path, tt.args.beforeEx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImgFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
