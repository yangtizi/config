package yamlconfig

import (
	"errors"
	"io/ioutil"

	"github.com/yangtizi/go/ioutils"
	"github.com/yangtizi/log/zaplog"
	"gopkg.in/yaml.v3"
)

// YMALParsing 解析YAML文件
func YMALParsing(strPath string, v interface{}) error {
	if !ioutils.File.Exists(strPath) {
		zaplog.Errorf("yaml需要加载配置，但是找不到文件 %s", strPath)
		return errors.New("找不到文件 " + strPath)
	}

	data, err := ioutil.ReadFile(strPath)

	if err != nil {
		zaplog.Errorf("读取文件 [%s] 错误 [%v]", strPath, err)
		return err
	}

	// 解包
	err = yaml.Unmarshal(data, v)
	if err != nil {
		zaplog.Errorf("解包文件 [%s] 错误 [%v]", strPath, err)
	}

	return err
}
