import { useMutation, useQuery } from "@tanstack/react-query"
import styles from "../../style/issuedbooks.module.css"
import { useState } from "react"
import { getIssuesBooks } from "../../react-queries/api/books"
import { addReturnRequest } from "../../react-queries/api/requestTracking"

function IssuedBooks() {
  const [payload , setPayload] = useState({})
const returnRequestMutation = useMutation({
  mutationFn:addReturnRequest,
  mutationKey :"return request",
  onSuccess:()=>{
    console.log("succedd return")
  }
})
  const today = new Date().toISOString().split("T")[0] // Get today's date in YYYY-MM-DD format

  const handleReturn = (book) => {
    const confirmReturn = window.confirm(`Are you sure you want to return "${book.bookTitle}"?`)
    if (confirmReturn) {
      // setIssuedBooks(issuedBooks.filter((book) => book.id !== bookId))
      payload.bookId = book.id
      payload.requestType = "return"
      returnRequestMutation.mutate(payload)

      console.log(`Returned book: ${book.bookTitle}`)
    }
  }

  const {data:issuedBooks} = useQuery ({
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
          { issuedBooks && issuedBooks.map((book) => (
            <tr key={book.id}>
              <td>{book.title}</td>
              <td className={book.expectedReturnDate < today ? styles.overdue : ""}>{book.expected_return_date}</td>
              <td>
                <button 
                  className={styles.returnButton} 
                  onClick={() => handleReturn({id:book.id, bookTitle:book.title})}
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
