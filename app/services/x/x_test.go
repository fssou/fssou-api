package x

import (
	"testing"
)

func TestMe(t *testing.T) {
	x := X{}
	me, err := x.Me()
	if err != nil {
		t.Errorf("Me() retornou um erro: %v", err)
	}
	expectedUsername := "fssouoficial"
	if me.Data.Username != expectedUsername {
		t.Errorf("Expected username %v, but got %v", expectedUsername, me.Data.Username)
	}
}
