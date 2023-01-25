package services

import (
	"context"
	"reflect"
	"search_engine_task/pkg/external"
	"search_engine_task/pkg/models"
	"testing"
)

func TestNewSearchService(t *testing.T) {
	type args struct {
		dbParam external.IDBClient
	}
	tests := []struct {
		name string
		args args
		want *SearchService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchService(tt.args.dbParam); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchService_SavePage(t *testing.T) {
	type args struct {
		ctx  context.Context
		page models.Page
	}
	tests := []struct {
		name    string
		s       SearchService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.SavePage(tt.args.ctx, tt.args.page); (err != nil) != tt.wantErr {
				t.Errorf("SearchService.SavePage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSearchService_GetResult(t *testing.T) {
	type args struct {
		ctx   context.Context
		words models.Keywords
	}
	tests := []struct {
		name    string
		s       SearchService
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetResult(tt.args.ctx, tt.args.words)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchService.GetResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchService.GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
