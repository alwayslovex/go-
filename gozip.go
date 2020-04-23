package main

import (
	"fmt"
	encryzip "github.com/yeka/zip"
	"io/ioutil"
	"os"
	"time"
)

//filenames原始文件列表，zipfilename 压缩包名称，passwd 解压密码
func EncryFile(filenames []string, zipfilename, passwd string) error {
	zipfile, err := os.Create(zipfilename)
	if err != nil {
		return err
	}
	fzip := encryzip.NewWriter(zipfile)
	defer fzip.Close()
	for _, v := range filenames {
		t1 := time.Now().Unix()
		w, err := fzip.Encrypt(v, passwd, encryzip.AES128Encryption)
		if err != nil {
			return err
		}
		bd, err := ioutil.ReadFile(v)
		if err != nil {
			return err
		}
		w.Write(bd)
		t2 := time.Now().Unix()
		fmt.Println(t2 - t1)
	}
	return fzip.Flush()
}

func main() {

}
