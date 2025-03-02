import React from 'react'
import { Outlet } from 'react-router-dom';
import { getUserInfo } from '../utils/auth/getUserInfo';
import PageNotFound from './shared/PageNotFound';

function Home({requiredRole}) {

    const user = getUserInfo()
    console.log(requiredRole)
    console.log(user)
     if (!user) {
        return <PageNotFound />
     }
     const isInclude = requiredRole.includes(user.Role)
     console.log(isInclude)
 return isInclude ? <Outlet /> : <PageNotFound />

}

export default Home