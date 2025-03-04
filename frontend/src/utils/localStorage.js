export const  setUserInfo = (accessToken)  =>{
   
    setLocalStorage("token" , accessToken)
}

export const setLocalStorage = (key ,token) =>{

    localStorage.setItem(key  , token)
}

export const getInfoFromStorage = (key) =>{

    const token = localStorage.getItem(key)
   
    return token
}
    