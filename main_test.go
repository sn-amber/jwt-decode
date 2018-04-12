package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_base64ToJSON(t *testing.T) {
	type args struct {
		str     string
		utcFlag bool
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "signing method",
			args: args{
				str:     "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9",
				utcFlag: false,
			},
			want: map[string]interface{}{
				"alg": "RS256",
				"typ": "JWT",
			},
			wantErr: false,
		}, {
			name: "test utc flag on",
			args: args{
				str:     "eyJleHAiOjE1MjM1NDY5NzYsIm5hbWUiOiJmb28iLCJraW5kIjoibmlsIn0",
				utcFlag: true,
			},
			want: map[string]interface{}{
				"exp":  time.Date(2018, 4, 12, 17, 29, 36, 0, time.Local),
				"kind": "nil",
				"name": "foo",
			},
			wantErr: false,
		}, {
			name: "test utc flag off",
			args: args{
				str:     "eyJleHAiOjE1MjM1NDY5NzYsIm5hbWUiOiJmb28iLCJraW5kIjoibmlsIn0",
				utcFlag: false,
			},
			want: map[string]interface{}{
				"exp":  1523546976.0,
				"kind": "nil",
				"name": "foo",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base64ToJSON(tt.args.str, tt.args.utcFlag)
			if (err != nil) != tt.wantErr {
				t.Errorf("base64ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got, ok := got.(map[string]interface{}); ok {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("base64ToJSON() = %v, want %v", got, tt.want)
				}
			} else {
				t.Errorf("couldn't convert parsed result to map[string]interface{}")
			}
		})
	}
}
