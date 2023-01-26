package external

import (
	"search_engine_task/pkg/dbconn"
	"search_engine_task/pkg/models"
	"testing"

	"github.com/magiconair/properties/assert"
	"golang.org/x/net/context"
)

func TestMongoDB_GetAllCollection(t *testing.T) {
	// INCOMPLETE
}

func TestMongoDB_InsertOnePage(t *testing.T) {	
	db := NewMongoDB(dbconn.Dbconn())

	input := models.Page{
		Title:    "Page 10",
		Keywords: []string{"wrd1", "wrd2"},
	}

	var ctx context.Context

	err := db.InsertOnePage(ctx, input)

	assert.Equal(t, nil, err)

}
