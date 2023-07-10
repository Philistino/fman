package fileutils

// func TestCopy(t *testing.T) {
// 	fs := afero.NewMemMapFs()

// 	// Create a source file
// 	src := "/path/to/source/file.txt"
// 	err := fs.MkdirAll(filepath.Dir(src), os.ModePerm)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	srcFile, err := fs.Create(src)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = srcFile.WriteString("Hello, world!")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = srcFile.Close()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Test copying a file
// 	dst := "/path/to/dest/file.txt"
// 	err = Copy(fs, src, dst)
// 	log.Println(afero.Exists(fs, dst))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	dstFile, err := fs.Open(dst)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer dstFile.Close()
// 	dstContent, err := afero.ReadAll(dstFile)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if string(dstContent) != "Hello, world!" {
// 		t.Errorf("Expected destination file content to be 'Hello, world!', but got '%s'", string(dstContent))
// 	}
// }

// func TestCopyForDir(t *testing.T) {
// 	fs := afero.NewMemMapFs()

// 	// Test copying a directory
// 	srcDir := "/path/to/source/dir"
// 	err := fs.MkdirAll(srcDir, os.ModePerm)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	srcFile, err := fs.Create(filepath.Join(srcDir, "file2.txt"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = srcFile.WriteString("Hello again!")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = srcFile.Close()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	dstDir := "/path/to/destination/dir"
// 	err = Copy(fs, srcDir, dstDir)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	dstFile2, err := fs.Open(filepath.Join(dstDir, "file2.txt"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer dstFile2.Close()
// 	dstContent2, err := afero.ReadAll(dstFile2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if string(dstContent2) != "Hello again!" {
// 		t.Errorf("Expected destination file content to be 'Hello again!', but got '%s'", string(dstContent2))
// 	}

// 	// Test copying a non-existent file
// 	err = Copy(fs, "/path/to/non-existent/file.txt", "/path/to/destination/file.txt")
// 	if !errors.Is(err, os.ErrNotExist) {
// 		t.Errorf("Expected error to be os.ErrNotExist, but got %v", err)
// 	}

// 	// Test copying from or to the virtual root directory
// 	err = Copy(fs, "/", "/path/to/destination/file.txt")
// 	if !errors.Is(err, os.ErrInvalid) {
// 		t.Errorf("Expected error to be os.ErrInvalid, but got %v", err)
// 	}
// 	err = Copy(fs, "/path/to/source/file.txt", "/")
// 	if !errors.Is(err, os.ErrInvalid) {
// 		t.Errorf("Expected error to be os.ErrInvalid, but got %v", err)
// 	}

// 	// Test copying to the same file
// 	err = Copy(fs, srcDir, srcDir)
// 	if !errors.Is(err, os.ErrInvalid) {
// 		t.Errorf("Expected error to be os.ErrInvalid, but got %v", err)
// 	}
// }

// func TestCopySlashes(t *testing.T) {
// 	fs := afero.NewMemMapFs()
// 	err := Copy(fs, "", "/")
// 	if !errors.Is(err, os.ErrInvalid) {
// 		t.Errorf("Expected error to be os.ErrInvalid, but got %v", err)
// 	}
// 	err = Copy(fs, "/", "")
// 	if !errors.Is(err, os.ErrInvalid) {
// 		t.Errorf("Expected error to be os.ErrInvalid, but got %v", err)
// 	}
// }
