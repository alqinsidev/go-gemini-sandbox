package repository

import (
	"database/sql"
	"fmt"

	"github.com/alqinsidev/go-gemini-sandbox/src/domain"
	_error "github.com/alqinsidev/go-gemini-sandbox/src/error"
	"github.com/gin-gonic/gin"
)

func (r *Repository) InsertInformation(ctx *gin.Context, payload domain.InformationBodyRequest) (id int64, err error) {
	query := `INSERT INTO informations (question,answer) VALUES (?,?)`
	res, err := r.db.ExecContext(ctx, query, payload.Question, payload.Answer)
	if err != nil {
		err = _error.ErrDB
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		err = _error.ErrDB
		return
	}

	return
}
func (r *Repository) EditInformation(ctx *gin.Context, payload domain.InformationBodyRequest, id int64) (isSuccess bool, err error) {
	query := `UPDATE informations SET question=?,answer=? WHERE id=?`
	res, err := r.db.ExecContext(ctx, query, payload.Question, payload.Answer, id)
	if err != nil {
		fmt.Println("err: ", err.Error())
		err = _error.ErrDB
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		err = _error.ErrDB
		return
	}

	return affected > 0, nil
}

func (r *Repository) DeleteInformation(ctx *gin.Context, id int64) (success bool, err error) {
	query := `DELETE FROM informations WHERE id = ?`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		err = _error.ErrDB
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		err = _error.ErrDB
		return
	}

	return affected > 0, nil
}

func (r *Repository) GetInformations(ctx *gin.Context) (informations []domain.Information, err error) {
	query := `SELECT id, answer, question FROM informations`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	informations = make([]domain.Information, 0)
	for rows.Next() {
		information := new(domain.Information)

		if err = rows.Scan(&information.ID, &information.Answer, &information.Question); err != nil {
			return
		}
		informations = append(informations, *information)
	}
	return
}

func (r *Repository) GetInformationByID(ctx *gin.Context, id int64) (information domain.Information, err error) {

	query := `SELECT id, answer, question FROM informations WHERE id = ?`
	err = r.db.QueryRowContext(ctx, query, id).Scan(&information.ID, &information.Answer, &information.Question)
	if err == sql.ErrNoRows {
		err = _error.ErrNotFound
	}
	return
}
