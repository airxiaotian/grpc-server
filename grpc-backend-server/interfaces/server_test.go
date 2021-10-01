package interfaces

import (
	"context"
	"strconv"
	"strings"
	"time"

	"net"
	"testing"

	"git.paylabo.com/c002/harp/backend-purchase/domain/model"
	"git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	ifu "git.paylabo.com/c002/harp/backend-purchase/infra/infra_utils"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"

	pb "git.paylabo.com/c002/harp/backend-purchase/interfaces/proto/git.paylabo.com/c002/harp"
	"git.paylabo.com/c002/harp/backend-purchase/mock/mock_repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func TestServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	errMessageInvalidArgument := "rpc error: code = InvalidArgument desc = invalid argument"
	defer ctrl.Finish()

	contractManagerDetailRepository := mock_repository.NewMockContractManagerDetailRepository(ctrl)
	orderDetailsRepository := mock_repository.NewMockOrderDetailsRepository(ctrl)
	orderTypeRepository := mock_repository.NewMockOrderTypeRepository(ctrl)
	quotationRepository := mock_repository.NewMockQuotationRepository(ctrl)
	quotationItemsRepository := mock_repository.NewMockQuotationItemsRepository(ctrl)
	harpSysSequenceRepository := mock_repository.NewMockHarpSysSequenceRepository(ctrl)
	orderItemsRepository := mock_repository.NewMockOrderItemsRepository(ctrl)
	orderRepository := mock_repository.NewMockOrderRepository(ctrl)
	quotationDetailRepository := mock_repository.NewMockQuotationDetailRepository(ctrl)
	itemUnitsRepository := mock_repository.NewMockItemUnitsRepository(ctrl)
	orderStatesRepository := mock_repository.NewMockOrderStatesRepository(ctrl)
	projectCostDetailRepository := mock_repository.NewMockProjectCostDetailRepository(ctrl)
	quotationHistoryRepository := mock_repository.NewMockQuotationHistoryRepository(ctrl)
	acceptanceDetailsRepository := mock_repository.NewMockAcceptanceDetailsRepository(ctrl)
	orderHistoryRepository := mock_repository.NewMockOrderHistoryRepository(ctrl)

	listener := bufconn.Listen(1024 * 1024)
	server := NewServer(ServerParams{
		OrderTypeRepository:             orderTypeRepository,
		QuotationRepository:             quotationRepository,
		OrderRepository:                 orderRepository,
		OrderDetailsRepository:          orderDetailsRepository,
		OrderItemsRepository:            orderItemsRepository,
		OrderStatesRepository:           orderStatesRepository,
		ItemUnitsRepository:             itemUnitsRepository,
		QuotationsRepository:            quotationRepository,
		QuotationItemsRepository:        quotationItemsRepository,
		QuotationDetailRepository:       quotationDetailRepository,
		QuotationHistoryRepository:      quotationHistoryRepository,
		ContractManagerDetailRepository: contractManagerDetailRepository,
		HarpSysSequenceRepository:       harpSysSequenceRepository,
		ProjectCostDetailRepository:     projectCostDetailRepository,
		AcceptanceDetailsRepository:     acceptanceDetailsRepository,
		OrderHistoryRepository:          orderHistoryRepository,
	})
	go func() {
		_ = server.Serve(listener)
	}()
	defer server.GracefulStop()

	dialer := func(ctx context.Context, address string) (net.Conn, error) {
		return listener.Dial()
	}
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	assert.Nil(t, err)

	orderTypeClient := pb.NewOrderTypeClient(conn)
	orderClient := pb.NewOrderClient(conn)
	orderDetailClient := pb.NewOrderDetailClient(conn)
	orderItemClient := pb.NewOrderItemClient(conn)
	orderStateClient := pb.NewOrderStateClient(conn)
	itemUnitClient := pb.NewItemUnitClient(conn)
	quotationClient := pb.NewQuotationClient(conn)
	quotationItemClient := pb.NewQuotationItemClient(conn)
	acceptanceDetailClient := pb.NewAcceptanceDetailClient(conn)
	quotationDetailClient := pb.NewQuotationDetailClient(conn)
	quotationHistoryClient := pb.NewQuotationHistoryClient(conn)
	contractManagerDetailClient := pb.NewContractManagerDetailClient(conn)
	harpSysSequenceClient := pb.NewHarpSysSequenceClient(conn)
	projectCostDetailClient := pb.NewProjectCostDetailClient(conn)
	orderHistoryClient := pb.NewOrderHistoryClient(conn)

	t.Run("OrderItem", func(t *testing.T) {

		t.Run("CountOrderItem", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				request := &pb.FilterOrderItemsRequest{
					Ids:             []string{"1", "2"},
					OrderDetailsIds: []string{"1", "2"},
					OrdersIds:       []string{"1", "2"},
					OrderDetailsNos: []string{"1", "2"},
				}
				orderItemsRepository.EXPECT().CountOrderItems(gomock.Any(), gomock.Any()).Return(int64(2), nil)
				res, err := orderItemClient.CountOrderItems(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(2), res.Count)
			})
			t.Run("full_conditions", func(t *testing.T) {
				request := &pb.FilterOrderItemsRequest{Ids: []string{"1"}}
				orderItemsRepository.EXPECT().CountOrderItems(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				res, err := orderItemClient.CountOrderItems(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(1), res.Count)
			})
		})

		t.Run("ListOrderItems", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {

				orderDetailsNo := "6"
				orderItems := []*model.OrderItem{
					{
						Id:             "1",
						OrdersId:       "2",
						OrderDetailsId: "3",
						ProductName:    "mockedProductName",
						OrderQuantity:  "4",
						OrderPrice:     "5",
						OrderDetailsNo: &orderDetailsNo,
					},
				}

				request := &pb.ListOrderItemsRequest{}
				orderItemsRepository.EXPECT().ListOrderItems(gomock.Any(), gomock.Any()).Return(orderItems, nil)
				res, err := orderItemClient.ListOrderItems(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, len(orderItems), len(res.OrderItems))
				assert.Equal(t, int32(1), res.OrderItems[0].Id.Value)
				assert.Equal(t, int32(2), res.OrderItems[0].OrdersId.Value)
				assert.Equal(t, int32(3), res.OrderItems[0].OrderDetailsId.Value)
				assert.Equal(t, "mockedProductName", res.OrderItems[0].ProductName)
				assert.Equal(t, int32(4), res.OrderItems[0].OrderQuantity.Value)
				assert.Equal(t, int32(5), res.OrderItems[0].OrderPrice.Value)
				assert.Equal(t, int32(6), res.OrderItems[0].OrderDetailsNo.Value)
			})

			t.Run("null_fields", func(t *testing.T) {
				orderItems := []*model.OrderItem{{
					Id:             "1",
					OrdersId:       "2",
					OrderDetailsId: "3",
					ProductName:    "mockedProductName",
					OrderQuantity:  "4",
					OrderPrice:     "5",
					OrderDetailsNo: nil,
				}}
				request := &pb.ListOrderItemsRequest{}
				orderItemsRepository.EXPECT().ListOrderItems(gomock.Any(), gomock.Any()).Return(orderItems, nil)
				res, err := orderItemClient.ListOrderItems(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, len(orderItems), len(res.OrderItems))
				assert.Equal(t, int32(1), res.OrderItems[0].Id.Value)
				assert.Equal(t, int32(2), res.OrderItems[0].OrdersId.Value)
				assert.Equal(t, int32(3), res.OrderItems[0].OrderDetailsId.Value)
				assert.Equal(t, "mockedProductName", res.OrderItems[0].ProductName)
				assert.Equal(t, int32(4), res.OrderItems[0].OrderQuantity.Value)
				assert.Equal(t, int32(5), res.OrderItems[0].OrderPrice.Value)
				assert.Nil(t, res.OrderItems[0].OrderDetailsNo)
			})

			t.Run("full_conditions", func(t *testing.T) {
				orderItems := []*model.OrderItem{{
					Id:             "1",
					OrdersId:       "2",
					OrderDetailsId: "3",
					ProductName:    "mockedProductName",
					OrderQuantity:  "4",
					OrderPrice:     "5",
					OrderDetailsNo: nil,
				}}
				request := &pb.ListOrderItemsRequest{
					Ids:             []string{"1", "2"},
					Limit:           10,
					Offset:          10,
					OrderDetailsIds: []string{"1", "2"},
					OrdersIds:       []string{"1", "2"},
					OrderDetailsNos: []string{"1", "2"},
				}
				orderItemsRepository.EXPECT().ListOrderItems(gomock.Any(), gomock.Any()).Return(orderItems, nil)
				res, err := orderItemClient.ListOrderItems(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, len(orderItems), len(res.OrderItems))
				assert.Equal(t, int32(1), res.OrderItems[0].Id.Value)
				assert.Equal(t, int32(2), res.OrderItems[0].OrdersId.Value)
				assert.Equal(t, int32(3), res.OrderItems[0].OrderDetailsId.Value)
				assert.Equal(t, "mockedProductName", res.OrderItems[0].ProductName)
				assert.Equal(t, int32(4), res.OrderItems[0].OrderQuantity.Value)
				assert.Equal(t, int32(5), res.OrderItems[0].OrderPrice.Value)
				assert.Nil(t, res.OrderItems[0].OrderDetailsNo)
			})
		})

		t.Run("GetOrderItem", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				orderDetailsNo := "6"
				orderItem := &model.OrderItem{
					Id:             "1",
					OrdersId:       "2",
					OrderDetailsId: "3",
					ProductName:    "mockedProductName",
					OrderQuantity:  "4",
					OrderPrice:     "5",
					OrderDetailsNo: &orderDetailsNo,
				}
				request := &pb.GetOrderItemRequest{Id: "1"}
				orderItemsRepository.EXPECT().GetOrderItem(gomock.Any(), gomock.Any()).Return(orderItem, nil)
				res, err := orderItemClient.GetOrderItem(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, int32(2), res.OrdersId.Value)
				assert.Equal(t, int32(3), res.OrderDetailsId.Value)
				assert.Equal(t, "mockedProductName", res.ProductName)
				assert.Equal(t, int32(4), res.OrderQuantity.Value)
				assert.Equal(t, int32(5), res.OrderPrice.Value)
				assert.Equal(t, int32(6), res.OrderDetailsNo.Value)
			})
			t.Run("by_id_null_fields", func(t *testing.T) {
				orderItem := &model.OrderItem{
					Id:             "1",
					OrdersId:       "2",
					OrderDetailsId: "3",
					ProductName:    "mockedProductName",
					OrderQuantity:  "4",
					OrderPrice:     "5",
					OrderDetailsNo: nil,
				}
				request := &pb.GetOrderItemRequest{Id: "1"}
				orderItemsRepository.EXPECT().GetOrderItem(gomock.Any(), gomock.Any()).Return(orderItem, nil)
				res, err := orderItemClient.GetOrderItem(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, int32(2), res.OrdersId.Value)
				assert.Equal(t, int32(3), res.OrderDetailsId.Value)
				assert.Equal(t, "mockedProductName", res.ProductName)
				assert.Equal(t, int32(4), res.OrderQuantity.Value)
				assert.Equal(t, int32(5), res.OrderPrice.Value)
				assert.Nil(t, res.OrderDetailsNo)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetOrderItemRequest{Id: ""}
				res, err := orderItemClient.GetOrderItem(context.Background(), request)
				assert.NotNil(t, err)
				assert.Nil(t, res)
			})
		})

		t.Run("DeleteOrderItem", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteOrderItemRequest{Id: "1"}
				orderItemsRepository.EXPECT().DeleteOrderItem(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderItemClient.DeleteOrderItem(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error_id_unspecified", func(t *testing.T) {
				request := &pb.DeleteOrderItemRequest{}
				res, err := orderItemClient.DeleteOrderItem(context.Background(), request)
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), errMessageInvalidArgument)
				assert.Nil(t, res)
			})
		})

		t.Run("UpdateOrderItem", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				request := &pb.UpdateOrderItemRequest{
					Id:             "1",
					OrdersId:       "2",
					OrderDetailsId: "3",
					ProductName:    "mockedProductName",
					OrderQuantity:  "4",
					OrderPrice:     "5",
					OrderDetailsNo: &wrappers.Int32Value{Value: 6}}

				affectedRows := int64(1)
				orderItemsRepository.EXPECT().UpdateOrderItem(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderItemClient.UpdateOrderItem(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})

			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateOrderItemRequest{
					Id:             "",
					OrdersId:       "2",
					OrderDetailsId: "3",
					ProductName:    "mockedProductName",
					OrderQuantity:  "4",
					OrderPrice:     "5",
					OrderDetailsNo: nil}
				res, err := orderItemClient.UpdateOrderItem(context.Background(), request)
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), errMessageInvalidArgument)
				assert.Nil(t, res)
			})
		})

		t.Run("CreateOrderItem", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.CreateOrderItemRequest{
					OrdersId:       "2",
					OrderDetailsId: "3",
					ProductName:    "mockedProductName",
					OrderQuantity:  "4",
					OrderPrice:     "5",
					OrderDetailsNo: &wrappers.Int32Value{Value: 6}}
				orderItemsRepository.EXPECT().CreateOrderItem(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderItemClient.CreateOrderItem(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_failed", func(t *testing.T) {
				request := &pb.CreateOrderItemRequest{
					OrdersId:       "",
					OrderDetailsId: "",
					OrderQuantity:  "",
					OrderPrice:     "",
					OrderDetailsNo: &wrappers.Int32Value{Value: 6}}
				res, err := orderItemClient.CreateOrderItem(context.Background(), request)
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), errMessageInvalidArgument)
				assert.Nil(t, res)
			})
		})

		t.Run("SumOrderItemsPrice", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				sum := int64(3)
				request := &pb.FilterOrderItemsRequest{}
				orderItemsRepository.EXPECT().SumOrderItemsPrice(gomock.Any(), gomock.Any()).Return(sum, nil)
				res, err := orderItemClient.SumOrderItemsPrice(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, sum, res.Sum)
			})
			t.Run("full_conditions", func(t *testing.T) {
				sum := int64(3)
				request := &pb.FilterOrderItemsRequest{
					Ids:             []string{"1", "2"},
					OrderDetailsIds: []string{"1", "2"},
					OrdersIds:       []string{"1", "2"},
					OrderDetailsNos: []string{"1", "2"},
				}
				orderItemsRepository.EXPECT().SumOrderItemsPrice(gomock.Any(), gomock.Any()).Return(sum, nil)
				res, err := orderItemClient.SumOrderItemsPrice(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, sum, res.Sum)
			})
		})

		t.Run("SumOrderItemsQuantity", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				sum := int64(3)
				request := &pb.FilterOrderItemsRequest{}
				orderItemsRepository.EXPECT().SumOrderItemsQuantity(gomock.Any(), gomock.Any()).Return(sum, nil)
				res, err := orderItemClient.SumOrderItemsQuantity(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, sum, res.Sum)
			})
			t.Run("full_conditions", func(t *testing.T) {
				sum := int64(3)
				request := &pb.FilterOrderItemsRequest{
					Ids:             []string{"1", "2"},
					OrderDetailsIds: []string{"1", "2"},
					OrdersIds:       []string{"1", "2"},
					OrderDetailsNos: []string{"1", "2"},
				}
				orderItemsRepository.EXPECT().SumOrderItemsQuantity(gomock.Any(), gomock.Any()).Return(sum, nil)
				res, err := orderItemClient.SumOrderItemsQuantity(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, sum, res.Sum)
			})
		})

	})
	t.Run("ContractManagerDetail", func(t *testing.T) {
		t.Run("ListContractManagerDetail", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				staffsID := "123"
				gExuserID := "123"
				contractManagerDetails := []*model.ContractManagerDetail{{
					ID:                  "1",
					OrdersID:            "123",
					ManagerPurchaseID:   "123",
					ContractManagerType: "mockContractManagerType",
					DeputyType:          "mockDeputyType",
					GInoutType:          "mockGInoutType",
					StaffsID:            &staffsID,
					GExuserID:           &gExuserID,
				}, {
					ID:                  "2",
					OrdersID:            "123",
					ManagerPurchaseID:   "123",
					ContractManagerType: "mockContractManagerType",
					DeputyType:          "mockDeputyType",
					GInoutType:          "mockGInoutType",
					StaffsID:            &staffsID,
					GExuserID:           &gExuserID,
				},
				}
				request := &pb.ListContractManagerDetailsRequest{}
				contractManagerDetailRepository.EXPECT().ListContractManagerDetails(gomock.Any(), gomock.Any()).Return(contractManagerDetails, nil)
				res, err := contractManagerDetailClient.ListContractManagerDetails(context.Background(), request)
				//error
				assert.Nil(t, err)
				//length
				assert.Equal(t, 2, len(res.ContractManagerDetails))
				//ContractManagerDetails[0]
				assert.Equal(t, int32(1), res.ContractManagerDetails[0].Id.Value)
				assert.Equal(t, int32(123), res.ContractManagerDetails[0].OrdersId.Value)
				assert.Equal(t, int32(123), res.ContractManagerDetails[0].ManagerPurchaseId.Value)
				assert.Equal(t, "mockContractManagerType", res.ContractManagerDetails[0].ContractManagerType)
				assert.Equal(t, "mockDeputyType", res.ContractManagerDetails[0].DeputyType)
				assert.Equal(t, "mockGInoutType", res.ContractManagerDetails[0].GInoutType)
				assert.Equal(t, "123", res.ContractManagerDetails[0].StaffsId)
				assert.Equal(t, int32(123), res.ContractManagerDetails[0].GExuserId.Value)
				//ContractManagerDetails[1]
				assert.Equal(t, int32(2), res.ContractManagerDetails[1].Id.Value)
				assert.Equal(t, int32(123), res.ContractManagerDetails[1].OrdersId.Value)
				assert.Equal(t, int32(123), res.ContractManagerDetails[1].ManagerPurchaseId.Value)
				assert.Equal(t, "mockContractManagerType", res.ContractManagerDetails[1].ContractManagerType)
				assert.Equal(t, "mockDeputyType", res.ContractManagerDetails[1].DeputyType)
				assert.Equal(t, "mockGInoutType", res.ContractManagerDetails[1].GInoutType)
				assert.Equal(t, "123", res.ContractManagerDetails[1].StaffsId)
				assert.Equal(t, int32(123), res.ContractManagerDetails[1].GExuserId.Value)

			})
			t.Run("full_conditions", func(t *testing.T) {
				staffsID := "123"
				gExuserID := "123"
				contractManagerDetails := []*model.ContractManagerDetail{{
					ID:                  "1",
					OrdersID:            "123",
					ManagerPurchaseID:   "123",
					ContractManagerType: "mockContractManagerType",
					DeputyType:          "mockDeputyType",
					GInoutType:          "mockGInoutType",
					StaffsID:            &staffsID,
					GExuserID:           &gExuserID,
				}, {
					ID:                  "2",
					OrdersID:            "123",
					ManagerPurchaseID:   "123",
					ContractManagerType: "mockContractManagerType",
					DeputyType:          "mockDeputyType",
					GInoutType:          "mockGInoutType",
					StaffsID:            &staffsID,
					GExuserID:           &gExuserID,
				},
				}
				request := &pb.ListContractManagerDetailsRequest{
					Ids:    []string{"1", "2"},
					Limit:  int64(10),
					Offset: int64(0),
				}
				contractManagerDetailRepository.EXPECT().ListContractManagerDetails(gomock.Any(), gomock.Any()).Return(contractManagerDetails, nil)
				res, err := contractManagerDetailClient.ListContractManagerDetails(context.Background(), request)
				//error
				assert.Nil(t, err)
				//length
				assert.Equal(t, 2, len(res.ContractManagerDetails))
				//ContractManagerDetails[0]
				assert.Equal(t, int32(1), res.ContractManagerDetails[0].Id.Value)
				assert.Equal(t, int32(123), res.ContractManagerDetails[0].OrdersId.Value)
				assert.Equal(t, int32(123), res.ContractManagerDetails[0].ManagerPurchaseId.Value)
				assert.Equal(t, "mockContractManagerType", res.ContractManagerDetails[0].ContractManagerType)
				assert.Equal(t, "mockDeputyType", res.ContractManagerDetails[0].DeputyType)
				assert.Equal(t, "mockGInoutType", res.ContractManagerDetails[0].GInoutType)
				assert.Equal(t, "123", res.ContractManagerDetails[0].StaffsId)
				assert.Equal(t, int32(123), res.ContractManagerDetails[0].GExuserId.Value)
				//ContractManagerDetails[1]
				assert.Equal(t, int32(2), res.ContractManagerDetails[1].Id.Value)
				assert.Equal(t, int32(123), res.ContractManagerDetails[1].OrdersId.Value)
				assert.Equal(t, int32(123), res.ContractManagerDetails[1].ManagerPurchaseId.Value)
				assert.Equal(t, "mockContractManagerType", res.ContractManagerDetails[1].ContractManagerType)
				assert.Equal(t, "mockDeputyType", res.ContractManagerDetails[1].DeputyType)
				assert.Equal(t, "mockGInoutType", res.ContractManagerDetails[1].GInoutType)
				assert.Equal(t, "123", res.ContractManagerDetails[1].StaffsId)
				assert.Equal(t, int32(123), res.ContractManagerDetails[1].GExuserId.Value)

			})
		})
		t.Run("GetContractManagerDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetContractManagerDetailRequest{}
				res, err := contractManagerDetailClient.GetContractManagerDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("by_id", func(t *testing.T) {
				staffsID := "123"
				gExuserID := "123"
				contractManagerDetail := &model.ContractManagerDetail{
					ID:                  "1",
					OrdersID:            "123",
					ManagerPurchaseID:   "123",
					ContractManagerType: "mockContractManagerType",
					DeputyType:          "mockDeputyType",
					GInoutType:          "mockGInoutType",
					StaffsID:            &staffsID,
					GExuserID:           &gExuserID,
				}
				request := &pb.GetContractManagerDetailRequest{Id: "1"}
				contractManagerDetailRepository.EXPECT().GetContractManagerDetail(gomock.Any(), gomock.Any()).Return(contractManagerDetail, nil)
				res, err := contractManagerDetailClient.GetContractManagerDetail(context.Background(), request)

				//error
				assert.Nil(t, err)
				//ContractManagerDetail
				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, int32(123), res.OrdersId.Value)
				assert.Equal(t, int32(123), res.ManagerPurchaseId.Value)
				assert.Equal(t, "mockContractManagerType", res.ContractManagerType)
				assert.Equal(t, "mockDeputyType", res.DeputyType)
				assert.Equal(t, "mockGInoutType", res.GInoutType)
				assert.Equal(t, "123", res.StaffsId)
				assert.Equal(t, int32(123), res.GExuserId.Value)
			})
		})
		t.Run("CreateContractManagerDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.CreateContractManagerDetailRequest{
					OrdersId:            "",
					ManagerPurchaseId:   "",
					ContractManagerType: strings.Repeat("A", 2),
					DeputyType:          strings.Repeat("A", 2),
					GInoutType:          strings.Repeat("A", 2),
					StaffsId:            "10",
					GExuserId:           ifu.ToInt32Value("10"),
				}
				res, err := contractManagerDetailClient.CreateContractManagerDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)

				request := &pb.CreateContractManagerDetailRequest{
					Id:                  "123",
					OrdersId:            "123",
					ManagerPurchaseId:   "123",
					ContractManagerType: "A",
					DeputyType:          "A",
					GInoutType:          "A",
					StaffsId:            "1",
					GExuserId:           &wrappers.Int32Value{Value: 2},
				}
				contractManagerDetailRepository.EXPECT().CreateContractManagerDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := contractManagerDetailClient.CreateContractManagerDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)

			})
		})
		t.Run("UpdateContractManagerDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateContractManagerDetailRequest{
					Id:                  "",
					OrdersId:            "",
					ManagerPurchaseId:   "",
					ContractManagerType: strings.Repeat("A", 2),
					DeputyType:          strings.Repeat("A", 2),
					GInoutType:          strings.Repeat("A", 2),
					StaffsId:            "10",
					GExuserId:           ifu.ToInt32Value("10"),
				}
				res, err := contractManagerDetailClient.UpdateContractManagerDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)

				request := &pb.UpdateContractManagerDetailRequest{
					Id:                  "123",
					OrdersId:            "123",
					ManagerPurchaseId:   "123",
					ContractManagerType: "A",
					DeputyType:          "A",
					GInoutType:          "A",
					StaffsId:            "1",
					GExuserId:           &wrappers.Int32Value{Value: 2},
				}
				contractManagerDetailRepository.EXPECT().UpdateContractManagerDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := contractManagerDetailClient.UpdateContractManagerDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
		})
		t.Run("DeleteContractManagerDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.DeleteContractManagerDetailRequest{}
				res, err := contractManagerDetailClient.DeleteContractManagerDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("by_id", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteContractManagerDetailRequest{Id: "1"}
				contractManagerDetailRepository.EXPECT().DeleteContractManagerDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := contractManagerDetailClient.DeleteContractManagerDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
		})
		t.Run("CountContractManagerDetail", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				request := &pb.ListContractManagerDetailsRequest{}
				contractManagerDetailRepository.EXPECT().CountContractManagerDetails(gomock.Any(), gomock.Any()).Return(int64(2), nil)
				res, err := contractManagerDetailClient.CountContractManagerDetails(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(2), res.Count)
			})
			t.Run("full_conditions", func(t *testing.T) {
				request := &pb.ListContractManagerDetailsRequest{
					Ids:    []string{"1", "2", "3"},
					Limit:  int64(10),
					Offset: int64(0),
				}
				contractManagerDetailRepository.EXPECT().CountContractManagerDetails(gomock.Any(), gomock.Any()).Return(int64(3), nil)
				res, err := contractManagerDetailClient.CountContractManagerDetails(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(3), res.Count)
			})
		})
	})
	t.Run("OrderDetail", func(t *testing.T) {
		now := time.Now()
		nowTimestamp := ifu.ToTimestamp(now)
		t.Run("ListOrderDetail", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				orderDetailsNo := "2"
				cancelQuantity := "3"
				orderUnitClassification := "6"
				quotationsID := "10"
				quotationDetailsID := "11"
				orderDetails := []*model.OrderDetail{{
					Id:                                "1",
					Remarks:                           "MockedRemarks",
					ProductName:                       "MockedProductName",
					OrderDetailsNo:                    &orderDetailsNo,
					CancelQuantity:                    &cancelQuantity,
					ConfigurationManagementTargetFlag: "a",
					AcceptanceScheduledDate:           now,
					OrderQuantity:                     "5",
					OrderUnitClassification:           &orderUnitClassification,
					OrderUnitPrice:                    "7",
					OrdersId:                          "8",
					QuotationsId:                      &quotationsID,
					Specifications:                    "MockedSpecifications",
					QuotationDetailsId:                &quotationDetailsID,
				}}

				request := &pb.ListOrderDetailsRequest{}
				orderDetailsRepository.EXPECT().ListOrderDetails(gomock.Any(), gomock.Any()).Return(orderDetails, nil)
				res, err := orderDetailClient.ListOrderDetails(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, len(orderDetails), len(res.OrderDetails))
				assert.Equal(t, int32(1), res.OrderDetails[0].Id.Value)
				assert.Equal(t, "MockedRemarks", res.OrderDetails[0].Remarks)
				assert.Equal(t, "MockedProductName", res.OrderDetails[0].ProductName)
				assert.Equal(t, int32(2), res.OrderDetails[0].OrderDetailsNo.Value)
				assert.Equal(t, float64(3), res.OrderDetails[0].CancelQuantity.Value)
				assert.Equal(t, "a", res.OrderDetails[0].ConfigurationManagementTargetFlag)
				assert.Equal(t, now.UTC(), res.OrderDetails[0].AcceptanceScheduledDate.AsTime().UTC())
				assert.Equal(t, float64(5), res.OrderDetails[0].OrderQuantity.Value)
				assert.Equal(t, int32(6), res.OrderDetails[0].OrderUnitClassification.Value)
				assert.Equal(t, int32(7), res.OrderDetails[0].OrderUnitPrice.Value)
				assert.Equal(t, int32(8), res.OrderDetails[0].OrdersId.Value)
				assert.Equal(t, int32(10), res.OrderDetails[0].QuotationsId.Value)
				assert.Equal(t, "MockedSpecifications", res.OrderDetails[0].Specifications)
				assert.Equal(t, int32(11), res.OrderDetails[0].QuotationDetailsId.Value)
			})
			t.Run("full_conditions", func(t *testing.T) {
				orderDetailsNo := "2"
				cancelQuantity := "3"
				orderUnitClassification := "6"
				quotationsID := "10"
				quotationDetailsID := "11"
				orderDetails := []*model.OrderDetail{{
					Id:                                "1",
					Remarks:                           "MockedRemarks",
					ProductName:                       "MockedProductName",
					OrderDetailsNo:                    &orderDetailsNo,
					CancelQuantity:                    &cancelQuantity,
					ConfigurationManagementTargetFlag: "a",
					AcceptanceScheduledDate:           now,
					OrderQuantity:                     "5",
					OrderUnitClassification:           &orderUnitClassification,
					OrderUnitPrice:                    "7",
					OrdersId:                          "8",
					QuotationsId:                      &quotationsID,
					Specifications:                    "MockedSpecifications",
					QuotationDetailsId:                &quotationDetailsID,
				}}

				request := &pb.ListOrderDetailsRequest{
					Ids:                                []string{"1", "2"},
					ConfigurationManagementTargetFlags: []string{"1", "2"},
					//発注行番
					OrderDetailsNos: []string{"1", "2"},
					//発注単位区分
					OrderUnitClassifications: []string{"1", "2"},
					//発注見出し_i_d
					OrdersIds: []string{"1", "2"},
					//プロジェクト_i_d
					//見積発注行番
					QuotationDetailsId:          "1",
					QuotationsIds:               []string{"1", "2"},
					Limit:                       1,
					Offset:                      1,
					AcceptanceScheduledDateFrom: nowTimestamp,
					AcceptanceScheduledDateTo:   nowTimestamp,
					OrderBy: &pb.ListOrderDetailsOrderBy{
						Id:                                pb.SortEnum_ASC,
						CancelQuantity:                    pb.SortEnum_ASC,
						ConfigurationManagementTargetFlag: pb.SortEnum_ASC,
						AcceptanceScheduledDate:           pb.SortEnum_ASC,
						OrderDetailsNo:                    pb.SortEnum_ASC,
						OrderQuantity:                     pb.SortEnum_ASC,
						OrderUnitClassification:           pb.SortEnum_ASC,
						OrderUnitPrice:                    pb.SortEnum_ASC,
						OrdersId:                          pb.SortEnum_ASC,
						ProductName:                       pb.SortEnum_ASC,
						QuotationDetailsId:                pb.SortEnum_ASC,
						QuotationsId:                      pb.SortEnum_ASC,
						Remarks:                           pb.SortEnum_ASC,
						Specifications:                    pb.SortEnum_ASC,
					},
				}
				orderDetailsRepository.EXPECT().ListOrderDetails(gomock.Any(), gomock.Any()).Return(orderDetails, nil)
				res, err := orderDetailClient.ListOrderDetails(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, len(orderDetails), len(res.OrderDetails))
				assert.Equal(t, int32(1), res.OrderDetails[0].Id.Value)
				assert.Equal(t, "MockedRemarks", res.OrderDetails[0].Remarks)
				assert.Equal(t, "MockedProductName", res.OrderDetails[0].ProductName)
				assert.Equal(t, int32(2), res.OrderDetails[0].OrderDetailsNo.Value)
				assert.Equal(t, float64(3), res.OrderDetails[0].CancelQuantity.Value)
				assert.Equal(t, "a", res.OrderDetails[0].ConfigurationManagementTargetFlag)
				assert.Equal(t, now.UTC(), res.OrderDetails[0].AcceptanceScheduledDate.AsTime().UTC())
				assert.Equal(t, float64(5), res.OrderDetails[0].OrderQuantity.Value)
				assert.Equal(t, int32(6), res.OrderDetails[0].OrderUnitClassification.Value)
				assert.Equal(t, int32(7), res.OrderDetails[0].OrderUnitPrice.Value)
				assert.Equal(t, int32(8), res.OrderDetails[0].OrdersId.Value)
				assert.Equal(t, int32(10), res.OrderDetails[0].QuotationsId.Value)
				assert.Equal(t, "MockedSpecifications", res.OrderDetails[0].Specifications)
				assert.Equal(t, int32(11), res.OrderDetails[0].QuotationDetailsId.Value)
			})

		})
		t.Run("GetOrderDetail", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				orderDetailsNo := "2"
				cancelQuantity := "3"
				orderUnitClassification := "6"
				quotationsID := "10"
				quotationDetailsID := "11"
				orderDetail := &model.OrderDetail{
					Id:                                "1",
					Remarks:                           "MockedRemarks",
					ProductName:                       "MockedProductName",
					OrderDetailsNo:                    &orderDetailsNo,
					CancelQuantity:                    &cancelQuantity,
					ConfigurationManagementTargetFlag: "a",
					AcceptanceScheduledDate:           now,
					OrderQuantity:                     "5",
					OrderUnitClassification:           &orderUnitClassification,
					OrderUnitPrice:                    "7",
					OrdersId:                          "8",
					QuotationsId:                      &quotationsID,
					Specifications:                    "MockedSpecifications",
					QuotationDetailsId:                &quotationDetailsID,
				}
				request := &pb.GetOrderDetailRequest{Id: "1"}
				orderDetailsRepository.EXPECT().GetOrderDetail(gomock.Any(), gomock.Any()).Return(orderDetail, nil)
				res, err := orderDetailClient.GetOrderDetail(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, "MockedRemarks", res.Remarks)
				assert.Equal(t, "MockedProductName", res.ProductName)
				assert.Equal(t, int32(2), res.OrderDetailsNo.Value)
				assert.Equal(t, float64(3), res.CancelQuantity.Value)
				assert.Equal(t, "a", res.ConfigurationManagementTargetFlag)
				assert.Equal(t, now.UTC(), res.AcceptanceScheduledDate.AsTime().UTC())
				assert.Equal(t, float64(5), res.OrderQuantity.Value)
				assert.Equal(t, int32(6), res.OrderUnitClassification.Value)
				assert.Equal(t, int32(7), res.OrderUnitPrice.Value)
				assert.Equal(t, int32(8), res.OrdersId.Value)
				assert.Equal(t, int32(10), res.QuotationsId.Value)
				assert.Equal(t, "MockedSpecifications", res.Specifications)
				assert.Equal(t, int32(11), res.QuotationDetailsId.Value)
			})

			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetOrderDetailRequest{Id: ""}
				res, err := orderDetailClient.GetOrderDetail(context.Background(), request)
				assert.NotNil(t, err)
				assert.Nil(t, res)
			})
		})
		t.Run("CreateOrderDetail", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)

				request := &pb.CreateOrderDetailRequest{
					//取消数量
					CancelQuantity: &wrappers.DoubleValue{Value: 2},
					//構成管理対象フラグ
					ConfigurationManagementTargetFlag: "m",
					//検収予定日
					AcceptanceScheduledDate: ifu.ToTimestamp(now),
					//発注行番
					OrderDetailsNo: &wrappers.Int32Value{Value: 4},
					//発注数量
					OrderQuantity: "5",
					//発注単位区分
					OrderUnitClassification: &wrappers.Int32Value{Value: 6},
					//発注単価
					OrderUnitPrice: "7",
					//string
					OrdersId: "8",
					//品名
					ProductName: "MockedProductName",
					//見積発注行番
					QuotationDetailsId: &wrappers.Int32Value{Value: 10},
					//発注見積iD
					QuotationsId: &wrappers.Int32Value{Value: 11},
					//摘要
					Remarks: "MockedRemarks",
					//仕様等
					Specifications: "MockedSpecifications",
				}
				orderDetailsRepository.EXPECT().CreateOrderDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderDetailClient.CreateOrderDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})

			t.Run("validation_error", func(t *testing.T) {
				request := &pb.CreateOrderDetailRequest{
					ConfigurationManagementTargetFlag: "",
					OrderQuantity:                     "",
					OrderUnitPrice:                    "",
					OrdersId:                          "",
				}
				res, err := orderDetailClient.CreateOrderDetail(context.Background(), request)
				assert.NotNil(t, err)
				assert.Nil(t, res)
			})
		})
		t.Run("UpdateOrderDetail", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)

				request := &pb.UpdateOrderDetailRequest{
					Id: "1",
					//取消数量
					CancelQuantity: &wrappers.DoubleValue{Value: 2},
					//構成管理対象フラグ
					ConfigurationManagementTargetFlag: "m",
					//検収予定日
					AcceptanceScheduledDate: ifu.ToTimestamp(now),
					//発注行番
					OrderDetailsNo: &wrappers.Int32Value{Value: 4},
					//発注数量
					OrderQuantity: "5",
					//発注単位区分
					OrderUnitClassification: &wrappers.Int32Value{Value: 6},
					//発注単価
					OrderUnitPrice: "7",
					//string
					OrdersId: "8",
					//品名
					ProductName: "MockedProductName",
					//見積発注行番
					QuotationDetailsId: &wrappers.Int32Value{Value: 10},
					//発注見積iD
					QuotationsId: &wrappers.Int32Value{Value: 11},
					//摘要
					Remarks: "MockedRemarks",
					//仕様等
					Specifications: "MockedSpecifications",
				}
				orderDetailsRepository.EXPECT().UpdateOrderDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderDetailClient.UpdateOrderDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateOrderDetailRequest{
					Id:                                "",
					ConfigurationManagementTargetFlag: "",
					OrderQuantity:                     "",
					OrderUnitPrice:                    "",
					OrdersId:                          "",
				}
				res, err := orderDetailClient.UpdateOrderDetail(context.Background(), request)
				assert.NotNil(t, err)
				assert.Nil(t, res)
			})
		})
		t.Run("DeleteOrderDetail", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteOrderDetailRequest{Id: "1"}
				orderDetailsRepository.EXPECT().DeleteOrderDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderDetailClient.DeleteOrderDetail(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})

			t.Run("validation_error_id_unspecified", func(t *testing.T) {
				request := &pb.DeleteOrderDetailRequest{Id: ""}
				res, err := orderDetailClient.DeleteOrderDetail(context.Background(), request)
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), errMessageInvalidArgument)
				assert.Nil(t, res)
			})
		})
		t.Run("DeleteOrderDetails", func(t *testing.T) {
			t.Run("by_ordersId", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteOrderDetailsRequest{OrdersId: "1"}
				orderDetailsRepository.EXPECT().DeleteOrderDetails(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderDetailClient.DeleteOrderDetails(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})

			t.Run("validation_error_orders_id_unspecified", func(t *testing.T) {
				request := &pb.DeleteOrderDetailsRequest{OrdersId: ""}
				res, err := orderDetailClient.DeleteOrderDetails(context.Background(), request)
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), errMessageInvalidArgument)
				assert.Nil(t, res)
			})
		})
		t.Run("CreateOrderDetails", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(2)

				OrderDetails := make([]*pb.CreateOrderDetailRequest, 0)
				detail1 := &pb.CreateOrderDetailRequest{
					//取消数量
					CancelQuantity: &wrappers.DoubleValue{Value: 3},
					//構成管理対象フラグ
					ConfigurationManagementTargetFlag: "m",
					//検収予定日
					AcceptanceScheduledDate: ifu.ToTimestamp(now),
					//発注行番
					OrderDetailsNo: &wrappers.Int32Value{Value: 5},
					//発注数量
					OrderQuantity: "6",
					//発注単位区分
					OrderUnitClassification: &wrappers.Int32Value{Value: 7},
					//発注単価
					OrderUnitPrice: "8",
					//string
					OrdersId: "9",
					//品名
					ProductName: "MockedProductName",
					//見積発注行番
					QuotationDetailsId: &wrappers.Int32Value{Value: 11},
					//発注見積iD
					QuotationsId: &wrappers.Int32Value{Value: 12},
					//摘要
					Remarks: "MockedRemarks",
					//仕様等
					Specifications: "MockedSpecifications",
				}
				detail2 := &pb.CreateOrderDetailRequest{
					//取消数量
					CancelQuantity: &wrappers.DoubleValue{Value: 4},
					//構成管理対象フラグ
					ConfigurationManagementTargetFlag: "m",
					//検収予定日
					AcceptanceScheduledDate: ifu.ToTimestamp(now),
					//発注行番
					OrderDetailsNo: &wrappers.Int32Value{Value: 6},
					//発注数量
					OrderQuantity: "7",
					//発注単位区分
					OrderUnitClassification: &wrappers.Int32Value{Value: 8},
					//発注単価
					OrderUnitPrice: "9",
					//string
					OrdersId: "10",
					//品名
					ProductName: "MockedProductName",
					//見積発注行番
					QuotationDetailsId: &wrappers.Int32Value{Value: 11},
					//発注見積iD
					QuotationsId: &wrappers.Int32Value{Value: 12},
					//摘要
					Remarks: "MockedRemarks",
					//仕様等
					Specifications: "MockedSpecifications",
				}
				OrderDetails = append(OrderDetails, detail1, detail2)

				request := &pb.CreateOrderDetailsRequest{
					OrderDetails: OrderDetails,
				}

				orderDetailsRepository.EXPECT().CreateOrderDetails(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderDetailClient.CreateOrderDetails(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})

			t.Run("validation_error", func(t *testing.T) {
				OrderDetails := make([]*pb.CreateOrderDetailRequest, 0)
				detail1 := &pb.CreateOrderDetailRequest{
					ConfigurationManagementTargetFlag: "",
					OrderQuantity:                     "",
					OrderUnitPrice:                    "",
					OrdersId:                          "",
				}
				detail2 := &pb.CreateOrderDetailRequest{
					ConfigurationManagementTargetFlag: "",
					OrderQuantity:                     "",
					OrderUnitPrice:                    "",
					OrdersId:                          "",
				}
				OrderDetails = append(OrderDetails, detail1, detail2)

				request := &pb.CreateOrderDetailsRequest{
					OrderDetails: OrderDetails,
				}

				res, err := orderDetailClient.CreateOrderDetails(context.Background(), request)
				assert.NotNil(t, err)
				assert.Nil(t, res)
			})
		})
	})
	t.Run("OrderType", func(t *testing.T) {
		t.Run("ListOrderType", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				now := time.Now()
				OrderType := &model.OrderType{
					ID:        "1",
					TypeValue: "10",
					Name:      "物品購入（ハードウェア）",
					Remarks:   "remarks",
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: now,
				}
				orderTypeRepository.EXPECT().ListOrderTypes(gomock.Any(), gomock.Any()).Return([]*model.OrderType{OrderType, OrderType}, nil)
				request := &pb.ListOrderTypesRequest{}
				res, err := orderTypeClient.ListOrderTypes(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, 2, len(res.OrderTypes))

				firstOrderType := res.OrderTypes[0]
				assert.Equal(t, "1", firstOrderType.Id)
				assert.Equal(t, "10", firstOrderType.TypeValue)
				assert.Equal(t, "物品購入（ハードウェア）", firstOrderType.Name)
				assert.Equal(t, now.UTC(), firstOrderType.CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), firstOrderType.UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), firstOrderType.DeletedAt.AsTime().UTC())

				secondOrderType := res.OrderTypes[1]
				assert.Equal(t, "1", secondOrderType.Id)
				assert.Equal(t, "10", secondOrderType.TypeValue)
				assert.Equal(t, "物品購入（ハードウェア）", secondOrderType.Name)
				assert.Equal(t, now.UTC(), secondOrderType.CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), secondOrderType.UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), secondOrderType.DeletedAt.AsTime().UTC())

			})
		})
		t.Run("GetOrderType", func(t *testing.T) {
			t.Run("validation_error_typeValue", func(t *testing.T) {
				request := &pb.GetOrderTypeRequest{}
				res, err := orderTypeClient.GetOrderType(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})

			t.Run("typeValue_specified", func(t *testing.T) {
				now := time.Now()
				OrderType := &model.OrderType{
					ID:        "1",
					TypeValue: "10",
					Name:      "物品購入（ハードウェア）",
					Remarks:   "remarks",
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: now,
				}
				orderTypeRepository.EXPECT().GetOrderType(gomock.Any(), gomock.Any()).Return(OrderType, nil)
				request := &pb.GetOrderTypeRequest{
					TypeValue: "10",
				}
				res, err := orderTypeClient.GetOrderType(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, "1", res.Id)
				assert.Equal(t, "10", res.TypeValue)
				assert.Equal(t, "物品購入（ハードウェア）", res.Name)
				assert.Equal(t, now.UTC(), res.CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.DeletedAt.AsTime().UTC())
			})
		})
	})
	t.Run("Quotation", func(t *testing.T) {

		// assert.Nil(t, "DONT KNOW WHY BUT QUOTATION TEST CASE FREEZES, BYPASS IT")
		pointer := "1"
		now := time.Now()
		quotation := &model.Quotation{
			ID:                         1,
			QuotationNo:                "quotationNo",
			VersionNumber:              &pointer,
			SuppliersID:                1,
			CompanyGroupClassification: "1",
			Subject:                    "subject1",
			SupplierQuotationNo:        "1",
			RequestOrganizationID:      "1",
			RequestDate:                now,
			RequestBy:                  "0000002",
			Remarks:                    "remarks",
			QuotationEffectiveDate:     now,
			QuotationInvalidDate:       now,
			JiraNo:                     "jiraNo1",
			OrderClassification:        1,
			OrdersID:                   &pointer,
			QuotationStatus:            &pointer,
		}

		t.Run("ListQuotation", func(t *testing.T) {

			quotations := []*model.Quotation{quotation, quotation}

			t.Run("all", func(t *testing.T) {
				request := &pb.ListQuotationsRequest{}
				quotationRepository.EXPECT().ListQuotations(gomock.Any(), gomock.Any()).Return(quotations, nil)
				res, err := quotationClient.ListQuotations(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.Quotations[0].GetId().Value)
				assert.Equal(t, "quotationNo", res.Quotations[0].GetQuotationNo())
				assert.Equal(t, int32(1), (*res.Quotations[0].GetVersionNumber()).Value)
				assert.Equal(t, int32(1), res.Quotations[0].GetSuppliersId().Value)
				assert.Equal(t, "1", res.Quotations[0].GetCompanyGroupClassification())
				assert.Equal(t, "subject1", res.Quotations[0].GetSubject())
				assert.Equal(t, "1", res.Quotations[0].GetSupplierQuotationNo())
				assert.Equal(t, "1", res.Quotations[0].GetRequestOrganizationId())
				assert.Equal(t, now.UTC(), res.Quotations[0].GetRequestDate().AsTime().UTC())
				assert.Equal(t, "0000002", res.Quotations[0].GetRequestBy())
				assert.Equal(t, "remarks", res.Quotations[0].GetRemarks())
				assert.Equal(t, now.UTC(), res.Quotations[0].GetQuotationEffectiveDate().AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Quotations[0].GetQuotationInvalidDate().AsTime().UTC())
				assert.Equal(t, "jiraNo1", res.Quotations[0].GetJiraNo())
				assert.Equal(t, int32(1), res.Quotations[0].GetOrderClassification().Value)
				assert.Equal(t, int32(1), (*res.Quotations[0].GetOrdersId()).Value)
				assert.Equal(t, int32(1), (*res.Quotations[0].GetQuotationStatus()).Value)

				assert.Equal(t, int32(1), res.Quotations[1].GetId().Value)
				assert.Equal(t, "quotationNo", res.Quotations[1].GetQuotationNo())
				assert.Equal(t, int32(1), (*res.Quotations[1].GetVersionNumber()).Value)
				assert.Equal(t, int32(1), res.Quotations[1].GetSuppliersId().Value)
				assert.Equal(t, "1", res.Quotations[1].GetCompanyGroupClassification())
				assert.Equal(t, "subject1", res.Quotations[1].GetSubject())
				assert.Equal(t, "1", res.Quotations[1].GetSupplierQuotationNo())
				assert.Equal(t, "1", res.Quotations[1].GetRequestOrganizationId())
				assert.Equal(t, now.UTC(), res.Quotations[1].GetRequestDate().AsTime().UTC())
				assert.Equal(t, "0000002", res.Quotations[1].GetRequestBy())
				assert.Equal(t, "remarks", res.Quotations[1].GetRemarks())
				assert.Equal(t, now.UTC(), res.Quotations[1].GetQuotationEffectiveDate().AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Quotations[1].GetQuotationInvalidDate().AsTime().UTC())
				assert.Equal(t, "jiraNo1", res.Quotations[1].GetJiraNo())
				assert.Equal(t, int32(1), res.Quotations[1].GetOrderClassification().Value)
				assert.Equal(t, int32(1), (*res.Quotations[1].GetOrdersId()).Value)
				assert.Equal(t, int32(1), (*res.Quotations[1].GetQuotationStatus()).Value)
			})
			t.Run("full_conditions", func(t *testing.T) {
				nowtime := make([]*timestamp.Timestamp, 0)
				nowtime = append(nowtime, ifu.ToTimestamp(time.Now()))

				request := &pb.ListQuotationsRequest{
					Ids:                    []int32{1, 2, 3, 4},
					SuppliersIds:           []int32{1, 2, 3, 4},
					Subject:                "subject",
					RequestOrganizationIds: []string{"1", "2", "3", "4"},
					RequestBys:             []string{"0000001", "0000002", "0000003", "0000004"},
					RequestDate:            nowtime,
					OrderClassification:    []int32{1, 2, 3, 4},
					OrderBy: &pb.ListQuotationsRequestOrderBy{
						Id:                  pb.SortEnum_ASC,
						SuppliersId:         pb.SortEnum_ASC,
						OrderClassification: pb.SortEnum_ASC,
						Subject:             pb.SortEnum_ASC,
						RequestDate:         pb.SortEnum_ASC,
					},
				}
				quotationRepository.EXPECT().ListQuotations(gomock.Any(), gomock.Any()).Return(quotations, nil)
				res, err := quotationClient.ListQuotations(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.Quotations[0].GetId().Value)
				assert.Equal(t, "quotationNo", res.Quotations[0].GetQuotationNo())
				assert.Equal(t, int32(1), (*res.Quotations[0].GetVersionNumber()).Value)
				assert.Equal(t, int32(1), res.Quotations[0].GetSuppliersId().Value)
				assert.Equal(t, "1", res.Quotations[0].GetCompanyGroupClassification())
				assert.Equal(t, "subject1", res.Quotations[0].GetSubject())
				assert.Equal(t, "1", res.Quotations[0].GetSupplierQuotationNo())
				assert.Equal(t, "1", res.Quotations[0].GetRequestOrganizationId())
				assert.Equal(t, now.UTC(), res.Quotations[0].GetRequestDate().AsTime().UTC())
				assert.Equal(t, "0000002", res.Quotations[0].GetRequestBy())
				assert.Equal(t, "remarks", res.Quotations[0].GetRemarks())
				assert.Equal(t, now.UTC(), res.Quotations[0].GetQuotationEffectiveDate().AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Quotations[0].GetQuotationInvalidDate().AsTime().UTC())
				assert.Equal(t, "jiraNo1", res.Quotations[0].GetJiraNo())
				assert.Equal(t, int32(1), res.Quotations[0].GetOrderClassification().Value)
				assert.Equal(t, int32(1), (*res.Quotations[0].GetOrdersId()).Value)
				assert.Equal(t, int32(1), (*res.Quotations[0].GetQuotationStatus()).Value)

				assert.Equal(t, int32(1), res.Quotations[1].GetId().Value)
				assert.Equal(t, "quotationNo", res.Quotations[1].GetQuotationNo())
				assert.Equal(t, int32(1), (*res.Quotations[1].GetVersionNumber()).Value)
				assert.Equal(t, int32(1), res.Quotations[1].GetSuppliersId().Value)
				assert.Equal(t, "1", res.Quotations[1].GetCompanyGroupClassification())
				assert.Equal(t, "subject1", res.Quotations[1].GetSubject())
				assert.Equal(t, "1", res.Quotations[1].GetSupplierQuotationNo())
				assert.Equal(t, "1", res.Quotations[1].GetRequestOrganizationId())
				assert.Equal(t, now.UTC(), res.Quotations[1].GetRequestDate().AsTime().UTC())
				assert.Equal(t, "0000002", res.Quotations[1].GetRequestBy())
				assert.Equal(t, "remarks", res.Quotations[1].GetRemarks())
				assert.Equal(t, now.UTC(), res.Quotations[1].GetQuotationEffectiveDate().AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Quotations[1].GetQuotationInvalidDate().AsTime().UTC())
				assert.Equal(t, "jiraNo1", res.Quotations[1].GetJiraNo())
				assert.Equal(t, int32(1), res.Quotations[1].GetOrderClassification().Value)
				assert.Equal(t, int32(1), (*res.Quotations[1].GetOrdersId()).Value)
				assert.Equal(t, int32(1), (*res.Quotations[1].GetQuotationStatus()).Value)
			})
		})
		t.Run("GetQuotation", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				request := &pb.GetQuotationRequest{
					Id: "1",
				}
				quotationRepository.EXPECT().GetQuotation(gomock.Any(), gomock.Any()).Return(quotation, nil)
				res, err := quotationClient.GetQuotation(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.GetId().Value)
				assert.Equal(t, "quotationNo", res.GetQuotationNo())
				assert.Equal(t, int32(1), (*res.GetVersionNumber()).Value)
				assert.Equal(t, int32(1), res.GetSuppliersId().Value)
				assert.Equal(t, "1", res.GetCompanyGroupClassification())
				assert.Equal(t, "subject1", res.GetSubject())
				assert.Equal(t, "1", res.GetSupplierQuotationNo())
				assert.Equal(t, "1", res.GetRequestOrganizationId())
				assert.Equal(t, now.UTC(), res.GetRequestDate().AsTime().UTC())
				assert.Equal(t, "0000002", res.GetRequestBy())
				assert.Equal(t, "remarks", res.GetRemarks())
				assert.Equal(t, now.UTC(), res.GetQuotationEffectiveDate().AsTime().UTC())
				assert.Equal(t, now.UTC(), res.GetQuotationInvalidDate().AsTime().UTC())
				assert.Equal(t, "jiraNo1", res.GetJiraNo())
				assert.Equal(t, int32(1), res.GetOrderClassification().Value)
				assert.Equal(t, int32(1), (*res.GetOrdersId()).Value)
				assert.Equal(t, int32(1), (*res.GetQuotationStatus()).Value)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetQuotationRequest{}
				res, err := quotationClient.GetQuotation(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("CreateQuotation", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				quotationRepository.EXPECT().CreateQuotation(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				request := &pb.CreateQuotationRequest{
					QuotationNo:                "quotationNo1",              //発注見積NO
					VersionNumber:              ifu.ToInt32Value(1),         //版番号
					SuppliersId:                "1",                         //仕入先ID
					CompanyGroupClassification: "1",                         //企業グループ区分
					Subject:                    "subject1",                  //件名
					SupplierQuotationNo:        "1",                         //仕入先見積№
					RequestOrganizationId:      "1",                         //依頼元組織ID
					RequestDate:                ifu.ToTimestamp(time.Now()), //依頼日
					RequestBy:                  "0000002",                   //依頼者ID
					Remarks:                    "remarks1",                  //摘要
					QuotationEffectiveDate:     ifu.ToTimestamp(time.Now()), //見積発行日
					QuotationInvalidDate:       ifu.ToTimestamp(time.Now()), //見積失効日
					JiraNo:                     "jiraNo1",                   //JIRAチケット番号
					OrderClassification:        "1",                         //発注区分
					OrdersId:                   ifu.ToInt32Value(1),         //発注ID
					QuotationStatus:            ifu.ToInt32Value(1),
				}

				sequence := model.HarpSysSequence{
					ID:           "1",
					Key:          "1",
					SequenceName: "quotations",
					Value:        "1",
				}
				harpSysSequenceRepository.EXPECT().GetHarpSysSequence(gomock.Any(), gomock.Any()).Return(&sequence, nil)

				res, err := quotationClient.CreateQuotation(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(1), res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				sequence := model.HarpSysSequence{
					ID:           "1",
					Key:          "1",
					SequenceName: "quotations",
					Value:        "1",
				}
				harpSysSequenceRepository.EXPECT().GetHarpSysSequence(gomock.Any(), gomock.Any()).Return(&sequence, nil)
				request := &pb.CreateQuotationRequest{
					QuotationNo:                "quotationNo1",              //発注見積NO
					VersionNumber:              ifu.ToInt32Value(1),         //版番号
					SuppliersId:                "1a",                        //仕入先ID
					CompanyGroupClassification: "1",                         //企業グループ区分
					Subject:                    strings.Repeat("1", 31),     //件名
					SupplierQuotationNo:        "1+",                        //仕入先見積№
					RequestOrganizationId:      "1a",                        //依頼元組織ID
					RequestDate:                ifu.ToTimestamp(time.Now()), //依頼日
					Remarks:                    strings.Repeat("1", 301),    //摘要
					QuotationEffectiveDate:     ifu.ToTimestamp(time.Now()), //見積発行日
					QuotationInvalidDate:       ifu.ToTimestamp(time.Now()), //見積失効日
					JiraNo:                     "jiraNo1+",                  //JIRAチケット番号
					OrdersId:                   ifu.ToInt32Value(1),         //発注ID
					QuotationStatus:            ifu.ToInt32Value(1),
				}

				res, err := quotationClient.CreateQuotation(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("UpdateQuotation", func(t *testing.T) {

			t.Run("success", func(t *testing.T) {
				request := &pb.UpdateQuotationRequest{
					Id:                         "1",
					QuotationNo:                "quotationNo1",              //発注見積NO
					VersionNumber:              ifu.ToInt32Value(1),         //版番号
					SuppliersId:                "1",                         //仕入先ID
					CompanyGroupClassification: "1",                         //企業グループ区分
					Subject:                    "subject1",                  //件名
					SupplierQuotationNo:        "1",                         //仕入先見積№
					RequestOrganizationId:      "1",                         //依頼元組織ID
					RequestDate:                ifu.ToTimestamp(time.Now()), //依頼日
					RequestBy:                  "0000002",                   //依頼者ID
					Remarks:                    "remarks1",                  //摘要
					QuotationEffectiveDate:     ifu.ToTimestamp(time.Now()), //見積発行日
					QuotationInvalidDate:       ifu.ToTimestamp(time.Now()), //見積失効日
					JiraNo:                     "jiraNo1",                   //JIRAチケット番号
					OrderClassification:        "1",                         //発注区分
					OrdersId:                   ifu.ToInt32Value(1),         //発注ID
					QuotationStatus:            ifu.ToInt32Value(1),
				}

				quotationRepository.EXPECT().UpdateQuotation(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				res, err := quotationClient.UpdateQuotation(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(1), res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateQuotationRequest{
					QuotationNo:                "quotationNo1",              //発注見積NO
					VersionNumber:              ifu.ToInt32Value(1),         //版番号
					SuppliersId:                "1a",                        //仕入先ID
					CompanyGroupClassification: "1",                         //企業グループ区分
					Subject:                    strings.Repeat("1", 31),     //件名
					SupplierQuotationNo:        "1+",                        //仕入先見積№
					RequestOrganizationId:      "1a",                        //依頼元組織ID
					RequestDate:                ifu.ToTimestamp(time.Now()), //依頼日
					Remarks:                    strings.Repeat("1", 301),    //摘要
					QuotationEffectiveDate:     ifu.ToTimestamp(time.Now()), //見積発行日
					QuotationInvalidDate:       ifu.ToTimestamp(time.Now()), //見積失効日
					JiraNo:                     "jiraNo1+",                  //JIRAチケット番号
					OrdersId:                   ifu.ToInt32Value(1),         //発注ID
					QuotationStatus:            ifu.ToInt32Value(1),
				}

				res, err := quotationClient.UpdateQuotation(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("DeleteQuotation", func(t *testing.T) {
			request := &pb.DeleteQuotationRequest{
				Id: "1",
			}
			t.Run("success", func(t *testing.T) {
				quotationDetailRepository.EXPECT().DeleteQuotationDetails(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				quotationHistoryRepository.EXPECT().DeleteQuotationHistories(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				quotationRepository.EXPECT().DeleteQuotation(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				res, err := quotationClient.DeleteQuotation(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(1), res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.DeleteQuotationRequest{}
				res, err := quotationClient.DeleteQuotation(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("CountQuotations", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				request := &pb.ListQuotationsRequest{}
				quotationRepository.EXPECT().CountQuotations(gomock.Any(), gomock.Any()).Return(int64(2), nil)
				res, err := quotationClient.CountQuotations(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(2), res.Count)
			})
			t.Run("full_conditions", func(t *testing.T) {
				request := &pb.ListQuotationsRequest{
					Ids:                    []int32{1, 2, 3},
					SuppliersIds:           []int32{1, 2, 3},
					Subject:                "subject",
					RequestOrganizationIds: []string{"1", "2", "3"},
					RequestBys:             []string{"0000001", "0000002", "0000003"},
					Limit:                  int64(10),
					Offset:                 int64(0),
				}
				quotationRepository.EXPECT().CountQuotations(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				res, err := quotationClient.CountQuotations(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(1), res.Count)
			})
		})
	})
	t.Run("QuotationItem", func(t *testing.T) {

		quotationItem := &model.QuotationItem{
			ID:             1,
			VersionNumber:  1,
			OrdersID:       1,
			OrderDetailsID: 1,
			ProductName:    "productName1",
			OrderQuantity:  1,
			OrderPrice:     1,
		}

		t.Run("ListQuotationItem", func(t *testing.T) {
			quotationItems := []*model.QuotationItem{quotationItem, quotationItem}
			t.Run("by_id", func(t *testing.T) {
				request := &pb.ListQuotationItemsRequest{
					QuotationDetailsId: 1,
				}
				quotationItemsRepository.EXPECT().ListQuotationItems(gomock.Any(), gomock.Any()).Return(quotationItems, nil)
				res, err := quotationItemClient.ListQuotationItems(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.QuotationItems[0].GetId().Value)
				assert.Equal(t, int32(1), res.QuotationItems[0].GetOrderDetailsId().Value)
				assert.Equal(t, int32(1), res.QuotationItems[0].GetVersionNumber().Value)
				assert.Equal(t, "productName1", res.QuotationItems[0].GetProductName())
				assert.Equal(t, int32(1), res.QuotationItems[0].GetOrderQuantity().Value)
				assert.Equal(t, int32(1), res.QuotationItems[0].GetOrderPrice().Value)

				assert.Equal(t, int32(1), res.QuotationItems[1].GetId().Value)
				assert.Equal(t, int32(1), res.QuotationItems[1].GetOrderDetailsId().Value)
				assert.Equal(t, int32(1), res.QuotationItems[1].GetVersionNumber().Value)
				assert.Equal(t, "productName1", res.QuotationItems[1].GetProductName())
				assert.Equal(t, int32(1), res.QuotationItems[1].GetOrderQuantity().Value)
				assert.Equal(t, int32(1), res.QuotationItems[1].GetOrderPrice().Value)

			})
		})
		t.Run("GetQuotationItem", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				request := &pb.GetQuotationItemRequest{
					Id: "1",
				}

				quotationItemsRepository.EXPECT().GetQuotationItem(gomock.Any(), gomock.Any()).Return(quotationItem, nil)
				res, err := quotationItemClient.GetQuotationItem(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.GetId().Value)
				assert.Equal(t, int32(1), res.GetOrderDetailsId().Value)
				assert.Equal(t, int32(1), res.GetVersionNumber().Value)
				assert.Equal(t, "productName1", res.GetProductName())
				assert.Equal(t, int32(1), res.GetOrderQuantity().Value)
				assert.Equal(t, int32(1), res.GetOrderPrice().Value)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetQuotationItemRequest{}
				res, err := quotationItemClient.GetQuotationItem(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
	})
	t.Run("HarpSysSequence", func(t *testing.T) {

		t.Run("GetHarpSysSequence", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetHarpSysSequenceRequest{}
				res, err := harpSysSequenceClient.GetHarpSysSequence(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("by_key_and_sequenceName", func(t *testing.T) {
				harpSysSequence := &model.HarpSysSequence{
					ID:           "1",
					Key:          "mockKey",
					SequenceName: "mockSequenceName",
					Value:        "123",
				}
				request := &pb.GetHarpSysSequenceRequest{SequenceName: "mockSequenceName", Key: "mockKey"}
				harpSysSequenceRepository.EXPECT().GetHarpSysSequence(gomock.Any(), gomock.Any()).Return(harpSysSequence, nil)
				res, err := harpSysSequenceClient.GetHarpSysSequence(context.Background(), request)

				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, "mockKey", res.Key)
				assert.Equal(t, "mockSequenceName", res.SequenceName)
				assert.Equal(t, int32(123), res.Value.Value)
			})
		})
	})
	t.Run("Orders", func(t *testing.T) {

		// assert.Nil(t, "DONT KNOW WHY BUT ORDERS TEST CASE FREEZES, BYPASS IT")
		t.Run("ListOrders", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				derivationSourceOrderID := "1"
				quotationsID := "1"
				orderApprovalStaffsID := "1"
				now := time.Now()

				Order := &model.Order{
					ID:                         int64(1),
					OrderNo:                    "20H0001",
					SuppliersID:                "1",
					CompanyGroupType:           "A",
					Subject:                    "subject",
					RequestOrganizationID:      "1",
					RequestDate:                now,
					RequestBy:                  "0000002",
					ApprovalFile:               "file://file.txt",
					DerivationSourceOrderID:    &derivationSourceOrderID,
					Remarks:                    "remarks",
					SuperiorApprovalDate:       now,
					PurchasingDeptApprovalDate: now,
					OrderIssueDate:             now,
					FinalAcceptanceDate:        now,
					AcceptanceCompletedDate:    now,
					CancelDate:                 now,
					OrderCaseCd:                "10",
					OrderStatus:                "20",
					JiraNo:                     "CTC-0101",
					QuotationsID:               &quotationsID,
					OrderApprovalStaffsID:      orderApprovalStaffsID,
					ProjectsID:                 "0000001",
					ProjectCostID:              "0000002",
					CostTypesID:                "0000003",
				}
				orderRepository.EXPECT().ListOrders(gomock.Any(), gomock.Any()).Return([]*model.Order{Order, Order}, nil)

				request := &pb.ListOrdersRequest{}
				res, err := orderClient.ListOrders(context.Background(), request)

				assert.Nil(t, err)
				assert.Equal(t, 2, len(res.Orders))
				firstOrder := res.Orders[0]
				assert.Equal(t, int32(1), firstOrder.Id.Value)
				assert.Equal(t, "20H0001", firstOrder.OrderNo)
				assert.Equal(t, int32(1), firstOrder.SuppliersId.Value)
				assert.Equal(t, "1", firstOrder.RequestOrganizationId)
				assert.Equal(t, "A", firstOrder.CompanyGroupType)
				assert.Equal(t, "subject", firstOrder.Subject)
				assert.Equal(t, now.UTC(), firstOrder.RequestDate.AsTime().UTC())
				assert.Equal(t, "0000002", firstOrder.RequestBy)
				assert.Equal(t, "file://file.txt", firstOrder.ApprovalFile)
				assert.Equal(t, int32(1), firstOrder.DerivationSourceOrderId.Value)
				assert.Equal(t, now.UTC(), firstOrder.SuperiorApprovalDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), firstOrder.PurchasingDeptApprovalDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), firstOrder.OrderIssueDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), firstOrder.FinalAcceptanceDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), firstOrder.AcceptanceCompletedDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), firstOrder.CancelDate.AsTime().UTC())
				assert.Equal(t, int32(10), firstOrder.OrderCaseCd.Value)
				assert.Equal(t, int32(20), firstOrder.OrderStatus.Value)
				assert.Equal(t, "CTC-0101", firstOrder.JiraNo)
				assert.Equal(t, int32(1), firstOrder.QuotationsId.Value)
				assert.Equal(t, "1", firstOrder.OrderApprovalStaffsId)
				assert.Equal(t, "0000001", firstOrder.ProjectsId)
				assert.Equal(t, "0000002", firstOrder.ProjectCostId)
				assert.Equal(t, "0000003", firstOrder.CostTypesId)

				secondOrder := res.Orders[1]
				assert.Equal(t, int32(1), secondOrder.Id.Value)
				assert.Equal(t, "20H0001", secondOrder.OrderNo)
				assert.Equal(t, int32(1), secondOrder.SuppliersId.Value)
				assert.Equal(t, "1", secondOrder.RequestOrganizationId)
				assert.Equal(t, "A", secondOrder.CompanyGroupType)
				assert.Equal(t, "subject", secondOrder.Subject)
				assert.Equal(t, now.UTC(), secondOrder.RequestDate.AsTime().UTC())
				assert.Equal(t, "0000002", secondOrder.RequestBy)
				assert.Equal(t, "file://file.txt", secondOrder.ApprovalFile)
				assert.Equal(t, int32(1), secondOrder.DerivationSourceOrderId.Value)
				assert.Equal(t, now.UTC(), secondOrder.SuperiorApprovalDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), secondOrder.PurchasingDeptApprovalDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), secondOrder.OrderIssueDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), secondOrder.FinalAcceptanceDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), secondOrder.AcceptanceCompletedDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), secondOrder.CancelDate.AsTime().UTC())
				assert.Equal(t, int32(10), secondOrder.OrderCaseCd.Value)
				assert.Equal(t, int32(20), secondOrder.OrderStatus.Value)
				assert.Equal(t, "CTC-0101", secondOrder.JiraNo)
				assert.Equal(t, int32(1), secondOrder.QuotationsId.Value)
				assert.Equal(t, "1", secondOrder.OrderApprovalStaffsId)
				assert.Equal(t, "0000001", secondOrder.ProjectsId)
				assert.Equal(t, "0000002", secondOrder.ProjectCostId)
				assert.Equal(t, "0000003", secondOrder.CostTypesId)
			})
		})
		t.Run("GetOrder", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetOrderRequest{}
				res, err := orderClient.GetOrder(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})

			t.Run("by_id", func(t *testing.T) {
				derivationSourceOrderID := "1"
				quotationsID := "1"
				orderApprovalStaffsID := "1"
				now := time.Now()

				Order := &model.Order{
					ID:                         int64(1),
					OrderNo:                    "20H0001",
					SuppliersID:                "1",
					CompanyGroupType:           "A",
					Subject:                    "subject",
					RequestOrganizationID:      "1",
					RequestDate:                now,
					RequestBy:                  "0000002",
					ApprovalFile:               "file://file.txt",
					DerivationSourceOrderID:    &derivationSourceOrderID,
					Remarks:                    "remarks",
					SuperiorApprovalDate:       now,
					PurchasingDeptApprovalDate: now,
					OrderIssueDate:             now,
					FinalAcceptanceDate:        now,
					AcceptanceCompletedDate:    now,
					CancelDate:                 now,
					OrderCaseCd:                "10",
					OrderStatus:                "20",
					JiraNo:                     "CTC-0101",
					QuotationsID:               &quotationsID,
					OrderApprovalStaffsID:      orderApprovalStaffsID,
					ProjectsID:                 "0000001",
					ProjectCostID:              "0000002",
					CostTypesID:                "0000003",
				}
				orderRepository.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Return(Order, nil)

				request := &pb.GetOrderRequest{
					Id: "1",
				}
				res, err := orderClient.GetOrder(context.Background(), request)

				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, "20H0001", res.OrderNo)
				assert.Equal(t, int32(1), res.SuppliersId.Value)
				assert.Equal(t, "1", res.RequestOrganizationId)
				assert.Equal(t, "A", res.CompanyGroupType)
				assert.Equal(t, "subject", res.Subject)
				assert.Equal(t, now.UTC(), res.RequestDate.AsTime().UTC())
				assert.Equal(t, "0000002", res.RequestBy)
				assert.Equal(t, "file://file.txt", res.ApprovalFile)
				assert.Equal(t, int32(1), res.DerivationSourceOrderId.Value)
				assert.Equal(t, now.UTC(), res.SuperiorApprovalDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.PurchasingDeptApprovalDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.OrderIssueDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.FinalAcceptanceDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.AcceptanceCompletedDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.CancelDate.AsTime().UTC())
				assert.Equal(t, int32(10), res.OrderCaseCd.Value)
				assert.Equal(t, int32(20), res.OrderStatus.Value)
				assert.Equal(t, "CTC-0101", res.JiraNo)
				assert.Equal(t, int32(1), res.QuotationsId.Value)
				assert.Equal(t, "1", res.OrderApprovalStaffsId)
				assert.Equal(t, "0000001", res.ProjectsId)
				assert.Equal(t, "0000002", res.ProjectCostId)
				assert.Equal(t, "0000003", res.CostTypesId)
			})
		})
		t.Run("CreateOrders", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.CreateOrderRequest{
					OrderNo:                    strings.Repeat("1", 9),
					SuppliersId:                "",
					CompanyGroupType:           "",
					Subject:                    strings.Repeat("あ", 31),
					RequestOrganizationId:      "",
					RequestDate:                nil,
					RequestBy:                  "",
					ApprovalFile:               strings.Repeat("F", 201),
					DerivationSourceOrderId:    ifu.ToInt32Value("10"),
					Remarks:                    strings.Repeat("R", 301),
					SuperiorApprovalDate:       nil,
					PurchasingDeptApprovalDate: nil,
					OrderIssueDate:             nil,
					FinalAcceptanceDate:        nil,
					AcceptanceCompletedDate:    nil,
					CancelDate:                 nil,
					OrderCaseCd:                "",
					OrderStatus:                "",
					JiraNo:                     strings.Repeat("J", 21),
					QuotationsId:               nil,
					OrderApprovalStaffsId:      "",
				}
				sequence := model.HarpSysSequence{
					ID:           "1",
					Key:          "1",
					SequenceName: "orders",
					Value:        "1",
				}
				harpSysSequenceRepository.EXPECT().GetHarpSysSequence(gomock.Any(), gomock.Any()).Return(&sequence, nil)
				res, err := orderClient.CreateOrder(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})

			t.Run("success", func(t *testing.T) {
				now := time.Now()
				request := &pb.CreateOrderRequest{
					OrderNo:                    strings.Repeat("1", 8),
					SuppliersId:                "10",
					CompanyGroupType:           "A",
					Subject:                    strings.Repeat("あ", 30),
					RequestOrganizationId:      "10",
					RequestDate:                ifu.ToTimestamp(now),
					RequestBy:                  "0000002",
					ApprovalFile:               "file://file.txt",
					DerivationSourceOrderId:    ifu.ToInt32Value("10"),
					Remarks:                    strings.Repeat("R", 300),
					SuperiorApprovalDate:       ifu.ToTimestamp(now),
					PurchasingDeptApprovalDate: ifu.ToTimestamp(now),
					OrderIssueDate:             ifu.ToTimestamp(now),
					FinalAcceptanceDate:        ifu.ToTimestamp(now),
					AcceptanceCompletedDate:    ifu.ToTimestamp(now),
					CancelDate:                 ifu.ToTimestamp(now),
					OrderCaseCd:                "10",
					OrderStatus:                "10",
					JiraNo:                     strings.Repeat("J", 20),
					QuotationsId:               ifu.ToInt32Value("10"),
					OrderApprovalStaffsId:      "10",
				}

				pointer := "10"
				Order := &model.Order{
					OrderNo:                    strings.Repeat("1", 8),
					SuppliersID:                "10",
					CompanyGroupType:           "A",
					Subject:                    "subject",
					RequestOrganizationID:      "1",
					RequestDate:                now,
					RequestBy:                  "0000002",
					ApprovalFile:               "file://file.txt",
					DerivationSourceOrderID:    &pointer,
					Remarks:                    strings.Repeat("R", 300),
					SuperiorApprovalDate:       now,
					PurchasingDeptApprovalDate: now,
					OrderIssueDate:             now,
					FinalAcceptanceDate:        now,
					AcceptanceCompletedDate:    now,
					CancelDate:                 now,
					OrderCaseCd:                "10",
					OrderStatus:                "10",
					JiraNo:                     strings.Repeat("J", 20),
					QuotationsID:               &pointer,
					OrderApprovalStaffsID:      "10",
				}

				returning := repository.CreateOrderReturning{AffectedRows: 1, Order: Order}

				orderRepository.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(returning, nil)

				sequence := model.HarpSysSequence{
					ID:           "1",
					Key:          "1",
					SequenceName: "orders",
					Value:        "1",
				}
				harpSysSequenceRepository.EXPECT().GetHarpSysSequence(gomock.Any(), gomock.Any()).Return(&sequence, nil)

				res, err := orderClient.CreateOrder(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(1), res.AffectedRows)

				assert.Equal(t, int32(0), res.Order.Id.Value)
				assert.Equal(t, "11111111", res.Order.OrderNo)
				assert.Equal(t, int32(10), res.Order.SuppliersId.Value)
				assert.Equal(t, "1", res.Order.RequestOrganizationId)
				assert.Equal(t, "A", res.Order.CompanyGroupType)
				assert.Equal(t, "subject", res.Order.Subject)
				assert.Equal(t, now.UTC(), res.Order.RequestDate.AsTime().UTC())
				assert.Equal(t, "0000002", res.Order.RequestBy)
				assert.Equal(t, "file://file.txt", res.Order.ApprovalFile)
				assert.Equal(t, int32(10), res.Order.DerivationSourceOrderId.Value)
				assert.Equal(t, now.UTC(), res.Order.SuperiorApprovalDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Order.PurchasingDeptApprovalDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Order.OrderIssueDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Order.FinalAcceptanceDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Order.AcceptanceCompletedDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.Order.CancelDate.AsTime().UTC())
				assert.Equal(t, int32(10), res.Order.OrderCaseCd.Value)
				assert.Equal(t, int32(10), res.Order.OrderStatus.Value)
				assert.Equal(t, "JJJJJJJJJJJJJJJJJJJJ", res.Order.JiraNo)
				assert.Equal(t, int32(10), res.Order.QuotationsId.Value)
				assert.Equal(t, "10", res.Order.OrderApprovalStaffsId)
			})
		})
		t.Run("UpdateOrders", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateOrderRequest{
					Id:                         "",
					OrderNo:                    strings.Repeat("1", 9),
					SuppliersId:                "",
					Subject:                    strings.Repeat("あ", 31),
					RequestOrganizationId:      "",
					RequestDate:                nil,
					RequestBy:                  "",
					ApprovalFile:               strings.Repeat("F", 201),
					DerivationSourceOrderId:    ifu.ToInt32Value("10"),
					Remarks:                    strings.Repeat("R", 301),
					SuperiorApprovalDate:       nil,
					PurchasingDeptApprovalDate: nil,
					OrderIssueDate:             nil,
					FinalAcceptanceDate:        nil,
					AcceptanceCompletedDate:    nil,
					CancelDate:                 nil,
					OrderCaseCd:                "",
					OrderStatus:                "",
					JiraNo:                     strings.Repeat("J", 21),
					QuotationsId:               nil,
					OrderApprovalStaffsId:      "",
				}
				res, err := orderClient.UpdateOrder(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				now := time.Now()
				request := &pb.UpdateOrderRequest{
					Id:                         "10",
					OrderNo:                    strings.Repeat("1", 8),
					SuppliersId:                "10",
					Subject:                    strings.Repeat("あ", 30),
					RequestOrganizationId:      "10",
					RequestDate:                ifu.ToTimestamp(now),
					RequestBy:                  "0000002",
					ApprovalFile:               strings.Repeat("F", 200),
					DerivationSourceOrderId:    ifu.ToInt32Value("10"),
					Remarks:                    strings.Repeat("R", 300),
					SuperiorApprovalDate:       ifu.ToTimestamp(now),
					PurchasingDeptApprovalDate: ifu.ToTimestamp(now),
					OrderIssueDate:             ifu.ToTimestamp(now),
					FinalAcceptanceDate:        ifu.ToTimestamp(now),
					AcceptanceCompletedDate:    ifu.ToTimestamp(now),
					CancelDate:                 ifu.ToTimestamp(now),
					OrderCaseCd:                "10",
					OrderStatus:                "10",
					JiraNo:                     strings.Repeat("J", 20),
					QuotationsId:               ifu.ToInt32Value("10"),
					OrderApprovalStaffsId:      "10",
				}
				orderRepository.EXPECT().UpdateOrder(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				res, err := orderClient.UpdateOrder(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, int64(1), res.AffectedRows)
			})
		})
		t.Run("UpdateOrderProjectCostInfo", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateOrderProjectCostInfoRequest{
					Id:            "",
					ProjectCostId: "",
				}
				res, err := orderClient.UpdateOrderProjectCostInfo(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				request := &pb.UpdateOrderProjectCostInfoRequest{
					Id:            "1",
					ProjectsId:    "0000001",
					ProjectCostId: "0000002",
					CostTypesId:   "0000003",
				}
				orderRepository.EXPECT().UpdateOrderProjectCostInfo(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				res, err := orderClient.UpdateOrderProjectCostInfo(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(1), res.AffectedRows)
			})
		})
		t.Run("DeleteOrders", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.DeleteOrderRequest{}
				res, err := orderClient.DeleteOrder(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				request := &pb.DeleteOrderRequest{
					Id: "1",
				}
				derivationSourceOrderID := "1"
				quotationsID := "1"
				orderApprovalStaffsID := "1"
				now := time.Now()
				Order := &model.Order{
					ID:                         int64(1),
					OrderNo:                    "20H0001",
					SuppliersID:                "1",
					CompanyGroupType:           "A",
					Subject:                    "subject",
					RequestOrganizationID:      "1",
					RequestDate:                now,
					RequestBy:                  "0000002",
					ApprovalFile:               "file://file.txt",
					DerivationSourceOrderID:    &derivationSourceOrderID,
					Remarks:                    "remarks",
					SuperiorApprovalDate:       now,
					PurchasingDeptApprovalDate: now,
					OrderIssueDate:             now,
					FinalAcceptanceDate:        now,
					AcceptanceCompletedDate:    now,
					CancelDate:                 now,
					OrderCaseCd:                "10",
					OrderStatus:                "20",
					JiraNo:                     "CTC-0101",
					QuotationsID:               &quotationsID,
					OrderApprovalStaffsID:      orderApprovalStaffsID,
				}
				orderRepository.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Return(Order, nil)
				orderDetailsRepository.EXPECT().DeleteOrderDetails(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				quotationRepository.EXPECT().UpdateQuotationForDeleteOrder(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				quotationDetailRepository.EXPECT().UpdateQuotationDetailForDeleteOrder(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				orderRepository.EXPECT().DeleteOrder(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				res, err := orderClient.DeleteOrder(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, int64(1), res.AffectedRows)
			})
		})
		t.Run("CountOrder", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				request := &pb.ListOrdersRequest{}
				orderRepository.EXPECT().CountOrders(gomock.Any(), gomock.Any()).Return(int64(2), nil)
				res, err := orderClient.CountOrders(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(2), res.Count)
			})
			t.Run("full_conditions", func(t *testing.T) {
				request := &pb.ListOrdersRequest{
					Ids:                    []int32{1, 2, 3},
					SuppliersIds:           []int32{1, 2, 3},
					Subject:                "subject",
					RequestOrganizationIds: []string{"1", "2", "3"},
					RequestBys:             []string{"0000001", "0000002", "0000003"},
					OrderCaseCds:           []int32{1, 2, 3},
					OrderStatuses:          []int32{1, 2, 3},
					ProjectsIds:            []string{"0000001", "0000002", "0000003"},
					ProjectCostIds:         []string{"0000001", "0000002", "0000003"},
					Limit:                  int64(10),
					Offset:                 int64(0),
					OrderBy: &pb.ListOrdersRequestOrderBy{
						Id:                  pb.SortEnum_ASC,
						SuppliersId:         pb.SortEnum_ASC,
						OrderStatus:         pb.SortEnum_ASC,
						OrderCaseCd:         pb.SortEnum_ASC,
						Subject:             pb.SortEnum_ASC,
						RequestDate:         pb.SortEnum_ASC,
						FinalAcceptanceDate: pb.SortEnum_ASC,
					},
				}
				orderRepository.EXPECT().CountOrders(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				res, err := orderClient.CountOrders(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(1), res.Count)
			})
		})
		t.Run("SumNearestTwoMonthsAmount", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				nearestTwoMonthsAmount := []*model.NearestTwoMonthsAmount{{
					Date:     "202102",
					MonthSum: int64(492100),
				}, {
					Date:     "202103",
					MonthSum: int64(774500),
				},
				}
				request := &pb.SumNearestTwoMonthsAmountRequest{}

				orderRepository.EXPECT().SumNearestTwoMonthsAmount(gomock.Any(), gomock.Any()).Return(nearestTwoMonthsAmount, nil)
				res, err := orderClient.SumNearestTwoMonthsAmount(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int(2), len(res.NearestTwoMonthsAmount))
				//nearestTwoMonthsAmount[0]
				assert.Equal(t, "202102", res.NearestTwoMonthsAmount[0].Date)
				assert.Equal(t, int64(492100), res.NearestTwoMonthsAmount[0].MonthSum)
				//nearestTwoMonthsAmount[1]
				assert.Equal(t, "202103", res.NearestTwoMonthsAmount[1].Date)
				assert.Equal(t, int64(774500), res.NearestTwoMonthsAmount[1].MonthSum)
			})
		})
	})
	t.Run("QuotationDetail", func(t *testing.T) {
		VersionNumber := "1"
		OrderDetailsID := "1"
		commonQuotationDetail := &model.QuotationDetail{
			ID:                 "1",
			QuotationsID:       "1",
			VersionNumber:      &VersionNumber,
			OrderDetailsID:     &OrderDetailsID,
			ProductName:        "ProductName",
			Specifications:     "specifications",
			OrderQuantity:      "100",
			OrderPrice:         "1000",
			UnitClassification: "10",
		}
		t.Run("ListQuotationDetail", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				quotationDetails := []*model.QuotationDetail{commonQuotationDetail}
				request := &pb.ListQuotationDetailsRequest{}
				quotationDetailRepository.EXPECT().ListQuotationDetails(gomock.Any(), gomock.Any()).Return(quotationDetails, nil)
				res, err := quotationDetailClient.ListQuotationDetails(context.Background(), request)
				assert.Nil(t, err)

				quotationDetailResponse :=
					[]*pb.QuotationDetailResponse{
						{
							Id:                 ifu.ToInt32Value(commonQuotationDetail.ID),
							QuotationsId:       ifu.ToInt32Value(commonQuotationDetail.QuotationsID),
							VersionNumber:      ifu.ToInt32Value(commonQuotationDetail.VersionNumber),
							OrderDetailsId:     ifu.ToInt32Value(commonQuotationDetail.OrderDetailsID),
							ProductName:        commonQuotationDetail.ProductName,
							Specifications:     commonQuotationDetail.Specifications,
							OrderQuantity:      ifu.ToDoubleValue(commonQuotationDetail.OrderQuantity),
							OrderPrice:         ifu.ToInt32Value(commonQuotationDetail.OrderPrice),
							UnitClassification: ifu.ToInt32Value(commonQuotationDetail.UnitClassification),
						},
					}
				assert.Equal(t, quotationDetailResponse, res.QuotationDetails)
			})
			t.Run("full_conditions", func(t *testing.T) {
				quotationDetails := []*model.QuotationDetail{commonQuotationDetail, commonQuotationDetail}
				request := &pb.ListQuotationDetailsRequest{
					Ids:          []string{"1", "2"},
					QuotationsId: "1",
					Limit:        int64(10),
					Offset:       int64(0)}
				quotationDetailRepository.EXPECT().ListQuotationDetails(gomock.Any(), gomock.Any()).Return(quotationDetails, nil)
				res, err := quotationDetailClient.ListQuotationDetails(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, int32(1), res.QuotationDetails[0].GetId().Value)
				assert.Equal(t, int32(1), res.QuotationDetails[0].GetQuotationsId().Value)
				assert.Equal(t, int32(1), res.QuotationDetails[0].GetVersionNumber().Value)
				assert.Equal(t, int32(1), res.QuotationDetails[0].GetOrderDetailsId().Value)
				assert.Equal(t, "ProductName", res.QuotationDetails[0].GetProductName())
				assert.Equal(t, "specifications", res.QuotationDetails[0].GetSpecifications())
				assert.Equal(t, float64(100), res.QuotationDetails[0].GetOrderQuantity().Value)
				assert.Equal(t, int32(1000), res.QuotationDetails[0].GetOrderPrice().Value)
				assert.Equal(t, int32(10), res.QuotationDetails[0].GetUnitClassification().Value)

				assert.Equal(t, int32(1), res.QuotationDetails[1].GetId().Value)
				assert.Equal(t, int32(1), res.QuotationDetails[1].GetQuotationsId().Value)
				assert.Equal(t, int32(1), res.QuotationDetails[1].GetVersionNumber().Value)
				assert.Equal(t, int32(1), res.QuotationDetails[1].GetOrderDetailsId().Value)
				assert.Equal(t, "ProductName", res.QuotationDetails[1].GetProductName())
				assert.Equal(t, "specifications", res.QuotationDetails[1].GetSpecifications())
				assert.Equal(t, float64(100), res.QuotationDetails[1].GetOrderQuantity().Value)
				assert.Equal(t, int32(1000), res.QuotationDetails[1].GetOrderPrice().Value)
				assert.Equal(t, int32(10), res.QuotationDetails[1].GetUnitClassification().Value)
			})
		})
		t.Run("GetQuotationDetail", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				request := &pb.GetQuotationDetailRequest{Id: commonQuotationDetail.ID}
				quotationDetailRepository.EXPECT().GetQuotationDetail(gomock.Any(), gomock.Any()).Return(commonQuotationDetail, nil)
				res, err := quotationDetailClient.GetQuotationDetail(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, int32(1), res.QuotationsId.Value)
				assert.Equal(t, int32(1), res.VersionNumber.Value)
				assert.Equal(t, int32(1), res.OrderDetailsId.Value)
				assert.Equal(t, "ProductName", res.ProductName)
				assert.Equal(t, "specifications", res.Specifications)
				assert.Equal(t, float64(100), res.OrderQuantity.Value)
				assert.Equal(t, int32(1000), res.OrderPrice.Value)
				assert.Equal(t, int32(10), res.UnitClassification.Value)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetQuotationDetailRequest{}
				res, err := quotationDetailClient.GetQuotationDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("CreateQuotationDetail", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)

				request := &pb.CreateQuotationDetailRequest{
					QuotationsId:       "2",
					VersionNumber:      &wrappers.Int32Value{Value: 1},
					OrderDetailsId:     &wrappers.Int32Value{Value: 2},
					ProductName:        "createProductName",
					Specifications:     "createSpecifications",
					OrderQuantity:      "200",
					OrderPrice:         "2000",
					UnitClassification: "11",
				}

				quotationDetailRepository.EXPECT().CreateQuotationDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := quotationDetailClient.CreateQuotationDetail(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.CreateQuotationDetailRequest{
					QuotationsId:       "",
					ProductName:        "",
					Specifications:     strings.Repeat("S", 301),
					OrderQuantity:      "",
					OrderPrice:         "",
					UnitClassification: "",
				}
				res, err := quotationDetailClient.CreateQuotationDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("UpdateQuotationDetail", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)

				request := &pb.UpdateQuotationDetailRequest{
					Id:                 "2",
					QuotationsId:       "2",
					VersionNumber:      &wrappers.Int32Value{Value: 2},
					OrderDetailsId:     &wrappers.Int32Value{Value: 2},
					ProductName:        "updateProductName",
					Specifications:     "updateSpecifications",
					OrderQuantity:      "222",
					OrderPrice:         "2222",
					UnitClassification: "11",
				}

				quotationDetailRepository.EXPECT().UpdateQuotationDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := quotationDetailClient.UpdateQuotationDetail(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateQuotationDetailRequest{
					Id:                 "",
					QuotationsId:       "",
					ProductName:        "",
					Specifications:     strings.Repeat("S", 301),
					OrderQuantity:      "",
					OrderPrice:         "",
					UnitClassification: "",
				}
				res, err := quotationDetailClient.UpdateQuotationDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("DeleteQuotationDetail", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteQuotationDetailRequest{Id: "1"}
				quotationDetailRepository.EXPECT().DeleteQuotationDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := quotationDetailClient.DeleteQuotationDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.DeleteQuotationDetailRequest{}
				res, err := quotationDetailClient.DeleteQuotationDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("DeleteQuotationDetails", func(t *testing.T) {
			t.Run("by_quotationsId", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteQuotationDetailsRequest{QuotationsId: "1"}
				quotationDetailRepository.EXPECT().DeleteQuotationDetails(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := quotationDetailClient.DeleteQuotationDetails(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.DeleteQuotationDetailsRequest{}
				res, err := quotationDetailClient.DeleteQuotationDetails(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("CreateQuotationDetails", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(2)

				QuotationDetails := make([]*pb.CreateQuotationDetailRequest, 0)
				detail1 := &pb.CreateQuotationDetailRequest{
					QuotationsId:       "3",
					VersionNumber:      &wrappers.Int32Value{Value: 3},
					OrderDetailsId:     &wrappers.Int32Value{Value: 2},
					ProductName:        "createProductName",
					Specifications:     "createSpecifications",
					OrderQuantity:      "333",
					OrderPrice:         "333",
					UnitClassification: "22",
				}
				detail2 := &pb.CreateQuotationDetailRequest{
					QuotationsId:       "4",
					VersionNumber:      &wrappers.Int32Value{Value: 4},
					OrderDetailsId:     &wrappers.Int32Value{Value: 3},
					ProductName:        "createProductName",
					Specifications:     "createSpecifications",
					OrderQuantity:      "444",
					OrderPrice:         "444",
					UnitClassification: "33",
				}
				QuotationDetails = append(QuotationDetails, detail1, detail2)

				request := &pb.CreateQuotationDetailsRequest{
					QuotationDetails: QuotationDetails,
				}

				quotationDetailRepository.EXPECT().CreateQuotationDetails(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := quotationDetailClient.CreateQuotationDetails(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {

				QuotationDetails := make([]*pb.CreateQuotationDetailRequest, 0)
				detail1 := &pb.CreateQuotationDetailRequest{
					QuotationsId:       "",
					ProductName:        "",
					Specifications:     strings.Repeat("S", 302),
					OrderQuantity:      "",
					OrderPrice:         "",
					UnitClassification: "",
				}
				detail2 := &pb.CreateQuotationDetailRequest{
					QuotationsId:       "",
					ProductName:        "",
					Specifications:     strings.Repeat("S", 303),
					OrderQuantity:      "",
					OrderPrice:         "",
					UnitClassification: "",
				}
				QuotationDetails = append(QuotationDetails, detail1, detail2)

				request := &pb.CreateQuotationDetailsRequest{
					QuotationDetails: QuotationDetails,
				}

				res, err := quotationDetailClient.CreateQuotationDetails(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
	})
	t.Run("ItemUnits", func(t *testing.T) {

		t.Run("ListItemUnits", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				now := time.Now()
				ItemUnit := &model.ItemUnit{
					ID:        1,
					TypeValue: 11,
					Name:      "単位：本、冊、枚",
					Remarks:   "mocked REMARKS",
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: now,
				}
				itemUnits := []*model.ItemUnit{ItemUnit, ItemUnit}
				request := &pb.ListItemUnitsRequest{}
				itemUnitsRepository.EXPECT().ListItemUnits(gomock.Any(), gomock.Any()).Return(itemUnits, nil)
				res, err := itemUnitClient.ListItemUnits(context.Background(), request)
				assert.Nil(t, err)
				// Length
				assert.Len(t, res.ItemUnits, 2)
				// Object 1
				assert.Equal(t, int32(1), res.ItemUnits[0].Id.Value)
				assert.Equal(t, int32(11), res.ItemUnits[0].TypeValue.Value)
				assert.Equal(t, "単位：本、冊、枚", res.ItemUnits[0].Name)
				assert.Equal(t, "mocked REMARKS", res.ItemUnits[0].Remarks)
				assert.Equal(t, now.UTC(), res.ItemUnits[0].CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.ItemUnits[0].UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.ItemUnits[0].DeletedAt.AsTime().UTC())
				// Object 2
				assert.Equal(t, int32(1), res.ItemUnits[1].Id.Value)
				assert.Equal(t, int32(11), res.ItemUnits[1].TypeValue.Value)
				assert.Equal(t, "単位：本、冊、枚", res.ItemUnits[1].Name)
				assert.Equal(t, "mocked REMARKS", res.ItemUnits[1].Remarks)
				assert.Equal(t, now.UTC(), res.ItemUnits[1].CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.ItemUnits[1].UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.ItemUnits[1].DeletedAt.AsTime().UTC())
			})
		})
		t.Run("GetItemUnit", func(t *testing.T) {
			t.Run("getByTypeValue", func(t *testing.T) {
				now := time.Now()
				itemUnit := &model.ItemUnit{
					ID:        3,
					TypeValue: 13,
					Name:      "単位：袋、巻",
					Remarks:   "mocked REMARKS2",
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: now,
				}
				request := &pb.GetItemUnitRequest{TypeValue: strconv.FormatInt(itemUnit.TypeValue, 10)}
				itemUnitsRepository.EXPECT().GetItemUnit(gomock.Any(), gomock.Any()).Return(itemUnit, nil)
				res, err := itemUnitClient.GetItemUnit(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(3), res.Id.Value)
				assert.Equal(t, int32(13), res.TypeValue.Value)
				assert.Equal(t, "単位：袋、巻", res.Name)
				assert.Equal(t, "mocked REMARKS2", res.Remarks)
				assert.Equal(t, now.UTC(), res.CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.DeletedAt.AsTime().UTC())
			})
		})
	})
	t.Run("OrderState", func(t *testing.T) {
		now := time.Now()
		t.Run("ListOrderState", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				orderStates := []*model.OrderState{{
					Id:        "1",
					TypeValue: "2",
					Name:      "MockedOrderState",
					Remarks:   "MockedRemark",
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: now,
				}}
				request := &pb.ListOrderStatesRequest{}
				orderStatesRepository.EXPECT().ListOrderStates(gomock.Any(), gomock.Any()).Return(orderStates, nil)
				res, err := orderStateClient.ListOrderStates(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, len(orderStates), len(res.OrderStates))
				assert.Equal(t, int32(1), res.OrderStates[0].Id.Value)
				assert.Equal(t, int32(2), res.OrderStates[0].TypeValue.Value)
				assert.Equal(t, "MockedOrderState", res.OrderStates[0].Name)
				assert.Equal(t, "MockedRemark", res.OrderStates[0].Remarks)
				assert.Equal(t, now.UTC(), res.OrderStates[0].CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.OrderStates[0].UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.OrderStates[0].DeletedAt.AsTime().UTC())
			})

			t.Run("full_conditions", func(t *testing.T) {
				orderStates := []*model.OrderState{{
					Id:        "1",
					TypeValue: "2",
					Name:      "MockedOrderState",
					Remarks:   "MockedRemark",
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: now,
				}}
				request := &pb.ListOrderStatesRequest{
					OrderBy: &pb.ListOrderStatesOrderBy{
						Id:        pb.SortEnum_ASC,
						TypeValue: pb.SortEnum_ASC,
						Name:      pb.SortEnum_ASC,
						Remarks:   pb.SortEnum_ASC,
						CreatedAt: pb.SortEnum_ASC,
						UpdatedAt: pb.SortEnum_ASC,
						DeletedAt: pb.SortEnum_ASC,
					},
				}
				orderStatesRepository.EXPECT().ListOrderStates(gomock.Any(), gomock.Any()).Return(orderStates, nil)
				res, err := orderStateClient.ListOrderStates(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, len(orderStates), len(res.OrderStates))
				assert.Equal(t, int32(1), res.OrderStates[0].Id.Value)
				assert.Equal(t, int32(2), res.OrderStates[0].TypeValue.Value)
				assert.Equal(t, "MockedOrderState", res.OrderStates[0].Name)
				assert.Equal(t, "MockedRemark", res.OrderStates[0].Remarks)
				assert.Equal(t, now.UTC(), res.OrderStates[0].CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.OrderStates[0].UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.OrderStates[0].DeletedAt.AsTime().UTC())
			})
		})
		t.Run("GetOrderState", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				orderState := &model.OrderState{
					Id:        "1",
					TypeValue: "2",
					Name:      "MockedOrderState",
					Remarks:   "MockedRemark",
					CreatedAt: now,
					UpdatedAt: now,
					DeletedAt: now,
				}
				request := &pb.GetOrderStateRequest{TypeValue: orderState.TypeValue}
				orderStatesRepository.EXPECT().GetOrderState(gomock.Any(), gomock.Any()).Return(orderState, nil)
				res, err := orderStateClient.GetOrderState(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, int32(2), res.TypeValue.Value)
				assert.Equal(t, "MockedOrderState", res.Name)
				assert.Equal(t, "MockedRemark", res.Remarks)
				assert.Equal(t, now.UTC(), res.CreatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.UpdatedAt.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.DeletedAt.AsTime().UTC())
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetOrderStateRequest{TypeValue: ""}
				res, err := orderStateClient.GetOrderState(context.Background(), request)
				assert.NotNil(t, err)
				assert.Nil(t, res)
			})
		})
	})
	t.Run("ProjectCostDetail", func(t *testing.T) {
		t.Run("ListProjectCostDetail", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				StockingCost := "123"
				ProjectCostDetail := &model.ProjectCostDetail{
					ID:                  "1",
					OrdersID:            "123",
					OrderDetailsID:      "123",
					AcceptanceDetailsID: "123",
					CostEntryYymm:       "202012",
					ProjectsID:          "123",
					CostTypesID:         "123",
					StockingCost:        &StockingCost,
				}

				projectCostDetails := []*model.ProjectCostDetail{ProjectCostDetail, ProjectCostDetail}
				request := &pb.ListProjectCostDetailsRequest{}
				projectCostDetailRepository.EXPECT().ListProjectCostDetails(gomock.Any(), gomock.Any()).Return(projectCostDetails, nil)
				res, err := projectCostDetailClient.ListProjectCostDetails(context.Background(), request)
				//error
				assert.Nil(t, err)
				//length
				assert.Equal(t, 2, len(res.ProjectCostDetails))
				//ProjectCostDetails[0]
				assert.Equal(t, int32(1), res.ProjectCostDetails[0].Id.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[0].OrdersId.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[0].OrderDetailsId.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[0].AcceptanceDetailsId.Value)
				assert.Equal(t, "202012", res.ProjectCostDetails[0].CostEntryYymm)
				assert.Equal(t, "123", res.ProjectCostDetails[0].ProjectsId)
				assert.Equal(t, "123", res.ProjectCostDetails[0].CostTypesId)
				assert.Equal(t, int32(123), res.ProjectCostDetails[0].StockingCost.Value)
				//ProjectCostDetails[1]
				assert.Equal(t, int32(1), res.ProjectCostDetails[1].Id.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[1].OrdersId.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[1].OrderDetailsId.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[1].AcceptanceDetailsId.Value)
				assert.Equal(t, "202012", res.ProjectCostDetails[1].CostEntryYymm)
				assert.Equal(t, "123", res.ProjectCostDetails[1].ProjectsId)
				assert.Equal(t, "123", res.ProjectCostDetails[1].CostTypesId)
				assert.Equal(t, int32(123), res.ProjectCostDetails[1].StockingCost.Value)
			})
			t.Run("full_conditions", func(t *testing.T) {
				StockingCost := "123"
				ProjectCostDetail := &model.ProjectCostDetail{
					ID:                  "1",
					OrdersID:            "123",
					OrderDetailsID:      "123",
					AcceptanceDetailsID: "123",
					CostEntryYymm:       "202012",
					ProjectsID:          "123",
					CostTypesID:         "123",
					StockingCost:        &StockingCost,
				}

				projectCostDetails := []*model.ProjectCostDetail{ProjectCostDetail, ProjectCostDetail}
				request := &pb.ListProjectCostDetailsRequest{
					Ids:    []string{"1", "2"},
					Limit:  int64(10),
					Offset: int64(0),
				}
				projectCostDetailRepository.EXPECT().ListProjectCostDetails(gomock.Any(), gomock.Any()).Return(projectCostDetails, nil)
				res, err := projectCostDetailClient.ListProjectCostDetails(context.Background(), request)
				//error
				assert.Nil(t, err)
				//length
				assert.Equal(t, 2, len(res.ProjectCostDetails))
				//ProjectCostDetails[0]
				assert.Equal(t, int32(1), res.ProjectCostDetails[0].Id.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[0].OrdersId.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[0].OrderDetailsId.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[0].AcceptanceDetailsId.Value)
				assert.Equal(t, "202012", res.ProjectCostDetails[0].CostEntryYymm)
				assert.Equal(t, "123", res.ProjectCostDetails[0].ProjectsId)
				assert.Equal(t, "123", res.ProjectCostDetails[0].CostTypesId)
				assert.Equal(t, int32(123), res.ProjectCostDetails[0].StockingCost.Value)
				//ProjectCostDetails[1]
				assert.Equal(t, int32(1), res.ProjectCostDetails[1].Id.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[1].OrdersId.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[1].OrderDetailsId.Value)
				assert.Equal(t, int32(123), res.ProjectCostDetails[1].AcceptanceDetailsId.Value)
				assert.Equal(t, "202012", res.ProjectCostDetails[1].CostEntryYymm)
				assert.Equal(t, "123", res.ProjectCostDetails[1].ProjectsId)
				assert.Equal(t, "123", res.ProjectCostDetails[1].CostTypesId)
				assert.Equal(t, int32(123), res.ProjectCostDetails[1].StockingCost.Value)
			})
		})
		t.Run("GetProjectCostDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetProjectCostDetailRequest{}
				res, err := projectCostDetailClient.GetProjectCostDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("by_id", func(t *testing.T) {
				StockingCost := "123"
				projectCostDetail := &model.ProjectCostDetail{
					ID:                  "1",
					OrdersID:            "123",
					OrderDetailsID:      "123",
					AcceptanceDetailsID: "123",
					CostEntryYymm:       "202012",
					ProjectsID:          "123",
					CostTypesID:         "123",
					StockingCost:        &StockingCost,
				}
				request := &pb.GetProjectCostDetailRequest{Id: "123"}
				projectCostDetailRepository.EXPECT().GetProjectCostDetail(gomock.Any(), gomock.Any()).Return(projectCostDetail, nil)
				res, err := projectCostDetailClient.GetProjectCostDetail(context.Background(), request)

				//error
				assert.Nil(t, err)
				//ProjectCostDetail
				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, int32(123), res.OrdersId.Value)
				assert.Equal(t, int32(123), res.OrderDetailsId.Value)
				assert.Equal(t, int32(123), res.AcceptanceDetailsId.Value)
				assert.Equal(t, "202012", res.CostEntryYymm)
				assert.Equal(t, "123", res.ProjectsId)
				assert.Equal(t, "123", res.CostTypesId)
				assert.Equal(t, int32(123), res.StockingCost.Value)
			})
		})
		t.Run("CreateProjectCostDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.CreateProjectCostDetailRequest{
					OrdersId:            "",
					OrderDetailsId:      "",
					AcceptanceDetailsId: "",
					CostEntryYymm:       strings.Repeat("A", 7),
					ProjectsId:          "",
					CostTypesId:         "",
					StockingCost:        ifu.ToInt32Value("10"),
				}
				res, err := projectCostDetailClient.CreateProjectCostDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.CreateProjectCostDetailRequest{
					Id:                  "123",
					OrdersId:            "123",
					OrderDetailsId:      "123",
					AcceptanceDetailsId: "123",
					CostEntryYymm:       "202012",
					ProjectsId:          "123",
					CostTypesId:         "123",
					StockingCost:        &wrappers.Int32Value{Value: 1},
				}
				projectCostDetailRepository.EXPECT().CreateProjectCostDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := projectCostDetailClient.CreateProjectCostDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)

			})
		})
		t.Run("UpdateProjectCostDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateProjectCostDetailRequest{
					Id:                  "",
					OrdersId:            "",
					OrderDetailsId:      "",
					AcceptanceDetailsId: "",
					CostEntryYymm:       strings.Repeat("A", 7),
					ProjectsId:          "",
					CostTypesId:         "",
					StockingCost:        ifu.ToInt32Value("10"),
				}
				res, err := projectCostDetailClient.UpdateProjectCostDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.UpdateProjectCostDetailRequest{
					Id:                  "123",
					OrdersId:            "123",
					OrderDetailsId:      "123",
					AcceptanceDetailsId: "123",
					CostEntryYymm:       "202012",
					ProjectsId:          "123",
					CostTypesId:         "123",
					StockingCost:        &wrappers.Int32Value{Value: 1},
				}
				projectCostDetailRepository.EXPECT().UpdateProjectCostDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := projectCostDetailClient.UpdateProjectCostDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
		})
		t.Run("DeleteProjectCostDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.DeleteProjectCostDetailRequest{}
				res, err := projectCostDetailClient.DeleteProjectCostDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteProjectCostDetailRequest{Id: "123"}
				projectCostDetailRepository.EXPECT().DeleteProjectCostDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := projectCostDetailClient.DeleteProjectCostDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
		})
		t.Run("CountProjectCostDetail", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				request := &pb.ListProjectCostDetailsRequest{}
				projectCostDetailRepository.EXPECT().CountProjectCostDetails(gomock.Any(), gomock.Any()).Return(int64(2), nil)
				res, err := projectCostDetailClient.CountProjectCostDetails(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(2), res.Count)
			})
			t.Run("full_conditions", func(t *testing.T) {
				request := &pb.ListProjectCostDetailsRequest{
					Ids:    []string{"1", "2", "3"},
					Limit:  int64(10),
					Offset: int64(0),
				}
				projectCostDetailRepository.EXPECT().CountProjectCostDetails(gomock.Any(), gomock.Any()).Return(int64(3), nil)
				res, err := projectCostDetailClient.CountProjectCostDetails(context.Background(), request)
				assert.Nil(t, err)
				assert.Equal(t, int64(3), res.Count)
			})
		})

	})
	t.Run("QuotationHistory", func(t *testing.T) {
		now := time.Now()
		commonQuotationHistory := &model.QuotationHistory{
			ID:            "1",
			QuotationsID:  "1",
			QuotationData: "{\"id\": 1, \"jiraNo\": \"666666\", \"remarks\": \"摘要です。\", \"subject\": \"1\", \"supplier\": {\"id\": 1, \"name\": \"文房具\", \"__typename\": \"harp_suppliers\"}, \"requestBy\": \"0000002\", \"requester\": {\"id\": 1, \"kanjiName\": \"test\", \"__typename\": \"libra_staffs\"}, \"quotationNo\": \"\", \"requestDate\": \"2020-09-17\", \"suppliersId\": 6, \"organization\": {\"id\": 1, \"name\": \"ＩＴサービス・ペイメント事業本部\", \"__typename\": \"gemini_organizations\"}, \"versionNumber\": 1, \"quotationsDetails\": [], \"orderClassification\": 10, \"supplierQuotationNo\": \"\", \"quotationInvalidDate\": \"2020-09-17\", \"requestOrganizationId\": 1, \"quotationEffectiveDate\": \"2020-09-17\", \"companyGroupClassification\": \"\"}",
			CreateDate:    now,
		}
		t.Run("ListQuotationHistories", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				quotationHistories := []*model.QuotationHistory{commonQuotationHistory}
				request := &pb.ListQuotationHistoriesRequest{}
				quotationHistoryRepository.EXPECT().ListQuotationHistories(gomock.Any(), gomock.Any()).Return(quotationHistories, nil)
				res, err := quotationHistoryClient.ListQuotationHistories(context.Background(), request)
				assert.Nil(t, err)

				quotationHistoryResponse := []*pb.QuotationHistoryResponse{
					{
						Id:            ifu.ToInt32Value(commonQuotationHistory.ID),
						QuotationsId:  ifu.ToInt32Value(commonQuotationHistory.QuotationsID),
						QuotationData: commonQuotationHistory.QuotationData,
						CreateDate:    ifu.ToTimestamp(commonQuotationHistory.CreateDate),
					},
				}

				assert.Equal(t, quotationHistoryResponse, res.QuotationHistories)

			})
			t.Run("full_conditions", func(t *testing.T) {
				quotationHistories := []*model.QuotationHistory{commonQuotationHistory, commonQuotationHistory}
				request := &pb.ListQuotationHistoriesRequest{
					Ids:          []string{"1", "2"},
					QuotationsId: "1",
					Limit:        int64(10),
					Offset:       int64(0),
				}
				quotationHistoryRepository.EXPECT().ListQuotationHistories(gomock.Any(), gomock.Any()).Return(quotationHistories, nil)
				res, err := quotationHistoryClient.ListQuotationHistories(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, int32(1), res.QuotationHistories[0].GetId().Value)
				assert.Equal(t, int32(1), res.QuotationHistories[0].GetQuotationsId().Value)
				assert.Equal(t, "{\"id\": 1, \"jiraNo\": \"666666\", \"remarks\": \"摘要です。\", \"subject\": \"1\", \"supplier\": {\"id\": 1, \"name\": \"文房具\", \"__typename\": \"harp_suppliers\"}, \"requestBy\": \"0000002\", \"requester\": {\"id\": 1, \"kanjiName\": \"test\", \"__typename\": \"libra_staffs\"}, \"quotationNo\": \"\", \"requestDate\": \"2020-09-17\", \"suppliersId\": 6, \"organization\": {\"id\": 1, \"name\": \"ＩＴサービス・ペイメント事業本部\", \"__typename\": \"gemini_organizations\"}, \"versionNumber\": 1, \"quotationsDetails\": [], \"orderClassification\": 10, \"supplierQuotationNo\": \"\", \"quotationInvalidDate\": \"2020-09-17\", \"requestOrganizationId\": 1, \"quotationEffectiveDate\": \"2020-09-17\", \"companyGroupClassification\": \"\"}", res.QuotationHistories[0].GetQuotationData())
				assert.Equal(t, now.UTC(), res.QuotationHistories[0].GetCreateDate().AsTime().UTC())

				assert.Equal(t, int32(1), res.QuotationHistories[1].GetId().Value)
				assert.Equal(t, int32(1), res.QuotationHistories[1].GetQuotationsId().Value)
				assert.Equal(t, "{\"id\": 1, \"jiraNo\": \"666666\", \"remarks\": \"摘要です。\", \"subject\": \"1\", \"supplier\": {\"id\": 1, \"name\": \"文房具\", \"__typename\": \"harp_suppliers\"}, \"requestBy\": \"0000002\", \"requester\": {\"id\": 1, \"kanjiName\": \"test\", \"__typename\": \"libra_staffs\"}, \"quotationNo\": \"\", \"requestDate\": \"2020-09-17\", \"suppliersId\": 6, \"organization\": {\"id\": 1, \"name\": \"ＩＴサービス・ペイメント事業本部\", \"__typename\": \"gemini_organizations\"}, \"versionNumber\": 1, \"quotationsDetails\": [], \"orderClassification\": 10, \"supplierQuotationNo\": \"\", \"quotationInvalidDate\": \"2020-09-17\", \"requestOrganizationId\": 1, \"quotationEffectiveDate\": \"2020-09-17\", \"companyGroupClassification\": \"\"}", res.QuotationHistories[1].GetQuotationData())
				assert.Equal(t, now.UTC(), res.QuotationHistories[1].GetCreateDate().AsTime().UTC())
			})
		})
		t.Run("GetQuotationHistory", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				request := &pb.GetQuotationHistoryRequest{Id: commonQuotationHistory.ID}
				quotationHistoryRepository.EXPECT().GetQuotationHistory(gomock.Any(), gomock.Any()).Return(commonQuotationHistory, nil)
				res, err := quotationHistoryClient.GetQuotationHistory(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, int32(1), res.QuotationsId.Value)
				assert.Equal(t, "{\"id\": 1, \"jiraNo\": \"666666\", \"remarks\": \"摘要です。\", \"subject\": \"1\", \"supplier\": {\"id\": 1, \"name\": \"文房具\", \"__typename\": \"harp_suppliers\"}, \"requestBy\": \"0000002\", \"requester\": {\"id\": 1, \"kanjiName\": \"test\", \"__typename\": \"libra_staffs\"}, \"quotationNo\": \"\", \"requestDate\": \"2020-09-17\", \"suppliersId\": 6, \"organization\": {\"id\": 1, \"name\": \"ＩＴサービス・ペイメント事業本部\", \"__typename\": \"gemini_organizations\"}, \"versionNumber\": 1, \"quotationsDetails\": [], \"orderClassification\": 10, \"supplierQuotationNo\": \"\", \"quotationInvalidDate\": \"2020-09-17\", \"requestOrganizationId\": 1, \"quotationEffectiveDate\": \"2020-09-17\", \"companyGroupClassification\": \"\"}", res.QuotationData)
				assert.Equal(t, now.UTC(), res.CreateDate.AsTime().UTC())
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetQuotationHistoryRequest{}
				res, err := quotationHistoryClient.GetQuotationHistory(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("CreateQuotationHistory", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)

				request := &pb.CreateQuotationHistoryRequest{
					QuotationsId:  "2",
					QuotationData: "{\"id\": 2, \"jiraNo\": \"777777\", \"remarks\": \"摘要です。\", \"subject\": \"1\", \"supplier\": {\"id\": 1, \"name\": \"文房具\", \"__typename\": \"harp_suppliers\"}, \"requestBy\": \"0000002\", \"requester\": {\"id\": 1, \"kanjiName\": \"test\", \"__typename\": \"libra_staffs\"}, \"quotationNo\": \"\", \"requestDate\": \"2020-09-17\", \"suppliersId\": 6, \"organization\": {\"id\": 1, \"name\": \"ＩＴサービス・ペイメント事業本部\", \"__typename\": \"gemini_organizations\"}, \"versionNumber\": 1, \"quotationsDetails\": [], \"orderClassification\": 10, \"supplierQuotationNo\": \"\", \"quotationInvalidDate\": \"2020-09-17\", \"requestOrganizationId\": 1, \"quotationEffectiveDate\": \"2020-09-17\", \"companyGroupClassification\": \"\"}",
					CreateDate:    ifu.ToTimestamp(time.Now()),
				}
				quotationHistoryRepository.EXPECT().CreateQuotationHistory(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := quotationHistoryClient.CreateQuotationHistory(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.CreateQuotationHistoryRequest{
					QuotationsId:  "",
					QuotationData: "",
					CreateDate:    nil,
				}
				res, err := quotationHistoryClient.CreateQuotationHistory(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("DeleteQuotationHistories", func(t *testing.T) {
			t.Run("by_quotationsId", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteQuotationHistoriesRequest{QuotationsId: "1"}
				quotationHistoryRepository.EXPECT().DeleteQuotationHistories(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := quotationHistoryClient.DeleteQuotationHistories(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.DeleteQuotationHistoriesRequest{}
				res, err := quotationHistoryClient.DeleteQuotationHistories(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
	})
	t.Run("AcceptanceDetails", func(t *testing.T) {
		t.Run("ListAcceptanceDetails", func(t *testing.T) {
			t.Run("all", func(t *testing.T) {
				acceptanceQuantity := "1"
				acceptanceAmount := "1"
				returnQuantity := "1"
				returnPrice := "1"
				suppliersID := "1"
				moduleUnregister := "1"
				projectsId := "1"
				projectCostId := "1"
				scheduledAcceptanceQuantity := "1"
				scheduledAcceptanceAmount := "1"
				now := time.Now()
				acceptanceDetails := []*model.AcceptanceDetail{
					{
						ID:                          "1",
						OrdersId:                    "1",
						OrderDetailsId:              "1",
						ScheduledAcceptanceDate:     now,
						ActualAcceptanceDate:        now,
						AcceptanceQuantity:          &acceptanceQuantity,
						AcceptanceAmount:            &acceptanceAmount,
						ReturnQuantity:              &returnQuantity,
						ReturnPrice:                 &returnPrice,
						ApprovalDate:                now,
						ApprovalBy:                  "1",
						Remarks:                     "テスト",
						SuppliersId:                 &suppliersID,
						ScheduledAcceptanceYymm:     "202009",
						ActualAcceptanceYymm:        "202009",
						ModuleUnregister:            &moduleUnregister,
						ProjectsId:                  projectsId,
						ProjectCostId:               projectCostId,
						ScheduledAcceptanceQuantity: &scheduledAcceptanceQuantity,
						ScheduledAcceptanceAmount:   &scheduledAcceptanceAmount,
					},
				}
				// AcceptanceDetails := []*model.AcceptanceDetail{acceptanceDetails}
				request := &pb.ListAcceptanceDetailsRequest{}
				acceptanceDetailsRepository.EXPECT().ListAcceptanceDetails(gomock.Any(), gomock.Any()).Return(acceptanceDetails, nil)
				res, err := acceptanceDetailClient.ListAcceptanceDetails(context.Background(), request)
				assert.Nil(t, err)

				assert.Len(t, res.AcceptanceDetails, 1)
				assert.Equal(t, int32(1), res.AcceptanceDetails[0].Id.Value)
				assert.Equal(t, int32(1), res.AcceptanceDetails[0].OrdersId.Value)
				assert.Equal(t, int32(1), res.AcceptanceDetails[0].OrderDetailsId.Value)
				assert.Equal(t, now.UTC(), res.AcceptanceDetails[0].ScheduledAcceptanceDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.AcceptanceDetails[0].ActualAcceptanceDate.AsTime().UTC())
				assert.Equal(t, float64(1), res.AcceptanceDetails[0].AcceptanceQuantity.Value)
				assert.Equal(t, int32(1), res.AcceptanceDetails[0].AcceptanceAmount.Value)
				assert.Equal(t, float64(1), res.AcceptanceDetails[0].ReturnQuantity.Value)
				assert.Equal(t, int32(1), res.AcceptanceDetails[0].ReturnPrice.Value)
				assert.Equal(t, now.UTC(), res.AcceptanceDetails[0].ApprovalDate.AsTime().UTC())
				assert.Equal(t, "1", res.AcceptanceDetails[0].ApprovalBy)
				assert.Equal(t, acceptanceDetails[0].Remarks, res.AcceptanceDetails[0].Remarks)
				assert.Equal(t, int32(1), res.AcceptanceDetails[0].SuppliersId.Value)
				assert.Equal(t, acceptanceDetails[0].ScheduledAcceptanceYymm, res.AcceptanceDetails[0].ScheduledAcceptanceYymm)
				assert.Equal(t, acceptanceDetails[0].ActualAcceptanceYymm, res.AcceptanceDetails[0].ActualAcceptanceYymm)
				assert.Equal(t, int32(1), res.AcceptanceDetails[0].ModuleUnregister.Value)
				assert.Equal(t, "1", res.AcceptanceDetails[0].ProjectsId)
				assert.Equal(t, "1", res.AcceptanceDetails[0].ProjectCostId)
				assert.Equal(t, float64(1), res.AcceptanceDetails[0].ScheduledAcceptanceQuantity.Value)
				assert.Equal(t, int32(1), res.AcceptanceDetails[0].ScheduledAcceptanceAmount.Value)

			})
		})
		t.Run("GetAcceptanceDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetAcceptanceDetailRequest{}
				res, err := acceptanceDetailClient.GetAcceptanceDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})

			t.Run("success", func(t *testing.T) {
				acceptanceQuantity := "2"
				acceptanceAmount := "2"
				returnQuantity := "2"
				returnPrice := "2"
				suppliersID := "2"
				moduleUnregister := "2"
				projectsId := "1"
				projectCostId := "1"
				scheduledAcceptanceQuantity := "1"
				scheduledAcceptanceAmount := "1"
				now := time.Now()
				acceptanceDetail := &model.AcceptanceDetail{
					ID:                          "2",
					OrdersId:                    "2",
					OrderDetailsId:              "2",
					ScheduledAcceptanceDate:     now,
					ActualAcceptanceDate:        now,
					AcceptanceQuantity:          &acceptanceQuantity,
					AcceptanceAmount:            &acceptanceAmount,
					ReturnQuantity:              &returnQuantity,
					ReturnPrice:                 &returnPrice,
					ApprovalDate:                now,
					ApprovalBy:                  "2",
					Remarks:                     "テスト2",
					SuppliersId:                 &suppliersID,
					ScheduledAcceptanceYymm:     "202009",
					ActualAcceptanceYymm:        "202009",
					ModuleUnregister:            &moduleUnregister,
					ProjectsId:                  projectsId,
					ProjectCostId:               projectCostId,
					ScheduledAcceptanceQuantity: &scheduledAcceptanceQuantity,
					ScheduledAcceptanceAmount:   &scheduledAcceptanceAmount,
				}
				request := &pb.GetAcceptanceDetailRequest{Id: acceptanceDetail.ID}
				acceptanceDetailsRepository.EXPECT().GetAcceptanceDetail(gomock.Any(), gomock.Any()).Return(acceptanceDetail, nil)
				res, err := acceptanceDetailClient.GetAcceptanceDetail(context.Background(), request)

				assert.Nil(t, err)

				assert.Equal(t, int32(2), res.Id.Value)
				assert.Equal(t, int32(2), res.OrdersId.Value)
				assert.Equal(t, int32(2), res.OrderDetailsId.Value)
				assert.Equal(t, now.UTC(), res.ScheduledAcceptanceDate.AsTime().UTC())
				assert.Equal(t, now.UTC(), res.ActualAcceptanceDate.AsTime().UTC())
				assert.Equal(t, float64(2), res.AcceptanceQuantity.Value)
				assert.Equal(t, int32(2), res.AcceptanceAmount.Value)
				assert.Equal(t, float64(2), res.ReturnQuantity.Value)
				assert.Equal(t, int32(2), res.ReturnPrice.Value)
				assert.Equal(t, now.UTC(), res.ApprovalDate.AsTime().UTC())
				assert.Equal(t, "2", res.ApprovalBy)
				assert.Equal(t, "テスト2", res.Remarks)
				assert.Equal(t, int32(2), res.SuppliersId.Value)
				assert.Equal(t, "202009", res.ScheduledAcceptanceYymm)
				assert.Equal(t, "202009", res.ActualAcceptanceYymm)
				assert.Equal(t, int32(2), res.ModuleUnregister.Value)
				assert.Equal(t, "1", res.ProjectsId)
				assert.Equal(t, "1", res.ProjectCostId)
				assert.Equal(t, float64(1), res.ScheduledAcceptanceQuantity.Value)
				assert.Equal(t, int32(1), res.ScheduledAcceptanceAmount.Value)

			})
		})
		t.Run("CreateAcceptanceDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.CreateAcceptanceDetailRequest{
					OrdersId:                    1234567,
					OrderDetailsId:              1234567,
					ScheduledAcceptanceDate:     nil,
					ActualAcceptanceDate:        nil,
					AcceptanceQuantity:          nil,
					AcceptanceAmount:            nil,
					ReturnQuantity:              nil,
					ReturnPrice:                 nil,
					ApprovalDate:                nil,
					ApprovalBy:                  "",
					Remarks:                     strings.Repeat("R", 301),
					SuppliersId:                 nil,
					ScheduledAcceptanceYymm:     strings.Repeat("R", 8),
					ActualAcceptanceYymm:        strings.Repeat("R", 8),
					ModuleUnregister:            nil,
					ProjectsId:                  "",
					ProjectCostId:               "",
					ScheduledAcceptanceQuantity: nil,
					ScheduledAcceptanceAmount:   nil,
				}
				res, err := acceptanceDetailClient.CreateAcceptanceDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})

			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)
				nowTimestamp, _ := ptypes.TimestampProto(time.Now())
				request := &pb.CreateAcceptanceDetailRequest{
					OrdersId:                    14,
					OrderDetailsId:              1,
					ScheduledAcceptanceDate:     nowTimestamp,
					ActualAcceptanceDate:        nowTimestamp,
					AcceptanceQuantity:          &wrappers.DoubleValue{Value: 1},
					AcceptanceAmount:            &wrappers.Int32Value{Value: 2},
					ReturnQuantity:              &wrappers.DoubleValue{Value: 3},
					ReturnPrice:                 &wrappers.Int32Value{Value: 4},
					ApprovalDate:                nowTimestamp,
					ApprovalBy:                  "5",
					Remarks:                     "テスト3",
					SuppliersId:                 &wrappers.Int32Value{Value: 6},
					ScheduledAcceptanceYymm:     "202009",
					ActualAcceptanceYymm:        "202009",
					ModuleUnregister:            &wrappers.Int32Value{Value: 7},
					ProjectsId:                  "1",
					ProjectCostId:               "1",
					ScheduledAcceptanceQuantity: &wrappers.DoubleValue{Value: 1},
					ScheduledAcceptanceAmount:   &wrappers.Int32Value{Value: 2},
				}
				acceptanceDetailsRepository.EXPECT().CreateAcceptanceDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := acceptanceDetailClient.CreateAcceptanceDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.Equal(t, affectedRows, res.AffectedRows)

			})
		})
		t.Run("UpdateAcceptanceDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateAcceptanceDetailRequest{
					Id:                          1234567,
					OrdersId:                    1234567,
					OrderDetailsId:              1234567,
					ScheduledAcceptanceDate:     nil,
					ActualAcceptanceDate:        nil,
					AcceptanceQuantity:          nil,
					AcceptanceAmount:            nil,
					ReturnQuantity:              nil,
					ReturnPrice:                 nil,
					ApprovalDate:                nil,
					ApprovalBy:                  "",
					Remarks:                     strings.Repeat("R", 301),
					SuppliersId:                 nil,
					ScheduledAcceptanceYymm:     strings.Repeat("R", 8),
					ActualAcceptanceYymm:        strings.Repeat("R", 8),
					ModuleUnregister:            nil,
					ProjectsId:                  "",
					ProjectCostId:               "",
					ScheduledAcceptanceQuantity: nil,
					ScheduledAcceptanceAmount:   nil,
				}
				res, err := acceptanceDetailClient.UpdateAcceptanceDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)
				nowTimestamp, _ := ptypes.TimestampProto(time.Now())
				request := &pb.UpdateAcceptanceDetailRequest{
					Id:                          3,
					OrdersId:                    15,
					OrderDetailsId:              1,
					ScheduledAcceptanceDate:     nowTimestamp,
					ActualAcceptanceDate:        nowTimestamp,
					AcceptanceQuantity:          &wrappers.DoubleValue{Value: 1},
					AcceptanceAmount:            &wrappers.Int32Value{Value: 2},
					ReturnQuantity:              &wrappers.DoubleValue{Value: 3},
					ReturnPrice:                 &wrappers.Int32Value{Value: 4},
					ApprovalDate:                nowTimestamp,
					ApprovalBy:                  "5",
					Remarks:                     "テスト4",
					SuppliersId:                 &wrappers.Int32Value{Value: 6},
					ScheduledAcceptanceYymm:     "202009",
					ActualAcceptanceYymm:        "202009",
					ModuleUnregister:            &wrappers.Int32Value{Value: 7},
					ProjectsId:                  "1",
					ProjectCostId:               "1",
					ScheduledAcceptanceQuantity: &wrappers.DoubleValue{Value: 1},
					ScheduledAcceptanceAmount:   &wrappers.Int32Value{Value: 2},
				}
				acceptanceDetailsRepository.EXPECT().UpdateAcceptanceDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := acceptanceDetailClient.UpdateAcceptanceDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
		})
		t.Run("DeleteAcceptanceDetail", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.DeleteAcceptanceDetailRequest{}
				res, err := acceptanceDetailClient.DeleteAcceptanceDetail(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.DeleteAcceptanceDetailRequest{Id: "1"}
				acceptanceDetailsRepository.EXPECT().DeleteAcceptanceDetail(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := acceptanceDetailClient.DeleteAcceptanceDetail(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
		})
		t.Run("UpdateAcceptanceDetailModuleUnregister", func(t *testing.T) {
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.UpdateAcceptanceDetailModuleUnregisterRequest{
					Id:               1234567,
					ModuleUnregister: nil,
				}
				res, err := acceptanceDetailClient.UpdateAcceptanceDetailModuleUnregister(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.Internal, status.Code())
			})
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)
				request := &pb.UpdateAcceptanceDetailModuleUnregisterRequest{
					Id:               3,
					ModuleUnregister: &wrappers.Int32Value{Value: 7},
				}
				acceptanceDetailsRepository.EXPECT().UpdateAcceptanceDetailModuleUnregister(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := acceptanceDetailClient.UpdateAcceptanceDetailModuleUnregister(context.Background(), request)

				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
		})
	})
	t.Run("OrderHistory", func(t *testing.T) {
		now := time.Now()
		commonOrderHistory := &model.OrderHistory{
			ID:         "1",
			OrdersID:   "1",
			OrderData:  "{}",
			CreateDate: now,
		}
		t.Run("ListOrderHistories", func(t *testing.T) {
			t.Run("full_conditions", func(t *testing.T) {
				orderHistories := []*model.OrderHistory{commonOrderHistory, commonOrderHistory}
				request := &pb.ListOrderHistoriesRequest{
					OrdersId: "1",
					Limit:    int64(10),
					Offset:   int64(0),
				}
				orderHistoryRepository.EXPECT().ListOrderHistories(gomock.Any(), gomock.Any()).Return(orderHistories, nil)
				res, err := orderHistoryClient.ListOrderHistories(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, int32(1), res.OrderHistories[0].GetId().Value)
				assert.Equal(t, int32(1), res.OrderHistories[0].GetOrdersId().Value)
				assert.Equal(t, "{}", res.OrderHistories[0].GetOrderData())
				assert.Equal(t, now.UTC(), res.OrderHistories[0].GetCreateDate().AsTime().UTC())

				assert.Equal(t, int32(1), res.OrderHistories[1].GetId().Value)
				assert.Equal(t, int32(1), res.OrderHistories[1].GetOrdersId().Value)
				assert.Equal(t, "{}", res.OrderHistories[1].GetOrderData())
				assert.Equal(t, now.UTC(), res.OrderHistories[1].GetCreateDate().AsTime().UTC())
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.ListOrderHistoriesRequest{}
				res, err := orderHistoryClient.ListOrderHistories(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("GetOrderHistory", func(t *testing.T) {
			t.Run("by_id", func(t *testing.T) {
				request := &pb.GetOrderHistoryRequest{Id: commonOrderHistory.ID}
				orderHistoryRepository.EXPECT().GetOrderHistory(gomock.Any(), gomock.Any()).Return(commonOrderHistory, nil)
				res, err := orderHistoryClient.GetOrderHistory(context.Background(), request)
				assert.Nil(t, err)

				assert.Equal(t, int32(1), res.Id.Value)
				assert.Equal(t, int32(1), res.OrdersId.Value)
				assert.Equal(t, "{}", res.OrderData)
				assert.Equal(t, now.UTC(), res.CreateDate.AsTime().UTC())
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.GetOrderHistoryRequest{}
				res, err := orderHistoryClient.GetOrderHistory(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
		t.Run("CreateOrderHistory", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				affectedRows := int64(1)

				request := &pb.CreateOrderHistoryRequest{
					OrdersId:     "2",
					OrderData:    "{}",
					Subject:      "subject",
					OrderCaseCd:  "2",
					SuppliersId:  "2",
					JiraNo:       "2NOOOJ",
					ApprovalFile: "c://",
					Remarks:      "remarks",
					// 発注明細データ更新
					ProductName:                       "name",
					OrderUnitPrice:                    "20",
					OrderQuantity:                     "2",
					OrderUnitClassification:           ifu.ToInt32Value("2"),
					AcceptanceScheduledDate:           ifu.ToTimestamp(now),
					ConfigurationManagementTargetFlag: "1",
					Specifications:                    "2",
					RemarksDetails:                    "remarks",
					// PRJコスト明細 - 発注データ 更新
					ProjectsId:    "2",
					ProjectCostId: "2",
					CostTypesId:   "2",
				}
				orderHistoryRepository.EXPECT().CreateOrderHistory(gomock.Any(), gomock.Any()).Return(affectedRows, nil)
				res, err := orderHistoryClient.CreateOrderHistory(context.Background(), request)
				assert.Nil(t, err)
				assert.EqualValues(t, affectedRows, res.AffectedRows)
			})
			t.Run("validation_error", func(t *testing.T) {
				request := &pb.CreateOrderHistoryRequest{
					OrdersId:  "",
					OrderData: "",
				}
				res, err := orderHistoryClient.CreateOrderHistory(context.Background(), request)
				assert.Nil(t, res)
				assert.NotNil(t, err)
				status := status.Convert(err)
				assert.NotNil(t, status)
				assert.Equal(t, codes.InvalidArgument, status.Code())
			})
		})
	})
}
