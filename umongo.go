package umongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func DebugQuery(query interface{}) {
	rendered, err := bson.MarshalExtJSON(query, false, false)
	if err != nil {
		fmt.Println("DEBUG: err marshaling", err)
	} else {
		fmt.Println("DEBUG: rendered", string(rendered))
	}
}
