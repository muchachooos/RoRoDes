import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { GET_AVAILABLE_USERS } from '../../utils/apiEndpointsConsts';


export const authAPI = createApi({
    reducerPath: 'auth/api',
    baseQuery: fetchBaseQuery({
        baseUrl: import.meta.env.VITE_BASE_URL,
        prepareHeaders(headers, api) {
            headers.set('Authorization', 'Token')
            return headers
        },
    }),
    
    endpoints: build => ({
        checkToken: build.query<{token: string}, {token: string}>({
            query: () => ({ url: GET_AVAILABLE_USERS }),
        })

    })
});
