package ehsm

import "testing"

func TestParseReference(t *testing.T) {
	tests := []struct {
		in      string
		wantKey string
		wantErr bool
	}{
		{
			in:      "ehsm://cosign",
			wantKey: "cosign",
			wantErr: false,
		},
		{
			in:      "foo://bar",
			wantErr: true,
		},
		{
			in:      "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			gotKey, err := parseReference(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("parseReference() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
