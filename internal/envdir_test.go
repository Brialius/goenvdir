package internal

import (
	"reflect"
	"sort"
	"testing"
)

func Test_mergeEnvVars(t *testing.T) {
	type args struct {
		envVars map[string]string
		environ []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"OK merge",
			args{
				map[string]string{
					"FILE1": "VALUE1", "FILE2": "VALUE2", "FILE3_EXISTS": "VALUE3", "EMPTY_NOT_EXISTS": "", "EMPTY_EXISTS": "",
				},
				[]string{
					"=WIN_SPECIFIC_KEY=A:", "OS1=1", "OS2=2", "FILE3_EXISTS=OLD_VALUE", "EMPTY_EXISTS=OLD_VALUE", "OS3=3",
				},
			},
			[]string{
				"FILE1=VALUE1", "FILE2=VALUE2", "FILE3_EXISTS=VALUE3", "OS1=1", "OS2=2", "OS3=3",
			},
		},
		{
			"OK merge empty",
			args{
				map[string]string{
					"FILE1": "VALUE1", "FILE2": "VALUE2", "FILE3_EXISTS": "VALUE3", "EMPTY_NOT_EXISTS": "", "EMPTY_EXISTS": "",
				},
				[]string{
					"=WIN_SPECIFIC_KEY=A:", "OS1=", "OS2=", "FILE3_EXISTS=OLD_VALUE", "EMPTY_EXISTS=OLD_VALUE", "OS3=3",
				},
			},
			[]string{
				"FILE1=VALUE1", "FILE2=VALUE2", "FILE3_EXISTS=VALUE3", "OS1=", "OS2=", "OS3=3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeEnvVars(tt.args.envVars, tt.args.environ)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeEnvVars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvVarsFromDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			"OK testdata",
			args{"testdata"},
			map[string]string{
				"EMPTY":         "",
				"FILE1":         "VALUE1",
				"FILE2":         "VALUE2",
				"SUBLEVELFILE1": "SUBLEVELVALUE1",
				"USER":          "",
				"USERNAME":      "",
			},
			false,
		},
		{
			"OK testdata DIR1",
			args{"testdata/DIR1"},
			map[string]string{
				"SUBLEVELFILE1": "SUBLEVELVALUE1",
			},
			false,
		},
		{
			"Fail not exists",
			args{"testdata/DIR2"},
			nil,
			true,
		},
		{
			"Fail not a dir",
			args{"testdata/FILE1"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getEnvVarsFromDir(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getEnvVarsFromDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEnvVarsFromDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
