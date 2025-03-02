import React, { useState } from 'react'
import { Form, Link, Navigate, useNavigate } from 'react-router-dom'
import { login } from '../../react-queries/api/auth'
import { useMutation } from '@tanstack/react-query'
import "../../style/login.module.css"
import styles from "../../style/login.module.css";


function Login() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const navigate = useNavigate()
  const handleSubmit = async (e) => {
    e.preventDefault()
    authMutation.mutate({
      email, password
    })
  }
  const authMutation = useMutation({
    mutationFn: login,
    onSuccess: () => {
      console.log("successfully loggedin")
      navigate("/")
    }
  })

  return (
    <div className={styles.authForm}>
  
<div className={styles.wrapper}>
    <form onSubmit={handleSubmit}>
      <h2>Login</h2>
        <div className={styles.inputField}>
        <input type="text" name='email' value={email} onChange={(e)=> setEmail(e.target.value)} required />
        <label>Enter your email</label>
      </div>
      <div className={styles.inputField}>
        <input type="password"  name='password' value={password} onChange={(e)=> setPassword(e.target.value)} required />
        <label>Enter your password</label>
      </div>
      <div className={styles.forget}>
        <label htmlFor="remember">
          <input type="checkbox" id="remember" />
          <p>Remember me</p>
        </label>
        <a href="#">Forgot password?</a>
      </div>
      <button type="submit">Log In</button>
      <div className= {styles.register}>

      <p>Don't have an account? <Link to="/signup">Register</Link></p>

      </div>
    </form>
  </div>



    </div>
  )
}

export default Login