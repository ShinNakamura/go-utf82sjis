// 参考: https://dev.classmethod.jp/articles/golang-iconv/
// 今のところほぼそのまま
package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	infGlob := os.Args[1] // glob が渡されたら展開して最初のファイルだけ採用
	outf := os.Args[2]    // glob 不可

	infs, err := filepath.Glob(infGlob)
	if err != nil {
		log.Fatal(err)
	}
	inBytes, err := ioutil.ReadFile(infs[0])
	if err != nil {
		log.Fatal(err)
	}

	// 元のファイルにCRが含まれていることは考慮しない。
	withCRLF := strings.ReplaceAll(string(inBytes), "\n", "\r\n")

	// 書き込み先ファイルを用意
	sjisFile, err := os.Create(outf)
	if err != nil {
		log.Fatal(err)
	}
	defer sjisFile.Close()

	// ShiftJISのエンコーダーを噛ませたWriterを作成する
	writer := transform.NewWriter(sjisFile, japanese.ShiftJIS.NewEncoder())

	utf8Reader := strings.NewReader(withCRLF)

	// 書き込み
	tee := io.TeeReader(utf8Reader, writer)
	s := bufio.NewScanner(tee)
	for s.Scan() {
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
