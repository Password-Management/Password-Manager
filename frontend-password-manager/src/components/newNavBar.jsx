import React, { useEffect, useState } from "react";

const Navbar = () => {
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
  if (isLoggedIn && userType == "master") {
    content = (
      <>
        <a href="#" className="hover:text-blue-400">
          CreateUser
        </a>
        <a href="#" className="hover:text-blue-400">
          ListUsers
        </a>
        <a href="#" className="hover:text-blue-400">
          EditKey
        </a>
        <a href="#" className="hover:text-blue-400">
          EditAlgorithm
        </a>
        <a href="#" className="hover:text-blue-400">
          GetInformation
        </a>
      </>
    );
  } else if (isLoggedIn && userType == "user") {
    content = (
      <>
        <a href="#" className="hover:text-blue-400">
          CreatePassword
        </a>
        <a href="#" className="hover:text-blue-400">
          ListWebsites
        </a>
        <a href="#" className="hover:text-blue-400">
          GetPasswordInformation
        </a>
      </>
    );
  } else {
    content = (
      <div className="hidden md:flex space-x-4">
        <a href="#" className="hover:text-blue-400">
          Pricing
        </a>
        <a href="#" className="hover:text-blue-400">
          About
        </a>
      </div>
    );
  }

  return (
    <nav className="flex items-center justify-between p-4 bg-gray-800 text-white">
      {/* Left section with links */}
      <div className="flex space-x-4">
        <a href="#" className="text-2xl hover:text-blue-400">
          KeyPass
        </a>
        {content}
      </div>

      {/* Right section with login button */}
      <div>
        {isLoggedIn ? (
          <button className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded">
            Logout
          </button>
        ) : (
          <button className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded">
            Login
          </button>
        )}
      </div>
    </nav>
  );
};

export default Navbar;
