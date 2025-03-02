"use client"

import { useState ,useEffect} from "react"
import styles from "../../style/books.module.css"
import { getAllBooks } from "../../react-queries/api/books"
import { useQuery } from "@tanstack/react-query"
import AddBookModal from "./AddBookModal"
function Books() {
const [showModal , setShowModal] = useState(false)
const changeModalState = (st) =>{
  setShowModal(st)
 }
  const {data:allBooks , isSuccess} = useQuery({
    queryKey:["books"],
    queryFn:getAllBooks

  })

  useEffect(()=>{
    console.log(allBooks)
  } , [isSuccess])


  const [searchTerm, setSearchTerm] = useState("")
  const [filter, setFilter] = useState("all")

  const filteredBooks =   allBooks &&  allBooks.filter((book) => {
    const matchesSearch =
     book.title.toLowerCase().includes( searchTerm.toLowerCase()) ||
      book.authors.toLowerCase().includes( searchTerm.toLowerCase())

    if (filter === "all") return matchesSearch
    if (filter === "available") return matchesSearch && book.availableCopies>0
    if (filter === "not available") return matchesSearch && book.availableCopies<=0

    return matchesSearch
  }) 

  return (
    <div className={styles.booksPage}>
      <div className={styles.pageHeader}>
        <h1 className={styles.pageTitle}>Book Management</h1>
        <button className={styles.addButton}  onClick={()=> setShowModal(true)} >Add New Book</button>
      </div>
     {showModal &&  <AddBookModal setShowModal = {changeModalState} />}

      <div className={styles.filterContainer}>
        <div className={styles.searchContainer}>
          <input
            type="text"
            className={styles.searchInput}
            placeholder="Search books by title or author..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>

        <div className={styles.filterItem}>
          <span className={styles.filterLabel}>Status:</span>
          <select className={styles.filterSelect} value={filter} onChange={(e) => setFilter(e.target.value)}>
            <option value="all">All Books</option>
            <option value="available">Available</option>
            <option value="not available">Not Available</option>
          </select>
        </div>
      </div>

      <div className={styles.bookGrid}>
        {filteredBooks &&  filteredBooks.map((book) => (
          <div className={styles.bookCard} key={book.id}>
            <div className={styles.bookCover}>
              <span>Book Cover</span>
            </div>
            <div className={styles.bookDetails}>
              <h3 className={styles.bookTitle}>{book.title}</h3>
              <p className={styles.bookAuthor}>by {book.authors}</p>
              <div className={styles.bookMeta}>
                <span>Publisher: {book.publisher}</span>
                <span>Available: {book.availableCopies}</span>
              </div>
              <div className={styles.bookMeta}>
                <span>ISBN: {book.isbn}</span>
                <span className={`${styles.badge} ${book.available ? styles.badgeSuccess : styles.badgeDanger}`}>
                  {book.available ? "Available" : "Issued"}
                </span>
              </div>
              <div className={styles.bookMeta}>
                <span>TotalCopies: {book.totalCopies}</span>
              </div>
              <div className={styles.actionButtons}>
                <button className={styles.editButton}>Edit</button>
                <button className={styles.deleteButton}>Delete</button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default Books

 