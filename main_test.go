package main

import (
	"testing"
)

func Test_base64ToJSON(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			args: args{
				str: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9",
			},
			want: map[string]string{
				"alg": "RS256",
				"typ": "JWT",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base64ToJSON(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("base64ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gott := got.(map[string]interface{})
			want := tt.want.(map[string]string)
			if !(gott["typ"] == want["typ"] || gott["alg"] == want["alg"]) {
				t.Errorf("base64ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
