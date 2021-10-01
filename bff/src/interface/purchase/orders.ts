import { gql, IResolvers } from 'apollo-server-koa'
import { OrderRepository, Order } from '../../infra/repository/ordersRepository'

export const ordersTypeDef = gql`
  """
  order
  """
  input OrdersOrderByInput {
    id: SortEnum
  }

  """
  発注新規入力
  """
  input OrdersInsertInput {
    """
    name
    """
    name: string
  }

  """
  発注更新入力
  """
  input OrdersSetInput {
    """
    name
    """
    name: date
  }

  type Order @key(fields: "id") {
    """
    id
    """
    id: Int!
    """
    name
    """
    name: string
  }

  type OrdersMutationResponse {
    # number of affected rows by the mutation
    affectedRows: Int!
    # return data of order mutation
    returning: Order!
  }

  extend type Query {
    """
    order(one)
    """
    order(id: Int!): Order
    """
    orders
    """
    orders(ids: [Int]): [Order!]
  }

  extend type Mutation {
    """
    create
    """
    createOrder(object: OrdersInsertInput): OrdersMutationResponse
  }
`
export const ordersResolver = (params: { orderRepository: OrderRepository }): IResolvers => {
  const { orderRepository } = params
  return {
    Query: {
      orders: async (source, args): Promise<Order[] | undefined> =>
        await orderRepository.listOrders(args),
    },
    Mutation: {},
    Order: {},
  }
}
