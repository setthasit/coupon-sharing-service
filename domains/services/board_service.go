package services

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/repositories"

	"github.com/gin-gonic/gin"
)

type BoardService interface {
	Find(ctx *gin.Context, boardUserID uint) ([]entities.Board, error)
	CreateNewBoard(ctx *gin.Context, newBoard *entities.BoardCreateNew) (*entities.Board, error)
}

type BoardServiceInstance struct {
	boardRepo       repositories.BoardRepository
	boardMemberRepo repositories.BoardMemberRepository
}

func NewBoardService(
	boardRepo repositories.BoardRepository,
	boardMemberRepo repositories.BoardMemberRepository,
) BoardService {
	return &BoardServiceInstance{
		boardRepo:       boardRepo,
		boardMemberRepo: boardMemberRepo,
	}
}

func (sv *BoardServiceInstance) Find(ctx *gin.Context, boardUserID uint) ([]entities.Board, error) {
	users, err := sv.boardRepo.Find(ctx, boardUserID)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (sv *BoardServiceInstance) CreateNewBoard(ctx *gin.Context, newBoard *entities.BoardCreateNew) (*entities.Board, error) {
	board, err := sv.boardRepo.Create(ctx, &entities.Board{
		Name:        newBoard.Name,
		BoardUserID: newBoard.BoardUserID,
		CreatedBy:   newBoard.BoardUserID,
		UpdatedBy:   newBoard.BoardUserID,
	})
	if err != nil {
		return nil, err
	}

	filledMembers := []entities.BoardMember{
		{BoardUserID: newBoard.BoardUserID, BoardID: board.ID},
	}
	for _, member := range newBoard.BoardMember {
		filledMembers = append(filledMembers, entities.BoardMember{
			BoardUserID: member.BoardUserID,
			BoardID:     board.ID,
		})
	}
	members, err := sv.boardMemberRepo.CreateBatch(ctx, filledMembers)
	if err != nil {
		return nil, err
	}
	board.BoardMember = members

	return board, nil
}
