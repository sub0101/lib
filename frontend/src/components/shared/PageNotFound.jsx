import { useNavigate } from "react-router-dom"
import styles from "../../style/notfound.module.css"

function PageNotFound() {
  const navigate = useNavigate()

  return (
    <div className={styles.notFound}>
      <h1>404</h1>
      <h2>Oops! Page Not Found</h2>
      <p>The page you are looking for doesn’t exist or you don’t have permission to access it.</p>
      <button onClick={() => navigate("/")} className={styles.homeButton}>
        Go Back Home
      </button>
    </div>
  )
}

export default PageNotFound
