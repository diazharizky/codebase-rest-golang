package mongo

import (
	"context"
	"time"

	"github.com/diazharizky/codebase-rest-golang/pkg/types"
)

func (mh *MongoHandler) GetOne(c *types.Foo, filter interface{}) error {
	// Will automatically create a collection if not available
	collection := mh.client.Database(mh.db).Collection("foo")
	ctx, cn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cn()
	err := collection.FindOne(ctx, filter).Decode(c)
	return err
}
