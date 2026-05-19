package productquestion

import (
	"context"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/associate_questions/entity"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/pgutil"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
)

func (r *AssociateQuestionItemRepository) Associate(
	ctx context.Context,
	pq *entity.AssociateQuestion,
) error {
	return r.q.AssociateQuestionItem(ctx, sqlc.AssociateQuestionItemParams{
		ItemCode:   pq.ItemCode,
		QuestionID: pq.QuestionID,
		Position:   int32(pq.Position),
		CreatedAt:  pgutil.ToPgTimestamptz(pq.CreatedAt),
	})
}

func (r *AssociateQuestionItemRepository) ExistsByItemAndQuestion(
	ctx context.Context,
	itemID int64,
	questionID int64,
) (bool, error) {
	return r.q.ExistsByItemAndQuestion(ctx, sqlc.ExistsByItemAndQuestionParams{
		ItemCode:   itemID,
		QuestionID: questionID,
	})
}

func (r *AssociateQuestionItemRepository) ExistsByItemAndPosition(
	ctx context.Context,
	itemID int64,
	position int,
) (bool, error) {
	return r.q.ExistsByItemAndPosition(ctx, sqlc.ExistsByItemAndPositionParams{
		ItemCode: itemID,
		Position: int32(position),
	})
}

func (r *AssociateQuestionItemRepository) GetByItemCode(
	ctx context.Context,
	itemCode int64,
) ([]entity.AssociateQuestionDetail, error) {
	rows, err := r.q.GetQuestionsByItemCode(ctx, itemCode)
	if err != nil {
		return nil, err
	}

	result := make([]entity.AssociateQuestionDetail, 0, len(rows))
	for _, row := range rows {
		var createdAt time.Time
		if row.CreatedAt.Valid {
			createdAt = row.CreatedAt.Time
		}
		result = append(result, entity.AssociateQuestionDetail{
			ItemCode:     row.ItemCode,
			QuestionID:   row.QuestionID,
			QuestionName: row.QuestionName,
			Position:     int(row.Position),
			CreatedAt:    createdAt,
		})
	}
	return result, nil
}

func (r *AssociateQuestionItemRepository) ListAll(
	ctx context.Context,
) ([]entity.ItemQuestionRow, error) {
	rows, err := r.q.ListAllItemQuestions(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]entity.ItemQuestionRow, 0, len(rows))
	for _, row := range rows {
		var createdAt time.Time
		if row.CreatedAt.Valid {
			createdAt = row.CreatedAt.Time
		}
		result = append(result, entity.ItemQuestionRow{
			ItemCode:         row.ItemCode,
			ItemBusinessCode: row.ItemBusinessCode,
			QuestionID:       row.QuestionID,
			QuestionName:     row.QuestionName,
			Position:         int(row.Position),
			CreatedAt:        createdAt,
		})
	}
	return result, nil
}
