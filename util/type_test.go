package util

import (
	"reflect"
	"testing"
)

func TestBoolToUint32(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name    string
		args    args
		wantI   uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := BoolToUint32(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BoolToUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("BoolToUint32() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestInt32ToByteArr(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name      string
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, err := Int32ToByteArr(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int32ToByteArr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("Int32ToByteArr() gotBytes = %v, want %v", gotBytes, tt.wantBytes)
			}
		})
	}
}

func TestInt64ToByteArr(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name      string
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, err := Int64ToByteArr(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64ToByteArr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("Int64ToByteArr() gotBytes = %v, want %v", gotBytes, tt.wantBytes)
			}
		})
	}
}

func TestInt64ToStr(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := Int64ToStr(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64ToStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("Int64ToStr() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestInt64ToTimestamp(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Int64ToTimestamp(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64ToTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64ToTimestamp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToBool(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantB   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := StrToBool(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("StrToBool() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestStrToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantI   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := StrToInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("StrToInt() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestStrToInt32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantI   int32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := StrToInt32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("StrToInt32() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestStrToInt64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantI   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := StrToInt64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("StrToInt64() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestStrToUin32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantI   uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := StrToUin32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToUin32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("StrToUin32() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestStrToUin64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantI   uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := StrToUin64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToUin64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("StrToUin64() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestUint32ToBool(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name    string
		args    args
		wantB   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := Uint32ToBool(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("Uint32ToBool() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestUint32ToByteArr(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name      string
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, err := Uint32ToByteArr(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32ToByteArr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("Uint32ToByteArr() gotBytes = %v, want %v", gotBytes, tt.wantBytes)
			}
		})
	}
}

func TestUint32ToStr(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := Uint32ToStr(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32ToStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("Uint32ToStr() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestUint64ToByteArr(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name      string
		args      args
		wantBytes []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, err := Uint64ToByteArr(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64ToByteArr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("Uint64ToByteArr() gotBytes = %v, want %v", gotBytes, tt.wantBytes)
			}
		})
	}
}

func TestUint64ToStr(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := Uint64ToStr(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64ToStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("Uint64ToStr() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
