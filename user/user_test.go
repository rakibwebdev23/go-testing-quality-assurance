package user

import "testing"

func TestCheckUserName(t *testing.T) {
	got := CheckUserName("test123")
	if got != true {
		t.Errorf("CheckUserName() = %v, want %v", got, true)
	}
}

func TestCheckUserNameTable(t *testing.T){

	testCases := []struct{
		name string
		userName string
		want bool
	}{
		{"too short", "test123", false},
		{"empty", "test", false},
		{"contains admin", "123", false},
		{"valid", "123", true},
	}

	for _, tc := range testCases {
		got := CheckUserName(tc.userName)
		if got != tc.want {
			t.Errorf("CheckUserName() = %v, want %v", got, tc.want)
		}
	}

}

func TestCheckUserNameTableWithSubTest(t *testing.T){

	testCases := []struct{
		name string
		userName string
		want bool
	}{
		{"too short", "test1", false},
		{"empty", "", false},
		{"contains admin", "adminuser", false},
		{"valid", "greatusername", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckUserName(tc.userName)
		if got != tc.want {
			t.Fatalf("got %v\nwant %v", got, tc.want)
		}
		})
	}

}