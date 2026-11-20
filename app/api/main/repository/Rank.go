package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IRankRepository interface {
	FindByIds([]interface{}) ([]model.BangumiListModel, error)
	IncrBangumi(string) error
	IncrEpisode(string) error
}

type RankRepository struct {
	Db *gorm.DB
}

func NewRankRepository(db *gorm.DB) IRankRepository {
	return &RankRepository{Db: db}
}

func (repository *RankRepository) FindByIds(ids []interface{}) ([]model.BangumiListModel, error) {
	//var list []model.BangumiListModel
	list := []model.BangumiListModel{}
	if ids != nil {
		repository.Db.Table("media").
			Clauses(clause.OrderBy{Expression: clause.Expr{SQL: "FIELD(media_id,?)", Vars: ids, WithoutParentheses: true}}).Find(&list)
	}

	//order := fmt.Sprintf("FIELD(media_id,%s)", strings.Join(ids, ","))
	//repository.Db.Table("media").
	//	Where("media_id in (?)", ids).
	//	Order(order).
	//	Find(&list)
	return list, nil
}

func (repository *RankRepository) IncrBangumi(id string) error {
	query := "UPDATE media SET views = views + 1 WHERE media_id = ?"
	repository.Db.Exec(query, id)
	return nil
}

func (repository *RankRepository) IncrEpisode(id string) error {
	query := "UPDATE episode SET views = views + 1 WHERE ep_id = ?"
	repository.Db.Exec(query, id)
	return nil
}
