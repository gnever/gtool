package gfile_test

import (
	"reflect"
	"testing"

	"github.com/gnever/gtool/gfile"
)

func TestMkdir(t *testing.T) {
	testDir := "/tmp/testdir"

	if gfile.Exists(testDir) {
		t.Fatalf("%s exists", testDir)
	}

	if err := gfile.Mkdir(testDir); err != nil {
		t.Fatal(err)
	}

	if !gfile.IsDir(testDir) {
		t.Errorf("%s not dir", testDir)
	}

	if err := gfile.Remove(testDir); err != nil {
		t.Errorf("remove %s failed", testDir)
	}

	if gfile.Exists(testDir) {
		t.Errorf("%s exists after remove", testDir)
	}
}

func TestCreate(t *testing.T) {
	f := "/tmp/testfile/name.ext"

	if gfile.Exists(f) {
		t.Fatalf("%s exists", f)
	}

	if err := gfile.Create(f); err != nil {
		t.Fatal(err)
	}

	if !gfile.IsFile(f) {
		t.Errorf("%s not a file", f)
	}

	if bn := gfile.Basename(f); bn != "name.ext" {
		t.Errorf("%s get basename is %s, Should be name.ext", f, bn)
	}

	if bn := gfile.Name(f); bn != "name" {
		t.Errorf("%s get basename is %s, Should be name", f, bn)
	}

	if bn := gfile.Ext(f); bn != ".ext" {
		t.Errorf("%s get ext is %s, Should be .ext", f, bn)
	}

	if bn := gfile.ExtName(f); bn != "ext" {
		t.Errorf("%s get ext is %s, Should be ext", f, bn)
	}

	if err := gfile.Remove(f); err != nil {
		t.Errorf("remove %s failed", f)
	}

	if gfile.Exists(f) {
		t.Errorf("%s exists after remove", f)
	}

}
func TestCopyfile(t *testing.T) {
	f := "/tmp/testfile/name.ext"
	dst := "/tmp/testfile-dst/dst.ext"

	if gfile.Exists(f) {
		t.Fatalf("%s exists", f)
	}

	if gfile.Exists(dst) {
		t.Fatalf("%s exists", f)
	}

	if err := gfile.Create(f); err != nil {
		t.Fatal(err)
	}

	if !gfile.IsFile(f) {
		t.Errorf("%s not a file", f)
	}

	gfile.Copy(f, dst)

	if !gfile.IsFile(dst) {
		t.Errorf("%s not a file after copy", dst)
	}

	if err := gfile.Remove(f); err != nil {
		t.Errorf("remove %s failed", f)
	}

	if err := gfile.Remove(dst); err != nil {
		t.Errorf("remove %s failed", f)
	}

	if gfile.Exists(f) {
		t.Errorf("%s exists after remove", f)
	}

	if gfile.Exists(dst) {
		t.Errorf("%s exists after remove", f)
	}

}

func TestCopyDir(t *testing.T) {
	f := "/tmp/testfile/dir"
	dst := "/tmp/testfile-dst/dir-dst"

	if gfile.Exists(f) {
		t.Fatalf("%s exists", f)
	}

	if gfile.Exists(dst) {
		t.Fatalf("%s exists", f)
	}

	if err := gfile.Mkdir(f); err != nil {
		t.Fatal(err)
	}

	if !gfile.IsDir(f) {
		t.Errorf("%s not a dir", f)
	}

	gfile.Copy(f, dst)

	if !gfile.IsDir(dst) {
		t.Errorf("%s not a file after copy", dst)
	}

	if err := gfile.Remove(f); err != nil {
		t.Errorf("remove %s failed", f)
	}

	if err := gfile.Remove(dst); err != nil {
		t.Errorf("remove %s failed", f)
	}

	if gfile.Exists(f) {
		t.Errorf("%s exists after remove", f)
	}

	if gfile.Exists(dst) {
		t.Errorf("%s exists after remove", f)
	}

}

func TestWriteLine(t *testing.T) {
	file := "/tmp/test-write.log"
	contents := `testas`

	if err := gfile.PutContents(file, contents); err != nil {
		t.Fatalf("put contents fail %s", err)
	}

	if ctx, err := gfile.GetContents(file); err != nil {
		t.Fatalf("get contents fail %s", err)
	} else if ctx != contents {
		t.Fatalf("get content is %s,  Should be %s", ctx, contents)
	}
}

func TestWriteBytes(t *testing.T) {
	file := "/tmp/test-write.log"
	contents := []byte(`testas`)

	if err := gfile.PutBytes(file, contents); err != nil {
		t.Fatalf("put contents fail %s", err)
	}

	if ctx, err := gfile.GetContents(file); err != nil {
		t.Fatalf("get contents fail %s", err)
	} else if ctx != string(contents) {
		t.Fatalf("get content is %s,  Should be %s", ctx, contents)
	}
}

func TestAppendStrings(t *testing.T) {
	file := "/tmp/test-write.log"
	list := []string{"a", "b", "c"}

	if gfile.IsFile(file) {
		if err := gfile.Remove(file); err != nil {
			t.Errorf("remove %s failed", file)
		}
	}

	for _, v := range list {
		err := gfile.PutContentsAppend(file, v+"\n")
		if err != nil {
			t.Fatalf("write %s to %s fail", v, file)
		}
	}

	getList := make([]string, 0)
	cf := func(line string) {
		getList = append(getList, line)
	}

	gfile.GetLinesByScan(file, cf)

	if reflect.DeepEqual(getList, list) == false {
		t.Errorf("get contents is %s,  Should be %s", getList, list)
	}
}

func TestAppendBytes(t *testing.T) {
	file := "/tmp/test-write.log"
	list := []string{"a", "b", "c"}

	if gfile.IsFile(file) {
		if err := gfile.Remove(file); err != nil {
			t.Errorf("remove %s failed", file)
		}
	}

	for _, v := range list {
		err := gfile.PutBytesAppend(file, []byte(v+"\n"))
		if err != nil {
			t.Fatalf("write %s to %s fail", v, file)
		}
	}

	getList := make([]string, 0)
	cf := func(line []byte) {
		getList = append(getList, string(line))
	}

	gfile.GetBytesByScan(file, cf)

	if reflect.DeepEqual(getList, list) == false {
		t.Errorf("get contents is %s,  Should be %s", getList, list)
	}
}
