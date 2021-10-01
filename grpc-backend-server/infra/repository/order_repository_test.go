package repository

import (
	"context"
	"testing"
	"time"

	"git.paylabo.com/c002/harp/backend-purchase/app/usecase/utils"
	"git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	ifu "git.paylabo.com/c002/harp/backend-purchase/infra/infra_utils"
	"git.paylabo.com/c002/harp/backend-purchase/testutil"
	"github.com/stretchr/testify/assert"
)

func TestOrderRepository(t *testing.T) {
	db, err := testutil.ConnectDatabaseWithGorm()
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
	defer db.Close()
	orderRepository := NewOrderRepository(db)
	t.Run("GetOrder", func(t *testing.T) {
		id := "1"
		Order, err := orderRepository.GetOrder(context.Background(), id)
		assert.NotEmpty(t, Order)
		assert.Nil(t, err)

		assert.NotNil(t, Order.ID)
		assert.Equal(t, int64(1), Order.ID)
		assert.Equal(t, "20H00005", Order.OrderNo)
		assert.Equal(t, "1", Order.SuppliersID)
		assert.Equal(t, "001", Order.CompanyGroupType)
		assert.Equal(t, "test order subject1", Order.Subject)
		assert.Equal(t, "0000001", Order.RequestOrganizationID)
		assert.Equal(t, "2020-09-01T00:00:00+09:00", Order.RequestDate.Format(time.RFC3339))
		assert.Equal(t, "1", Order.RequestBy)
		assert.Equal(t, "file://approval_file.txt", Order.ApprovalFile)
		assert.Nil(t, Order.DerivationSourceOrderID)
		assert.Equal(t, "", Order.Remarks)
		assert.Equal(t, true, Order.SuperiorApprovalDate.IsZero())
		assert.Equal(t, true, Order.PurchasingDeptApprovalDate.IsZero())
		assert.Equal(t, true, Order.OrderIssueDate.IsZero())
		assert.Equal(t, true, Order.AcceptanceCompletedDate.IsZero())
		assert.Equal(t, true, Order.CancelDate.IsZero())
		assert.Equal(t, "10", Order.OrderCaseCd)
		assert.Equal(t, "10", Order.OrderStatus)
		assert.Equal(t, "", Order.JiraNo)
		assert.Nil(t, Order.QuotationsID)
		assert.Equal(t, "", Order.OrderApprovalStaffsID)
		assert.Equal(t, "", Order.ProjectsID)
		assert.Equal(t, "", Order.ProjectCostID)
		assert.Equal(t, "", Order.CostTypesID)
	})
	t.Run("OrderRequesterAggregate", func(t *testing.T) {
		params := repository.GetOrderAggregateParams{}
		nodes, err := orderRepository.GetOrderRequesterAggregate(context.Background(), params)
		assert.Nil(t, err)
		assert.Equal(t, "1", nodes[0].RequestBy)
		assert.Equal(t, "2", nodes[1].RequestBy)
	})
	t.Run("OrderSupplierAggregate", func(t *testing.T) {
		params := repository.GetOrderAggregateParams{}
		nodes, err := orderRepository.GetOrderSupplierAggregate(context.Background(), params)
		assert.Nil(t, err)
		assert.Equal(t, "1", nodes[0].SuppliersID)
		assert.Equal(t, "2", nodes[1].SuppliersID)
	})
	t.Run("CountOrders", func(t *testing.T) {
		t.Run("all", func(t *testing.T) {
			params := repository.CountOrdersParams{}
			count, err := orderRepository.CountOrders(context.Background(), params)
			assert.Nil(t, err)
			assert.Equal(t, int64(15), count)
		})
		t.Run("full_conditions", func(t *testing.T) {
			params := repository.CountOrdersParams{
				FilterOrdersParams: repository.FilterOrdersParams{
					IDs:                    []int32{1, 2, 3},
					SuppliersIDs:           []int32{1, 2, 3},
					Subject:                "order",
					RequestOrganizationIDs: []string{"0000001"},
					RequestBys:             []string{"1", "2", "3"},
					OrderCaseCds:           []int32{10, 19},
					OrderStatuses:          []int32{10, 29},
					ProjectsIDs:            []string{},
					ProjectCostIDs:         []string{},
				},
			}
			count, err := orderRepository.CountOrders(context.Background(), params)
			assert.Nil(t, err)
			assert.Equal(t, int64(3), count)
		})
	})
	t.Run("ListOrders", func(t *testing.T) {
		t.Run("all", func(t *testing.T) {
			params := repository.ListOrdersParams{
				Offset: 0,
				Limit:  100,
			}
			Orders, err := orderRepository.ListOrders(context.Background(), params)
			assert.NotEmpty(t, Orders)
			assert.Nil(t, err)
			assert.Equal(t, 15, len(Orders))

			Order := Orders[0]
			assert.Equal(t, int64(1), Order.ID)
			assert.Equal(t, "20H00005", Order.OrderNo)
			assert.Equal(t, "1", Order.SuppliersID)
			assert.Equal(t, "001", Order.CompanyGroupType)
			assert.Equal(t, "test order subject1", Order.Subject)
			assert.Equal(t, "0000001", Order.RequestOrganizationID)
			assert.Equal(t, "2020-09-01T00:00:00+09:00", Order.RequestDate.Format(time.RFC3339))
			assert.Equal(t, "1", Order.RequestBy)
			assert.Equal(t, "file://approval_file.txt", Order.ApprovalFile)
			assert.Nil(t, Order.DerivationSourceOrderID)
			assert.Equal(t, "", Order.Remarks)
			assert.Equal(t, true, Order.SuperiorApprovalDate.IsZero())
			assert.Equal(t, true, Order.PurchasingDeptApprovalDate.IsZero())
			assert.Equal(t, true, Order.OrderIssueDate.IsZero())
			assert.Equal(t, true, Order.AcceptanceCompletedDate.IsZero())
			assert.Equal(t, true, Order.CancelDate.IsZero())
			assert.Equal(t, "10", Order.OrderCaseCd)
			assert.Equal(t, "10", Order.OrderStatus)
			assert.Equal(t, "", Order.JiraNo)
			assert.Nil(t, Order.QuotationsID)
			assert.Equal(t, "", Order.OrderApprovalStaffsID)
			assert.Equal(t, "", Order.ProjectsID)
			assert.Equal(t, "", Order.ProjectCostID)
			assert.Equal(t, "", Order.CostTypesID)

			Order = Orders[1]
			assert.Equal(t, int64(2), Order.ID)
			assert.Equal(t, "20H00006", Order.OrderNo)
			assert.Equal(t, "2", Order.SuppliersID)
			assert.Equal(t, "001", Order.CompanyGroupType)
			assert.Equal(t, "test order subject2", Order.Subject)
			assert.Equal(t, "0000001", Order.RequestOrganizationID)
			assert.Equal(t, "2020-09-02T00:00:00+09:00", Order.RequestDate.Format(time.RFC3339))
			assert.Equal(t, "2", Order.RequestBy)
			assert.Equal(t, "file://approval_file.txt", Order.ApprovalFile)
			assert.Nil(t, Order.DerivationSourceOrderID)
			assert.Equal(t, "", Order.Remarks)
			assert.Equal(t, true, Order.SuperiorApprovalDate.IsZero())
			assert.Equal(t, true, Order.PurchasingDeptApprovalDate.IsZero())
			assert.Equal(t, true, Order.OrderIssueDate.IsZero())
			assert.Equal(t, true, Order.AcceptanceCompletedDate.IsZero())
			assert.Equal(t, true, Order.CancelDate.IsZero())
			assert.Equal(t, "19", Order.OrderCaseCd)
			assert.Equal(t, "10", Order.OrderStatus)
			assert.Equal(t, "", Order.JiraNo)
			assert.Nil(t, Order.QuotationsID)
			assert.Equal(t, "", Order.OrderApprovalStaffsID)
			assert.Equal(t, "", Order.ProjectsID)
			assert.Equal(t, "", Order.ProjectCostID)
			assert.Equal(t, "", Order.CostTypesID)
		})
		t.Run("orderby_id_asc", func(t *testing.T) {
			params := repository.ListOrdersParams{}
			params.OrderBy = &repository.ListOrdersOrderBy{Id: utils.SortEnumASC}
			orders, err := orderRepository.ListOrders(context.Background(), params)
			assert.Nil(t, err)
			assert.Len(t, orders, 15)
			assert.Equal(t, int64(1), orders[0].ID)
			assert.Equal(t, int64(2), orders[1].ID)
			assert.Equal(t, int64(3), orders[2].ID)
			assert.Equal(t, int64(4), orders[3].ID)
			assert.Equal(t, int64(5), orders[4].ID)
			assert.Equal(t, int64(6), orders[5].ID)
			assert.Equal(t, int64(7), orders[6].ID)
			assert.Equal(t, int64(8), orders[7].ID)
			assert.Equal(t, int64(9), orders[8].ID)
			assert.Equal(t, int64(10), orders[9].ID)
			assert.Equal(t, int64(11), orders[10].ID)
			assert.Equal(t, int64(12), orders[11].ID)
			assert.Equal(t, int64(13), orders[12].ID)
			assert.Equal(t, int64(14), orders[13].ID)
			assert.Equal(t, int64(15), orders[14].ID)
		})
		t.Run("orderby_id_desc", func(t *testing.T) {
			params := repository.ListOrdersParams{}
			params.OrderBy = &repository.ListOrdersOrderBy{Id: utils.SortEnumDESC}
			orders, err := orderRepository.ListOrders(context.Background(), params)
			assert.Nil(t, err)
			assert.Len(t, orders, 15)
			assert.Equal(t, int64(15), orders[0].ID)
			assert.Equal(t, int64(14), orders[1].ID)
			assert.Equal(t, int64(13), orders[2].ID)
			assert.Equal(t, int64(12), orders[3].ID)
			assert.Equal(t, int64(11), orders[4].ID)
			assert.Equal(t, int64(10), orders[5].ID)
			assert.Equal(t, int64(9), orders[6].ID)
			assert.Equal(t, int64(8), orders[7].ID)
			assert.Equal(t, int64(7), orders[8].ID)
			assert.Equal(t, int64(6), orders[9].ID)
			assert.Equal(t, int64(5), orders[10].ID)
			assert.Equal(t, int64(4), orders[11].ID)
			assert.Equal(t, int64(3), orders[12].ID)
			assert.Equal(t, int64(2), orders[13].ID)
			assert.Equal(t, int64(1), orders[14].ID)
		})
		t.Run("full_conditions", func(t *testing.T) {
			params := repository.ListOrdersParams{
				Offset: 0,
				Limit:  100,
				FilterOrdersParams: repository.FilterOrdersParams{
					IDs:                    []int32{1, 2, 3, 4},
					SuppliersIDs:           []int32{1, 2, 3},
					Subject:                "order",
					RequestOrganizationIDs: []string{"0000001", "0000002"},
					RequestBys:             []string{"1", "2", "3"},
					OrderCaseCds:           []int32{10, 19},
					OrderStatuses:          []int32{10, 29},
					ProjectsIDs:            []string{},
					ProjectCostIDs:         []string{},
				},
			}
			Orders, err := orderRepository.ListOrders(context.Background(), params)
			assert.NotEmpty(t, Orders)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(Orders))
		})
	})
	t.Run("CreateOrder", func(t *testing.T) {
		now := time.Now()
		params := repository.CreateOrderParams{
			OrderNo:                    "20H99999",
			SuppliersID:                "999",
			CompanyGroupType:           "999",
			Subject:                    "CreateOrder",
			RequestOrganizationID:      "9",
			RequestDate:                ifu.ToTimestamp(now),
			RequestBy:                  "9",
			ApprovalFile:               "file://approval_file.txt",
			DerivationSourceOrderID:    nil,
			Remarks:                    "remarks",
			SuperiorApprovalDate:       ifu.ToTimestamp(now),
			PurchasingDeptApprovalDate: ifu.ToTimestamp(now),
			OrderIssueDate:             ifu.ToTimestamp(now),
			FinalAcceptanceDate:        ifu.ToTimestamp(now),
			AcceptanceCompletedDate:    ifu.ToTimestamp(now),
			CancelDate:                 ifu.ToTimestamp(now),
			OrderCaseCd:                "10",
			OrderStatus:                "10",
			JiraNo:                     "CTC-9999",
			QuotationsID:               nil,
			OrderApprovalStaffsID:      "",
		}
		response, err := orderRepository.CreateOrder(context.Background(), params)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), response.AffectedRows)

	})
	t.Run("UpdateOrder", func(t *testing.T) {
		now := time.Now().Truncate(24 * time.Hour)
		derivationSourceOrderID := "1"
		quotationsID := "1"
		orderApprovalStaffsID := "1"

		params := repository.UpdateOrderParams{
			ID: "1",
			CreateOrderParams: repository.CreateOrderParams{
				OrderNo:                    "20H99999",
				SuppliersID:                "999",
				CompanyGroupType:           "999",
				Subject:                    "UpdateOrder",
				RequestOrganizationID:      "9",
				RequestDate:                ifu.ToTimestamp(now),
				RequestBy:                  "9",
				ApprovalFile:               "file://approval_file.txt",
				DerivationSourceOrderID:    &derivationSourceOrderID,
				Remarks:                    "remarks",
				SuperiorApprovalDate:       ifu.ToTimestamp(now),
				PurchasingDeptApprovalDate: ifu.ToTimestamp(now),
				OrderIssueDate:             ifu.ToTimestamp(now),
				FinalAcceptanceDate:        ifu.ToTimestamp(now),
				AcceptanceCompletedDate:    ifu.ToTimestamp(now),
				CancelDate:                 ifu.ToTimestamp(now),
				OrderCaseCd:                "99",
				OrderStatus:                "99",
				JiraNo:                     "CTC-9999",
				QuotationsID:               &quotationsID,
				OrderApprovalStaffsID:      orderApprovalStaffsID,
			},
		}
		affectedRows, err := orderRepository.UpdateOrder(context.Background(), params)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), affectedRows)

		Order, _ := orderRepository.GetOrder(context.Background(), "1")

		assert.Equal(t, int64(1), Order.ID)
		assert.Equal(t, "20H99999", Order.OrderNo)
		assert.Equal(t, "999", Order.SuppliersID)
		assert.Equal(t, "999", Order.CompanyGroupType)
		assert.Equal(t, "UpdateOrder", Order.Subject)
		assert.Equal(t, "9", Order.RequestOrganizationID)
		assert.Equal(t, now.Format("2006-01-02"), Order.RequestDate.Format("2006-01-02"))
		assert.Equal(t, "9", Order.RequestBy)
		assert.Equal(t, "file://approval_file.txt", Order.ApprovalFile)
		assert.Equal(t, "1", *Order.DerivationSourceOrderID)
		assert.Equal(t, "remarks", Order.Remarks)
		assert.Equal(t, now.UTC(), Order.SuperiorApprovalDate.UTC())
		assert.Equal(t, now.UTC(), Order.PurchasingDeptApprovalDate.UTC())
		assert.Equal(t, now.UTC(), Order.OrderIssueDate.UTC())
		assert.Equal(t, now.UTC(), Order.AcceptanceCompletedDate.UTC())
		assert.Equal(t, now.UTC(), Order.CancelDate.UTC())
		assert.Equal(t, "99", Order.OrderCaseCd)
		assert.Equal(t, "99", Order.OrderStatus)
		assert.Equal(t, "CTC-9999", Order.JiraNo)
		assert.Equal(t, "1", *Order.QuotationsID)
		assert.Equal(t, "1", Order.OrderApprovalStaffsID)
	})
	t.Run("UpdateOrderProjectCostInfo", func(t *testing.T) {
		params := repository.UpdateOrderProjectCostInfoParams{
			ID:            "1",
			ProjectsID:    "0000001",
			ProjectCostID: "0000002",
			CostTypesID:   "0000003",
		}
		affectedRows, err := orderRepository.UpdateOrderProjectCostInfo(context.Background(), params)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), affectedRows)

		Order, _ := orderRepository.GetOrder(context.Background(), "1")

		assert.Equal(t, int64(1), Order.ID)
		assert.Equal(t, "0000001", Order.ProjectsID)
		assert.Equal(t, "0000002", Order.ProjectCostID)
		assert.Equal(t, "0000003", Order.CostTypesID)
	})
	t.Run("DeleteOrder", func(t *testing.T) {
		params := repository.DeleteOrderParams{
			ID: "15",
		}
		affectedRows, err := orderRepository.DeleteOrder(context.Background(), params)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), affectedRows)

		_, err1 := orderRepository.GetOrder(context.Background(), "15")
		assert.NotNil(t, err1)
	})
	t.Run("SumNearestTwoMonthsAmount", func(t *testing.T) {
		params := repository.SumNearestTwoMonthsAmountParams{
			RequestOrganizationID: "0000001",
		}
		sum, err := orderRepository.SumNearestTwoMonthsAmount(context.Background(), params)
		assert.Nil(t, err)
		assert.Len(t, sum, 2)
		//今月の発注金額と日付
		assert.Equal(t, int64(774500), sum[1].MonthSum)
		assert.Equal(t, "202103", sum[1].Date)
		//先月の発注金額と日付
		assert.Equal(t, int64(492100), sum[0].MonthSum)
		assert.Equal(t, "202102", sum[0].Date)
	})
}
