import React from 'react'
import { Navigate } from 'react-router-dom'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import loginStyle from '../../style/login.module.css'
import signupStyle from '../../style/signup.module.css'
export const UserRegistration = () => {

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [name, setName] = useState("")
  const [contact, setContact] = useState("")
  const [library, setLibrary] = useState("")
  const listLibrary  = ["Lib1"  ,"Lib2" , "Lib3" , "Lib4","Lib5"  ,"Lib6" , "Lib7" , "Lib8fsfdsfdsfdsfdsfdsfdsdffdfdfdfdfdffffffffffffffff"]
  const navigate = useNavigate()
  const handleSubmit = async (e) => {
    e.preventDefault()
    // authMutation.mutate({
    //   email, password
    // })
  }
  // const authMutation = useMutation({
  //   mutationFn: login,
  //   onSuccess: () => {
  //     console.log("successfully loggedin")
  //     navigate("/")
  //   }
  // })

  return (
    <div className={loginStyle.authForm}>

      <div className={loginStyle.wrapper}>
        <form onSubmit={handleSubmit}>
          <h2> Registration</h2>
          <div className={loginStyle.inputField}>
            <input name='name' value={password} onChange={(e) => setPassword(e.target.value)} required />
            <label>Enter your Name</label>
          </div>
          <div className={loginStyle.inputField}>
            <input type="text" name='email' value={email} onChange={(e) => setEmail(e.target.value)} required />
            <label>Enter your email</label>
          </div>

          <div className={loginStyle.inputField}>
            <input type="text" name='contact' value={password} onChange={(e) => setPassword(e.target.value)} required />
            <label>Enter your contact</label>
          </div>
          <div className={loginStyle.inputField}>
            <input type="password" name='password' value={password} onChange={(e) => setPassword(e.target.value)} required />
            <label>Enter your password</label>
          </div>
          <div className={`${loginStyle.inputField}`} >
          <input list ="libraries"  name='library' value={library} onChange={(e) => setLibrary(e.target.value)} required />
          <datalist className={signupStyle.dropDown} id="libraries">
             
          {listLibrary.map((item, inx) => (
    <option key={inx} value={item} />
  ))}
            </datalist>
            <label>Select Your Library</label>

          </div>


          <button type="submit">Register</button>
          <div class="register">

          </div>
        </form>
      </div>


    </div>
  )
}

