package umongo

import (
	"fmt"

	"github.com/dunv/ulog"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	ulog.AddReplaceFunction("github.com/dunv/umongo.ModelService.EnsureIndexes", "umongo.EnsureIndexes")
}

func DebugQuery(query interface{}) {
	rendered, err := bson.MarshalExtJSON(query, false, false)
	if err != nil {
		fmt.Println("DEBUG: err marshaling", err)
	} else {
		fmt.Println("DEBUG: rendered", string(rendered))
	}
}
