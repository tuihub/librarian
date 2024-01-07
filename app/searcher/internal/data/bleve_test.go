package data_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/app/searcher/internal/data"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/model"

	"github.com/blevesearch/bleve/v2"
)

func Test_bleveSearcherRepo_SearchID(t *testing.T) {
	type Document struct {
		Name        string
		Description string
	}
	docs := []*Document{
		{
			Name: "Celeste",
			Description: "Help Madeline survive her inner demons on her journey to the top of " +
				"Celeste Mountain, in this super-tight platformer from the creators of TowerFall. " +
				"Brave hundreds of hand-crafted challenges, uncover devious secrets, and piece " +
				"together the mystery of the mountain.",
		},
		{
			Name: "Hollow Knight",
			Description: "Forge your own path in Hollow Knight! An epic action adventure through " +
				"a vast ruined kingdom of insects and heroes. Explore twisting caverns, battle " +
				"tainted creatures and befriend bizarre bugs, all in a classic, hand-drawn 2D style.",
		},
		{
			Name: "Disco Elysium",
			Description: "Disco Elysium - The Final Cut is a groundbreaking role playing game. " +
				"You’re a detective with a unique skill system at your disposal and a whole city " +
				"to carve your path across. Interrogate unforgettable characters, crack murders or" +
				" take bribes. Become a hero or an absolute disaster of a human being.",
		},
	}
	mapping := bleve.NewIndexMapping()
	dbPath := "bleve.db"
	index, err := bleve.Open(dbPath)
	if err != nil {
		if !errors.Is(err, bleve.ErrorIndexPathDoesNotExist) {
			t.Error(err)
			return
		} else {
			index, err = bleve.New(dbPath, mapping)
			if err != nil {
				t.Error(err)
				return
			}
		}
	}
	if err != nil {
		t.Error(err)
	}
	indexMap := make(map[biz.Index]bleve.Index)
	indexMap[biz.IndexGeneral] = index
	r, err := data.NewSearcherRepo(
		indexMap,
		nil, nil,
	)
	if err != nil {
		t.Error(err)
	}
	for i := range docs {
		var str []byte
		str, err = libcodec.Marshal(libcodec.JSON, &docs[i])
		if err != nil {
			return
		}
		if err = r.DescribeID(context.Background(), model.InternalID(i), biz.IndexGeneral, false, string(str)); err != nil {
			t.Errorf("DescribeID() error = %v", err)
		}
	}
	ids, err := r.SearchID(context.Background(),
		biz.IndexGeneral,
		model.Paging{
			PageSize: 10,
			PageNum:  1,
		}, "your")
	if err != nil {
		t.Errorf("SearchID() error = %v", err)
	}
	t.Log(ids)
}
