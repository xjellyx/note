package main

import (
	"bytes"
	"os"
)

func tail(filename string, n int) (lines []string, err error) {
	var (
		f    os.FileInfo
		size int64
		fi   *os.File
	)

	if f, err = os.Stat(filename); err != nil {
		log.Errorln(err)
		return
	}
	size = f.Size()
	if fi, err = os.Open(filename); err != nil {
		log.Errorln(err)
		return
	}
	defer fi.Close()
	var (
		b         = make([]byte, defaultBufSize)
		readSize  = int64(defaultBufSize)
		lineNum   = n
		bTail     = bytes.NewBuffer([]byte{})
		seekStart = size
		flag      = true
	)
	for flag {
		// 直接从文件头部开始
		if seekStart < defaultBufSize {
			readSize = seekStart
			seekStart = 0
		} else { // 每次从开始位置减去读取的字节大小读取数据；从文件尾部开始读取,第一次 size-readSize
			seekStart -= readSize
		}
		if _, err = fi.Seek(seekStart, os.SEEK_SET); err != nil {
			log.Errorln(err)
			return
		}
		mm, _err := fi.Read(b)
		if _err != nil {
			err = _err
			log.Errorln(err)
			return
		}
		if mm > 0 {
			j := mm
			// 读取每个字节，以\n为一行
			for i := mm - 1; i >= 0; i-- {
				if b[i] == '\n' {
					bLine := bytes.NewBuffer([]byte{})
					bLine.Write(b[i+1 : j])
					j = i
					if bTail.Len() > 0 {
						bLine.Write(bTail.Bytes())
						bTail.Reset()
					}

					if (lineNum == n && bLine.Len() > 0) || lineNum < n { //skip last "\n"
						lines = append(lines, bLine.String())
						lineNum--
					}
					if lineNum == 0 {
						flag = false
						break
					}
				}
			}
			if flag && j > 0 {
				if seekStart == 0 {
					bLine := bytes.NewBuffer([]byte{})
					bLine.Write(b[:j])
					if bTail.Len() > 0 {
						bLine.Write(bTail.Bytes())
						bTail.Reset()
					}
					lines = append(lines, bLine.String())
					flag = false
				} else {
					bb := make([]byte, bTail.Len())
					copy(bb, bTail.Bytes())
					bTail.Reset()
					bTail.Write(b[:j])
					bTail.Write(bb)
				}
			}
		}
	}

	return
}
