import axios from "axios";
import { getToken } from "../utils/auth/getUserInfo";
const BASE_URL = import.meta.env.VITE_BASE_URL

export const unAuthAxios = axios.create({
    baseURL: BASE_URL,
    headers :{
        'Content-Type': 'application/json',
    },
    withCredentials:true
})


export const authAxios = axios.create ({
    baseURL: BASE_URL,
    headers :{
        'Content-Type': 'application/json',
    },
    withCredentials:true
})

authAxios.interceptors.request.use(
    config =>{
        const token  = getToken()
     
        if(token) {
            config.headers["Authorization"]= `Bearer ${token}`
        }
        return config
    },(error)=> Promise.reject(error)
)
