import React, { useEffect, useState , useRef} from "react";
import logo from "../assets/K-2.png";
import { Link } from "react-router-dom";
import { FaUser } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import { IoExit } from "react-icons/io5";
import { UserLogout } from "../services/loginService/loginService";

const Navbar = () => {
  let navigate = useNavigate();
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const userType = localStorage.getItem("navbarUserType");
  const dropdownRef = useRef(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [profileOption, setProfileOption] = useState(false);
  

  const handleClickOutside = (event) => {
    if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
      setProfileOption(false);
    }
  };

  useEffect(() => {
    const isLoggedCheck = localStorage.getItem("isLoggedIn");
    if (isLoggedCheck === "true") {
      setIsLoggedIn(true);
    }
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const handleLogout = async () => {
    console.log("Handle logout is called")
    let id = "";
    if (userType === "master") {
      id = localStorage.getItem("masterId");
    } else {
      id = localStorage.getItem("userId");
    }
    console.log("id = ", id)
    const response = await UserLogout(id);
    console.log(response.result);
  };

  const toggleProfileDropDown = () => {
    setProfileOption(!profileOption);
  };
  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const handlePage = () => {
    const isLoggedCheck = localStorage.getItem("isLoggedIn") === "true";

    if (!isLoggedCheck || !userType) {
      navigate("/");
    } else if (isLoggedCheck && userType === "user") {
      navigate("/userhomepage");
    } else if (isLoggedCheck && userType === "master") {
      navigate("/masterhomepage");
    }
  };

  let content;
  if (isLoggedIn && userType === "master") {
    content = (
      <>
        <Link to={"/master/adduser"}>
          <li>
            <button className="hover:text-gray-500">User Information</button>
          </li>
        </Link>
        <Link to={"/master/editconfig"}>
          <li>
            <button className="hover:text-gray-500">Config</button>
          </li>
        </Link>
        <Link to={"/master/info"}>
          <li>
            <button className="hover:text-gray-500">
              Personal Information
            </button>
          </li>
        </Link>
      </>
    );
  } else if (isLoggedIn && userType === "user") {
    content = (
      <>
        <Link to={"/user/addpassword"}>
          <li>
            <button className="hover:text-gray-500">Add Credentials</button>
          </li>
        </Link>
        <Link to={"/user/info"}>
          <li>
            <button className="hover:text-gray-500">Account Information</button>
          </li>
        </Link>
      </>
    );
  } 

  let button;
  if (isLoggedIn) {
    button = (
      <div key="loggedIn" className="hidden md:flex">
        <FaUser
          className="text-xl cursor-pointer"
          onClick={toggleProfileDropDown}
        />
      </div>
    );
  } else {
    button = (
      <div key="loggedOut">
        <div className="hidden md:flex">
          <Link to="/upgradeproduct">
            <button className="bg-[#AA7DFF] text-white px-5 py-2 rounded-full hover:bg-[#C49DFF]">
              Upgrade Product
            </button>
          </Link>
        </div>
      </div>
    );
  }

  return (
    <header className="fixed top-0 left-0 right-0 bg-white w-full z-10">
      <nav className="flex justify-between items-center px-4 md:px-8 h-16">
        <div className="flex items-center">
          <img
            className="w-16 cursor-pointer"
            src={logo}
            alt="Custom Logo"
            onClick={handlePage}
          />

          <span className="ml-2 text-xl font-bold text-black">KeyPass</span>
        </div>

        {/* Menu Links (hidden on small screens, visible on medium+) */}
        <div className="hidden md:flex space-x-6">
          <ul className="flex md:flex-row flex-col md:items-center md:gap-[5vw] gap-2">
            {content}
          </ul>
        </div>

        {button}

        {profileOption && (
          <div 
          ref={dropdownRef}
          className="absolute right-0 mt-20 w-48 h-12 bg-white border border-gray-300 rounded-lg shadow-lg z-10">
            <ul className="py-1">
              <li>
                <button
                  onClick={() => {
                    setProfileOption(false);
                    handleLogout();
                    localStorage.setItem("path", "logout");
                    navigate("/reset");
                  }}
                  className="w-full px-4 py-2 text-gray-700 hover:bg-blue-200 flex items-center space-x-2"
                >
                  <IoExit className="text-xl" />
                  Logout
                </button>
              </li>
            </ul>
          </div>
        )}

        <div className="md:hidden flex items-center">
          <button onClick={toggleMenu}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              strokeWidth={2}
              className={`w-8 h-8 text-gray-800 transition duration-300 ease-in-out ${
                isMenuOpen ? "transform rotate-45" : ""
              }`}
            >
              {isMenuOpen ? (
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M6 18L18 6M6 6l12 12"
                />
              ) : (
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M4 6h16M4 12h16m-7 6h7"
                />
              )}
            </svg>
          </button>
        </div>
      </nav>

      {/* Mobile Menu (only visible when isMenuOpen is true) */}
      {isMenuOpen && (
        <div className="md:hidden bg-white shadow-lg flex flex-col space-y-4 px-4 py-6 absolute w-full top-16 z-10">
          <ul className="flex flex-col space-y-2">
            {content}
            {!isLoggedIn ? (
              <Link to="/upgradeproduct">
                <button className="bg-[#AA7DFF] text-white px-5 py-2 rounded-full hover:bg-[#C49DFF]">
                  Upgrade Product
                </button>
              </Link>
            ) : (
              <></>
            )}
          </ul>
        </div>
      )}
    </header>
  );
};

export default Navbar;
