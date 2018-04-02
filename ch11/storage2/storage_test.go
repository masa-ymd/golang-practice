package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	// 他のテストで元に戻せるよう、notifyUserを一時保存し、deferでもとに戻す
	saved := notifyUser
	defer func() { notifyUser = saved }()

	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	usage["joe@example.org"] = 980000000

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}
