package repository

import (
	"context"
	"fmt"
	"time"

	"git.paylabo.com/c002/harp/backend-purchase/domain/domainerror"
	"git.paylabo.com/c002/harp/backend-purchase/domain/model"
	"git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	"git.paylabo.com/c002/harp/backend-purchase/infra"
	ifu "git.paylabo.com/c002/harp/backend-purchase/infra/infra_utils"
	"github.com/jinzhu/gorm"
)

// OrderRepository はstruct
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository はOrderRepositoryを新規する。
func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

// ListOrders は発注一覧を取得する。
func (o *OrderRepository) ListOrders(ctx context.Context, params repository.ListOrdersParams) ([]*model.Order, error) {
	db := ifu.ApplyOrderBys(o.db, params.OrderBy, "id")
	db = filterListOrderQuery(db, params)
	var Orders []*model.Order
	if err := db.Find(&Orders).Error; err != nil {
		return nil, domainerror.NewInternalServerError(err.Error(), err)
	}
	return Orders, nil
}

func getThisMonthUnacceptedOrdersWhereCondition(ids []string) string {
	if len(ids) == 0 {
		return ""
	}
	return ids[0]
}

// ListUnacceptedOrdersThisMonth は未検修発注一覧を取得する。
func (o *OrderRepository) ListUnacceptedOrdersThisMonth(ctx context.Context, params repository.ListOrdersParams) ([]*model.Order, error) {
	var Orders []*model.Order
	sql := "select * from orders inner join (SELECT DISTINCT(orders_id) FROM acceptance_details WHERE DATE_FORMAT(scheduled_acceptance_date,'%Y%m') = DATE_FORMAT(now(),'%Y%m') AND actual_acceptance_date IS NULL) as acceptances_this_month on orders.id = acceptances_this_month.orders_id where request_organization_id='" +
		getThisMonthUnacceptedOrdersWhereCondition(params.RequestOrganizationIDs) +
		"' order by request_date desc"

	if err := o.db.Raw(sql).Scan(&Orders).Error; err != nil {
		return nil, domainerror.NewInternalServerError(err.Error(), err)
	}
	for _, o := range Orders {
		o.AfterFind()
	}
	return Orders, nil
}

// GetOrder はidで発注データを取得する。
func (o *OrderRepository) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	Order := &model.Order{}
	if err := o.db.Where("id = ?", id).Find(Order).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, domainerror.NewNotFoundError(err.Error(), err)
		}
		return nil, domainerror.NewInternalServerError(err.Error(), err)
	}
	return Order, nil
}

// CountOrders は発注をカウントする。
func (o *OrderRepository) CountOrders(ctx context.Context, params repository.CountOrdersParams) (int64, error) {
	db := filterOrderQuery(o.db, params.FilterOrdersParams)
	var count int64
	db.Model(&model.Order{}).Count(&count)
	return count, nil
}

// CreateOrder は発注を作成する。
func (o *OrderRepository) CreateOrder(ctx context.Context, params repository.CreateOrderParams) (repository.CreateOrderReturning, error) {
	Order := &model.Order{
		OrderNo:                    params.OrderNo,
		SuppliersID:                params.SuppliersID,
		CompanyGroupType:           params.CompanyGroupType,
		Subject:                    params.Subject,
		RequestOrganizationID:      params.RequestOrganizationID,
		RequestDate:                infra.ToTime(params.RequestDate),
		RequestBy:                  params.RequestBy,
		ApprovalFile:               params.ApprovalFile,
		DerivationSourceOrderID:    params.DerivationSourceOrderID,
		Remarks:                    params.Remarks,
		SuperiorApprovalDate:       infra.ToTime(params.SuperiorApprovalDate),
		PurchasingDeptApprovalDate: infra.ToTime(params.PurchasingDeptApprovalDate),
		OrderIssueDate:             infra.ToTime(params.OrderIssueDate),
		FinalAcceptanceDate:        infra.ToTime(params.FinalAcceptanceDate),
		AcceptanceCompletedDate:    infra.ToTime(params.AcceptanceCompletedDate),
		CancelDate:                 infra.ToTime(params.CancelDate),
		OrderCaseCd:                params.OrderCaseCd,
		OrderStatus:                params.OrderStatus,
		JiraNo:                     params.JiraNo,
		QuotationsID:               params.QuotationsID,
		OrderApprovalStaffsID:      params.OrderApprovalStaffsID,
		UpdatedAt:                  time.Now(),
	}
	db := o.db.Create(&Order)
	if err := db.Error; err != nil {
		return repository.CreateOrderReturning{
			AffectedRows: db.RowsAffected,
			Order:        nil,
		}, err
	}

	return repository.CreateOrderReturning{
		AffectedRows: db.RowsAffected,
		Order:        Order,
	}, nil
}

