package maps

import "testing"

func Test_directiroy_Search(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    directiroy
		args args
		want string
	}{
		// TODO: Add test cases.
		{"directiroy_Search", directiroy{"eng": "英语"}, args{"eng"}, "英语"},
		{"directiroy_Search", directiroy{"num": "123409000000"}, args{"num"}, "123409000000"},
		{"directiroy_Search", directiroy{"特殊字符": "~!@#$%^&*()_+{}|:’<>?+-‘"}, args{"特殊字符"}, "~!@#$%^&*()_+{}|:’<>?+-‘"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Search(tt.args.key); got != tt.want {
				t.Errorf("directiroy.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directiroy_Add(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		d    directiroy
		args args
		want string
	}{
		// TODO: Add test cases.
		{"directiroy_Add", directiroy{"eng": "英语"}, args{key: "Chinese", value: "语文"}, "Chinese"},
		{"directiroy_Add", directiroy{}, args{key: "数字", value: "12432432"}, "数字"},
		{"directiroy_Add", directiroy{}, args{key: "特殊字符", value: "《》？：“{}"}, "特殊字符"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Add(tt.args.key, tt.args.value)
			if got := tt.d.Search(tt.want); got == "" {
				t.Errorf("directiroy.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directiroy_Del(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    directiroy
		args args
	}{
		// TODO: Add test cases.
		{"directiroy_Del", directiroy{"eng": "英语", "Chinese": "国文"}, args{"Chinese"}},
		{"directiroy_Del", directiroy{"num": "123", "float": "123.0123468748"}, args{"num"}},
		{"directiroy_Del", directiroy{"word": "!@#$#@%$^$%&^%", "t": "<>?{}_+)"}, args{"word"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Del(tt.args.key)
			if got := tt.d.Search(tt.args.key); got == "" {
				t.Errorf("directiroy.Del() = %v", got)
			}
		})
	}
}

func Test_directiroy_Update(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		d    directiroy
		args args
		want string
	}{
		// TODO: Add test cases.
		{"directiroy_Add", directiroy{"Chinese": "Chinese"}, args{key: "Chinese", value: "语文"}, "语文"},
		{"directiroy_Add", directiroy{"数字": "数字"}, args{key: "数字", value: "12432432"}, "12432432"},
		{"directiroy_Add", directiroy{"特殊字符": "特殊字符"}, args{key: "特殊字符", value: "《》？：“{}"}, "《》？：“{}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Update(tt.args.key, tt.args.value)
			if got := tt.d.Search(tt.want); got == "" {
				t.Errorf("directiroy.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
