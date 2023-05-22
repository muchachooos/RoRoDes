import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { IAvaliableUsers } from './lobbyAPITypes';
import { CREATE_GAME, GET_AVAILABLE_USERS } from '../../utils/apiEndpointsConsts';
import { RootState } from '../../app/store';


export const lobbyAPI = createApi({
    reducerPath: 'lobby/api',
    baseQuery: fetchBaseQuery({
        baseUrl: import.meta.env.VITE_BASE_URL,
        prepareHeaders(headers, api) {
            headers.set('Authorization', 'Token');
            //TODO: handle token
            headers.set('mode', "no-cors",
            );
            return headers;
        },
    }),

    endpoints: build => ({
        getAvaliableUsers: build.query<IAvaliableUsers, {}>({
            query: () => ({ url: GET_AVAILABLE_USERS }),
        }),

        createGame: build.mutation<string, undefined>({
            query: () => ({
                url: CREATE_GAME,
                method: 'POST',
            }),
        })

    })
});
