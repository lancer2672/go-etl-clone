package csv

import (
	"github.com/Breeze0806/go-etl/config"
	"github.com/Breeze0806/go-etl/datax/plugin/writer"
)

func init() {
	var err error
	maker := &maker{}
	if err = writer.RegisterWriter(maker); err != nil {
		panic(err)
	}
}

var pluginConfig = `{
    "name" : "csvwriter",
    "developer":"Breeze0806",
    "creator":"csv",
    "description":""
}`

//NewWriterFromString create writer
func NewWriterFromString(plugin string) (wr writer.Writer, err error) {
	w := &Writer{}
	if w.pluginConf, err = config.NewJSONFromString(plugin); err != nil {
		return nil, err
	}
	wr = w
	return
}

type maker struct{}

func (m *maker) Default() (writer.Writer, error) {
	return NewWriterFromString(pluginConfig)
}