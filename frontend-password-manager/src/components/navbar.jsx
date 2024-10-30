import React, { useEffect, useState } from "react";
import logo from "../assets/K-2.png";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

const Navbar = () => {
  let navigate = useNavigate();
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const userType = localStorage.getItem("navbarUserType");
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  let content;
  useEffect(() => {
    const isLoggedCheck = localStorage.getItem("isLoggedIn");
    console.log(isLoggedCheck);
    if (isLoggedCheck) {
      setIsLoggedIn(true);
    }
  }, []);
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
        <li>
          <button className="hover:text-gray-500">Add Password</button>
        </li>
        <li>
          <button className="hover:text-gray-500">List Websites</button>
        </li>
        <li>
          <button className="hover:text-gray-500">Get Passwords</button>
        </li>
      </>
    );
  } else {
    content = (
      <>
        <li>
          <Link to={"/about"}>
            <button className="hover:text-gray-500">About</button>
          </Link>
        </li>
        <li>
          <Link to={"/price"}>
            <button className="hover:text-gray-500">Pricing</button>
          </Link>
        </li>
      </>
    );
  }

  return (
    <header className="fixed top-0 left-0 right-0 bg-white w-full z-10">
      <nav className="flex justify-between items-center px-4 md:px-8 h-16">
        {/* Logo Section */}
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
        {/* Request Product Button (always visible) */}
        <div className="hidden md:flex">
          <Link to="/requestproduct">
            <button className="bg-[#AA7DFF] text-white px-5 py-2 rounded-full hover:bg-[#C49DFF]">
              Request Product
            </button>
          </Link>
        </div>

        {/* Hamburger Menu (only visible on small screens) */}
        <div className="md:hidden flex items-center">
          <button onClick={toggleMenu}>
            {/* Simple Hamburger Icon */}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              strokeWidth={2}
              className={`w-8 h-8 text-gray-800 transition duration-300 ease-in-out ${
                isMenuOpen ? "transform rotate-45" : "" // Add transform for open state
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
            <Link to="/requestproduct">
              <button className="bg-[#AA7DFF] text-white px-5 py-2 rounded-full hover:bg-[#C49DFF]">
                Request Product
              </button>
            </Link>
          </ul>
        </div>
      )}
    </header>
  );
};

export default Navbar;
