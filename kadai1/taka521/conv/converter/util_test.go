package converter

import (
	"testing"

	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv/constant"
)

func Test_convertedPath(t *testing.T) {
	type args struct {
		filePath string
		to       constant.Extension
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{
			name:  "フルパス、ディレクトリパス、ファイル名が返ること",
			args:  args{filePath: "/tmp/avatar.png", to: constant.ExtensionGif},
			want:  "/tmp/converted/avatar.gif",
			want1: "/tmp/converted",
			want2: "avatar.gif",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := convertedPath(tt.args.filePath, tt.args.to)
			if got != tt.want {
				t.Errorf("convertedPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("convertedPath() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("convertedPath() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
