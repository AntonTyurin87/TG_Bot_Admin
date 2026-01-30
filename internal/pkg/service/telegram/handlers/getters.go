package telegram

//
//import (
//	"TG_Bot_Admin/internal/pkg/domain/entity"
//	"TG_Bot_Admin/internal/pkg/domain/usecase"
//	"context"
//	"fmt"
//	"log"
//	"sync"
//	"time"
//)
//
//const (
//	CacheTTL    = 5 * time.Minute // Время жизни кэша
//	GRPCTimeout = 10 * time.Second
//)
//
//// Структуры данных
//type PeriodGroup struct {
//	ID   int32  `json:"id"`
//	Name string `json:"name"`
//}
//
//type ItemGroup struct {
//	ID   int32  `json:"id"`
//	Name string `json:"name"`
//}
//
//type SearchData struct {
//	Regions      []*entity.Region `json:"regions"`
//	PeriodGroups []PeriodGroup    `json:"period_groups"` // TODO Заменить на entity
//	ItemGroups   []ItemGroup      `json:"item_groups"`   // TODO Заменить на entity
//}
//
//// Кэши
//var (
//	regionsCache     []*entity.Region
//	regionsCacheTime time.Time
//
//	periodsCache     []PeriodGroup
//	periodsCacheTime time.Time
//
//	itemsCache     []ItemGroup
//	itemsCacheTime time.Time
//
//	cacheMutex sync.RWMutex
//)
//
//// Метод getSearchData получает все необходимые данные для поиска
//func (h *Handler) getSearchData(ctx context.Context) (*SearchData, error) {
//	// Создаем SearchData для заполнения
//	searchData := &SearchData{}
//
//	var mu sync.Mutex
//	var wg sync.WaitGroup
//	errors := make([]error, 0)
//
//	// 1. Получаем регионы
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		regions, err := h.getAllRegionsFromGRPC(ctx)
//		mu.Lock()
//		if err != nil {
//			errors = append(errors, fmt.Errorf("regions: %w", err))
//		} else {
//			searchData.Regions = regions
//		}
//		mu.Unlock()
//	}()
//
//	// 2. Получаем группы периодов
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		periodGroups, err := h.getAllPeriodGroupsFromGRPC(ctx)
//		mu.Lock()
//		if err != nil {
//			errors = append(errors, fmt.Errorf("period groups: %w", err))
//		} else {
//			searchData.PeriodGroups = periodGroups
//		}
//		mu.Unlock()
//	}()
//
//	// 3. Получаем группы источников
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		itemGroups, err := h.getAllItemGroupsFromGRPC(ctx)
//		mu.Lock()
//		if err != nil {
//			errors = append(errors, fmt.Errorf("item groups: %w", err))
//		} else {
//			searchData.ItemGroups = itemGroups
//		}
//		mu.Unlock()
//	}()
//
//	// Ждем завершения всех goroutine
//	wg.Wait()
//
//	// Логируем ошибки, если есть
//	if len(errors) > 0 {
//		log.Printf("Partial errors in getSearchData: %v", errors)
//	}
//
//	// Если регионы не загрузились - это критическая ошибка
//	if len(searchData.Regions) == 0 {
//		return nil, fmt.Errorf("failed to load regions: %v", errors)
//	}
//
//	log.Printf("Search data loaded: %d regions, %d period groups, %d item groups",
//		len(searchData.Regions), len(searchData.PeriodGroups), len(searchData.ItemGroups))
//
//	return searchData, nil
//}
//
//// Получение регионов через gRPC (заглушка)
//func (h *Handler) getAllRegionsFromGRPC(ctx context.Context) ([]*entity.Region, error) {
//	cacheMutex.RLock()
//	// Проверяем кэш (актуален в течение CacheTTL минут)
//	if time.Since(regionsCacheTime) < CacheTTL && len(regionsCache) > 0 {
//		regions := regionsCache
//		cacheMutex.RUnlock()
//		return regions, nil
//	}
//	cacheMutex.RUnlock()
//
//	// Здесь должен быть реальный gRPC вызов
//	regions, err := h.usecase.GetAllRegions(ctx, &usecase.GetAllRegionsRequest{})
//	if err != nil {
//		return nil, err
//	}
//
//	// Обновляем кэш
//	cacheMutex.Lock()
//	regionsCache = regions.GetRegions()
//	regionsCacheTime = time.Now()
//	cacheMutex.Unlock()
//
//	return regions.GetRegions(), nil
//}
//
//// Получение групп периодов через gRPC (заглушка)
//func (h *Handler) getAllPeriodGroupsFromGRPC(ctx context.Context) ([]PeriodGroup, error) {
//	cacheMutex.RLock()
//	// Проверяем кэш
//	if time.Since(periodsCacheTime) < CacheTTL && len(periodsCache) > 0 {
//		periods := periodsCache
//		cacheMutex.RUnlock()
//		return periods, nil
//	}
//	cacheMutex.RUnlock()
//
//	ctx, cancel := context.WithTimeout(ctx, GRPCTimeout)
//	defer cancel()
//
//	// Здесь должен быть реальный gRPC вызов
//	// periodGroups, err := h.usecase.GetAllPeriodGroups(ctx)
//
//	// Заглушка TODO Заменить на реальный вызов!
//	periodGroups := []PeriodGroup{
//		{ID: 1, Name: "6-8 века н.э."},
//		{ID: 2, Name: "9-11 века н.э."},
//		{ID: 3, Name: "12-13 века н.э."},
//	}
//
//	// Обновляем кэш
//	cacheMutex.Lock()
//	periodsCache = periodGroups
//	periodsCacheTime = time.Now()
//	cacheMutex.Unlock()
//
//	return periodGroups, nil
//}
//
//// Получение групп источников через gRPC (заглушка)
//func (h *Handler) getAllItemGroupsFromGRPC(ctx context.Context) ([]ItemGroup, error) {
//	cacheMutex.RLock()
//	// Проверяем кэш
//	if time.Since(itemsCacheTime) < CacheTTL && len(itemsCache) > 0 {
//		items := itemsCache
//		cacheMutex.RUnlock()
//		return items, nil
//	}
//	cacheMutex.RUnlock()
//
//	ctx, cancel := context.WithTimeout(ctx, GRPCTimeout)
//	defer cancel()
//
//	// Здесь должен быть реальный gRPC вызов
//	// itemGroups, err := h.usecase.GetAllItemGroups(ctx)
//
//	// Заглушка TODO Заменить на реальный вызов!
//	itemGroups := []ItemGroup{
//		{ID: 1, Name: "Костюм"},
//		{ID: 2, Name: "Воинское снаряжение"},
//		{ID: 3, Name: "Быт"},
//		{ID: 4, Name: "Ремесло"},
//		{ID: 5, Name: "Пища"},
//		{ID: 6, Name: "Жилища"},
//	}
//
//	// Обновляем кэш
//	cacheMutex.Lock()
//	itemsCache = itemGroups
//	itemsCacheTime = time.Now()
//	cacheMutex.Unlock()
//
//	return itemGroups, nil
//}
