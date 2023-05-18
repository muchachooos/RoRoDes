import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import counterReducer from "../features/counter/counterSlice";
import { lobbyAPI } from "../features/lobby/lobbyAPI";
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';
import { authAPI } from "../features/auth/authAPI";
import { gameAPI } from "../features/game/gameAPI";


export const store = configureStore({
  reducer: {
    [lobbyAPI.reducerPath]: lobbyAPI.reducer,
    [authAPI.reducerPath]: authAPI.reducer,
    [gameAPI.reducerPath]: gameAPI.reducer,
    counter: counterReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(lobbyAPI.middleware).concat(authAPI.middleware).concat(lobbyAPI.middleware)
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;