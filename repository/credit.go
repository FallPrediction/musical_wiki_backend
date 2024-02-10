package repository

import (
	"musical_wiki/global"
	"musical_wiki/models"
)

type CreditRepository struct{}

func (repository *CreditRepository) IndexByActorId(actorId string) ([]models.Credit, error) {
	var credits []models.Credit
	err := global.Db.Where("actor_id = ?", actorId).Find(&credits).Error
	return credits, err
}

func (repository *CreditRepository) Store(credit *models.Credit) error {
	// Check if actor exists
	var actor models.Actor
	err := global.Db.Where("id = ?", credit.ActorId).First(&actor).Error
	if err != nil {
		return err
	}
	return global.Db.Create(&credit).Error
}

func (repository *CreditRepository) Update(id string, credit *models.Credit) error {
	var model models.Credit
	err := global.Db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return err
	}
	err = global.Db.Model(&model).Updates(credit).Error
	return err
}

func (repository *CreditRepository) Destroy(id string) error {
	var model models.Credit
	err := global.Db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return err
	}
	err = global.Db.Delete(&model).Error
	return err
}
