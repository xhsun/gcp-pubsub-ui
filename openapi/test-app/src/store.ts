import { configureStore } from '@reduxjs/toolkit'
import {pubsubui} from './client'

export const store = configureStore({
    reducer: {
        [pubsubui.reducerPath]: pubsubui.reducer
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(pubsubui.middleware),
})

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch