package leetcode

import (
	"problems-everyday-bot/model/response"
	"reflect"
	"testing"
)

func TestGetTodayQuestion(t *testing.T) {
	tests := []struct {
		name    string
		want    *response.GetQuestionOfTodayResp
		wantErr bool
	}{
		{
			name:    "test",
			want:    nil,
			wantErr: false,
		},
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTodayQuestion()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTodayQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTodayQuestion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTodayQuestionDetail(t *testing.T) {
	tests := []struct {
		name    string
		want    *response.GetQuestionOfTodayDetailResp
		wantErr bool
	}{
		{
			name:    "test",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTodayQuestionDetail()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTodayQuestionDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTodayQuestionDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
