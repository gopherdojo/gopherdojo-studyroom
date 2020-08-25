package conv

import "testing"

func Test_checkOpt(t *testing.T) {
	type args struct {
		before string
		after  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "-aが違う",
			args:    args{after: "d"},
			wantErr: true,
		},
		{
			name:    "-bが違う",
			args:    args{before: "d"},
			wantErr: true,
		},
		{
			name: "-a, -b 両方違う",
			args: args{
				before: "d",
				after:  "d",
			},
			wantErr: true,
		},
		{
			name:    "-bのデフォルトと同じ拡張子を-aで指定",
			args:    args{after: "jpeg"},
			wantErr: false,
		},
		{
			name:    "-aのデフォルトと同じ拡張子を-bで指定",
			args:    args{before: "png"},
			wantErr: false,
		},
		{
			name:    "-aが大文字で正しい",
			args:    args{after: "GIF"},
			wantErr: false,
		},
		{
			name:    "-aが大文字で正しくない",
			args:    args{after: "D"},
			wantErr: true,
		},
		{
			name:    "-bが大文字で正しい",
			args:    args{before: "GIF"},
			wantErr: false,
		},
		{
			name:    "-bが大文字で正しくない",
			args:    args{before: "D"},
			wantErr: true,
		},
		{
			name: "-a, -b 両方大文字で違う",
			args: args{
				before: "D",
				after:  "D",
			},
			wantErr: true,
		},
		{
			name: "-a, -b 両方大文字で正しい",
			args: args{
				before: "JPEG",
				after:  "GIF",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkOpt(tt.args.before, tt.args.after); (err != nil) != tt.wantErr {
				t.Errorf("checkOpt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
