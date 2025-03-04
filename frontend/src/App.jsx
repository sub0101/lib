import React from 'react'
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from './components/Auth/Login'
import ProtectedLogin from './components/Auth/ProtectedLogin';
import PrivateOutlet from './components/shared/PrivateOutlet';
import AdminPage from './components/Admin/AdminPage';
import Home from './components/Home';
import ReaderPage from './components/reader/ReaderPage';
import { UserRegistration } from './components/Auth/UserRegistration';
import Contactus from './components/reader/Contactus';

import Users from './components/Admin/users';
import Books from './components/Admin/books';
import ReaderDashboard from './components/reader/ReaderDashBoard';
import BookTracking from './components/reader/tracking';
import IssueNewBook from './components/reader/issue-book';
import IssuedBooks from './components/reader/IssuedBooks';
import Dashboard from './components/Admin/DashBoard';
import ViewBooks from './components/reader/books';
import ManageRequests from './components/Admin/ManageRequests';
function App() {
  return (
    <Router >
      <Routes>
        <Route path='/login' element={
          <ProtectedLogin>
            <Login />
          </ProtectedLogin>
        }
        />
         <Route path='/signup' element={
<ProtectedLogin >
< UserRegistration/>
</ProtectedLogin>
         
      
        }
        />

        <Route element={<PrivateOutlet />} >
          <Route element={<Home requiredRole={["admin", "owner"]} />} >
            <Route path="/admin" element={<AdminPage />} >
            <Route path='' element ={<Dashboard />} />
            <Route path='users' element ={<Users/>} />
            <Route path='books' element={<Books />} />
            <Route path='manage_request' element={<ManageRequests />} />
            </Route>  
          </Route>
        </Route>

        <Route element={<PrivateOutlet />} >
          <Route element={<Home requiredRole={["reader"]} />} >
            <Route path="" element={<ReaderPage />} >
            <Route path='' element = {<ReaderDashboard />} />
            <Route path='/contact' element={<Contactus />} />
            <Route path='tracking' element={<BookTracking />} />
            <Route path='/issued' element={<IssuedBooks />} />
            <Route path='/books' element={<ViewBooks />} />
            </Route>
          </Route>
        </Route>

      </Routes>
    </Router>
  )
}

export default App