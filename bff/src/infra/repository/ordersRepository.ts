import { HarpMutationResponse } from '..'
import { IOrderClient } from '../proto/order_grpc_pb'

import { ListOrdersRequest } from '../proto/order_pb'

export type Order = {
  id: string
  name: string
}

type ListOrdersParams = {
  ids?: number
}

type CreateOrderParams = {
  object: {
    id: string
    name: string
  }
}

export interface OrderRepository {
  listOrders(params?: ListOrdersParams): Promise<Order[]>
  createOrder(params: CreateOrderParams): Promise<MutationResponseWithReturning | null>
}

type MutationResponseWithReturning = HarpMutationResponse & { returning: Order | undefined }

export class OrderRepository implements OrderRepository {
  private readonly client: IOrderClient

  constructor(params: { client: IOrderClient }) {
    const { client } = params
    this.client = client
  }

  public listOrders(params: ListOrdersParams = {}): Promise<Order[] | undefined> {
    const request = new ListOrdersRequest()
    if (params.ids) request.addIds(params.ids)

    return new Promise((resolve, reject) => {
      this.client.listOrders(request, (error, response) => {
        if (error) reject(error)
        else resolve(undefined) // result
      })
    })
  }
}
