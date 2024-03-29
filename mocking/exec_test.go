package mocking

import (
	"errors"
	"testing"
)

func TestDoSomeStuff(t *testing.T) {
	tests := []struct {
		name       string
		DoStuff    error
		ThrowError error
		wantErr    bool
	}{
		{"base-case", nil, nil, false},
		{"DoStuff error", errors.New("failed"), nil, true},
		{"ThrowError error", nil, errors.New("failed"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//An example of mocking an interface with our mock struct
			d := MockDoStuffer{}
			d.MockDoStuff = func(string) error {
				return tt.DoStuff
			}

			//Mocking a function that is declared as a variable will
			//not work for func A(), muct be var A=func()
			defer patch(&ThrowError, func() error { return tt.ThrowError }).Restore()

			if err := DoSomeStuff(&d); (err != nil) != tt.wantErr {
				t.Errorf("DoSomeStuff () error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
