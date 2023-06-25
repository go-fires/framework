package jet

import "testing"

func TestURLPathGenerator(t *testing.T) {
	type args struct {
		service string
		method  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				service: "Jet\\Test\\Service",
				method:  "Test",
			},
			want: "/jet_test_service/Test",
		},
		{
			name: "",
			args: args{
				service: "JET\\UserInfoService",
				method:  "GetUserInfo",
			},
			want: "/j_e_t_user_info/GetUserInfo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &URLPathGenerator{}
			if got := d.Generate(tt.args.service, tt.args.method); got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
