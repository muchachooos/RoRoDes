import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { IAvaliableUsers } from './lobbyAPITypes';
import { GET_AVAILABLE_USERS } from '../../utils/apiEndpointsConsts';
import { RootState } from '../../app/store';


export const lobbyAPI = createApi({
    reducerPath: 'lobby/api',
    baseQuery: fetchBaseQuery({
        baseUrl: import.meta.env.VITE_BASE_URL,
        prepareHeaders(headers, api) {
            headers.set('Authorization', 'Token')
            //TODO: handle token
            return headers
        },
    }),
    
    endpoints: build => ({
        getAvaliableUsers: build.query<IAvaliableUsers, {}>({
            query: () => ({ url: GET_AVAILABLE_USERS }),
        })

    })
});
