"use client"

import { useState } from "react"
import styles from "../../style/issuenewbook.module.css"

function IssueBookModal({ onClose }) {
  const [bookTitle, setBookTitle] = useState("")
  const [bookId, setBookId] = useState("")

  const handleSubmit = (e) => {
    e.preventDefault()
    // Here you would typically send a request to your backend to issue the book
    console.log("Issuing book:", { bookTitle, bookId })
    // Close the modal
    onClose()
  }

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modal}>
        <h2>Issue New Book</h2>
        <form onSubmit={handleSubmit}>
          <div className={styles.formGroup}>
            <label htmlFor="bookTitle">Book Title:</label>
            <input
              type="text"
              id="bookTitle"
              value={bookTitle}
              onChange={(e) => setBookTitle(e.target.value)}
              required
            />
          </div>
          <div className={styles.formGroup}>
            <label htmlFor="bookId">Book ID:</label>
            <input type="text" id="bookId" value={bookId} onChange={(e) => setBookId(e.target.value)} required />
          </div>
          <div className={styles.buttonGroup}>
            <button type="submit" className={styles.submitButton}>
              Issue Book
            </button>
            <button type="button" className={styles.cancelButton} onClick={onClose}>
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

export default IssueBookModal

