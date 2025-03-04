import React from 'react';
import NavBar from '../shared/NavBar';
import { Outlet } from 'react-router-dom';
import SideBar from '../shared/SideBar';
import styles from "../../style/readerpage.module.css"
import ReaderDashboard from './ReaderDashBoard';

const Dashboard = () => <div>Reader Dashboard</div>
const BookTracking = () => <div>Book Tracking</div>
const IssuedBooks = () => <div>Currently Issued Books</div>
const IssueNewBook = () => <div>Issue New Book</div>

function ReaderPage() {
  return (
    <div className={styles.readerPage}>
      <SideBar userRole={"reader"}/>
   
      <div className={styles.mainContent}>
        <NavBar />
        <div className={styles.contentArea}>
          <Outlet />
        </div>
      </div>
    </div>
  )
}

export default ReaderPage

