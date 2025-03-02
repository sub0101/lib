import { useQuery } from "@tanstack/react-query"
import styles from "../../style/requeststracking.module.css"
import { getAllRequest } from "../../react-queries/api/requestTracking"

function RequestTracking() {

  const  {data:requests ,isSuccess} = useQuery({
    queryKey:["requests"],
    queryFn:getAllRequest
  })

  // const requests = [
  //   { id: 1, title: "To Kill a Mockingbird", status: "Pending", type: "Issue" },
  //   { id: 2, title: "1984", status: "Approved", type: "Return" },
  //   { id: 3, title: "Pride and Prejudice", status: "Rejected", type: "Issue" },
  // ]

  return (
    <div className={styles.requestTracking}>
      <h1>Request Tracking</h1>
      <table className={styles.requestTable}>
        <thead>
          <tr>
            <th>Book Title</th>
            <th>Request Type</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          { requests && requests.map((request) => (
            <tr key={request.ID}>
              <td>{request.Book.title}</td>
              <td>
                <span className={`${styles.requestType} ${styles[request.requestType.toLowerCase()]}`}>
                  {request.requestType}
                </span>
              </td>
              <td>
                <span className={`${styles.status} ${styles[request.requestStatus.toLowerCase()]}`}>
                  {request.requestStatus}
                </span>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default RequestTracking
