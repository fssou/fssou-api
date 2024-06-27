package x

import (
	"context"
	"testing"
)

func TestMe(t *testing.T) {
	x, err := New(context.Background())
	if err != nil {
		t.Errorf("New() retornou um erro: %v", err)
	}
	me, err := x.Me()
	if err != nil {
		t.Errorf("Me() retornou um erro: %v", err)
	}
	expectedUsername := "fssouoficial"
	if me.Data.Username != expectedUsername {
		t.Errorf("Expected username %v, but got %v", expectedUsername, me.Data.Username)
	}
}
