package cache

import (
	"sync"
	"testing"
	"time"
)

func TestCache_Query(t *testing.T) {
	type fields struct {
		requestd time.Duration
		timeout  time.Duration
		value    string
		RWMutex  sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Query longer than timeout 1",
			fields: fields{timeout: 300 * time.Millisecond, requestd: 500 * time.Millisecond, value: "cached"},
			want:   "cached",
		},
		{
			name:   "Query longer than timeout 2",
			fields: fields{timeout: 1 * time.Millisecond, requestd: 2 * time.Millisecond, value: "cached"},
			want:   "cached",
		},
		{
			name:   "Query longer than timeout 3",
			fields: fields{timeout: 100 * time.Millisecond, requestd: 123 * time.Millisecond, value: "cached"},

			want: "cached",
		},
		{
			name:   "Query longer than timeout 4",
			fields: fields{timeout: 1 * time.Second, requestd: 2 * time.Second, value: "cached"},
			want:   "cached",
		},
		{
			name:   "Query shorter than timeout 5",
			fields: fields{timeout: 1 * time.Second, requestd: 500 * time.Millisecond, value: "cached"},
			want:   "QueryValue",
		},
		{
			name:   "Query shorter than timeout 6",
			fields: fields{timeout: 2 * time.Millisecond, requestd: 500 * time.Microsecond, value: "cached"},
			want:   "QueryValue",
		},
		{
			name:   "Query shorter than timeout 7",
			fields: fields{timeout: 400 * time.Millisecond, requestd: 300 * time.Millisecond, value: "cached"},
			want:   "QueryValue",
		},
		{
			name:   "Query shorter than timeout 8",
			fields: fields{timeout: 1 * time.Second, requestd: 1 * time.Millisecond, value: "cached"},
			want:   "QueryValue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				requestd: tt.fields.requestd,
				timeout:  tt.fields.timeout,
				value:    tt.fields.value,
				RWMutex:  tt.fields.RWMutex,
			}
			if got := c.Query(); got != tt.want {
				t.Errorf("Cache.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
