import { React, useEffect } from "react";
import HomePageImage from "../assets/Home.png";
import { Link } from "react-router-dom";

const Home = () => {
  useEffect(() => {
    localStorage.setItem("userType", "user");
    localStorage.setItem("navbarUserType", "");
    localStorage.setItem("isLoggedIn", false);
  }, []);

  return (
    <>
      <section className="min-h-screen flex items-center justify-center pt-16">
        <div className="container flex flex-col-reverse lg:flex-row items-center gap-10">
          <div className="flex flex-1 flex-col items-center lg:items-start">
            <h2 className="text-black text-3xl md:text-4xl lg:text-5xl text-center lg:text-left mb-6">
              KeyPass
            </h2>
            <p className="text-gray-600 text-lg text-center lg:text-left mb-6">
              At KeyPass, your data security is our top priority. We use
              advanced RSA encryption to keep your passwords safe and secure.
            </p>
            <div className="flex justify-center flex-wrap gap-6">
              <Link to={"/login"}>
                <button
                  type="button"
                  className="bg-[#AA7DFF] hover:bg-[#C49DFF] text-white px-4 py-2 rounded transition duration-300"
                >
                  User Login
                </button>
              </Link>
              <Link to={"/login"}>
                <button
                  onClick={() => localStorage.setItem("userType", "master")}
                  type="button"
                  className="bg-white text-purple-600 border border-purple-300 px-4 py-2 rounded hover:bg-[#C49DFF] hover:text-white transition duration-300"
                >
                  Master Login
                </button>
              </Link>
            </div>
          </div>
          <div className="flex justify-center flex-1 mb-10 md:mb-16 lg:mb-0 z-10">
            <img
              className="w-5/6 h-auto sm:w-3/4 md:w-full"
              src={HomePageImage}
              alt="HomePage"
            />
          </div>
        </div>
      </section>
    </>
  );
};

export default Home;