// UpdateOrder は発注を更新する。
func (o *OrderRepository) UpdateOrder(ctx context.Context, params repository.UpdateOrderParams) (int64, error) {
	if _, err := o.GetOrder(ctx, params.ID); err != nil {
		return 0, nil
	}

	Order := &model.Order{
		ID:                         ifu.ToInt64(params.ID),
		OrderNo:                    params.OrderNo,
		SuppliersID:                params.SuppliersID,
		CompanyGroupType:           params.CompanyGroupType,
		Subject:                    params.Subject,
		RequestOrganizationID:      params.RequestOrganizationID,
		RequestDate:                infra.ToTime(params.RequestDate),
		RequestBy:                  params.RequestBy,
		ApprovalFile:               params.ApprovalFile,
		DerivationSourceOrderID:    params.DerivationSourceOrderID,
		Remarks:                    params.Remarks,
		SuperiorApprovalDate:       infra.ToTime(params.SuperiorApprovalDate),
		PurchasingDeptApprovalDate: infra.ToTime(params.PurchasingDeptApprovalDate),
		OrderIssueDate:             infra.ToTime(params.OrderIssueDate),
		FinalAcceptanceDate:        infra.ToTime(params.FinalAcceptanceDate),
		AcceptanceCompletedDate:    infra.ToTime(params.AcceptanceCompletedDate),
		CancelDate:                 infra.ToTime(params.CancelDate),
		OrderCaseCd:                params.OrderCaseCd,
		OrderStatus:                params.OrderStatus,
		JiraNo:                     params.JiraNo,
		QuotationsID:               params.QuotationsID,
		OrderApprovalStaffsID:      params.OrderApprovalStaffsID,
	}
	return getOrderAffectedRowsResponse(o.db.Omit("projects_id", "project_cost_id", "cost_types_id").Save(&Order))
}

// UpdateOrderProjectCostInfo は発注データの<プロジェクト費用情報>部分のみ更新する。
func (o *OrderRepository) UpdateOrderProjectCostInfo(ctx context.Context, params repository.UpdateOrderProjectCostInfoParams) (int64, error) {
	if _, err := o.GetOrder(ctx, params.ID); err != nil {
		return 0, nil
	}

	Order := &model.Order{
		ID: ifu.ToInt64(params.ID),
	}

	return getOrderAffectedRowsResponse(o.db.Model(&Order).Updates(map[string]interface{}{
		"projects_id":     params.ProjectsID,
		"project_cost_id": params.ProjectCostID,
		"cost_types_id":   params.CostTypesID,
	}))
}

// DeleteOrder は発注を削除する。
func (o *OrderRepository) DeleteOrder(ctx context.Context, params repository.DeleteOrderParams) (int64, error) {
	//delete orders by id
	Order := &model.Order{
		ID: ifu.ToInt64(params.ID),
	}
	return getOrderAffectedRowsResponse(o.db.Delete(&Order))
}

func (o *OrderRepository) GetOrderRequesterAggregate(ctx context.Context, params repository.GetOrderAggregateParams) ([]*model.Order, error) {
	var Orders []*model.Order
	if err := o.db.Select("DISTINCT request_by").Order("request_by").Find(&Orders).Error; err != nil {
		return nil, domainerror.NewInternalServerError(err.Error(), err)
	}
	return Orders, nil
}
func (o *OrderRepository) GetOrderSupplierAggregate(ctx context.Context, params repository.GetOrderAggregateParams) ([]*model.Order, error) {
	var Orders []*model.Order
	if err := o.db.Select("DISTINCT suppliers_id").Order("suppliers_id").Find(&Orders).Error; err != nil {
		return nil, domainerror.NewInternalServerError(err.Error(), err)
	}
	return Orders, nil
}

