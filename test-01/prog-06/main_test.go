package main

import "testing"

func TestAll(t *testing.T) {

	type con struct {
		main   string
		check  string
		answer bool
	}

	testcon := []con{
		con{"I am here", "here", true},
		con{"I am here", "where", false},
		con{"You are amazing", "are", true},
		con{"Where are you", "you", true},
		con{"is it true", "true", true},
	}

	for _, v := range testcon {
		if conCheck(v.main, v.check) != v.answer {
			t.Error("Expecting", v.answer, "but got", conCheck(v.main, v.check))
		}
	}

}
