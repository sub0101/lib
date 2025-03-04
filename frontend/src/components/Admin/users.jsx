"use client"

import { useEffect, useState } from "react"
import styles from "../../style/users.module.css"
import { useQuery } from "@tanstack/react-query"
import { getAllUsers } from "../../react-queries/api/users"

function Users() {
  const { data: allUsers, isSuccess } = useQuery({
    queryKey: ["users"],
    queryFn: getAllUsers,
  })

  useEffect(() => {
    console.log(allUsers)
  }, [isSuccess])

  const [searchTerm, setSearchTerm] = useState("")
  const [selectedUser, setSelectedUser] = useState(null)
  const [showEditModal, setShowEditModal] = useState(false)
  const [showDeleteModal, setShowDeleteModal] = useState(false)
  const [newRole, setNewRole] = useState("reader") // Default to "reader"

  const filteredUsers = allUsers?.filter(
    (user) =>
      user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.email.toLowerCase().includes(searchTerm.toLowerCase())
  )

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
    // API call to update role
    // For now, just console.log the change
    console.log(`User ${selectedUser.name} role updated to ${newRole}`)
    // Close modal after submission
    setShowEditModal(false)
  }

  const handleDeleteConfirm = () => {
    // API call to delete user
    // For now, just console.log the deletion
    console.log(`User ${selectedUser.name} deleted`)
    // Close modal after deletion
    setShowDeleteModal(false)
  }

  return (
    <div className={styles.usersPage}>
      <div className={styles.pageHeader}>
        <h1 className={styles.pageTitle}>User Management</h1>
        <button className={styles.addButton}>Add New User</button>
      </div>

      <div className={styles.searchContainer}>
        <input
          type="text"
          className={styles.searchInput}
          placeholder="Search users by name or email..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>

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
              {filteredUsers?.map((user) => (
                <tr key={user.ID}>
                  <td>{user.ID}</td>
                  <td>{user.name}</td>
                  <td>{user.email}</td>
                  <td>
                    <span
                      className={`${styles.badge} ${user.role === "faculty" ? styles.badgeSuccess : styles.badgeWarning}`}
                    >
                      {user.role}
                    </span>
                  </td>
                  <td>{1}</td>
                  <td>{new Date(Date.parse(user.CreatedAt)).toString()}</td>
                  <td>
                    <div className={styles.actionButtons}>
                      <button
                        className={styles.editButton}
                        onClick={() => handleEditClick(user)}
                      >
                        Edit
                      </button>
                      <button
                        className={styles.deleteButton}
                        onClick={() => handleDeleteClick(user)}
                      >
                        Delete
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      {/* Edit Role Modal */}
      {showEditModal && (
        <div className={styles.modal}>
          <div className={styles.modalContent}>
            <h2>Edit User Role</h2>
            <label>
              Role:
              <select
                value={newRole}
                onChange={(e) => setNewRole(e.target.value)}
              >
                <option value="reader">Reader</option>
                <option value="admin">Admin</option>
              </select>
            </label>
            <div className={styles.modalActions}>
              <button className={styles.cancelButton} onClick={() => setShowEditModal(false)}>
                Cancel
              </button>
              <button className={styles.okButton} onClick={handleEditSubmit}>
                OK
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Delete Confirmation Modal */}
      {showDeleteModal && (
        <div className={styles.modal}>
          <div className={styles.modalContent}>
            <h2>Are you sure you want to delete this user?</h2>
            <div className={styles.modalActions}>
              <button className={styles.cancelButton} onClick={() => setShowDeleteModal(false)}>
                Cancel
              </button>
              <button className={styles.okButton} onClick={handleDeleteConfirm}>
                OK
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

export default Users
