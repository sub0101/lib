import { data } from "react-router-dom"
import { authAxios } from "../axios"

const URL="users/"
export const getAllUsers = async() =>{

    const response =  await authAxios.get(URL)
    console.log(response)
    return response.data.data

}

export const updateUser= async(payload) =>{
    const response = await authAxios.patch(`${URL}make_admin` , payload)
    return response.data.data
}