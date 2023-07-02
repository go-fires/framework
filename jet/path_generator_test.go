package jet

import "testing"

func TestURLPathGenerator_Generate(t *testing.T) {
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

func TestURLPathGenerator_snake(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"", "FiresGOComponent", "fires_g_o_component"},
		{"", "FiresGoComponent", "fires_go_component"},
		{"", "FiresGoComponent", "fires_go_component"},
		{"", "Fires Go Component", "fires_go_component"},
		{"", "Fires    Go      Component   ", "fires_go_component"},
		{"", "FiresGoComponent", "fires__go__component"},
		{"", "FiresGoComponent_", "fires_go_component_"},
		{"", "fires go Component", "fires_go_component"},
		{"", "fires go MoreComponent", "fires_go_more_component"},
		{"", "foo-bar", "foo-bar"},
		{"", "Foo-Bar", "foo-_bar"},
		{"", "Foo_Bar", "foo__bar"},
		{"", "ŻółtaŁódka", "żółtałódka"},
	}

	g := &URLPathGenerator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.snake(tt.args); got != tt.want {
				t.Errorf("snake(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
