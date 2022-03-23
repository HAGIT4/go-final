package service

import (
	"strconv"
	"sync"

	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

type OrderToWrite struct {
	Number  int
	Accural float32
	Status  string
	Action  string
	UserID  int
}

func (sv *BonusService) ProcessOrders() (err error) {
	for range sv.updateTicker.C {
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

		for _, order := range readyOrders {
			o := order
			if err := sv.updateOrder(o.Number, o.Status, o.Accural); err != nil {
				return err
			}
			if err := sv.addSumToUserBalance(o.UserID, o.Accural); err != nil {
				return err
			}
		}
	}
	return nil
}

func (sv *BonusService) updateOrder(number int, status string, accrual float32) (err error) {
	dbReq := modelStorage.UpdateOrderRequest{
		Number:  number,
		Status:  status,
		Accural: int(accrual * 100),
	}
	if _, err := sv.storage.UpdateOrder(&dbReq); err != nil {
		return err
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
	order, _ := strconv.Atoi(resp.Order)
	orderToWrite = &OrderToWrite{
		Number:  order,
		Accural: resp.Accural,
		Status:  resp.Status,
		Action:  resp.Action,
		UserID:  orderToProcess.UserID,
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
