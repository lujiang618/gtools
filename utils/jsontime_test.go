package utils

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestJSONTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		tr      JSONTime
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		tr      *JSONTime
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("JSONTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONTime_Value(t *testing.T) {
	tests := []struct {
		name    string
		tr      JSONTime
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONTime.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONTime.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONTime_Scan(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		tr      *JSONTime
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Scan(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("JSONTime.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
