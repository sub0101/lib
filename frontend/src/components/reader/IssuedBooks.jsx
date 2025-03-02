import { useQuery } from "@tanstack/react-query"
import styles from "../../style/issuedbooks.module.css"
import { useState } from "react"
import { getIssuesBooks } from "../../react-queries/api/books"

function IssuedBooks() {
  const [issuedBooks, setIssuedBooks] = useState([
    { id: 1, title: "The Great Gatsby", dueDate: "2025-03-10" },
    { id: 2, title: "Moby Dick", dueDate: "2025-03-15" },
    { id: 3, title: "War and Peace", dueDate: "2025-03-05" }, // Example overdue book
  ])

  const today = new Date().toISOString().split("T")[0] // Get today's date in YYYY-MM-DD format

  const handleReturn = (bookId, bookTitle) => {
    const confirmReturn = window.confirm(`Are you sure you want to return "${bookTitle}"?`)
    if (confirmReturn) {
      setIssuedBooks(issuedBooks.filter((book) => book.id !== bookId))
      console.log(`Returned book: ${bookTitle}`)
    }
  }

  const {data:issuedBook} = useQuery ({
    queryKey:["issued"],
    queryFn:getIssuesBooks
  })
  return (
    <div className={styles.issuedBooks}>
      <h1>Issued Books</h1>
      <table className={styles.bookTable}>
        <thead>
          <tr>
            <th>Title</th>
            <th>Due Date</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {issuedBooks.map((book) => (
            <tr key={book.id}>
              <td>{book.title}</td>
              <td className={book.dueDate < today ? styles.overdue : ""}>{book.dueDate}</td>
              <td>
                <button 
                  className={styles.returnButton} 
                  onClick={() => handleReturn(book.id, book.title)}
                >
                  Return Book
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default IssuedBooks