// SumNearestTwoMonthsAmountはを先月と今月の発注金額を集計する。
func (o *OrderRepository) SumNearestTwoMonthsAmount(ctx context.Context, params repository.SumNearestTwoMonthsAmountParams) ([]*model.NearestTwoMonthsAmount, error) {
	nearestTwoMonthsAmount := make([]*model.NearestTwoMonthsAmount, 0)
	if params.RequestOrganizationID != "" {
		sqlFormat := `
		SELECT
			month_sequence.yyyymm AS yyyymm,
			IFNULL(CAST(SUM(order_amount_sum) AS UNSIGNED), 0) AS month_sum
	FROM
	(	
		SELECT 
			DATE_FORMAT(DATE_SUB(CURRENT_DATE(), INTERVAL @num := @num +1 MONTH), '%%Y%%m') yyyymm 
				FROM orders, 
			(SELECT @num := -1) num LIMIT 2) AS month_sequence 
			
			LEFT JOIN
			(
				SELECT
					orders.id,
					IFNULL(SUM(	
						order_details.order_quantity * order_details.order_unit_price
						), 0) AS order_amount_sum,
					DATE_FORMAT(orders.order_issue_date, '%%Y%%m') as yyyymm
				FROM
					orders
						LEFT JOIN order_details ON order_details.orders_id = orders.id
				WHERE
					request_organization_id = '%s'
					AND order_issue_date IS NOT NULL
					AND DATE_FORMAT(order_issue_date, '%%Y%%m') = DATE_FORMAT(NOW(), '%%Y%%m')
					OR DATE_FORMAT(order_issue_date, '%%Y%%m') = DATE_FORMAT(CURDATE() - INTERVAL 1 MONTH, '%%Y%%m')
				GROUP BY
					orders.id) AS month_order
					ON month_sequence.yyyymm = month_order.yyyymm
		GROUP BY
			yyyymm
					`
		sqlTemplate := fmt.Sprintf(sqlFormat, params.RequestOrganizationID)
		if err := o.db.Raw(sqlTemplate).Find(&nearestTwoMonthsAmount).Error; err != nil {
			return nil, domainerror.NewInternalServerError(err.Error(), err)
		}
	}
	return nearestTwoMonthsAmount, nil
}

func getOrderAffectedRowsResponse(db *gorm.DB) (int64, error) {
	if err := db.Error; err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}

func filterListOrderQuery(db *gorm.DB, params repository.ListOrdersParams) *gorm.DB {
	if params.Limit != 0 {
		db = db.Limit(params.Limit)
	}
	if params.Offset != 0 {
		db = db.Offset(params.Offset)
	}
	return filterOrderQuery(db, params.FilterOrdersParams)
}

func filterOrderQuery(db *gorm.DB, params repository.FilterOrdersParams) *gorm.DB {
	if len(params.IDs) != 0 {
		db = db.Where("id IN (?)", params.IDs)
	}
	if len(params.OrderCaseCds) != 0 {
		db = db.Where("order_case_cd IN (?)", params.OrderCaseCds)
	}
	if len(params.RequestOrganizationIDs) != 0 {
		db = db.Where("request_organization_id IN (?)", params.RequestOrganizationIDs)
	}
	if len(params.RequestBys) != 0 {
		db = db.Where("request_by IN (?)", params.RequestBys)
	}
	if len(params.SuppliersIDs) != 0 {
		db = db.Where("suppliers_id IN (?)", params.SuppliersIDs)
	}
	if len(params.OrderStatuses) != 0 {
		db = db.Where("order_status IN (?)", params.OrderStatuses)
	}
	if len(params.Subject) != 0 {
		db = db.Where("subject LIKE ?", "%"+params.Subject+"%")
	}
	if len(params.ProjectsIDs) != 0 {
		db = db.Where("projects_id IN (?)", params.ProjectsIDs)
	}
	if len(params.ProjectCostIDs) != 0 {
		db = db.Where("project_cost_id IN (?)", params.ProjectCostIDs)
	}
	return db
}

