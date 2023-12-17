package repository

import (
	"mtg-collector/pkg/model"
)

type InMemoryRepository struct {
	Cards []model.Card
}

func NewInMemoryCardRepository(cards []model.Card) GetCardRepository {
	return &InMemoryRepository{
		Cards: cards,
	}
}

func (m *InMemoryRepository) GetCard(uniqueNumber int) (model.Card, error) {
	return m.Cards[uniqueNumber], nil
}
