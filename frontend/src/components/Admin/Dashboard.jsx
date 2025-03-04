import { useEffect, useState } from "react"
import { getAllBooks } from "../../react-queries/api/books"
import styles from "../../style/dashboard.module.css"
import { useQuery } from "@tanstack/react-query"
function Dashboard() {

  const {data:booksQuery,isSuccess} =  useQuery({
    queryFn:getAllBooks,
    queryKey:["books"],
    
  })
 
 const [totalCount , setTotalCount] = useState(0)
 const [AvailableCount ,setAvailableCount ] = useState(0)

 useEffect(() => {
  if (isSuccess) {
    const totalBooks = booksQuery.map((book) => book.total_copies).reduce((sum, count) => sum + count, 0);
    const availableCopies = booksQuery.map((book) => book.available_copies).reduce((sum, count) => sum + count, 0);

    setTotalCount(totalBooks);
    setAvailableCount(availableCopies);
    console.log(isSuccess)
  }
}, [isSuccess]);
  return (
    <div className={styles.dashboard}>
      <h1 className={styles.pageTitle}>Dashboard</h1>

      <div className={styles.statsContainer}>
        <div className={`${styles.statCard} ${styles.books}`}>
          <h3>Total Books</h3>
          <div className={styles.statValue}>{totalCount}</div>
        </div>

        <div className={`${styles.statCard} ${styles.available}`}>
          <h3>Available Books</h3>
          <div className={styles.statValue}>{AvailableCount}</div>
        </div>

        <div className={`${styles.statCard} ${styles.issued}`}>
          <h3>Issued Books</h3>
          <div className={styles.statValue}>150</div>
        </div>

        <div className={`${styles.statCard} ${styles.users}`}>
          <h3>Total Users</h3>
          <div className={styles.statValue}>320</div>
        </div>
      </div>

      <div className={styles.recentActivities}>
        <h2>Recent Activities</h2>
        <div className={styles.tableContainer}>
          <table className={styles.table}>
            <thead>
              <tr>
                <th>Activity</th>
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>John Doe borrowed "The Great Gatsby"</td>
                <td>2023-06-15</td>
              </tr>
              <tr>
                <td>Jane Smith returned "To Kill a Mockingbird"</td>
                <td>2023-06-14</td>
              </tr>
              <tr>
                <td>New book added: "The Catcher in the Rye"</td>
                <td>2023-06-13</td>
              </tr>
              <tr>
                <td>New user registered: Robert Johnson</td>
                <td>2023-06-12</td>
              </tr>
              <tr>
                <td>Emily Davis borrowed "1984"</td>
                <td>2023-06-11</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}

export default Dashboard

