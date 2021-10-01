// package: com.paylabo.c002.harp.v1
// file: order.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as enums_pb from "./enums_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";

export class OrderResponse extends jspb.Message { 

    hasId(): boolean;
    clearId(): void;
    getId(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setId(value?: google_protobuf_wrappers_pb.Int32Value): OrderResponse;

    getOrderNo(): string;
    setOrderNo(value: string): OrderResponse;


    hasSuppliersId(): boolean;
    clearSuppliersId(): void;
    getSuppliersId(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setSuppliersId(value?: google_protobuf_wrappers_pb.Int32Value): OrderResponse;

    getCompanyGroupType(): string;
    setCompanyGroupType(value: string): OrderResponse;

    getSubject(): string;
    setSubject(value: string): OrderResponse;

    getRequestOrganizationId(): string;
    setRequestOrganizationId(value: string): OrderResponse;


    hasRequestDate(): boolean;
    clearRequestDate(): void;
    getRequestDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setRequestDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;

    getRequestBy(): string;
    setRequestBy(value: string): OrderResponse;

    getApprovalFile(): string;
    setApprovalFile(value: string): OrderResponse;


    hasDerivationSourceOrderId(): boolean;
    clearDerivationSourceOrderId(): void;
    getDerivationSourceOrderId(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setDerivationSourceOrderId(value?: google_protobuf_wrappers_pb.Int32Value): OrderResponse;

    getRemarks(): string;
    setRemarks(value: string): OrderResponse;


    hasSuperiorApprovalDate(): boolean;
    clearSuperiorApprovalDate(): void;
    getSuperiorApprovalDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setSuperiorApprovalDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;


    hasPurchasingDeptApprovalDate(): boolean;
    clearPurchasingDeptApprovalDate(): void;
    getPurchasingDeptApprovalDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setPurchasingDeptApprovalDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;


    hasOrderIssueDate(): boolean;
    clearOrderIssueDate(): void;
    getOrderIssueDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setOrderIssueDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;


    hasFinalAcceptanceDate(): boolean;
    clearFinalAcceptanceDate(): void;
    getFinalAcceptanceDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setFinalAcceptanceDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;


    hasAcceptanceCompletedDate(): boolean;
    clearAcceptanceCompletedDate(): void;
    getAcceptanceCompletedDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setAcceptanceCompletedDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;


    hasCancelDate(): boolean;
    clearCancelDate(): void;
    getCancelDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCancelDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;


    hasOrderCaseCd(): boolean;
    clearOrderCaseCd(): void;
    getOrderCaseCd(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setOrderCaseCd(value?: google_protobuf_wrappers_pb.Int32Value): OrderResponse;


    hasOrderStatus(): boolean;
    clearOrderStatus(): void;
    getOrderStatus(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setOrderStatus(value?: google_protobuf_wrappers_pb.Int32Value): OrderResponse;

    getJiraNo(): string;
    setJiraNo(value: string): OrderResponse;


    hasQuotationsId(): boolean;
    clearQuotationsId(): void;
    getQuotationsId(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setQuotationsId(value?: google_protobuf_wrappers_pb.Int32Value): OrderResponse;

    getOrderApprovalStaffsId(): string;
    setOrderApprovalStaffsId(value: string): OrderResponse;


    hasUpdatedAt(): boolean;
    clearUpdatedAt(): void;
    getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;

    getProjectsId(): string;
    setProjectsId(value: string): OrderResponse;

    getProjectCostId(): string;
    setProjectCostId(value: string): OrderResponse;

    getCostTypesId(): string;
    setCostTypesId(value: string): OrderResponse;


    hasCostingStartDate(): boolean;
    clearCostingStartDate(): void;
    getCostingStartDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCostingStartDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;


    hasCostingEndDate(): boolean;
    clearCostingEndDate(): void;
    getCostingEndDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCostingEndDate(value?: google_protobuf_timestamp_pb.Timestamp): OrderResponse;


    hasCostingPolicy(): boolean;
    clearCostingPolicy(): void;
    getCostingPolicy(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setCostingPolicy(value?: google_protobuf_wrappers_pb.Int32Value): OrderResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OrderResponse.AsObject;
    static toObject(includeInstance: boolean, msg: OrderResponse): OrderResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OrderResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OrderResponse;
    static deserializeBinaryFromReader(message: OrderResponse, reader: jspb.BinaryReader): OrderResponse;
}

export namespace OrderResponse {
    export type AsObject = {
        id?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        orderNo: string,
        suppliersId?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        companyGroupType: string,
        subject: string,
        requestOrganizationId: string,
        requestDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        requestBy: string,
        approvalFile: string,
        derivationSourceOrderId?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        remarks: string,
        superiorApprovalDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        purchasingDeptApprovalDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        orderIssueDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        finalAcceptanceDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        acceptanceCompletedDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        cancelDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        orderCaseCd?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        orderStatus?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        jiraNo: string,
        quotationsId?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        orderApprovalStaffsId: string,
        updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        projectsId: string,
        projectCostId: string,
        costTypesId: string,
        costingStartDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        costingEndDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        costingPolicy?: google_protobuf_wrappers_pb.Int32Value.AsObject,
    }
}

export class ListOrdersRequest extends jspb.Message { 
    clearIdsList(): void;
    getIdsList(): Array<number>;
    setIdsList(value: Array<number>): ListOrdersRequest;
    addIds(value: number, index?: number): number;

    clearSuppliersIdsList(): void;
    getSuppliersIdsList(): Array<number>;
    setSuppliersIdsList(value: Array<number>): ListOrdersRequest;
    addSuppliersIds(value: number, index?: number): number;

    getSubject(): string;
    setSubject(value: string): ListOrdersRequest;

    clearRequestOrganizationIdsList(): void;
    getRequestOrganizationIdsList(): Array<string>;
    setRequestOrganizationIdsList(value: Array<string>): ListOrdersRequest;
    addRequestOrganizationIds(value: string, index?: number): string;

    clearRequestBysList(): void;
    getRequestBysList(): Array<string>;
    setRequestBysList(value: Array<string>): ListOrdersRequest;
    addRequestBys(value: string, index?: number): string;

    clearOrderCaseCdsList(): void;
    getOrderCaseCdsList(): Array<number>;
    setOrderCaseCdsList(value: Array<number>): ListOrdersRequest;
    addOrderCaseCds(value: number, index?: number): number;

    clearOrderStatusesList(): void;
    getOrderStatusesList(): Array<number>;
    setOrderStatusesList(value: Array<number>): ListOrdersRequest;
    addOrderStatuses(value: number, index?: number): number;

    clearProjectsIdsList(): void;
    getProjectsIdsList(): Array<string>;
    setProjectsIdsList(value: Array<string>): ListOrdersRequest;
    addProjectsIds(value: string, index?: number): string;

    clearProjectCostIdsList(): void;
    getProjectCostIdsList(): Array<string>;
    setProjectCostIdsList(value: Array<string>): ListOrdersRequest;
    addProjectCostIds(value: string, index?: number): string;

    getLimit(): number;
    setLimit(value: number): ListOrdersRequest;

    getOffset(): number;
    setOffset(value: number): ListOrdersRequest;


    hasOrderBy(): boolean;
    clearOrderBy(): void;
    getOrderBy(): ListOrdersRequestOrderBy | undefined;
    setOrderBy(value?: ListOrdersRequestOrderBy): ListOrdersRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListOrdersRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListOrdersRequest): ListOrdersRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListOrdersRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListOrdersRequest;
    static deserializeBinaryFromReader(message: ListOrdersRequest, reader: jspb.BinaryReader): ListOrdersRequest;
}

export namespace ListOrdersRequest {
    export type AsObject = {
        idsList: Array<number>,
        suppliersIdsList: Array<number>,
        subject: string,
        requestOrganizationIdsList: Array<string>,
        requestBysList: Array<string>,
        orderCaseCdsList: Array<number>,
        orderStatusesList: Array<number>,
        projectsIdsList: Array<string>,
        projectCostIdsList: Array<string>,
        limit: number,
        offset: number,
        orderBy?: ListOrdersRequestOrderBy.AsObject,
    }
}

export class ListOrdersRequestOrderBy extends jspb.Message { 
    getId(): enums_pb.SortEnum;
    setId(value: enums_pb.SortEnum): ListOrdersRequestOrderBy;

    getSuppliersId(): enums_pb.SortEnum;
    setSuppliersId(value: enums_pb.SortEnum): ListOrdersRequestOrderBy;

    getOrderStatus(): enums_pb.SortEnum;
    setOrderStatus(value: enums_pb.SortEnum): ListOrdersRequestOrderBy;

    getOrderCaseCd(): enums_pb.SortEnum;
    setOrderCaseCd(value: enums_pb.SortEnum): ListOrdersRequestOrderBy;

    getSubject(): enums_pb.SortEnum;
    setSubject(value: enums_pb.SortEnum): ListOrdersRequestOrderBy;

    getRequestDate(): enums_pb.SortEnum;
    setRequestDate(value: enums_pb.SortEnum): ListOrdersRequestOrderBy;

    getFinalAcceptanceDate(): enums_pb.SortEnum;
    setFinalAcceptanceDate(value: enums_pb.SortEnum): ListOrdersRequestOrderBy;

    getUpdatedAt(): enums_pb.SortEnum;
    setUpdatedAt(value: enums_pb.SortEnum): ListOrdersRequestOrderBy;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListOrdersRequestOrderBy.AsObject;
    static toObject(includeInstance: boolean, msg: ListOrdersRequestOrderBy): ListOrdersRequestOrderBy.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListOrdersRequestOrderBy, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListOrdersRequestOrderBy;
    static deserializeBinaryFromReader(message: ListOrdersRequestOrderBy, reader: jspb.BinaryReader): ListOrdersRequestOrderBy;
}

export namespace ListOrdersRequestOrderBy {
    export type AsObject = {
        id: enums_pb.SortEnum,
        suppliersId: enums_pb.SortEnum,
        orderStatus: enums_pb.SortEnum,
        orderCaseCd: enums_pb.SortEnum,
        subject: enums_pb.SortEnum,
        requestDate: enums_pb.SortEnum,
        finalAcceptanceDate: enums_pb.SortEnum,
        updatedAt: enums_pb.SortEnum,
    }
}

export class OrdersResponse extends jspb.Message { 
    clearOrdersList(): void;
    getOrdersList(): Array<OrderResponse>;
    setOrdersList(value: Array<OrderResponse>): OrdersResponse;
    addOrders(value?: OrderResponse, index?: number): OrderResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OrdersResponse.AsObject;
    static toObject(includeInstance: boolean, msg: OrdersResponse): OrdersResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OrdersResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OrdersResponse;
    static deserializeBinaryFromReader(message: OrdersResponse, reader: jspb.BinaryReader): OrdersResponse;
}

export namespace OrdersResponse {
    export type AsObject = {
        ordersList: Array<OrderResponse.AsObject>,
    }
}

export class GetOrderRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): GetOrderRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetOrderRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetOrderRequest): GetOrderRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetOrderRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetOrderRequest;
    static deserializeBinaryFromReader(message: GetOrderRequest, reader: jspb.BinaryReader): GetOrderRequest;
}

export namespace GetOrderRequest {
    export type AsObject = {
        id: string,
    }
}

export class CreateOrderRequest extends jspb.Message { 
    getOrderNo(): string;
    setOrderNo(value: string): CreateOrderRequest;

    getSuppliersId(): string;
    setSuppliersId(value: string): CreateOrderRequest;

    getCompanyGroupType(): string;
    setCompanyGroupType(value: string): CreateOrderRequest;

    getSubject(): string;
    setSubject(value: string): CreateOrderRequest;

    getRequestOrganizationId(): string;
    setRequestOrganizationId(value: string): CreateOrderRequest;


    hasRequestDate(): boolean;
    clearRequestDate(): void;
    getRequestDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setRequestDate(value?: google_protobuf_timestamp_pb.Timestamp): CreateOrderRequest;

    getRequestBy(): string;
    setRequestBy(value: string): CreateOrderRequest;

    getApprovalFile(): string;
    setApprovalFile(value: string): CreateOrderRequest;


    hasDerivationSourceOrderId(): boolean;
    clearDerivationSourceOrderId(): void;
    getDerivationSourceOrderId(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setDerivationSourceOrderId(value?: google_protobuf_wrappers_pb.Int32Value): CreateOrderRequest;

    getRemarks(): string;
    setRemarks(value: string): CreateOrderRequest;


    hasSuperiorApprovalDate(): boolean;
    clearSuperiorApprovalDate(): void;
    getSuperiorApprovalDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setSuperiorApprovalDate(value?: google_protobuf_timestamp_pb.Timestamp): CreateOrderRequest;


    hasPurchasingDeptApprovalDate(): boolean;
    clearPurchasingDeptApprovalDate(): void;
    getPurchasingDeptApprovalDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setPurchasingDeptApprovalDate(value?: google_protobuf_timestamp_pb.Timestamp): CreateOrderRequest;


    hasOrderIssueDate(): boolean;
    clearOrderIssueDate(): void;
    getOrderIssueDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setOrderIssueDate(value?: google_protobuf_timestamp_pb.Timestamp): CreateOrderRequest;


    hasFinalAcceptanceDate(): boolean;
    clearFinalAcceptanceDate(): void;
    getFinalAcceptanceDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setFinalAcceptanceDate(value?: google_protobuf_timestamp_pb.Timestamp): CreateOrderRequest;


    hasAcceptanceCompletedDate(): boolean;
    clearAcceptanceCompletedDate(): void;
    getAcceptanceCompletedDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setAcceptanceCompletedDate(value?: google_protobuf_timestamp_pb.Timestamp): CreateOrderRequest;


    hasCancelDate(): boolean;
    clearCancelDate(): void;
    getCancelDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCancelDate(value?: google_protobuf_timestamp_pb.Timestamp): CreateOrderRequest;

    getOrderCaseCd(): string;
    setOrderCaseCd(value: string): CreateOrderRequest;

    getOrderStatus(): string;
    setOrderStatus(value: string): CreateOrderRequest;

    getJiraNo(): string;
    setJiraNo(value: string): CreateOrderRequest;


    hasQuotationsId(): boolean;
    clearQuotationsId(): void;
    getQuotationsId(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setQuotationsId(value?: google_protobuf_wrappers_pb.Int32Value): CreateOrderRequest;

    getOrderApprovalStaffsId(): string;
    setOrderApprovalStaffsId(value: string): CreateOrderRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateOrderRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateOrderRequest): CreateOrderRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateOrderRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateOrderRequest;
    static deserializeBinaryFromReader(message: CreateOrderRequest, reader: jspb.BinaryReader): CreateOrderRequest;
}

export namespace CreateOrderRequest {
    export type AsObject = {
        orderNo: string,
        suppliersId: string,
        companyGroupType: string,
        subject: string,
        requestOrganizationId: string,
        requestDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        requestBy: string,
        approvalFile: string,
        derivationSourceOrderId?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        remarks: string,
        superiorApprovalDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        purchasingDeptApprovalDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        orderIssueDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        finalAcceptanceDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        acceptanceCompletedDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        cancelDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        orderCaseCd: string,
        orderStatus: string,
        jiraNo: string,
        quotationsId?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        orderApprovalStaffsId: string,
    }
}

export class UpdateOrderRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UpdateOrderRequest;

    getOrderNo(): string;
    setOrderNo(value: string): UpdateOrderRequest;

    getSuppliersId(): string;
    setSuppliersId(value: string): UpdateOrderRequest;

    getCompanyGroupType(): string;
    setCompanyGroupType(value: string): UpdateOrderRequest;

    getSubject(): string;
    setSubject(value: string): UpdateOrderRequest;

    getRequestOrganizationId(): string;
    setRequestOrganizationId(value: string): UpdateOrderRequest;


    hasRequestDate(): boolean;
    clearRequestDate(): void;
    getRequestDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setRequestDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderRequest;

    getRequestBy(): string;
    setRequestBy(value: string): UpdateOrderRequest;

    getApprovalFile(): string;
    setApprovalFile(value: string): UpdateOrderRequest;


    hasDerivationSourceOrderId(): boolean;
    clearDerivationSourceOrderId(): void;
    getDerivationSourceOrderId(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setDerivationSourceOrderId(value?: google_protobuf_wrappers_pb.Int32Value): UpdateOrderRequest;

    getRemarks(): string;
    setRemarks(value: string): UpdateOrderRequest;


    hasSuperiorApprovalDate(): boolean;
    clearSuperiorApprovalDate(): void;
    getSuperiorApprovalDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setSuperiorApprovalDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderRequest;


    hasPurchasingDeptApprovalDate(): boolean;
    clearPurchasingDeptApprovalDate(): void;
    getPurchasingDeptApprovalDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setPurchasingDeptApprovalDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderRequest;


    hasOrderIssueDate(): boolean;
    clearOrderIssueDate(): void;
    getOrderIssueDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setOrderIssueDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderRequest;


    hasFinalAcceptanceDate(): boolean;
    clearFinalAcceptanceDate(): void;
    getFinalAcceptanceDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setFinalAcceptanceDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderRequest;


    hasAcceptanceCompletedDate(): boolean;
    clearAcceptanceCompletedDate(): void;
    getAcceptanceCompletedDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setAcceptanceCompletedDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderRequest;


    hasCancelDate(): boolean;
    clearCancelDate(): void;
    getCancelDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCancelDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderRequest;

    getOrderCaseCd(): string;
    setOrderCaseCd(value: string): UpdateOrderRequest;

    getOrderStatus(): string;
    setOrderStatus(value: string): UpdateOrderRequest;

    getJiraNo(): string;
    setJiraNo(value: string): UpdateOrderRequest;


    hasQuotationsId(): boolean;
    clearQuotationsId(): void;
    getQuotationsId(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setQuotationsId(value?: google_protobuf_wrappers_pb.Int32Value): UpdateOrderRequest;

    getOrderApprovalStaffsId(): string;
    setOrderApprovalStaffsId(value: string): UpdateOrderRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateOrderRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateOrderRequest): UpdateOrderRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateOrderRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateOrderRequest;
    static deserializeBinaryFromReader(message: UpdateOrderRequest, reader: jspb.BinaryReader): UpdateOrderRequest;
}

export namespace UpdateOrderRequest {
    export type AsObject = {
        id: string,
        orderNo: string,
        suppliersId: string,
        companyGroupType: string,
        subject: string,
        requestOrganizationId: string,
        requestDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        requestBy: string,
        approvalFile: string,
        derivationSourceOrderId?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        remarks: string,
        superiorApprovalDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        purchasingDeptApprovalDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        orderIssueDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        finalAcceptanceDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        acceptanceCompletedDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        cancelDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        orderCaseCd: string,
        orderStatus: string,
        jiraNo: string,
        quotationsId?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        orderApprovalStaffsId: string,
    }
}

export class UpdateOrderProjectCostInfoRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UpdateOrderProjectCostInfoRequest;

    getProjectsId(): string;
    setProjectsId(value: string): UpdateOrderProjectCostInfoRequest;

    getProjectCostId(): string;
    setProjectCostId(value: string): UpdateOrderProjectCostInfoRequest;

    getCostTypesId(): string;
    setCostTypesId(value: string): UpdateOrderProjectCostInfoRequest;


    hasCostingStartDate(): boolean;
    clearCostingStartDate(): void;
    getCostingStartDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCostingStartDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderProjectCostInfoRequest;


    hasCostingEndDate(): boolean;
    clearCostingEndDate(): void;
    getCostingEndDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCostingEndDate(value?: google_protobuf_timestamp_pb.Timestamp): UpdateOrderProjectCostInfoRequest;


    hasCostingPolicy(): boolean;
    clearCostingPolicy(): void;
    getCostingPolicy(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setCostingPolicy(value?: google_protobuf_wrappers_pb.Int32Value): UpdateOrderProjectCostInfoRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateOrderProjectCostInfoRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateOrderProjectCostInfoRequest): UpdateOrderProjectCostInfoRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateOrderProjectCostInfoRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateOrderProjectCostInfoRequest;
    static deserializeBinaryFromReader(message: UpdateOrderProjectCostInfoRequest, reader: jspb.BinaryReader): UpdateOrderProjectCostInfoRequest;
}

export namespace UpdateOrderProjectCostInfoRequest {
    export type AsObject = {
        id: string,
        projectsId: string,
        projectCostId: string,
        costTypesId: string,
        costingStartDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        costingEndDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        costingPolicy?: google_protobuf_wrappers_pb.Int32Value.AsObject,
    }
}

export class DeleteOrderRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): DeleteOrderRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteOrderRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteOrderRequest): DeleteOrderRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteOrderRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteOrderRequest;
    static deserializeBinaryFromReader(message: DeleteOrderRequest, reader: jspb.BinaryReader): DeleteOrderRequest;
}

