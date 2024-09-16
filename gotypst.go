package gotypst

import (
	"archive/zip"
	"bytes"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

//go:embed assets.zip
var assets_zipped []byte

//go:embed fonts/*
var fonts_efs embed.FS

var bin_path string

func init() {
	dir := gotypstDir()

	name := runtime.GOARCH + "-" + runtime.GOOS

	bin_path = path.Join(dir, name)
	if _, err := os.Stat(bin_path); err != nil {
		zr := bytes.NewReader(assets_zipped)
		zip_fs, err := zip.NewReader(zr, int64(len(assets_zipped)))
		if err != nil {
			log.Fatal(err)
		}
		fi, err := zip_fs.Open("assets/" + name)
		if err != nil {
			log.Fatal(err)
		}
		bts, err := io.ReadAll(fi)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(bin_path, bts, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	font_rd, _ := fonts_efs.ReadDir("fonts")
	fonts_dir := fontsDir()
	os.Mkdir(fonts_dir, 0755)
	for _, v := range font_rd {
		if !strings.HasSuffix(v.Name(), ".ttf") {
			continue
		}
		dest_path := path.Join(fonts_dir, v.Name())
		if _, err := os.Stat(dest_path); err != nil {
			fn, _ := fonts_efs.ReadFile(path.Join("fonts", v.Name()))
			err := os.WriteFile(dest_path, fn, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func gotypstDir() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		dir = os.TempDir()
		err = nil
	}
	dir = path.Join(dir, "gotypst")
	os.Mkdir(dir, 0755)
	return dir
}
func fontsDir() string {
	dir := gotypstDir()
	dir = path.Join(dir, "fonts")
	os.Mkdir(dir, 0755)
	return dir
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
		return []byte{}, err
	}
	_, err = temp_typ.Write(bytes)
	if err != nil {
		return []byte{}, err
	}
	temp_pdf_name := strings.TrimSuffix(temp_typ.Name(), ".typ") + ".pdf"
	cmd := make([]string, 0)
	cmd = append(cmd, "compile")
	cmd = append(cmd, temp_typ.Name())
	cmd = append(cmd, options...)
	cmd = append(cmd, "--font-path")
	cmd = append(cmd, fontsDir())
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
