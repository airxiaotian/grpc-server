import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb'
import {
  BoolValue,
  DoubleValue,
  FloatValue,
  Int32Value,
  Int64Value,
  StringValue,
} from 'google-protobuf/google/protobuf/wrappers_pb'
import { SortEnum } from './proto/purchase/enums_pb'

export type HarpMutationResponse = {
  affectedRows: number
}

export const convertMutationResponse = (response: any): HarpMutationResponse => {
  return {
    affectedRows: response && response.getAffectedRows(),
  }
}

export const convertSortEnum = (orderByString: string): SortEnum => {
  switch (orderByString) {
    case 'asc':
      return SortEnum.ASC
    case 'desc':
      return SortEnum.DESC
    default:
      return SortEnum.NONE
  }
}

export const toStringArray = (ns: number[]): string[] => ns.map(String)
export const toNumberArray = (ns: string[]): number[] => ns.map(Number)
export const toInt32Value = (s: string | number | undefined): Int32Value | undefined => {
  if (s === undefined) return undefined
  const v = new Int32Value()
  v.setValue(Number(s))
  return v
}
export const toInt64Value = (s: string | number | undefined): Int64Value | undefined => {
  if (s === undefined) return undefined
  const v = new Int64Value()
  v.setValue(Number(s))
  return v
}
export const toFloatValue = (s: string | number | undefined): FloatValue | undefined => {
  if (s === undefined) return undefined
  const v = new FloatValue()
  v.setValue(Number(s))
  return v
}

export const toDoubleValue = (s: string | number | undefined): DoubleValue | undefined => {
  if (s === undefined) return undefined
  const v = new DoubleValue()
  v.setValue(Number(s))
  return v
}

export const toStringValue = (s: string | undefined): StringValue | undefined => {
  if (s === undefined) return undefined
  const v = new StringValue()
  v.setValue(s)
  return v
}
export const toBoolValue = (b: boolean | undefined): BoolValue | undefined => {
  if (b === undefined) return undefined
  const v = new BoolValue()
  v.setValue(b)
  return v
}
export const toTimestamp = (d: Date | undefined): Timestamp | undefined => {
  if (!d) return undefined
  const t = new Timestamp()
  t.fromDate(d)
  return t
}