export namespace DeleteOrderRequest {
    export type AsObject = {
        id: string,
    }
}

export class MutationOrderResponse extends jspb.Message { 
    getAffectedRows(): number;
    setAffectedRows(value: number): MutationOrderResponse;


    hasOrder(): boolean;
    clearOrder(): void;
    getOrder(): OrderResponse | undefined;
    setOrder(value?: OrderResponse): MutationOrderResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MutationOrderResponse.AsObject;
    static toObject(includeInstance: boolean, msg: MutationOrderResponse): MutationOrderResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MutationOrderResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MutationOrderResponse;
    static deserializeBinaryFromReader(message: MutationOrderResponse, reader: jspb.BinaryReader): MutationOrderResponse;
}

export namespace MutationOrderResponse {
    export type AsObject = {
        affectedRows: number,
        order?: OrderResponse.AsObject,
    }
}

export class CountOrdersResponse extends jspb.Message { 
    getCount(): number;
    setCount(value: number): CountOrdersResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CountOrdersResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CountOrdersResponse): CountOrdersResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CountOrdersResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CountOrdersResponse;
    static deserializeBinaryFromReader(message: CountOrdersResponse, reader: jspb.BinaryReader): CountOrdersResponse;
}

export namespace CountOrdersResponse {
    export type AsObject = {
        count: number,
    }
}

