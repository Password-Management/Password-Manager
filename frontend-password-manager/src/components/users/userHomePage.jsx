import { React, useEffect } from "react";
import { useNavigate } from "react-router-dom";
const UserHomePage = () => {
  let navigate = useNavigate();
  useEffect(() => {
    let getUserType = localStorage.getItem("userType");
    console.log(getUserType);
    if (getUserType === "master") {
      navigate("/error");
    }
  });

  return (
    <>
      <section className="bg-[#d1d1d1] min-h-screen flex flex-col items-center justify-center pt-16">
        <h2 className="text-2xl md:text-3xl lg:text-4xl font-bold mb-4">
          Welcome Back !
        </h2>
        <span className="text-Black">Get Back where you left off!!</span>
        <div className="flex flex-col md:flex-row gap-6 md:gap-8 mt-4">
          {/* First Card */}
          <div className="bg-white bg-opacity-80 shadow-lg rounded-lg p-4 md:p-6 w-full md:w-64 text-left">
            <h2 className="text-lg md:text-xl font-semibold mb-2 md:mb-4">Add Website</h2>
            <p className="text-sm md:text-md">Total Users: <span className="font-bold">10</span></p>
            <p className="text-sm md:text-md">User Left: <span className="font-bold">0</span></p>
            <button onClick={()=>{navigate("/user/addpassword")}}className="bg-[#AA7DFF] rounded-xl text-white py-2 px-2 mt-2 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300">Add Password</button>
          </div>

          {/* Second Card */}
          <div className="bg-white bg-opacity-80 shadow-lg rounded-lg p-4 md:p-6 w-full md:w-64 text-left">
            <h2 className="text-lg md:text-xl font-semibold mb-2 md:mb-4">Personal Information</h2>
            <p className="text-sm md:text-md">Special Key: <span className="font-bold">Updated Done</span></p>
            <p className="text-sm md:text-md">Algorithm Used: <span className="font-bold">RSA</span></p>
            <button onClick={()=>{navigate("/user/info")}}className="bg-[#AA7DFF] rounded-xl text-white py-2 px-2 mt-2 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300">User Information</button>
          </div>
        </div>
      </section>
    </>
  );
};

export default UserHomePage;
