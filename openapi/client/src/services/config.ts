/*
  This file includes a "dummy" client to allow the Typescript compiler to work. In order to import this package,
  you must have an alias for @services/config in your app that exports the emptySplitApi constant.
 */
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

// initialize an empty api service that we'll inject endpoints into later as needed
export const emptySplitApi = createApi({
    baseQuery: fetchBaseQuery({ baseUrl: '/' }),
    endpoints: () => ({}),
})