export class OrderAggregateRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OrderAggregateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: OrderAggregateRequest): OrderAggregateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OrderAggregateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OrderAggregateRequest;
    static deserializeBinaryFromReader(message: OrderAggregateRequest, reader: jspb.BinaryReader): OrderAggregateRequest;
}

export namespace OrderAggregateRequest {
    export type AsObject = {
    }
}

export class OrderAggregateResponse extends jspb.Message { 
    getSum(): number;
    setSum(value: number): OrderAggregateResponse;

    getCount(): number;
    setCount(value: number): OrderAggregateResponse;

    clearNodesList(): void;
    getNodesList(): Array<OrderResponse>;
    setNodesList(value: Array<OrderResponse>): OrderAggregateResponse;
    addNodes(value?: OrderResponse, index?: number): OrderResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OrderAggregateResponse.AsObject;
    static toObject(includeInstance: boolean, msg: OrderAggregateResponse): OrderAggregateResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OrderAggregateResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OrderAggregateResponse;
    static deserializeBinaryFromReader(message: OrderAggregateResponse, reader: jspb.BinaryReader): OrderAggregateResponse;
}

export namespace OrderAggregateResponse {
    export type AsObject = {
        sum: number,
        count: number,
        nodesList: Array<OrderResponse.AsObject>,
    }
}

