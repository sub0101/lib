"use client"

import { useState } from "react"
import styles from "../../style/issueBook.module.css"

function IssueNewBook() {
  const [bookTitle, setBookTitle] = useState("")
  const [bookId, setBookId] = useState("")

  const handleSubmit = (e) => {
    e.preventDefault()
    // Here you would typically send a request to your backend to issue the book
    console.log("Issuing book:", { bookTitle, bookId })
    // Reset form
    setBookTitle("")
    setBookId("")
  }

  return (
    <div className={styles.issueNewBook}>
      <h1>Issue New Book</h1>
      <form onSubmit={handleSubmit} className={styles.issueForm}>
        <div className={styles.formGroup}>
          <label htmlFor="bookTitle">Book Title:</label>
          <input type="text" id="bookTitle" value={bookTitle} onChange={(e) => setBookTitle(e.target.value)} required />
        </div>
        <div className={styles.formGroup}>
          <label htmlFor="bookId">Book ID:</label>
          <input type="text" id="bookId" value={bookId} onChange={(e) => setBookId(e.target.value)} required />
        </div>
        <button type="submit" className={styles.submitButton}>
          Issue Book
        </button>
      </form>
    </div>
  )
}

export default IssueNewBook

