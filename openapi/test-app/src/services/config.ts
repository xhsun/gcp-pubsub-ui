// Ref: https://redux-toolkit.js.org/rtk-query/usage/code-splitting
// Need to use the React-specific entry point to allow generating React hooks
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

// initialize an empty api service that we'll inject endpoints into later as needed
// This is how we split API into multiple files
export const emptySplitApi = createApi({
    reducerPath: 'emptySplitApi',
    baseQuery: fetchBaseQuery({
        baseUrl: 'http://localhost:50051',
        prepareHeaders: (headers) => {
            headers.set("content-type", "application/json")
            return headers;
        },
    }),
    endpoints: () => ({}),
});