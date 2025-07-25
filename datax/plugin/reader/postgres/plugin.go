package postgres

import (
	"github.com/Breeze0806/go-etl/config"
	"github.com/Breeze0806/go-etl/datax/plugin/reader"
)

func init() {
	var err error
	maker := &maker{}
	if err = reader.RegisterReader(maker); err != nil {
		panic(err)
	}
}

var pluginConfig = `{
    "name" : "postgresreader",
    "developer":"Breeze0806",
    "dialect":"postgres",
    "description":"use github.com/Breeze0806/go/database/pqto which base on github.com/lib/pq . database/sql DB execute select sql, retrieve data from the ResultSet. warn: The more you know about the database, the less problems you encounter."
}`

//NewReaderFromString create reader
func NewReaderFromString(plugin string) (rd reader.Reader, err error) {
	r := &Reader{}
	if r.pluginConf, err = config.NewJSONFromString(plugin); err != nil {
		return nil, err
	}
	rd = r
	return
}

type maker struct{}

func (m *maker) Default() (reader.Reader, error) {
	return NewReaderFromString(pluginConfig)
}
