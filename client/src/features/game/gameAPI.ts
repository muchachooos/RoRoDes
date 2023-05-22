
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { GET_GAME_FIELD } from '../../utils/apiEndpointsConsts';
import { IField } from './gameAPITypes';


export const gameAPI = createApi({
    reducerPath: 'game/api',
    baseQuery: fetchBaseQuery({
        baseUrl: import.meta.env.VITE_BASE_URL,
        prepareHeaders(headers, api) {
            headers.set('Authorization', 'Token')
            return headers
        },
    }),
    
    endpoints: build => ({
        getGameFields: build.query<IField, string>({
            query: () => ({ url: GET_GAME_FIELD }),
        })

    })
});


