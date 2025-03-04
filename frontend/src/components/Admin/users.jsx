"use client"

import { useEffect, useState } from "react"
import { useMutation, useQuery } from "@tanstack/react-query"
import { getAllUsers, updateUser } from "../../react-queries/api/users"
import styles from "../../style/users.module.css"


function Users() {
  const {
    data: allUsers,
    isLoading,
    isSuccess,
  } = useQuery({
    queryKey: ["users"],
    queryFn: getAllUsers,
  })

  useEffect(() => {
    console.log(allUsers)
  }, [allUsers])

  const [searchTerm, setSearchTerm] = useState("")
  const [selectedUser, setSelectedUser] = useState(null)
  const [showEditModal, setShowEditModal] = useState(false)
  const [showDeleteModal, setShowDeleteModal] = useState(false)
  const [newRole, setNewRole] = useState("reader")

  const filteredUsers = allUsers?.filter(
    (user) =>
      user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.email.toLowerCase().includes(searchTerm.toLowerCase()),
  )


  const updateMutation = useMutation({ 
    mutationFn: updateUser,
   mutationKey: ["updateUser"]
  })

  const handleEditClick = (user) => {
    setSelectedUser(user)
    setNewRole(user.role)
    setShowEditModal(true)
  }

  const handleDeleteClick = (user) => {
    setSelectedUser(user)
    setShowDeleteModal(true)
  }

  const handleEditSubmit = () => {
    updateMutation.mutate({ id: selectedUser.ID, role: newRole })

    console.log(`User ${selectedUser.name} role updated to ${newRole}`)
    setShowEditModal(false)
  }

  const handleDeleteConfirm = () => {
    // API call to delete user
    console.log(`User ${selectedUser.name} deleted`)
    setShowDeleteModal(false)
  }

  const formatDate = (dateString) => {
    const options = { year: "numeric", month: "short", day: "numeric" }
    return new Date(Date.parse(dateString)).toLocaleDateString(undefined, options)
  }

  return (
    <div className={styles.usersPage}>
      <div className={styles.pageHeader}>
        <div>
          <h1 className={styles.pageTitle}>User Management</h1>
          <p className={styles.pageSubtitle}>Manage library users and their roles</p>
        </div>
        <button className={styles.addButton}>
          <span className={styles.addIcon}>+</span>
          Add New User
        </button>
      </div>

      <div className={styles.searchContainer}>
        <svg
          className={styles.searchIcon}
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        >
          <circle cx="11" cy="11" r="8"></circle>
          <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
        </svg>
        <input
          type="text"
          className={styles.searchInput}
          placeholder="Search users by name or email..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>

      {isLoading ? (
        <div className={styles.loadingContainer}>
          <div className={styles.loadingSpinner}></div>
          <p>Loading users...</p>
        </div>
      ) : (
        <div className={styles.tableCard}>
          <div className={styles.tableContainer}>
            <table className={styles.table}>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Name</th>
                  <th>Email</th>
                  <th>Role</th>
                  <th>Books Issued</th>
                  <th>Join Date</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {filteredUsers?.length > 0 ? (
                  filteredUsers.map((user) => (
                    <tr key={user.ID}>
                      <td>{user.ID}</td>
                      <td className={styles.userName}>{user.name}</td>
                      <td className={styles.userEmail}>{user.email}</td>
                      <td>
                        <span
                          className={`${styles.badge} ${
                            user.role === "admin"
                              ? styles.badgePrimary
                              : user.role === "faculty"
                                ? styles.badgeSuccess
                                : styles.badgeWarning
                          }`}
                        >
                          {user.role}
                        </span>
                      </td>
                      <td className={styles.booksCount}>{1}</td>
                      <td>{formatDate(user.CreatedAt)}</td>
                      <td>
                        <div className={styles.actionButtons}>
                          <button
                            className={styles.editButton}
                            onClick={() => handleEditClick(user)}
                            aria-label={`Edit ${user.name}`}
                          >
                            <svg
                              xmlns="http://www.w3.org/2000/svg"
                              viewBox="0 0 24 24"
                              fill="none"
                              stroke="currentColor"
                              strokeWidth="2"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              className={styles.buttonIcon}
                            >
                              <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                              <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                            </svg>
                            Edit
                          </button>
                          <button
                            className={styles.deleteButton}
                            onClick={() => handleDeleteClick(user)}
                            aria-label={`Delete ${user.name}`}
                          >
                            <svg
                              xmlns="http://www.w3.org/2000/svg"
                              viewBox="0 0 24 24"
                              fill="none"
                              stroke="currentColor"
                              strokeWidth="2"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              className={styles.buttonIcon}
                            >
                              <polyline points="3 6 5 6 21 6"></polyline>
                              <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
                            </svg>
                            Delete
                          </button>
                        </div>
                      </td>
                    </tr>
                  ))
                ) : (
                  <tr>
                    <td colSpan="7" className={styles.noResults}>
                      No users found matching your search
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      )}

      {/* Edit Role Modal */}
      {showEditModal && (
        <>
          <div className={styles.modalOverlay} onClick={() => setShowEditModal(false)}></div>
          <div className={styles.modal}>
            <div className={styles.modalContent}>
              <div className={styles.modalHeader}>
                <h2 className={styles.modalTitle}>Edit User Role</h2>
                <button className={styles.modalClose} onClick={() => setShowEditModal(false)} aria-label="Close modal">
                  ×
                </button>
              </div>

              <div className={styles.modalBody}>
                <div className={styles.userInfo}>
                  <p className={styles.userName}>{selectedUser?.name}</p>
                  <p className={styles.userEmail}>{selectedUser?.email}</p>
                </div>

                <div className={styles.formGroup}>
                  <label className={styles.formLabel}>
                    Select Role:
                    <select className={styles.formSelect} value={newRole} onChange={(e) => setNewRole(e.target.value)}>
                      <option value="reader">Reader</option>
             
                      <option value="admin">Admin</option>
                    </select>
                  </label>
                </div>
              </div>

              <div className={styles.modalFooter}>
                <button className={styles.cancelButton} onClick={() => setShowEditModal(false)}>
                  Cancel
                </button>
                <button className={styles.confirmButton} onClick={handleEditSubmit}>
                  Update Role
                </button>
              </div>
            </div>
          </div>
        </>
      )}

      {/* Delete Confirmation Modal */}
      {showDeleteModal && (
        <>
          <div className={styles.modalOverlay} onClick={() => setShowDeleteModal(false)}></div>
          <div className={styles.modal}>
            <div className={styles.modalContent}>
              <div className={styles.modalHeader}>
                <h2 className={styles.modalTitle}>Confirm Deletion</h2>
                <button
                  className={styles.modalClose}
                  onClick={() => setShowDeleteModal(false)}
                  aria-label="Close modal"
                >
                  ×
                </button>
              </div>

              <div className={styles.modalBody}>
                <div className={styles.deleteWarning}>
                  <svg
                    className={styles.warningIcon}
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
                    <line x1="12" y1="9" x2="12" y2="13"></line>
                    <line x1="12" y1="17" x2="12.01" y2="17"></line>
                  </svg>
                  <p>Are you sure you want to delete this user?</p>
                </div>

                <div className={styles.userInfo}>
                  <p className={styles.userName}>{selectedUser?.name}</p>
                  <p className={styles.userEmail}>{selectedUser?.email}</p>
                </div>

                <p className={styles.deleteNote}>
                  This action cannot be undone. All data associated with this user will be permanently removed.
                </p>
              </div>

              <div className={styles.modalFooter}>
                <button className={styles.cancelButton} onClick={() => setShowDeleteModal(false)}>
                  Cancel
                </button>
                <button className={`${styles.confirmButton} ${styles.deleteConfirm}`} onClick={handleDeleteConfirm}>
                  Delete User
                </button>
              </div>
            </div>
          </div>
        </>
      )}
    </div>
  )
}

export default Users

