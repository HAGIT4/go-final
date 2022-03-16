package service

import (
	"sync"

	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

type OrderToWrite struct {
	Number  int
	Accural float32
	Status  string
	Action  string
}

func (sv *BonusService) ProcessOrders() (err error) {
	ordersForProcess, err := sv.getOrdersForProcess()
	if err != nil {
		return err
	}
	if err = sv.markNewWithProcessing(); err != nil {
		return err
	}
	var ordersProcessed []OrderToWrite
	m := sync.Mutex{}
	var wg sync.WaitGroup
	for _, order := range ordersForProcess {
		o := order
		wg.Add(1)
		go func() {
			defer wg.Done()
			orderToWrite := sv.processOrder(&o)
			m.Lock()
			ordersProcessed = append(ordersProcessed, *orderToWrite)
			m.Unlock()
		}()
	}
	wg.Wait()

	var readyOrders []OrderToWrite
	for _, order := range ordersProcessed {
		o := order
		if o.Action == "ok" {
			readyOrders = append(readyOrders, o)
		}
	}

	return nil
}

func (sv *BonusService) markNewWithProcessing() (err error) {
	dbReq := &modelStorage.MarkNewWithProcessingRequest{}
	_, err = sv.storage.MarkNewWithProcessing(dbReq)
	if err != nil {
		return err
	}
	return nil
}

func (sv *BonusService) processOrder(orderToProcess *modelStorage.ProcessedOrder) (orderToWrite *OrderToWrite) {
	resp := sv.accuralClient.GetOrderInfo(orderToProcess.Number)
	orderToWrite = &OrderToWrite{
		Number:  resp.Order,
		Accural: resp.Accural,
		Status:  resp.Status,
		Action:  resp.Action,
	}
	return orderToWrite
}

func (sv *BonusService) getOrdersForProcess() (orders []modelStorage.ProcessedOrder, err error) {
	dbReq := &modelStorage.GetOrdersForProcessRequest{}
	dbResp, err := sv.storage.GetOrdersForProcess(dbReq)
	if err != nil {
		return nil, err
	}
	orders = dbResp.ProcessedOrders
	return orders, nil
}
