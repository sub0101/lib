import { authAxios } from "../axios"

const URL = "request/"
export const getAllRequest = async() =>{
    const response =  await authAxios.get(URL)
    console.log(response.data)
    return response.data.data
}


export const AddRequest = async(payload) =>{
    const response = await  authAxios.post(URL , payload)
    console.log(response.data)
    return response.data.data
}

export const addReturnRequest =async(payload) =>{
    const response = await authAxios.post(URL ,payload)
    console.log(response.data)
    return response.data.data
}

export  const updateRequest = async (id ,payload) =>{

    const response   =  await  authAxios.patch(URL , payload , {
        params:id
    })
    console.log(response.data)
    return response.data.data
}

export const getAllAdminRequest = async() =>{
    const response =  await authAxios.get(`${URL}all`)
    console.log(response.data)
    return response.data.data
}
