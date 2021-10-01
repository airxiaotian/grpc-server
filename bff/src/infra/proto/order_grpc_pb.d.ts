// package: com.paylabo.c002.harp.v1
// file: order.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as order_pb from "./order_pb";
import * as enums_pb from "./enums_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";

interface IOrderService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listUnacceptedOrdersThisMonth: IOrderService_IListUnacceptedOrdersThisMonth;
    listOrders: IOrderService_IListOrders;
    countOrders: IOrderService_ICountOrders;
    getOrder: IOrderService_IGetOrder;
    createOrder: IOrderService_ICreateOrder;
    updateOrder: IOrderService_IUpdateOrder;
    updateOrderProjectCostInfo: IOrderService_IUpdateOrderProjectCostInfo;
    deleteOrder: IOrderService_IDeleteOrder;
    getOrderRequesterAggregate: IOrderService_IGetOrderRequesterAggregate;
    getOrderSupplierAggregate: IOrderService_IGetOrderSupplierAggregate;
    countOrdersWithGroupBy: IOrderService_ICountOrdersWithGroupBy;
    sumNearestTwoMonthsAmount: IOrderService_ISumNearestTwoMonthsAmount;
}

interface IOrderService_IListUnacceptedOrdersThisMonth extends grpc.MethodDefinition<order_pb.ListOrdersRequest, order_pb.OrdersResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/ListUnacceptedOrdersThisMonth"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.ListOrdersRequest>;
    requestDeserialize: grpc.deserialize<order_pb.ListOrdersRequest>;
    responseSerialize: grpc.serialize<order_pb.OrdersResponse>;
    responseDeserialize: grpc.deserialize<order_pb.OrdersResponse>;
}
interface IOrderService_IListOrders extends grpc.MethodDefinition<order_pb.ListOrdersRequest, order_pb.OrdersResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/ListOrders"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.ListOrdersRequest>;
    requestDeserialize: grpc.deserialize<order_pb.ListOrdersRequest>;
    responseSerialize: grpc.serialize<order_pb.OrdersResponse>;
    responseDeserialize: grpc.deserialize<order_pb.OrdersResponse>;
}
interface IOrderService_ICountOrders extends grpc.MethodDefinition<order_pb.ListOrdersRequest, order_pb.CountOrdersResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/CountOrders"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.ListOrdersRequest>;
    requestDeserialize: grpc.deserialize<order_pb.ListOrdersRequest>;
    responseSerialize: grpc.serialize<order_pb.CountOrdersResponse>;
    responseDeserialize: grpc.deserialize<order_pb.CountOrdersResponse>;
}
interface IOrderService_IGetOrder extends grpc.MethodDefinition<order_pb.GetOrderRequest, order_pb.OrderResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/GetOrder"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.GetOrderRequest>;
    requestDeserialize: grpc.deserialize<order_pb.GetOrderRequest>;
    responseSerialize: grpc.serialize<order_pb.OrderResponse>;
    responseDeserialize: grpc.deserialize<order_pb.OrderResponse>;
}
interface IOrderService_ICreateOrder extends grpc.MethodDefinition<order_pb.CreateOrderRequest, order_pb.MutationOrderResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/CreateOrder"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.CreateOrderRequest>;
    requestDeserialize: grpc.deserialize<order_pb.CreateOrderRequest>;
    responseSerialize: grpc.serialize<order_pb.MutationOrderResponse>;
    responseDeserialize: grpc.deserialize<order_pb.MutationOrderResponse>;
}
interface IOrderService_IUpdateOrder extends grpc.MethodDefinition<order_pb.UpdateOrderRequest, order_pb.MutationOrderResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/UpdateOrder"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.UpdateOrderRequest>;
    requestDeserialize: grpc.deserialize<order_pb.UpdateOrderRequest>;
    responseSerialize: grpc.serialize<order_pb.MutationOrderResponse>;
    responseDeserialize: grpc.deserialize<order_pb.MutationOrderResponse>;
}
interface IOrderService_IUpdateOrderProjectCostInfo extends grpc.MethodDefinition<order_pb.UpdateOrderProjectCostInfoRequest, order_pb.MutationOrderResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/UpdateOrderProjectCostInfo"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.UpdateOrderProjectCostInfoRequest>;
    requestDeserialize: grpc.deserialize<order_pb.UpdateOrderProjectCostInfoRequest>;
    responseSerialize: grpc.serialize<order_pb.MutationOrderResponse>;
    responseDeserialize: grpc.deserialize<order_pb.MutationOrderResponse>;
}
interface IOrderService_IDeleteOrder extends grpc.MethodDefinition<order_pb.DeleteOrderRequest, order_pb.MutationOrderResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/DeleteOrder"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.DeleteOrderRequest>;
    requestDeserialize: grpc.deserialize<order_pb.DeleteOrderRequest>;
    responseSerialize: grpc.serialize<order_pb.MutationOrderResponse>;
    responseDeserialize: grpc.deserialize<order_pb.MutationOrderResponse>;
}
interface IOrderService_IGetOrderRequesterAggregate extends grpc.MethodDefinition<order_pb.OrderAggregateRequest, order_pb.OrderAggregateResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/GetOrderRequesterAggregate"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.OrderAggregateRequest>;
    requestDeserialize: grpc.deserialize<order_pb.OrderAggregateRequest>;
    responseSerialize: grpc.serialize<order_pb.OrderAggregateResponse>;
    responseDeserialize: grpc.deserialize<order_pb.OrderAggregateResponse>;
}
interface IOrderService_IGetOrderSupplierAggregate extends grpc.MethodDefinition<order_pb.OrderAggregateRequest, order_pb.OrderAggregateResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/GetOrderSupplierAggregate"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.OrderAggregateRequest>;
    requestDeserialize: grpc.deserialize<order_pb.OrderAggregateRequest>;
    responseSerialize: grpc.serialize<order_pb.OrderAggregateResponse>;
    responseDeserialize: grpc.deserialize<order_pb.OrderAggregateResponse>;
}
interface IOrderService_ICountOrdersWithGroupBy extends grpc.MethodDefinition<order_pb.OrdersGroupByRequest, order_pb.OrdersGroupByResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/CountOrdersWithGroupBy"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.OrdersGroupByRequest>;
    requestDeserialize: grpc.deserialize<order_pb.OrdersGroupByRequest>;
    responseSerialize: grpc.serialize<order_pb.OrdersGroupByResponse>;
    responseDeserialize: grpc.deserialize<order_pb.OrdersGroupByResponse>;
}
interface IOrderService_ISumNearestTwoMonthsAmount extends grpc.MethodDefinition<order_pb.SumNearestTwoMonthsAmountRequest, order_pb.SumNearestTwoMonthsAmountResponse> {
    path: string; // "/com.paylabo.c002.harp.v1.Order/SumNearestTwoMonthsAmount"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<order_pb.SumNearestTwoMonthsAmountRequest>;
    requestDeserialize: grpc.deserialize<order_pb.SumNearestTwoMonthsAmountRequest>;
    responseSerialize: grpc.serialize<order_pb.SumNearestTwoMonthsAmountResponse>;
    responseDeserialize: grpc.deserialize<order_pb.SumNearestTwoMonthsAmountResponse>;
}

