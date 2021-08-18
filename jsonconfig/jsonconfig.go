package jsonconfig

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/yangtizi/go/ioutils"
	"github.com/yangtizi/log/zaplog"
)

//JSONParsing 解析json文件, 文件名, 解析的JSON结构体
func JSONParsing(strPath string, v interface{}) error {
	if !ioutils.File.Exists(strPath) {
		zaplog.Errorf("yaml需要加载配置，但是找不到文件 %s", strPath)
		return errors.New("找不到文件 " + strPath)
	}

	//创建一个新的buff
	buf := new(bytes.Buffer)
	//打开文件
	f, err := os.Open(strPath)

	if err != nil {
		zaplog.Errorf("打开文件 [%s] 错误 [%v]", strPath, err)
		return err
	}
	defer f.Close()

	//处理注释
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadSlice('\n')
		if err != nil {
			if len(line) > 0 {
				buf.Write(line)
			}
			break
		}

		strLine := strings.TrimLeft(string(line), "\t")  // 去掉TAB
		strLine = strings.TrimLeft(string(strLine), " ") // 去掉空格

		if !strings.HasPrefix(strLine, "//") {
			buf.Write(line)
		}
	}
	//解析json
	err = json.Unmarshal([]byte(buf.Bytes()), v)
	if err != nil {
		zaplog.Errorf("解包文件 [%s] 错误 [%v]", strPath, err)
	}

	return err
}
