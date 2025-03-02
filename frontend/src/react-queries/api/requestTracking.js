import { authAxios } from "../axios"

const URL = "request/"
export const getAllRequest = async() =>{
    const response =  await authAxios.get(URL)
    console.log(response.data)
    return response.data.data
}