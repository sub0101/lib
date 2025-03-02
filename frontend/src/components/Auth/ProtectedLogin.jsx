import React from 'react'
import { isLoggedin } from '../../utils/auth/getUserInfo'
import { Navigate } from 'react-router-dom'

function ProtectedLogin({children}) {
    const isLog =isLoggedin()
    console.log("this is protected lgoin")
console.log(isLog)

    if(isLog ) {
        return <Navigate to="/" replace />;
    }
    return children
//     if(isLog) {
//         if (isLoggedin()) {
//             return <Navigate to="/" replace />;
//           }
//           console.log("children")
//           return children;
//     }
}

export default ProtectedLogin