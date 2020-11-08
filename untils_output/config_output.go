package untils_output

import (
	"fmt"
	u "zx/untils"
	ms "zx/untils_model"
)

//读取配置文件

func Function(input ms.Input) (output ms.Output, err error) {
	fmt.Println("second .........,input:", input.Id)

	output.Database.Types, output.Database.Local, output.Database.Online, err = u.GetDatabase()
	if err != nil {
		return output, err
	}
	output.Self.Port, output.Self.Flag, output.Self.Tag, err = u.GetSelf()
	if err != nil {
		return output, err
	}

	return output, nil
}
