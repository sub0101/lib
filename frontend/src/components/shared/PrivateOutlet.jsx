import React, { useEffect } from 'react'
import { getUserInfo, isLoggedin } from '../../utils/auth/getUserInfo'
import { Navigate, Outlet, useNavigate } from 'react-router-dom'

function PrivateOutlet() {
    console.log("inside the private Outlet")
    const isLogged = isLoggedin()

    return  isLogged  ? <Outlet /> : <Navigate to= "/login" replace />
  
}

export default PrivateOutlet