import { ServiceError, status } from 'grpc'
import { GraphQLError, GraphQLFormattedError } from 'graphql'
import { UserInputError } from 'apollo-server-koa'
import { BadRequest, deserializeGoogleGrpcStatusDetails } from '@stackpath/node-grpc-error-details'
import { logger } from '../common/logger'

const isServiceError = (error: unknown): error is ServiceError =>
  error !== undefined &&
  error !== null &&
  typeof error === 'object' &&
  typeof (error as { code: unknown }).code === 'number' &&
  typeof (error as { metadata: unknown }).metadata === 'object' &&
  typeof (error as { details: unknown }).details === 'string'

const handleGrpcError = (
  error: ServiceError,
  wrappingError: GraphQLError
): GraphQLFormattedError => {
  if (error.code === status.INVALID_ARGUMENT) {
    const grpcErrorDetails = deserializeGoogleGrpcStatusDetails(error)
    const badRequest = grpcErrorDetails?.details.find((v) => v instanceof BadRequest) as BadRequest
    const validationErrors = badRequest
      .getFieldViolationsList()
      .reduce<Record<string, string>>((acc, x) => {
        acc[x.getField()] = x.getDescription()
        return acc
      }, {})
    return new UserInputError('validation error', { validationErrors })
  }
  return wrappingError
}

export const formatError = (error: GraphQLError): GraphQLFormattedError => {
  if (isServiceError(error.originalError)) return handleGrpcError(error.originalError, error)
  logger.error('graphql error', { error })
  return error
}
