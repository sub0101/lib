import NavBar from "../shared/NavBar"
import { Outlet, useNavigate, useLocation } from "react-router-dom"
import styles from "../../style/admindash.module.css"
import SideBar from "../shared/SideBar"

function AdminPage({ requiredRole }) {
  const navigate = useNavigate()
  const location = useLocation()
  const currentPath = location.pathname

  const isActive = (path) => {
    if (path === "/admin" && (currentPath === "/admin" || currentPath === "/admin/dashboard")) {
      return true
    }
    return currentPath.startsWith(path)
  }

  const navigateTo = (path) => {
    navigate(path)
  }

  return (
    <div className={styles.adminPage}>  
<SideBar userRole={"admin"} />
      <div className={styles.mainContent}>
        <NavBar />
        <div className={styles.contentArea}>
          <Outlet />
        </div>
      </div>
    </div>
  )
}

export default AdminPage

