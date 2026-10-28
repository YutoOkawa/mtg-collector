package repository

import (
	"errors"
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
	if len(m.Cards) <= uniqueNumber {
		return model.Card{}, errors.New("inputNumber exceeds the length of cards")
	}
	card := m.Cards[uniqueNumber]
	return card, nil
}
