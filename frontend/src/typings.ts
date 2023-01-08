import type { NumberSchema, StringSchema } from "yup";

export type TypeOf<T> = T extends any ? T : never;

export type YupSchemaType<T> = T extends string
  ? StringSchema
  : T extends number
  ? NumberSchema
  : never;

export type YupSchemaTypeMap<T> = {
  [K in keyof T]: YupSchemaType<T[K]>;
};

export interface LoginPayload {
  identity: string;
  password: string;
}