func (o *OrderRepository) CountOrdersWithGroupBy(ctx context.Context, params repository.CountOrdersWithGroupByParams) ([]*model.OrdersGroupBy, error) {
	ordersGroupBy := make([]*model.OrdersGroupBy, 0)
	// 月別(4ヶ月分)購買案件(発注)取得api
	switch params.GroupBy {
	case "purchasing_dept_approval_date":
		// Dashboard用
		if params.RequestOrganizationID != "" && params.RequestBy == "" {
			sqlFormat := `
						SELECT
							IfNull(monthly_count.count, 0) AS count,
							month_sequence.yyyymm AS value,
							'%s' AS groupBy
						FROM (
							SELECT
								LAST_DAY(DATE_SUB(CURRENT_DATE(), INTERVAL @num := @num + 1 MONTH)) yyyymm
							FROM
								orders,
								(
									SELECT
										@num := - 1) num
								LIMIT %d) AS month_sequence
							LEFT JOIN (
								SELECT
									count(1) AS count,
									yyyymm
								FROM (
									SELECT
										LAST_DAY(purchasing_dept_approval_date) AS yyyymm
									FROM
										orders
									WHERE
										request_organization_id = '%s'
										AND purchasing_dept_approval_date IS NOT NULL
										AND LAST_DAY(purchasing_dept_approval_date) BETWEEN LAST_DAY(DATE_SUB(CURRENT_DATE(), INTERVAL %d MONTH))
										AND LAST_DAY(CURRENT_DATE())) AS o
								GROUP BY
									yyyymm) AS monthly_count ON month_sequence.yyyymm = monthly_count.yyyymm
						`
			sqlTemplate := fmt.Sprintf(sqlFormat, params.GroupBy, params.RecentMonth, params.RequestOrganizationID, params.RecentMonth)
			if err := o.db.Raw(sqlTemplate).Find(&ordersGroupBy).Error; err != nil {
				return nil, domainerror.NewInternalServerError(err.Error(), err)
			}
		}
		// Mypage用
		if params.RequestOrganizationID == "" && params.RequestBy != "" {
			sqlFormat := `
						SELECT
							IfNull(monthly_count.count, 0) AS count,
							month_sequence.yyyymm AS value,
							'%s' AS groupBy
						FROM (
							SELECT
								LAST_DAY(DATE_SUB(CURRENT_DATE(), INTERVAL @num := @num + 1 MONTH)) yyyymm
							FROM
								orders,
								(
									SELECT
										@num := - 1) num
								LIMIT %d) AS month_sequence
							LEFT JOIN (
								SELECT
									count(1) AS count,
									yyyymm
								FROM (
									SELECT
										LAST_DAY(purchasing_dept_approval_date) AS yyyymm
									FROM
										orders
									WHERE
										request_by = '%s'
										AND purchasing_dept_approval_date IS NOT NULL
										AND LAST_DAY(purchasing_dept_approval_date) BETWEEN LAST_DAY(DATE_SUB(CURRENT_DATE(), INTERVAL %d MONTH))
										AND LAST_DAY(CURRENT_DATE())) AS o
								GROUP BY
									yyyymm) AS monthly_count ON month_sequence.yyyymm = monthly_count.yyyymm
						`
			sqlTemplate := fmt.Sprintf(sqlFormat, params.GroupBy, params.RecentMonth, params.RequestBy, params.RecentMonth)
			if err := o.db.Raw(sqlTemplate).Find(&ordersGroupBy).Error; err != nil {
				return nil, domainerror.NewInternalServerError(err.Error(), err)
			}
		}
	// 月別(4ヶ月分)購買案件(検収)取得api
	case "acceptance_completed_date":
		// Dashboard用
		if params.RequestOrganizationID != "" && params.RequestBy == "" {
			sqlFormat := `
						SELECT
							IfNull(monthly_count.count, 0) AS count,
							month_sequence.yyyymm AS value,
							'%s' AS groupBy
						FROM (
							SELECT
								LAST_DAY(DATE_SUB(CURRENT_DATE(), INTERVAL @num := @num + 1 MONTH)) yyyymm
							FROM
								orders,
								(
									SELECT
										@num := - 1) num
								LIMIT %d) AS month_sequence
							LEFT JOIN (
								SELECT
									count(1) AS count,
									yyyymm
								FROM (
									SELECT
										LAST_DAY(acceptance_completed_date) AS yyyymm
									FROM
										orders
									WHERE
										request_organization_id = '%s'
										AND acceptance_completed_date IS NOT NULL
										AND LAST_DAY(acceptance_completed_date) BETWEEN LAST_DAY(DATE_SUB(CURRENT_DATE(), INTERVAL %d MONTH))
										AND LAST_DAY(CURRENT_DATE())) AS o
								GROUP BY
									yyyymm) AS monthly_count ON month_sequence.yyyymm = monthly_count.yyyymm
						`
			sqlTemplate := fmt.Sprintf(sqlFormat, params.GroupBy, params.RecentMonth, params.RequestOrganizationID, params.RecentMonth)
			if err := o.db.Raw(sqlTemplate).Find(&ordersGroupBy).Error; err != nil {
				return nil, domainerror.NewInternalServerError(err.Error(), err)
			}
		}
		// Mypage用
		if params.RequestOrganizationID == "" && params.RequestBy != "" {
			sqlFormat := `
						SELECT
							IfNull(monthly_count.count, 0) AS count,
							month_sequence.yyyymm AS value,
							'%s' AS groupBy
						FROM (
							SELECT
								LAST_DAY(DATE_SUB(CURRENT_DATE(), INTERVAL @num := @num + 1 MONTH)) yyyymm
							FROM
								orders,
								(
									SELECT
										@num := - 1) num
								LIMIT %d) AS month_sequence
							LEFT JOIN (
								SELECT
									count(1) AS count,
									yyyymm
								FROM (
									SELECT
										LAST_DAY(acceptance_completed_date) AS yyyymm
									FROM
										orders
									WHERE
										request_by = '%s'
										AND acceptance_completed_date IS NOT NULL
										AND LAST_DAY(acceptance_completed_date) BETWEEN LAST_DAY(DATE_SUB(CURRENT_DATE(), INTERVAL %d MONTH))
										AND LAST_DAY(CURRENT_DATE())) AS o
								GROUP BY
									yyyymm) AS monthly_count ON month_sequence.yyyymm = monthly_count.yyyymm
						`
			sqlTemplate := fmt.Sprintf(sqlFormat, params.GroupBy, params.RecentMonth, params.RequestBy, params.RecentMonth)
			if err := o.db.Raw(sqlTemplate).Find(&ordersGroupBy).Error; err != nil {
				return nil, domainerror.NewInternalServerError(err.Error(), err)
			}
		}
	// 発注状況内訳取得api
	case "order_status":
		// Dashboard用
		if params.RequestOrganizationID != "" && params.RequestBy == "" {
			sqlTemplate := "SELECT count(1)as count, value, '" +
				params.GroupBy +
				"' AS groupBy FROM(SELECT order_status AS value FROM orders WHERE request_organization_id = '" +
				params.RequestOrganizationID +
				"' )as o GROUP BY value"

			if err := o.db.Raw(sqlTemplate).Find(&ordersGroupBy).Error; err != nil {
				return nil, domainerror.NewInternalServerError(err.Error(), err)
			}
		}
		// Mypage用
		if params.RequestOrganizationID == "" && params.RequestBy != "" {
			sqlTemplate := "SELECT count(1)as count, value, '" +
				params.GroupBy +
				"' AS groupBy FROM(SELECT order_status AS value FROM orders WHERE request_by = '" +
				params.RequestBy +
				"' )as o GROUP BY value"

			if err := o.db.Raw(sqlTemplate).Find(&ordersGroupBy).Error; err != nil {
				return nil, domainerror.NewInternalServerError(err.Error(), err)
			}
		}
	// 発注区分内訳取得api
	case "order_case_cd":
		sqlTemplate := "SELECT count(1)as count, value, '" +
			params.GroupBy +
			"' AS groupBy FROM(SELECT order_case_cd AS value FROM orders WHERE request_organization_id = '" +
			params.RequestOrganizationID +
			"' )as o GROUP BY value"

		if err := o.db.Raw(sqlTemplate).Find(&ordersGroupBy).Error; err != nil {
			return nil, domainerror.NewInternalServerError(err.Error(), err)
		}
	}
	return ordersGroupBy, nil
}
