package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

const defaultRootFolderName = "lbnetwork"

type PathTransformFunc = func(string) PathKey

type PathKey struct {
	Pathname string
	Filename string
}

func (p PathKey) FirstPathName() string {
	paths := strings.Split(p.Pathname, "/")
	if len(paths) == 0 {
		return ""
	}
	return paths[0]
}

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s/%s", p.Pathname, p.Filename)
}

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{
		Filename: key,
		Pathname: key,
	}
}

// Content addressable storage
// Store data base on thier content rather then location
func CASPathTransform(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	// neat trick to convert array to slice
	hashStr := hex.EncodeToString(hash[:])
	blockSize := 5
	howManyBlocks := len(hashStr) / blockSize
	paths := make([]string, howManyBlocks)

	for i := 0; i < howManyBlocks; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Filename: hashStr,
	}
}

type StoreOpts struct {
	// Root folder name, containing all system folders/files
	Root              string
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if len(opts.Root) == 0 {
		opts.Root = defaultRootFolderName
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Write(id, key string, r io.Reader) (int64, error) {
	return s.writeStream(id, key, r)
}

func (s *Store) Clear() error {
	return os.RemoveAll(s.Root)
}

func (s *Store) writeStream(id, key string, r io.Reader) (int64, error) {
	f, err := s.openFileForWriting(id, key)
	if err != nil {
		return 0, err
	}
	return io.Copy(f, r)
}
func (s *Store) openFileForWriting(id, key string) (*os.File, error) {
	pathKey := s.PathTransformFunc(key)
	pathNameWithRoot := fmt.Sprintf("%s/%s/%s", s.Root, id, pathKey.Pathname)
	if err := os.MkdirAll(pathNameWithRoot, os.ModePerm); err != nil {
		return nil, err
	}
	fullPathWithRoot := fmt.Sprintf("%s/%s/%s", s.Root, id, pathKey.FullPath())
	return os.Create(fullPathWithRoot)
}