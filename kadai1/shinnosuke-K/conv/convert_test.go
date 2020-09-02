package conv

import "testing"

func Test_checkOpt(t *testing.T) {
	type args struct {
		ex string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "拡張子が違う（小文字）",
			args:    args{ex: "d"},
			wantErr: true,
		},
		{
			name:    "拡張子が違う（大文字）",
			args:    args{ex: "D"},
			wantErr: true,
		},
		{
			name:    "拡張子が正しい（png）",
			args:    args{ex: "png"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（jpg）",
			args:    args{ex: "jpg"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（jpeg）",
			args:    args{ex: "jpeg"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（gif）",
			args:    args{ex: "gif"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（PNG）",
			args:    args{ex: "PNG"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（JPG）",
			args:    args{ex: "JPG"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（JPEG）",
			args:    args{ex: "JPEG"},
			wantErr: false,
		},
		{
			name:    "拡張子が正しい（GIF）",
			args:    args{ex: "GIF"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkOpt(tt.args.ex); (err != nil) != tt.wantErr {
				t.Errorf("checkOpt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
