import { getInfoFromStorage } from "../localStorage"
import {jwtDecode} from "jwt-decode"
export const getUserInfo = () =>{

    const authToken   = getInfoFromStorage("token")
     if(authToken) {    
     
    
        const user =  jwtDecode(authToken)
        
       return user.payload
     }
     return ""
}

export const isLoggedin = () =>{
const token = getInfoFromStorage("token")
console.log(!!token)
 return !!token;
}

export const loggedOut  = ()=>{

    return localStorage.removeItem("token")
}

export const getToken = () =>{
  
   const  token = getInfoFromStorage("token")
   
    return token
}