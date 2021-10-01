import { gql } from 'apollo-server-koa'
import { GraphQLResolverMap, GraphQLSchemaModule } from 'apollo-graphql'
import { ordersTypeDef, ordersResolver } from './purchase/orders'
import { OrderRepository } from '../infra/repository/ordersRepository'

export const scalar = gql`
  scalar date
  scalar timestamp

  # response of any mutation
  type HarpMutationResponse {
    # number of affected rows by the mutation
    affectedRows: Int

    # data of the affected rows by the mutation
    # returning: [harp_orders!]!
  }
`

const returnOnError = <T>(executor: () => T, fallback: T): T => {
  try {
    return executor()
  } catch (e) {
    return fallback
  }
}

const parseDate = (value: unknown): Date | null =>
  returnOnError(() => (value == null ? null : new Date(value as never)), null)

export const configureGraphqlSchemaModules = (config: {
  orderRepository: OrderRepository
}): GraphQLSchemaModule[] => {
  const { orderRepository } = config

  return [
    {
      typeDefs: ordersTypeDef,
      resolvers: ordersResolver({
        orderRepository,
      }) as GraphQLResolverMap,
    },
  ]
}
