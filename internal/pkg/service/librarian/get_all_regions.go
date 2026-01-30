package librarian

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// GetAllRegions ...
func (l *librarian) GetAllRegions(ctx context.Context) ([]*entity.Region, error) {
	regions, err := l.getAllRegions(ctx, &lib.GetAllRegionsRequest{})
	if err != nil {
		return nil, fmt.Errorf("librarian getAllRegions: %w", err)
	}

	result := make([]*entity.Region, 0, len(regions.GetRegion()))

	//TODO сделать презентер

	for _, region := range regions.GetRegion() {
		var r entity.Region
		r.ID = region.GetId()
		r.NameRu = region.GetNameRu()
		r.Description = region.GetDescription()

		result = append(result, &r)
	}

	return result, nil
}

func (l *librarian) getAllRegions(ctx context.Context, req *lib.GetAllRegionsRequest) (*lib.GetAllRegionsResponse, error) {
	regions, err := l.librarianClient.GetAllRegions(ctx, &lib.GetAllRegionsRequest{})

	if err != nil {
		fmt.Println("Запрос сейчас тут") //TODO убрать
		fmt.Println("Ошибка - ", err)    //TODO убрать

		return nil, fmt.Errorf("librarian GetAllRegions: %w", err)
	}

	return regions, nil
}
