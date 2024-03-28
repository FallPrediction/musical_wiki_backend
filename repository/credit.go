package repository

import (
	"musical_wiki/models"

	"gorm.io/gorm"
)

type CreditRepository struct {
	db *gorm.DB
}

func (repository *CreditRepository) IndexByActorId(actorId string) ([]models.Credit, error) {
	var credits []models.Credit
	err := repository.db.Where("actor_id = ?", actorId).Find(&credits).Error
	return credits, err
}

func (repository *CreditRepository) Store(credit *models.Credit) error {
	// Check if actor exists
	var actor models.Actor
	err := repository.db.Where("id = ?", credit.ActorId).First(&actor).Error
	if err != nil {
		return err
	}
	return repository.db.Create(&credit).Error
}

func (repository *CreditRepository) Update(id string, credit *models.Credit) error {
	var model models.Credit
	err := repository.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return err
	}
	err = repository.db.Model(&model).Updates(credit).Error
	return err
}

func (repository *CreditRepository) Destroy(id string) error {
	var model models.Credit
	err := repository.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return err
	}
	err = repository.db.Delete(&model).Error
	return err
}

func NewCreditRepository(db *gorm.DB) CreditRepository {
	return CreditRepository{db: db}
}
