package gotypst

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

//go:embed assets
var static embed.FS

var bin_path string

func init() {
	dir, err := os.UserCacheDir()
	if err != nil {
		dir = os.TempDir()
		err = nil
	}
	dir = path.Join(dir, "gotypst")
	os.Mkdir(dir, 0755)
	if err != nil {
		dir = os.TempDir()
		err = nil
	}
	name := runtime.GOARCH + "-" + runtime.GOOS
	bin_path = path.Join(dir, name)
	if _, err := os.Stat(bin_path); err != nil {
		bts, err := static.ReadFile("assets/" + name)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(bin_path, bts, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func RawExec(arg ...string) (string, error) {
	outmesg := ""
	cmd := exec.Command(bin_path, arg...)
	b := new(strings.Builder)
	cmd.Stdout = b
	cmd.Stderr = b
	err := cmd.Start()
	outmesg += b.String()
	b.Reset()
	if err != nil {
		return outmesg, err
	}
	err = cmd.Wait()
	outmesg += b.String()
	b.Reset()
	if err != nil {
		return outmesg, err
	}
	outmesg += b.String()
	b.Reset()
	return outmesg, nil
}

func PDF(bytes []byte, options ...string) ([]byte, error) {
	temp_typ, err := os.CreateTemp(os.TempDir(), "*.typ")
	if err != nil {
		return []byte{}, nil
	}
	_, err = temp_typ.Write(bytes)
	if err != nil {
		return []byte{}, nil
	}
	temp_pdf_name := strings.TrimSuffix(temp_typ.Name(), ".typ") + ".pdf"
	cmd := make([]string, 0)
	cmd = append(cmd, "compile")
	cmd = append(cmd, temp_typ.Name())
	cmd = append(cmd, options...)
	cmd = append(cmd, temp_pdf_name)
	out, err := RawExec(cmd...)
	_ = os.Remove(temp_typ.Name())
	if err != nil {
		return []byte{}, fmt.Errorf("%v %v", out, err)
	}
	out_bytes, err := os.ReadFile(temp_pdf_name)
	_ = os.Remove(temp_pdf_name)
	return out_bytes, err
}
