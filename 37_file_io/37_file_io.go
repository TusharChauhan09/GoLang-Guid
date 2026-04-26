// 37_file_io.go
// Topic: File I/O — os, io, bufio packages
//
// READ WHOLE FILE
//   data, err := os.ReadFile("path")    // []byte
//
// WRITE WHOLE FILE
//   err := os.WriteFile("path", data, 0644)
//
// OPEN FILE
//   f, err := os.Open("path")             // read-only
//   f, err := os.Create("path")           // truncate or create, write-only
//   f, err := os.OpenFile("p", flags, mode)
//     flags: O_RDONLY, O_WRONLY, O_RDWR, O_APPEND, O_CREATE, O_TRUNC, O_EXCL
//   defer f.Close()
//
// READ FROM FILE
//   buf := make([]byte, 1024)
//   n, err := f.Read(buf)
//
// LINE-BY-LINE: bufio.Scanner
//   sc := bufio.NewScanner(f)
//   for sc.Scan() { line := sc.Text() }
//
// BUFFERED WRITE: bufio.Writer
//   w := bufio.NewWriter(f)
//   w.WriteString("...")
//   w.Flush()                              // important!
//
// FILE INFO
//   info, _ := os.Stat("path")
//   info.Size(), info.IsDir(), info.Mode(), info.ModTime()
//
// DIRECTORIES
//   os.Mkdir("d", 0755)
//   os.MkdirAll("a/b/c", 0755)
//   os.Remove("file")
//   os.RemoveAll("dir")
//   os.Rename(old, new)
//   entries, _ := os.ReadDir("dir")        // returns []DirEntry
//
// PATHS
//   filepath.Join("a", "b", "c.txt")       // OS-correct separator
//   filepath.Ext, Base, Dir, Abs, Walk
//
// io INTERFACES
//   io.Reader, io.Writer, io.Closer, io.ReaderAt, io.Seeker
//   io.Copy(dst, src)                      — universal copy
//
// Run: go run 37_file_io.go     (creates /tmp/go_guide_demo.txt)

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := os.TempDir()
	path := filepath.Join(dir, "go_guide_demo.txt")
	defer os.Remove(path)

	// Write whole file
	if err := os.WriteFile(path, []byte("line1\nline2\nline3\n"), 0644); err != nil {
		panic(err)
	}

	// Read whole file
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes:\n%s", len(data), data)

	// Append using OpenFile
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	f.WriteString("appended\n")
	f.Close()

	// Line-by-line scanning
	f, _ = os.Open(path)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fmt.Println("line:", sc.Text())
	}
	f.Close()

	// Buffered write
	f, _ = os.Create(path)
	w := bufio.NewWriter(f)
	for i := 0; i < 3; i++ {
		fmt.Fprintf(w, "n=%d\n", i)
	}
	w.Flush()
	f.Close()

	// io.Copy from string Reader to file
	f, _ = os.Create(path)
	io.Copy(f, strings.NewReader("copied via io.Copy"))
	f.Close()

	// File info
	info, _ := os.Stat(path)
	fmt.Println("size:", info.Size(), "mode:", info.Mode(), "isDir:", info.IsDir())

	// Read directory entries
	entries, _ := os.ReadDir(dir)
	for _, e := range entries[:min(3, len(entries))] {
		fmt.Println("entry:", e.Name(), "isDir:", e.IsDir())
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
