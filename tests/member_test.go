package apis

import (
	"go-member-api/test_data"
	"net/http"
	"testing"
)

func TestMember(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get a member", "GET", "/members/:id", "/members/client_member_id/1", "", GetUser, http.StatusOK, path + "/user_t1.json"},
		{"t2 - get a member not present", "GET", "/members/:id", "/members/client_member_id/9999", "", GetUser, http.StatusNotFound, ""},
	})
}
