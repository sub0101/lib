  import React from 'react'
  import styles from '../../style/admindash.module.css';
import { useNavigate } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
  function SideBar({userRole}) {
    const navigate = useNavigate();
    const location = useLocation();
    const currentPath = location.pathname;
    
    const isActive = (path) => {
      if (path === '/admin' && (currentPath === '/admin' || currentPath === '/admin/dashboard')) {
        return true;
      }
      return currentPath.startsWith(path);
    };
  
    const navigateTo = (path) => {
      navigate(path);
    };

    const renderReaderMenu = () => (
            <>
              <li className={isActive("/tracking") ? styles.active : ""} onClick={() => navigateTo("/tracking")}>
                <div className={styles.menuItem}>
                  <img src="/tracking_i.png" alt="Tracking" />
                  <span>Book Tracking</span>
                </div>
              </li>
              <li className={isActive("/issued") ? styles.active : ""} onClick={() => navigateTo("/issued")}>
                <div className={styles.menuItem}>
                  <img src="/issued_i.png" alt="Issued Books" />
                  <span>Issued Books</span>
                </div>
              </li>
            
            </>
          )

    return (
        <nav className={styles.sidebar}>
        <div className={styles.logo}>
          <h2>Library MS</h2>
        </div>
        {userRole=="admin" &&   <ul className={styles.menu}>
          <li 
            className={isActive('/admin') ? styles.active : ''}
            onClick={() => navigateTo('/admin')}
          >
            <div className={styles.menuItem}>
              <img src='/dashboard_i.png' alt="Dashboard" />
              <span>Dashboard</span>
            </div>
          </li>
          <li 
            className={isActive('/admin/users') ? styles.active : ''}
            onClick={() => navigateTo('/admin/users')}
          >
            <div className={styles.menuItem}>
              <img src='/user_i.png' alt="Users" />
              <span>Manage Users</span>
            </div>
          </li>
          <li 
            className={isActive('/admin/books') ? styles.active : ''}
            onClick={() => navigateTo('/admin/books')}
          >
            <div className={styles.menuItem}>
              <img src='/book_i.png' alt="Books" />
              <span>Manage Books</span>
            </div>
          </li>
          {/* <li 
            className={isActive('/admin/settings') ? styles.active : ''}
            onClick={() => navigateTo('/admin/settings')}
          >
            <div className={styles.menuItem}>
              <img src='/setting_i.png' alt="Settings" />
              <span>Settings</span>
            </div>
          </li> */}
          <li 
            className={isActive('/admin/manage_request') ? styles.active : ''}
            onClick={() => navigateTo('/admin/manage_request')}
          >
            <div className={styles.menuItem}>
              <img src='/setting_i.png' alt="Settings" />
              <span>Manage Request</span>
            </div>
          </li>
        </ul>}

        {userRole == "reader" && 
           <ul className={styles.menu}>
         <li 
           className={isActive('/') ? styles.active : ''}
           onClick={() => navigateTo('/')}
         >
           <div className={styles.menuItem}>
             <img src='/dashboard_i.png' alt="Dashboard" />
             <span>Dashboard</span>
           </div>
         </li>
         <li 
           className={isActive('/books') ? styles.active : ''}
           onClick={() => navigateTo('/books')}
         >
           <div className={styles.menuItem}>
             <img src='/user_i.png' alt="books" />
             <span>View Books</span>
           </div>
         </li>
         <li 
           className={isActive('/issued') ? styles.active : ''}
           onClick={() => navigateTo('/issued')}
         >
           <div className={styles.menuItem}>
             <img src='/user_i.png' alt="books" />
             <span>Manage Books</span>
           </div>
         </li>
         <li 
           className={isActive('/tracking') ? styles.active : ''}
           onClick={() => navigateTo('/tracking')}
         >
           <div className={styles.menuItem}>
             <img src='/book_i.png' alt="Books" />
             <span>Request Tracking</span>
           </div>
         </li>
       </ul>
        }
       
      </nav>
    )
  }
  
  export default SideBar
 
