
package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUserName(t *testing.T) {
	assert.True(t, CheckUserName("test123"))
}

func TestCheckUserNameTable(t *testing.T) {
	testCases := []struct {
		name     string
		userName string
		want     bool
	}{
		{"too short", "test123", true},
		{"empty", "test", false},
		{"contains admin", "123", false},
		{"valid", "123", false},
	}

	for _, tc := range testCases {
		got := CheckUserName(tc.userName)
		assert.Equal(t, tc.want, got, "test case: %s", tc.name)
	}
}

func TestCheckUserNameTableWithSubTest(t *testing.T) {
	testCases := []struct {
		name     string
		userName string
		want     bool
	}{
		{"too short", "test1", false},
		{"empty", "", false},
		{"contains admin", "adminuser", false},
		{"valid", "greatusername", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckUserName(tc.userName)
			assert.Equal(t, tc.want, got, "test case: %s", tc.name)
		})
	}
}

func TestLogin (t *testing.T) {
	err, ok := Login ("testusername", "testpassword")
	assert.NoError(t, err)
	assert.True(t, ok)
}
