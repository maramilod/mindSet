package command

import (
	"fmt"
	"mind-set/internal/utils"
	"mind-set/internal/utils/logger"
	"path/filepath"

	c "mind-set/config"

	"github.com/gohouse/converter"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "command",
		Short:   "Start command",
		Example: "go run main.go command -t='table name'",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {

			logger.Logger.Info(table)

		},
	}
	table string
)

func init() {
	Cmd.Flags().StringVarP(&table, "table", "t", "", "The model file table name")
}

func CreateModelStruct() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Config.Mysql.Username,
		c.Config.Mysql.Password,
		c.Config.Mysql.Host,
		c.Config.Mysql.Port,
		c.Config.Mysql.Database,
		c.Config.Mysql.Charset,
	)
	modelPath := filepath.Join(utils.GetRunPath(), "/internal/model/", table+".go")
	t2t := converter.NewTable2Struct()
	t2t.Config(&converter.T2tConfig{
		RmTagIfUcFirsted: false,
		TagToLower:       false,
		UcFirstOnly:      false,
		SeperatFile:      true,
	})
	err := t2t.
		Table(table).
		Prefix("").
		EnableJsonTag(true).
		PackageName("model").
		TagKey("gorm").
		RealNameMethod("TableName").
		SavePath(modelPath).
		Dsn(dsn).
		Run()

	fmt.Println(err)
}
