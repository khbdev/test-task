package service

import (
	"context"
	"strconv"

	"test-task/internal/models/elasticc"
	"test-task/internal/models/repostory"
	"test-task/internal/repostory/database"
	"test-task/internal/repostory/elastic"
)

type SlotService struct {
	slotRepo    *database.ShelfSlotRepository
	elasticRepo *elastic.SlotRepository
}

func NewSlotService(
	slotRepo *database.ShelfSlotRepository,
	elasticRepo *elastic.SlotRepository,
) *SlotService {

	return &SlotService{
		slotRepo:    slotRepo,
		elasticRepo: elasticRepo,
	}
}

func (s *SlotService) UpdateSlots(ctx context.Context, companyID string, slots []repostory.ShelfSlot,
) error {

	err := s.slotRepo.UpdateSlots(
		ctx,
		companyID,
		slots,
	)

	if err != nil {
		return err
	}

	updates := make(
		[]elasticc.ProductSlotUpdate,
		0,
		len(slots),
	)

	for _, slot := range slots {

		if slot.ProductID == nil {
			continue
		}

		updates = append(
			updates,
			elasticc.ProductSlotUpdate{
				ID:   *slot.ProductID,
				Slot: strconv.Itoa(slot.Slot),
			},
		)
	}

	return s.elasticRepo.UpdateProductSlots(
		ctx,
		updates,
	)
}

func (s *SlotService) GetSlots(
	ctx context.Context,
	companyID string, page int, limit int, search string,
) ([]repostory.ShelfSlotGET, int, error) {

	return s.slotRepo.GetSlots(
		ctx,
		companyID,
		page,
		limit,
		search,
	)
}
