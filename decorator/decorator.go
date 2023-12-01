package decorator

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Decorate(config *viper.Viper) {
	var (
		src  = config.GetString("source-dir")
		dest = config.GetString("dest-dir")
	)

	if err := decorate(src, dest); err != nil {
		panic(err)
	}
}

func decorate(src, dest string) error {
	if err := ensureFolder(src); err != nil {
		return err
	}

	if err := ensureFolder(dest); err != nil {
		return err
	}

	images, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, image := range images {
		fmt.Println(image.Name())
	}

	return nil
}
