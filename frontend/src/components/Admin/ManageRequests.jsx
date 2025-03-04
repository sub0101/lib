"use client"

import { useState } from "react"
import { useMutation, useQuery } from "@tanstack/react-query"
import { getAllAdminRequest, updateRequest } from "../../react-queries/api/requestTracking"
import styles from "../../style/managerequest.module.css"

const ManageRequests = () => {
  const [activeRow, setActiveRow] = useState(null)
const [payload , setPayload] = useState(null)
  const { data: requests, isLoading } = useQuery({
    queryFn: getAllAdminRequest,
    queryKey: ["getRequests"],
  })

  const requestMutation = useMutation({
    mutationFn: updateRequest,
    onMutate: (request) => {
     
    },
})
    
  const handleApprove = (id) => {
    alert(`Request with ID ${id} approved!`)

    setPayload({requestType:"accepted"})
    
    requestMutation.mutate(id ,payload )

  }

  const handleReject = (id) => {
    alert(`Request with ID ${id} rejected!`)
    setPayload({requestType:"rejected"})
    
    requestMutation.mutate(id ,payload )
  }

  const getStatusClass = (status) => {
    if (!status) return styles.pending
    return status.toLowerCase() === "approved" ? styles.approved : styles.rejected
  }

  return (
    <div className={styles.adminContainer}>
      <div className={styles.adminHeader}>
        <h1 className={styles.adminTitle}>Manage Requests</h1>
        <div className={styles.adminSubtitle}>Review and process pending book requests</div>
      </div>

      {isLoading ? (
        <div className={styles.loadingContainer}>
          <div className={styles.loadingSpinner}></div>
          <p>Loading requests...</p>
        </div>
      ) : requests && requests.length > 0 ? (
        <div className={styles.tableContainer}>
          <table className={styles.requestTable}>
            <thead>
              <tr>
                <th>Book Title</th>
                <th>Request ID</th>
                <th>Request Type</th>
                <th>Date</th>
                <th>Status/Actions</th>
              </tr>
            </thead>
            <tbody>
              {requests.map((request) => (
                <tr
                  key={request.req_id}
                  className={activeRow === request.req_id ? styles.activeRow : ""}
                  onMouseEnter={() => setActiveRow(request.req_id)}
                  onMouseLeave={() => setActiveRow(null)}
                >
                  <td className={styles.bookTitle}>{request.title}</td>
                  <td>{request.req_id}</td>
                  <td>
                    <span className={styles.requestType}>{request.request_type}</span>
                  </td>
                  <td>{new Date(request.request_date).toLocaleDateString()}</td>
                  {!request.approver_id ? (
                    <td className={styles.actionButtons}>
                      <button className={styles.approveButton} onClick={() => handleApprove(request.req_id)} >
                        Approve
                      </button>
                      <button className={styles.rejectButton} onClick={() => handleApprove(request.req_id)}>
                        Reject
                      </button>
                    </td>
                  ) : (
                    <td>
                      <span className={`${styles.statusBadge} ${getStatusClass(request.request_status)}`}>
                        {request.request_status}
                      </span>
                    </td>
                  )}
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : (
        <div className={styles.emptyState}>
          <div className={styles.emptyIcon}>📚</div>
          <h3>No requests found</h3>
          <p>There are currently no book requests to review</p>
        </div>
      )}
    </div>
  )
}

export default ManageRequests