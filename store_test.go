package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransforFunc(t *testing.T) {
	key := "ThisIsTestKey"
	pathKey := CASPathTransform(key)
	expectedFilename := "95c7925d6193cbb1cad26cc5c88d93cc76b35f22"
	expectedPathName := "95c79/25d61/93cbb/1cad2/6cc5c/88d93/cc76b/35f22"
	if pathKey.Pathname != expectedPathName {
		t.Errorf("have %s want %s", pathKey.Pathname, expectedPathName)
	}

	if pathKey.Filename != expectedFilename {
		t.Errorf("have %s want %s", pathKey.Filename, expectedFilename)
	}

}

func TestStore(t *testing.T) {
	s := newStore()
	id := generateID()

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("Some test file data.")

		if _, err := s.writeStream(id, key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransform,
	}
	return NewStore(opts)
}
