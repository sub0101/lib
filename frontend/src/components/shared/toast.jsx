"use client"

import { useState, useEffect } from "react"
import "../styles/toast.css"

const ToastNotification = ({ message, type = "info", duration = 3000, onClose }) => {
  const [visible, setVisible] = useState(true)

  useEffect(() => {
    const timer = setTimeout(() => {
      setVisible(false)
      setTimeout(() => {
        if (onClose) onClose()
      }, 300) // Wait for fade out animation
    }, duration)

    return () => clearTimeout(timer)
  }, [duration, onClose])

  const handleClose = () => {
    setVisible(false)
    setTimeout(() => {
      if (onClose) onClose()
    }, 300) // Wait for fade out animation
  }

  return (
    <div className={`toast ${type} ${visible ? "visible" : ""}`}>
      <div className="toast-content">
        <span className="toast-message">{message}</span>
        <button className="toast-close" onClick={handleClose}>
          ×
        </button>
      </div>
    </div>
  )
}


// <div className="app">
//       <h1>Toast Notification Demo</h1>

//       <div className="buttons">
//         <button onClick={() => showToast("success")}>Show Success</button>
//         <button onClick={() => showToast("error")}>Show Error</button>
//         <button onClick={() => showToast("info")}>Show Info</button>
//         <button onClick={() => showToast("warning")}>Show Warning</button>
//       </div>

//       {toast && <ToastNotification message={toast.message} type={toast.type} onClose={() => setToast(null)} />}
//     </div>
export default ToastNotification

