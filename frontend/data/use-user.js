import useSWR from "swr";
import {getUserInfo} from "../libs/api-user";
import jsCookie from "js-cookie";

export function useUser() {
    const {data, mutate, error} = useSWR('api_user', getUserInfo, {
        revalidateIfStale: false,
        revalidateOnFocus: false,
        revalidateOnReconnect: false
    })
    const loading = !data && !error
    const loggedOut = typeof window !== "undefined" && !jsCookie.get('token') || error && error.status === 403
    return {
        loading,
        loggedOut,
        user: data && data.data,
        mutate,
        error
    }
}