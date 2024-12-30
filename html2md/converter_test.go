package html2md

import "testing"

func TestSayHello(t *testing.T) {
	output := SayHello()
	want := "hello world"
	if want != output {
		t.Errorf("expected %v but got %v", want, output)
	}
}