export class OrdersGroupByRequest extends jspb.Message { 
    getRequestOrganizationId(): string;
    setRequestOrganizationId(value: string): OrdersGroupByRequest;

    getRequestBy(): string;
    setRequestBy(value: string): OrdersGroupByRequest;

    getRecentMonth(): number;
    setRecentMonth(value: number): OrdersGroupByRequest;

    getGroupBy(): string;
    setGroupBy(value: string): OrdersGroupByRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OrdersGroupByRequest.AsObject;
    static toObject(includeInstance: boolean, msg: OrdersGroupByRequest): OrdersGroupByRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OrdersGroupByRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OrdersGroupByRequest;
    static deserializeBinaryFromReader(message: OrdersGroupByRequest, reader: jspb.BinaryReader): OrdersGroupByRequest;
}

export namespace OrdersGroupByRequest {
    export type AsObject = {
        requestOrganizationId: string,
        requestBy: string,
        recentMonth: number,
        groupBy: string,
    }
}

export class OrdersGroupBy extends jspb.Message { 
    getCount(): number;
    setCount(value: number): OrdersGroupBy;

    getValue(): string;
    setValue(value: string): OrdersGroupBy;

    getGroupBy(): string;
    setGroupBy(value: string): OrdersGroupBy;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OrdersGroupBy.AsObject;
    static toObject(includeInstance: boolean, msg: OrdersGroupBy): OrdersGroupBy.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OrdersGroupBy, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OrdersGroupBy;
    static deserializeBinaryFromReader(message: OrdersGroupBy, reader: jspb.BinaryReader): OrdersGroupBy;
}

