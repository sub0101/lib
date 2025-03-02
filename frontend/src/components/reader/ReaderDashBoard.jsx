
"use client"

import { useState } from "react"
import styles from "../../style/readerdashboard.module.css"

import IssueBookModal from "./IssueNewBookModal"

function ReaderDashboard() {
  const [isModalOpen, setIsModalOpen] = useState(false)

  return (
    <div className={styles.dashboard}>
      <h1>Reader Dashboard</h1>
      <div className={styles.statsContainer}>
        <div className={styles.statCard}>
          <h2>Books Read</h2>
          <p>15</p>
        </div>
        <div className={styles.statCard}>
          <h2>Currently Borrowed</h2>
          <p>3</p>
        </div>
        <div className={styles.statCard}>
          <h2>Overdue Books</h2>
          <p>1</p>
        </div>
      </div>
      <button className={styles.issueButton} onClick={() => setIsModalOpen(true)}>
        Issue New Book
      </button>
      {isModalOpen && <IssueBookModal onClose={() => setIsModalOpen(false)} />}
    </div>
  )
}

export default ReaderDashboard

