import React, { useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import navStyle from "../../style/navbar.module.css"
import { FaCog, FaSignOutAlt, FaUserCircle } from "react-icons/fa"
import { getUserInfo, loggedOut } from "../../utils/auth/getUserInfo"
import { adminRoles } from "../../utils"

const NavBar = () => {
  const [isNavToggled, setToggle] = useState(false)
  const [isDropdownOpen, setDropdownOpen] = useState(false)
  const user = getUserInfo()
  const isAdmin = adminRoles.includes(user.Role)
const navigate = useNavigate()
  const handleLogout = () => {
    
    loggedOut()
navigate("/login")
    // Implement logout logic here (e.g., clearing session, redirecting)
  }

  return (
    <div>
      <nav className={navStyle.nav}>
        <div className={navStyle.logo}>
          <h1>Xenon Library</h1>
        </div>

        <ul>
          <li><Link to={ isAdmin ? "/admin" : "/" }>Home</Link></li>
          <li><Link to={"/contact"}>Contact Us</Link></li>
        </ul>

        {/* User Profile Icon */}
        <div className={navStyle.userProfile} onClick={() => setDropdownOpen(!isDropdownOpen)}>
          <FaUserCircle className={navStyle.userIcon} />
          {isDropdownOpen && (
            <div className={navStyle.dropdownMenu}>
              <Link to="/settings" className={navStyle.dropdownItem}>
                <FaCog className={navStyle.dropdownIcon} />
                Settings
              </Link>
              <button onClick={handleLogout} className={navStyle.dropdownItem}>
                <FaSignOutAlt className={navStyle.dropdownIcon} />
                Logout
              </button>
            </div>
          )}
        </div>

        {/* Hamburger Menu */}
        {/* <div className={navStyle.menu} onClick={() => setToggle((val) => !val)}>
          <p className={navStyle.line}></p>
          <p className={navStyle.line}></p>
          <p className={navStyle.line}></p>
        </div> */}
      </nav>

      {/* Mobile Menu */}
      {/* <div className={isNavToggled ? `${navStyle.active} ${navStyle.menuBar}` : navStyle.menuBar}>
        <ul>
          <li><Link to={"/"}>Home</Link></li>
          <li><Link to={"/store"}>Store</Link></li>
          <li><Link to={"/service"}>Service</Link></li>
          <li><Link to={"/contactUs"}>Contact Us</Link></li>
        </ul>
      </div> */}
    </div>
  )
}

export default NavBar
