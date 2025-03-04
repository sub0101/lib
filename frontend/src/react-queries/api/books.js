import axios from "axios";
import { authAxios } from "../axios"

const URL = "books/"
export const getAllBooks  =  async () =>{
    
 const   response = await authAxios.get(URL);
 console.log(response.data.data)
return response.data.data
}

export const getBook = (id) =>{

} 

export const addBook = async (payload) =>{
    console.log(payload)
    const response  = await authAxios.post(URL , payload)
    console.log(response.data)
    return response.data.data
}

export const  updateBook = (payload) =>{

}

export const deleteBook = (id) =>{

}

export const getIssuesBooks = async ()=>{
    const response = await authAxios.get("books/issued")
    console.log(response.data)
    return response.data.data
}