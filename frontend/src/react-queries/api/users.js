import { data } from "react-router-dom"
import { authAxios } from "../axios"

const URL="users/"
export const getAllUsers = async() =>{

    const response =  await authAxios.get(URL)
    console.log(response)
    return response.data.data

}