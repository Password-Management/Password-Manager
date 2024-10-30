import { React, useEffect } from "react";
import { useNavigate } from "react-router-dom";
const UserHomePage = () => {
  let navigate = useNavigate();
  useEffect(() => {
    let getUserType = localStorage.getItem("userType");
    console.log(getUserType)
    if (getUserType === "master") {
      navigate("/error");
    }
  });
  
  return (
    <>
      <section className="flex h-screen items-center justify-center">
        <h1>UserHomePage</h1>
      </section>
    </>
  );
};

export default UserHomePage;