export namespace OrdersGroupBy {
    export type AsObject = {
        count: number,
        value: string,
        groupBy: string,
    }
}

export class OrdersGroupByResponse extends jspb.Message { 
    clearOrdersGroupByList(): void;
    getOrdersGroupByList(): Array<OrdersGroupBy>;
    setOrdersGroupByList(value: Array<OrdersGroupBy>): OrdersGroupByResponse;
    addOrdersGroupBy(value?: OrdersGroupBy, index?: number): OrdersGroupBy;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OrdersGroupByResponse.AsObject;
    static toObject(includeInstance: boolean, msg: OrdersGroupByResponse): OrdersGroupByResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OrdersGroupByResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OrdersGroupByResponse;
    static deserializeBinaryFromReader(message: OrdersGroupByResponse, reader: jspb.BinaryReader): OrdersGroupByResponse;
}

export namespace OrdersGroupByResponse {
    export type AsObject = {
        ordersGroupByList: Array<OrdersGroupBy.AsObject>,
    }
}

export class SumNearestTwoMonthsAmountRequest extends jspb.Message { 
    getRequestOrganizationId(): string;
    setRequestOrganizationId(value: string): SumNearestTwoMonthsAmountRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SumNearestTwoMonthsAmountRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SumNearestTwoMonthsAmountRequest): SumNearestTwoMonthsAmountRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SumNearestTwoMonthsAmountRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SumNearestTwoMonthsAmountRequest;
    static deserializeBinaryFromReader(message: SumNearestTwoMonthsAmountRequest, reader: jspb.BinaryReader): SumNearestTwoMonthsAmountRequest;
}

