package models

import (
	"reflect"
	"testing"
)

func TestGetArticles(t *testing.T) {
	alist := GetArticles()
	if len(alist) != len(GetArticles()) {
		t.Fail()
	}
	if !reflect.DeepEqual(alist, GetArticles()) {
		t.Fail()
	}
}
