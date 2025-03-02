"use client"

import { useEffect, useState } from "react"
import styles from "../../style/users.module.css"
import { useQuery } from "@tanstack/react-query"
import { getAllUsers } from "../../react-queries/api/users"

function Users() {
  const {data:allUsers , isSuccess} = useQuery({
    queryKey:["users"],
    queryFn:getAllUsers

  })

  useEffect(()=>{
    console.log(allUsers)
  } , [isSuccess])

  const [users, setUsers] = useState([
    { id: 1, name: "John Doe", email: "john@example.com", role: "student", booksIssued: 2, joinDate: "2023-01-15" },
    { id: 2, name: "Jane Smith", email: "jane@example.com", role: "faculty", booksIssued: 3, joinDate: "2023-02-20" },
    {
      id: 3,
      name: "Robert Johnson",
      email: "robert@example.com",
      role: "student",
      booksIssued: 1,
      joinDate: "2023-03-10",
    },
    { id: 4, name: "Emily Davis", email: "emily@example.com", role: "student", booksIssued: 0, joinDate: "2023-04-05" },
    {
      id: 5,
      name: "Michael Wilson",
      email: "michael@example.com",
      role: "faculty",
      booksIssued: 4,
      joinDate: "2023-01-25",
    },
    { id: 6, name: "Sarah Brown", email: "sarah@example.com", role: "student", booksIssued: 2, joinDate: "2023-05-12" },
    {
      id: 7,
      name: "David Miller",
      email: "david@example.com",
      role: "student",
      booksIssued: 1,
      joinDate: "2023-03-22",
    },
    { id: 8, name: "Lisa Taylor", email: "lisa@example.com", role: "faculty", booksIssued: 0, joinDate: "2023-02-18" },
  ])

  const [searchTerm, setSearchTerm] = useState("")

  const filteredUsers = users.filter(
    (user) =>
      user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.email.toLowerCase().includes(searchTerm.toLowerCase()),
  )

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
              { allUsers && allUsers.map((user) => (
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
                  <td>{ new Date(Date.parse(user.CreatedAt)).toString()}</td>
                  <td>
                    <div className={styles.actionButtons}>
                      <button className={styles.editButton}>Edit</button>
                      <button className={styles.deleteButton}>Delete</button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}

export default Users