export namespace SumNearestTwoMonthsAmountRequest {
    export type AsObject = {
        requestOrganizationId: string,
    }
}

export class NearestTwoMonthsAmount extends jspb.Message { 
    getDate(): string;
    setDate(value: string): NearestTwoMonthsAmount;

    getMonthSum(): number;
    setMonthSum(value: number): NearestTwoMonthsAmount;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NearestTwoMonthsAmount.AsObject;
    static toObject(includeInstance: boolean, msg: NearestTwoMonthsAmount): NearestTwoMonthsAmount.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NearestTwoMonthsAmount, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NearestTwoMonthsAmount;
    static deserializeBinaryFromReader(message: NearestTwoMonthsAmount, reader: jspb.BinaryReader): NearestTwoMonthsAmount;
}

export namespace NearestTwoMonthsAmount {
    export type AsObject = {
        date: string,
        monthSum: number,
    }
}

export class SumNearestTwoMonthsAmountResponse extends jspb.Message { 
    clearNearestTwoMonthsAmountList(): void;
    getNearestTwoMonthsAmountList(): Array<NearestTwoMonthsAmount>;
    setNearestTwoMonthsAmountList(value: Array<NearestTwoMonthsAmount>): SumNearestTwoMonthsAmountResponse;
    addNearestTwoMonthsAmount(value?: NearestTwoMonthsAmount, index?: number): NearestTwoMonthsAmount;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SumNearestTwoMonthsAmountResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SumNearestTwoMonthsAmountResponse): SumNearestTwoMonthsAmountResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SumNearestTwoMonthsAmountResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SumNearestTwoMonthsAmountResponse;
    static deserializeBinaryFromReader(message: SumNearestTwoMonthsAmountResponse, reader: jspb.BinaryReader): SumNearestTwoMonthsAmountResponse;
}

export namespace SumNearestTwoMonthsAmountResponse {
    export type AsObject = {
        nearestTwoMonthsAmountList: Array<NearestTwoMonthsAmount.AsObject>,
    }
}
