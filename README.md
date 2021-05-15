# SJIS to UTF-8

## 参考

Go言語(golang)でUtf-8のファイルを`Shift_JIS`に変換する
https://dev.classmethod.jp/articles/golang-iconv/

改行コードを CRLF にする

## 準備

```sh
go mod init github.com/ShinNakamura/go-sjis2utf8
go get -u golang.org/x/text
mkdir bin
```


## usage

```sh
$0 /path/to/utf8File /path/to/sjisFile
```

`/path/to/utf8File` は存在しないとエラー。Glob可だが、展開後の最初のファイルしか採用しないことに注意

`/path/to/sjisFile` はファイルは新規作成／上書き。事前にディレクトリを作成しておくこと
