import axios from "axios";
import jsCookie from "js-cookie";

export const register = (uname, pass) => axios.post(`${process.env.WEB_API}/user/register`, {uname, pass}).then(res => res.data)
export const login = (uname, pass) => axios.post(`${process.env.WEB_API}/user/login`, {uname, pass}).then(res => res.data)
export const getUserInfo = () => {
    let token = jsCookie.get('token')
    if (token) {
        return axios.get(`${process.env.WEB_API}/user/info`, {params: {token: token}}).then(res => res.data)
    }
}

