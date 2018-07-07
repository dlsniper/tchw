package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "works",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, "/", nil)
					return r
				}(),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			w, r := tt.args.w, tt.args.r
			Home(w, r)
			wo := w.(*httptest.ResponseRecorder)
			if wo.Code != http.StatusOK {
				t.Logf("status mismatch. expected: %d, got: %d", http.StatusOK, wo.Code)
				t.Fail()
			}
		})
	}
}
