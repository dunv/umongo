package umongo

import "github.com/dunv/ulog"

func init() {
	ulog.AddReplaceFunction("github.com/dunv/umongo.ModelService.EnsureIndexes", "umongo.EnsureIndexes")
}
