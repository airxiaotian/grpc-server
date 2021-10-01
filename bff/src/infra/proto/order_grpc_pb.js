// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var order_pb = require('./order_pb.js');
var enums_pb = require('./enums_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');

function serialize_com_paylabo_c002_harp_v1_CountOrdersResponse(arg) {
  if (!(arg instanceof order_pb.CountOrdersResponse)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.CountOrdersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_CountOrdersResponse(buffer_arg) {
  return order_pb.CountOrdersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_CreateOrderRequest(arg) {
  if (!(arg instanceof order_pb.CreateOrderRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.CreateOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_CreateOrderRequest(buffer_arg) {
  return order_pb.CreateOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_DeleteOrderRequest(arg) {
  if (!(arg instanceof order_pb.DeleteOrderRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.DeleteOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_DeleteOrderRequest(buffer_arg) {
  return order_pb.DeleteOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_GetOrderRequest(arg) {
  if (!(arg instanceof order_pb.GetOrderRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.GetOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_GetOrderRequest(buffer_arg) {
  return order_pb.GetOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_ListOrdersRequest(arg) {
  if (!(arg instanceof order_pb.ListOrdersRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.ListOrdersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_ListOrdersRequest(buffer_arg) {
  return order_pb.ListOrdersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_MutationOrderResponse(arg) {
  if (!(arg instanceof order_pb.MutationOrderResponse)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.MutationOrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_MutationOrderResponse(buffer_arg) {
  return order_pb.MutationOrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_OrderAggregateRequest(arg) {
  if (!(arg instanceof order_pb.OrderAggregateRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.OrderAggregateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_OrderAggregateRequest(buffer_arg) {
  return order_pb.OrderAggregateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_OrderAggregateResponse(arg) {
  if (!(arg instanceof order_pb.OrderAggregateResponse)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.OrderAggregateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_OrderAggregateResponse(buffer_arg) {
  return order_pb.OrderAggregateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_OrderResponse(arg) {
  if (!(arg instanceof order_pb.OrderResponse)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.OrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_OrderResponse(buffer_arg) {
  return order_pb.OrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_OrdersGroupByRequest(arg) {
  if (!(arg instanceof order_pb.OrdersGroupByRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.OrdersGroupByRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_OrdersGroupByRequest(buffer_arg) {
  return order_pb.OrdersGroupByRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_OrdersGroupByResponse(arg) {
  if (!(arg instanceof order_pb.OrdersGroupByResponse)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.OrdersGroupByResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_OrdersGroupByResponse(buffer_arg) {
  return order_pb.OrdersGroupByResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_OrdersResponse(arg) {
  if (!(arg instanceof order_pb.OrdersResponse)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.OrdersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_OrdersResponse(buffer_arg) {
  return order_pb.OrdersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_SumNearestTwoMonthsAmountRequest(arg) {
  if (!(arg instanceof order_pb.SumNearestTwoMonthsAmountRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.SumNearestTwoMonthsAmountRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_SumNearestTwoMonthsAmountRequest(buffer_arg) {
  return order_pb.SumNearestTwoMonthsAmountRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_SumNearestTwoMonthsAmountResponse(arg) {
  if (!(arg instanceof order_pb.SumNearestTwoMonthsAmountResponse)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.SumNearestTwoMonthsAmountResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_SumNearestTwoMonthsAmountResponse(buffer_arg) {
  return order_pb.SumNearestTwoMonthsAmountResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_UpdateOrderProjectCostInfoRequest(arg) {
  if (!(arg instanceof order_pb.UpdateOrderProjectCostInfoRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.UpdateOrderProjectCostInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_UpdateOrderProjectCostInfoRequest(buffer_arg) {
  return order_pb.UpdateOrderProjectCostInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_paylabo_c002_harp_v1_UpdateOrderRequest(arg) {
  if (!(arg instanceof order_pb.UpdateOrderRequest)) {
    throw new Error('Expected argument of type com.paylabo.c002.harp.v1.UpdateOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_paylabo_c002_harp_v1_UpdateOrderRequest(buffer_arg) {
  return order_pb.UpdateOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var OrderService = exports.OrderService = {
  // *
// 今月未検収発注見出し一覧取得
listUnacceptedOrdersThisMonth: {
    path: '/com.paylabo.c002.harp.v1.Order/ListUnacceptedOrdersThisMonth',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.ListOrdersRequest,
    responseType: order_pb.OrdersResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_ListOrdersRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_ListOrdersRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_OrdersResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_OrdersResponse,
  },
  // *
// 発注見出し一覧取得
listOrders: {
    path: '/com.paylabo.c002.harp.v1.Order/ListOrders',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.ListOrdersRequest,
    responseType: order_pb.OrdersResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_ListOrdersRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_ListOrdersRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_OrdersResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_OrdersResponse,
  },
  // *
// 発注見出しカウント
countOrders: {
    path: '/com.paylabo.c002.harp.v1.Order/CountOrders',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.ListOrdersRequest,
    responseType: order_pb.CountOrdersResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_ListOrdersRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_ListOrdersRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_CountOrdersResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_CountOrdersResponse,
  },
  // *
// 発注見出し取得
getOrder: {
    path: '/com.paylabo.c002.harp.v1.Order/GetOrder',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.GetOrderRequest,
    responseType: order_pb.OrderResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_GetOrderRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_GetOrderRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_OrderResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_OrderResponse,
  },
  // *
// 発注見出し作成
createOrder: {
    path: '/com.paylabo.c002.harp.v1.Order/CreateOrder',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.CreateOrderRequest,
    responseType: order_pb.MutationOrderResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_CreateOrderRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_CreateOrderRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_MutationOrderResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_MutationOrderResponse,
  },
  // *
// 発注見出し更新
updateOrder: {
    path: '/com.paylabo.c002.harp.v1.Order/UpdateOrder',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.UpdateOrderRequest,
    responseType: order_pb.MutationOrderResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_UpdateOrderRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_UpdateOrderRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_MutationOrderResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_MutationOrderResponse,
  },
  // *
// 発注見出しのプロジェクト費用部分更新
updateOrderProjectCostInfo: {
    path: '/com.paylabo.c002.harp.v1.Order/UpdateOrderProjectCostInfo',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.UpdateOrderProjectCostInfoRequest,
    responseType: order_pb.MutationOrderResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_UpdateOrderProjectCostInfoRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_UpdateOrderProjectCostInfoRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_MutationOrderResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_MutationOrderResponse,
  },
  // *
// 発注見出し削除
deleteOrder: {
    path: '/com.paylabo.c002.harp.v1.Order/DeleteOrder',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.DeleteOrderRequest,
    responseType: order_pb.MutationOrderResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_DeleteOrderRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_DeleteOrderRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_MutationOrderResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_MutationOrderResponse,
  },
  getOrderRequesterAggregate: {
    path: '/com.paylabo.c002.harp.v1.Order/GetOrderRequesterAggregate',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.OrderAggregateRequest,
    responseType: order_pb.OrderAggregateResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_OrderAggregateRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_OrderAggregateRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_OrderAggregateResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_OrderAggregateResponse,
  },
  getOrderSupplierAggregate: {
    path: '/com.paylabo.c002.harp.v1.Order/GetOrderSupplierAggregate',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.OrderAggregateRequest,
    responseType: order_pb.OrderAggregateResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_OrderAggregateRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_OrderAggregateRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_OrderAggregateResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_OrderAggregateResponse,
  },
  // *
// 条件付きカウント
countOrdersWithGroupBy: {
    path: '/com.paylabo.c002.harp.v1.Order/CountOrdersWithGroupBy',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.OrdersGroupByRequest,
    responseType: order_pb.OrdersGroupByResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_OrdersGroupByRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_OrdersGroupByRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_OrdersGroupByResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_OrdersGroupByResponse,
  },
  // *
// 先月と今月の発注金額
sumNearestTwoMonthsAmount: {
    path: '/com.paylabo.c002.harp.v1.Order/SumNearestTwoMonthsAmount',
    requestStream: false,
    responseStream: false,
    requestType: order_pb.SumNearestTwoMonthsAmountRequest,
    responseType: order_pb.SumNearestTwoMonthsAmountResponse,
    requestSerialize: serialize_com_paylabo_c002_harp_v1_SumNearestTwoMonthsAmountRequest,
    requestDeserialize: deserialize_com_paylabo_c002_harp_v1_SumNearestTwoMonthsAmountRequest,
    responseSerialize: serialize_com_paylabo_c002_harp_v1_SumNearestTwoMonthsAmountResponse,
    responseDeserialize: deserialize_com_paylabo_c002_harp_v1_SumNearestTwoMonthsAmountResponse,
  },
};

exports.OrderClient = grpc.makeGenericClientConstructor(OrderService);