export const OrderService: IOrderService;

export interface IOrderServer {
    listUnacceptedOrdersThisMonth: grpc.handleUnaryCall<order_pb.ListOrdersRequest, order_pb.OrdersResponse>;
    listOrders: grpc.handleUnaryCall<order_pb.ListOrdersRequest, order_pb.OrdersResponse>;
    countOrders: grpc.handleUnaryCall<order_pb.ListOrdersRequest, order_pb.CountOrdersResponse>;
    getOrder: grpc.handleUnaryCall<order_pb.GetOrderRequest, order_pb.OrderResponse>;
    createOrder: grpc.handleUnaryCall<order_pb.CreateOrderRequest, order_pb.MutationOrderResponse>;
    updateOrder: grpc.handleUnaryCall<order_pb.UpdateOrderRequest, order_pb.MutationOrderResponse>;
    updateOrderProjectCostInfo: grpc.handleUnaryCall<order_pb.UpdateOrderProjectCostInfoRequest, order_pb.MutationOrderResponse>;
    deleteOrder: grpc.handleUnaryCall<order_pb.DeleteOrderRequest, order_pb.MutationOrderResponse>;
    getOrderRequesterAggregate: grpc.handleUnaryCall<order_pb.OrderAggregateRequest, order_pb.OrderAggregateResponse>;
    getOrderSupplierAggregate: grpc.handleUnaryCall<order_pb.OrderAggregateRequest, order_pb.OrderAggregateResponse>;
    countOrdersWithGroupBy: grpc.handleUnaryCall<order_pb.OrdersGroupByRequest, order_pb.OrdersGroupByResponse>;
    sumNearestTwoMonthsAmount: grpc.handleUnaryCall<order_pb.SumNearestTwoMonthsAmountRequest, order_pb.SumNearestTwoMonthsAmountResponse>;
}

