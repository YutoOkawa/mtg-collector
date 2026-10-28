package repository

import "mtg-collector/pkg/model"

type GetCardRepository interface {
	GetCard(uniqueNumber int) (model.Card, error)
}
