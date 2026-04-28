// 37_file_io.go
// Topic: File I/O — os, io, bufio packages
//
// ─────────────────────────────────────────────────────────────
// 1. io INTERFACES (the types everything implements)
// ─────────────────────────────────────────────────────────────
//   io.Reader   — Read(p []byte) (n int, err error)
//   io.Writer   — Write(p []byte) (n int, err error)
//   io.Closer   — Close() error
//   io.Seeker   — Seek(offset, whence)         (random access)
//   io.ReaderAt — ReadAt(p, off)               (concurrent reads)
//   *os.File satisfies all of the above.
//   Anything Reader→Writer can be piped via io.Copy.
//
// ─────────────────────────────────────────────────────────────
// 2. READ WHOLE FILE (one-shot, small files)
// ─────────────────────────────────────────────────────────────
//   data, err := os.ReadFile("path")    // []byte, file auto-closed
//
// ─────────────────────────────────────────────────────────────
// 3. WRITE WHOLE FILE (one-shot, overwrites)
// ─────────────────────────────────────────────────────────────
//   err := os.WriteFile("path", data, 0644)
//
// ─────────────────────────────────────────────────────────────
// 4. OPEN FILE — why & how
// ─────────────────────────────────────────────────────────────
//   Why open? ReadFile/WriteFile load whole file into memory.
//   For big files, streaming, appending, or random access you
//   need an *os.File handle.
//
//   f, err := os.Open("path")            // read-only
//   f, err := os.Create("path")          // create/truncate, write-only
//   f, err := os.OpenFile(p, flags, mode)
//     flags: O_RDONLY, O_WRONLY, O_RDWR, O_APPEND, O_CREATE, O_TRUNC, O_EXCL
//     mode : 0644 (rw-r--r--), 0755 (rwxr-xr-x)
//   defer f.Close()                       // ALWAYS close
//
// ─────────────────────────────────────────────────────────────
// 5. READ FROM OPEN FILE
// ─────────────────────────────────────────────────────────────
//   Raw bytes:
//     buf := make([]byte, 1024)
//     n, err := f.Read(buf)               // err == io.EOF at end
//
//   Line-by-line (bufio.Scanner):
//     sc := bufio.NewScanner(f)
//     for sc.Scan() { line := sc.Text() }
//
// ─────────────────────────────────────────────────────────────
// 6. WRITE TO OPEN FILE
// ─────────────────────────────────────────────────────────────
//   Direct:
//     f.WriteString("hi")
//     f.Write([]byte{...})
//
//   Buffered (bufio.Writer — fewer syscalls):
//     w := bufio.NewWriter(f)
//     w.WriteString("...")
//     w.Flush()                            // MUST flush before close
//
// ─────────────────────────────────────────────────────────────
// 7. STREAM: read one file → write another
// ─────────────────────────────────────────────────────────────
//   src, _ := os.Open("in.txt")
//   dst, _ := os.Create("out.txt")
//   io.Copy(dst, src)                     // universal Reader→Writer pipe
//   src.Close(); dst.Close()
//
// ─────────────────────────────────────────────────────────────
// 8. FILE INFO
// ─────────────────────────────────────────────────────────────
//   info, _ := os.Stat("path")
//   info.Size(), info.IsDir(), info.Mode(), info.ModTime()
//
// ─────────────────────────────────────────────────────────────
// 9. DIRECTORIES
// ─────────────────────────────────────────────────────────────
//   os.Mkdir("d", 0755)
//   os.MkdirAll("a/b/c", 0755)
//   os.Remove("file")
//   os.RemoveAll("dir")
//   os.Rename(old, new)
//   entries, _ := os.ReadDir("dir")        // []DirEntry
//
// ─────────────────────────────────────────────────────────────
// 10. PATHS
// ─────────────────────────────────────────────────────────────
//   filepath.Join("a", "b", "c.txt")       // OS-correct separator
//   filepath.Ext, Base, Dir, Abs, Walk
//
// Run: go run 37_file_io.go     (creates file in os.TempDir())

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

	// ── 2. Write whole file (seed it first so we can read it) ──
	if err := os.WriteFile(path, []byte("line1\nline2\nline3\n"), 0644); err != nil {
		panic(err)
	}

	// ── 3. Read whole file ──
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes:\n%s", len(data), data)

	// ── 4. Open file with flags (append mode) ──
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	f.WriteString("appended\n")
	f.Close()

	// ── 5. Read from open file: line-by-line ──
	f, _ = os.Open(path)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fmt.Println("line:", sc.Text())
	}
	f.Close()

	// ── 6. Write to open file: buffered ──
	f, _ = os.Create(path)
	w := bufio.NewWriter(f)
	for i := 0; i < 3; i++ {
		fmt.Fprintf(w, "n=%d\n", i)
	}
	w.Flush()
	f.Close()

	// stream
	
	srcFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if err := writer.WriteByte(b); err != nil {
			panic(err)
		}
	}

	// short streaming
	srcFile1, _ := os.Open("input.txt")
	distFile1, _ := os.Create("output.txt")
	io.Copy(distFile1, srcFile1)
	srcFile1.Close()
	distFile1.Close()

	// ── 7. Stream: copy from one Reader into a file Writer ──
	f, _ = os.Create(path)
	io.Copy(f, strings.NewReader("copied via io.Copy"))
	f.Close()

	// ── 7b. Stream file → file ──
	src, _ := os.Open(path)
	dstPath := filepath.Join(dir, "go_guide_demo_copy.txt")
	defer os.Remove(dstPath)
	dst, _ := os.Create(dstPath)
	io.Copy(dst, src)
	src.Close()
	dst.Close()

	// ── 8. File info ──
	info, _ := os.Stat(path)
	fmt.Println("size:", info.Size(), "mode:", info.Mode(), "isDir:", info.IsDir())

	// ── 9. Read directory entries ──
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
