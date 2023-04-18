package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"sync"
)

const fileBufferSize = 64 * 1024 // 微软建议使用 64KB，实测比 32KB 稍快，再增大差不多： https://learn.microsoft.com/en-us/previous-versions/windows/it-pro/windows-2000-server/cc938632(v=technet.10)

var fileBufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, fileBufferSize)
	},
}

func sumReader(h hash.Hash, f *os.File) (string, error) {
	buf := fileBufferPool.Get().([]byte)
	defer fileBufferPool.Put(buf)

	for { // 这里没有用 io.Copy()，因为它的缓存大小是硬编码的 32KB
		switch n, err := f.Read(buf); err { // 避免一次性加载全部文件到内存中，导致内存溢出
		case nil:
			h.Write(buf[:n])
			if n == fileBufferSize {
				continue
			}
			// 读文件不像 socket，不会读到一半没数据了，如果长度不足 fileBufferSize，说明已经读完了
			fallthrough
		case io.EOF:
			return hex.EncodeToString(h.Sum(nil)), nil
		default:
			return "", err
		}
	}
}

func Sum(h hash.Hash, filePath string) (string, error) {
	if info, err := os.Stat(filePath); err != nil || info.IsDir() {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return sumReader(h, file)
}

func SHA256(filePath string) (string, error) {
	return Sum(sha256.New(), filePath)
}

func MD5(filePath string) (string, error) {
	return Sum(md5.New(), filePath)
}
