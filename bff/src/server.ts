import Koa from 'koa'
import { ApolloServer } from 'apollo-server-koa'
import { buildFederatedSchema } from '@apollo/federation'
import { credentials } from 'grpc'
import { logger } from './common/logger'
import { envConfig } from './common/config'
import { configureGraphqlSchemaModules } from './interface'
import { formatError } from './interface/errorHandler'
import { OrderRepository } from './infra/repository/ordersRepository'
import { OrderClient } from './infra/proto/order_grpc_pb'

const { BACKEND_SERVER_ADDRESS, BFF_LISTEN_PORT } = envConfig

const orderRepository = new OrderRepository({
  client: new OrderClient(BACKEND_SERVER_ADDRESS, credentials.createInsecure()),
})

const server = new ApolloServer({
  schema: buildFederatedSchema(
    configureGraphqlSchemaModules({
      orderRepository,
    })
  ),
  formatError,
})
const app = new Koa()
server.applyMiddleware({ app })
const port = BFF_LISTEN_PORT
app.listen({ port }, () => logger.info(`server start at [localhost:${port}]`))
