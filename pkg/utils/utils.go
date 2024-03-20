package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

func EncodeMD5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}

	return false, err
}

func PathContains(path string) (bool, error, string, string) {
	//fmt.Println("path ", path)
	index := strings.LastIndex(path, "/")
	roots := path[:index]
	dir := path[index+1:]
	//fmt.Println("roots: ", roots)
	//fmt.Println("dir: ", dir)
	files, err := os.ReadDir(roots)
	if err != nil {
		return false, err, "", ""
	}
	for _, file := range files {
		if strings.Contains(file.Name(), dir) {
			return true, nil, fmt.Sprintf("%s/%s", roots, file.Name()), file.Name()
		}
	}
	return false, err, "", ""
}

func Copy(src, dst string) error {
	f, e := os.Stat(src)
	if e != nil {
		return e
	}

	if f.IsDir() {
		//复制文件夹
		if list, e := os.ReadDir(src); e == nil {
			for _, item := range list {
				if e = Copy(filepath.Join(src, item.Name()), filepath.Join(dst, item.Name())); e != nil {
					return e
				}
			}
		}
	} else {
		//复制文件
		p := filepath.Dir(dst)
		if _, e := os.Stat(p); e != nil {
			if e = os.MkdirAll(p, 0777); e != nil {
				return e
			}
		}
		//读取源文件
		file, e := os.Open(src)
		if e != nil {
			return e
		}
		defer file.Close()
		bufReader := bufio.NewReader(file)
		out, e := os.Create(dst)
		if e != nil {
			return e
		}
		defer out.Close()

		//将文件流和文件流对接
		_, e = io.Copy(out, bufReader)
	}

	return e
}

func NowTimeString() string {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	ts := fmt.Sprintf("%d%02d%02d%02d%02d%02d", year, month, day, hour, minute, second)
	return ts
}

// 创建验证码
func CreateCode(width int) string  {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}