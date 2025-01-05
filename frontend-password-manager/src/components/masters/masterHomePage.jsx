import { React, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { LoginAnimation } from "../animation";
import {
  GetAllUsers,
  GetMastersInfo,
} from "../../services/masterService/masterservice";
const MasterHomePage = () => {
  const [isAnimating, setIsAnimating] = useState(true);
  const [totalUser, setTotalUser] = useState(0);
  const [masterInfo, setMasterInfo] = useState({
    plan: "",
    count: 0,
    algorithm: "",
    email: "",
  });
  let navigate = useNavigate();

  const plans = {
    "basic plan": 10,
    "pro plan": 50,
    "premium plan": 100,
  };

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsAnimating(false);
      let getUserType = localStorage.getItem("userType");
      if (getUserType === "user") {
        navigate("/error");
      }
      getAlgorithm();
      getUserDetails();
    }, 1000);

    return () => clearTimeout(timer);
  }, [navigate]);

  const getAlgorithm = async () => {
    const response = await GetMastersInfo();
    console.log(response.result.algorithm);
    setMasterInfo((prevState) => ({
      ...prevState,
      algorithm: response.result.algorithm,
      plan: response.result.plan,
      count: response.result.count,
      email: response.result.email,
    }));
  };

  const getUserDetails = async () => {
    const response = await GetAllUsers();
    setTotalUser(response.result.length);
  };

  const handlerUserClick = () => {
    navigate("/master/adduser");
  };

  const handleConfigClick = () => {
    navigate("/master/editconfig");
  };

  const handleInfoClick = () => {
    navigate("/master/info");
  };

  return (
    <>
      {isAnimating ? (
        <section className="flex items-center justify-center min-h-screen bg-[#d1d1d1]">
          <LoginAnimation />
        </section>
      ) : (
        <section className="bg-[#8f8a81] min-h-screen flex flex-col items-center justify-center pt-16">
          <h2 className="text-2xl md:text-3xl lg:text-4xl font-bold mb-4">
            Welcome Back !
          </h2>
          <span className="text-black">Get Back where you left</span>
          <div className="flex flex-col md:flex-row gap-6 md:gap-8 mt-4">
            {/* First Card */}
            <div className="bg-white bg-opacity-80 shadow-lg rounded-lg p-4 md:p-6 w-full md:w-64 text-left flex flex-col justify-between">
              <div>
                <h2 className="text-lg md:text-xl font-semibold mb-2 md:mb-4">
                  User Numbers
                </h2>
                <p className="text-sm md:text-md">
                  Total Users: <span className="font-bold">{totalUser}</span>
                </p>
                <p className="text-sm md:text-md">
                  User Left:{" "}
                  <span className="font-bold">
                    {plans[masterInfo.plan.toLowerCase()] - totalUser}
                  </span>
                </p>
              </div>
              <button
                onClick={handlerUserClick}
                className="bg-[#AA7DFF] rounded-xl text-white py-2 px-2 mt-4 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300"
              >
                User Information
              </button>
            </div>

            {/* Second Card */}
            <div className="bg-white bg-opacity-80 shadow-lg rounded-lg p-4 md:p-6 w-full md:w-64 text-left flex flex-col justify-between">
              <div>
                <h2 className="text-lg md:text-xl font-semibold mb-2 md:mb-4">
                  Config Info
                </h2>
                <p className="text-sm md:text-md">
                  Special Key:{" "}
                  <span className="font-bold">
                    {masterInfo.count === 1 ? "Updated Done" : "Please update"}
                  </span>
                </p>
                <p className="text-sm md:text-md">
                  Algorithm Used:{" "}
                  <span className="font-bold">{masterInfo.algorithm}</span>
                </p>
              </div>
              <button
                onClick={handleConfigClick}
                className="bg-[#AA7DFF] rounded-xl text-white py-2 px-2 mt-4 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300"
              >
                Config Information
              </button>
            </div>

            {/* Third Card */}
            <div className="bg-white bg-opacity-80 shadow-lg rounded-lg p-4 md:p-6 w-full md:w-64 text-left flex flex-col justify-between">
              <div>
                <h2 className="text-lg md:text-xl font-semibold mb-2 md:mb-4">
                  Master Information
                </h2>
                <p className="text-sm md:text-md">
                  Plan: <span className="font-bold">{masterInfo.plan}</span>
                </p>
                <p className="text-sm md:text-md">
                  Email: <span className="font-bold">{masterInfo.email}</span>
                </p>
              </div>
              <button
                onClick={handleInfoClick}
                className="bg-[#AA7DFF] rounded-xl text-white py-2 px-2 mt-4 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300"
              >
                Master Information
              </button>
            </div>
          </div>
        </section>
      )}
    </>
  );
};

export default MasterHomePage;
