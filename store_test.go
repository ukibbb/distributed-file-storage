package main

import (
	"bytes"
	"fmt"
	"io"
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

	defer tearDown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("Some test file data.")

		if _, err := s.writeStream(id, key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); !ok {
			t.Errorf("expected to have key %s", key)
		}

		_, r, err := s.Read(id, key)
		if err != nil {
			t.Error(err)
		}
		b, _ := io.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("want %s have %s", data, b)
		}
		if err := s.Delete(id, key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); ok {
			t.Errorf("expected to NOT have key %s", key)
		}

	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransform,
	}
	return NewStore(opts)
}

func tearDown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
