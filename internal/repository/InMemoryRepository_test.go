package repository

import (
	"mtg-collector/pkg/model"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetCard(t *testing.T) {
	testCards := []model.Card{
		{
			Name:  "test Goblin",
			Color: model.RED,
			ManaCost: model.ManaCost{
				model.RED: "1",
			},
		},
	}

	inMemoryrepository := NewInMemoryCardRepository(testCards)

	cases := []struct {
		name string

		uniqueNumber int

		wantErr  bool
		wantCard model.Card
	}{
		{
			name: "ShouldGetCardSuccessfully",

			uniqueNumber: 0,

			wantErr: false,
			wantCard: model.Card{
				Name:  "test Goblin",
				Color: model.RED,
				ManaCost: model.ManaCost{
					model.RED: "1",
				},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got, err := inMemoryrepository.GetCard(c.uniqueNumber)

			if c.wantErr {
				if err == nil {
					t.Error("err is expected, but nil")
				}
			} else {
				if err != nil {
					t.Errorf("err is expected to be nil, but %#v", err)
				}

				if diff := cmp.Diff(c.wantCard, got); diff != "" {
					t.Errorf("GetCard is mismatch (-want +got): %#v", diff)
				}
			}
		})
	}
}