export interface IOrderClient {
    listUnacceptedOrdersThisMonth(request: order_pb.ListOrdersRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    listUnacceptedOrdersThisMonth(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    listUnacceptedOrdersThisMonth(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    listOrders(request: order_pb.ListOrdersRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    listOrders(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    listOrders(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    countOrders(request: order_pb.ListOrdersRequest, callback: (error: grpc.ServiceError | null, response: order_pb.CountOrdersResponse) => void): grpc.ClientUnaryCall;
    countOrders(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.CountOrdersResponse) => void): grpc.ClientUnaryCall;
    countOrders(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.CountOrdersResponse) => void): grpc.ClientUnaryCall;
    getOrder(request: order_pb.GetOrderRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrderResponse) => void): grpc.ClientUnaryCall;
    getOrder(request: order_pb.GetOrderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrderResponse) => void): grpc.ClientUnaryCall;
    getOrder(request: order_pb.GetOrderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrderResponse) => void): grpc.ClientUnaryCall;
    createOrder(request: order_pb.CreateOrderRequest, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    createOrder(request: order_pb.CreateOrderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    createOrder(request: order_pb.CreateOrderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    updateOrder(request: order_pb.UpdateOrderRequest, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    updateOrder(request: order_pb.UpdateOrderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    updateOrder(request: order_pb.UpdateOrderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    updateOrderProjectCostInfo(request: order_pb.UpdateOrderProjectCostInfoRequest, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    updateOrderProjectCostInfo(request: order_pb.UpdateOrderProjectCostInfoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    updateOrderProjectCostInfo(request: order_pb.UpdateOrderProjectCostInfoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    deleteOrder(request: order_pb.DeleteOrderRequest, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    deleteOrder(request: order_pb.DeleteOrderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    deleteOrder(request: order_pb.DeleteOrderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    getOrderRequesterAggregate(request: order_pb.OrderAggregateRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    getOrderRequesterAggregate(request: order_pb.OrderAggregateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    getOrderRequesterAggregate(request: order_pb.OrderAggregateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    getOrderSupplierAggregate(request: order_pb.OrderAggregateRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    getOrderSupplierAggregate(request: order_pb.OrderAggregateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    getOrderSupplierAggregate(request: order_pb.OrderAggregateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    countOrdersWithGroupBy(request: order_pb.OrdersGroupByRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersGroupByResponse) => void): grpc.ClientUnaryCall;
    countOrdersWithGroupBy(request: order_pb.OrdersGroupByRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersGroupByResponse) => void): grpc.ClientUnaryCall;
    countOrdersWithGroupBy(request: order_pb.OrdersGroupByRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersGroupByResponse) => void): grpc.ClientUnaryCall;
    sumNearestTwoMonthsAmount(request: order_pb.SumNearestTwoMonthsAmountRequest, callback: (error: grpc.ServiceError | null, response: order_pb.SumNearestTwoMonthsAmountResponse) => void): grpc.ClientUnaryCall;
    sumNearestTwoMonthsAmount(request: order_pb.SumNearestTwoMonthsAmountRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.SumNearestTwoMonthsAmountResponse) => void): grpc.ClientUnaryCall;
    sumNearestTwoMonthsAmount(request: order_pb.SumNearestTwoMonthsAmountRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.SumNearestTwoMonthsAmountResponse) => void): grpc.ClientUnaryCall;
}

export class OrderClient extends grpc.Client implements IOrderClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public listUnacceptedOrdersThisMonth(request: order_pb.ListOrdersRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    public listUnacceptedOrdersThisMonth(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    public listUnacceptedOrdersThisMonth(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    public listOrders(request: order_pb.ListOrdersRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    public listOrders(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    public listOrders(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersResponse) => void): grpc.ClientUnaryCall;
    public countOrders(request: order_pb.ListOrdersRequest, callback: (error: grpc.ServiceError | null, response: order_pb.CountOrdersResponse) => void): grpc.ClientUnaryCall;
    public countOrders(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.CountOrdersResponse) => void): grpc.ClientUnaryCall;
    public countOrders(request: order_pb.ListOrdersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.CountOrdersResponse) => void): grpc.ClientUnaryCall;
    public getOrder(request: order_pb.GetOrderRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrderResponse) => void): grpc.ClientUnaryCall;
    public getOrder(request: order_pb.GetOrderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrderResponse) => void): grpc.ClientUnaryCall;
    public getOrder(request: order_pb.GetOrderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrderResponse) => void): grpc.ClientUnaryCall;
    public createOrder(request: order_pb.CreateOrderRequest, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public createOrder(request: order_pb.CreateOrderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public createOrder(request: order_pb.CreateOrderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public updateOrder(request: order_pb.UpdateOrderRequest, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public updateOrder(request: order_pb.UpdateOrderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public updateOrder(request: order_pb.UpdateOrderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public updateOrderProjectCostInfo(request: order_pb.UpdateOrderProjectCostInfoRequest, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public updateOrderProjectCostInfo(request: order_pb.UpdateOrderProjectCostInfoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public updateOrderProjectCostInfo(request: order_pb.UpdateOrderProjectCostInfoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public deleteOrder(request: order_pb.DeleteOrderRequest, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public deleteOrder(request: order_pb.DeleteOrderRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public deleteOrder(request: order_pb.DeleteOrderRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.MutationOrderResponse) => void): grpc.ClientUnaryCall;
    public getOrderRequesterAggregate(request: order_pb.OrderAggregateRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    public getOrderRequesterAggregate(request: order_pb.OrderAggregateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    public getOrderRequesterAggregate(request: order_pb.OrderAggregateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    public getOrderSupplierAggregate(request: order_pb.OrderAggregateRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    public getOrderSupplierAggregate(request: order_pb.OrderAggregateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    public getOrderSupplierAggregate(request: order_pb.OrderAggregateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrderAggregateResponse) => void): grpc.ClientUnaryCall;
    public countOrdersWithGroupBy(request: order_pb.OrdersGroupByRequest, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersGroupByResponse) => void): grpc.ClientUnaryCall;
    public countOrdersWithGroupBy(request: order_pb.OrdersGroupByRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersGroupByResponse) => void): grpc.ClientUnaryCall;
    public countOrdersWithGroupBy(request: order_pb.OrdersGroupByRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.OrdersGroupByResponse) => void): grpc.ClientUnaryCall;
    public sumNearestTwoMonthsAmount(request: order_pb.SumNearestTwoMonthsAmountRequest, callback: (error: grpc.ServiceError | null, response: order_pb.SumNearestTwoMonthsAmountResponse) => void): grpc.ClientUnaryCall;
    public sumNearestTwoMonthsAmount(request: order_pb.SumNearestTwoMonthsAmountRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: order_pb.SumNearestTwoMonthsAmountResponse) => void): grpc.ClientUnaryCall;
    public sumNearestTwoMonthsAmount(request: order_pb.SumNearestTwoMonthsAmountRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: order_pb.SumNearestTwoMonthsAmountResponse) => void): grpc.ClientUnaryCall;
}
