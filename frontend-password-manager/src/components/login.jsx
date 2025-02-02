import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { MasterLogin, UserLogin } from "../services/loginService/loginService";
const LoginPage = () => {
  let navigate = useNavigate();
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState("");
  const [userData, setUserData] = useState({
    email: "",
    password: "",
  });
  const [specialKey, setSpecialKey] = useState("");
  const userType = localStorage.getItem("userType");
  const togglePasswordVisibility = () => {
    setShowPassword((prevState) => !prevState);
  };
  const SuucessMasterLogin = async (e) => {
    e.preventDefault()
    const response = await MasterLogin(specialKey.password)
    console.log(response)
    if (response.result.message.includes("Login of master successfull") || response.result.message.includes("Reloging in user")) {
      localStorage.setItem("isLoggedIn", true);
      localStorage.setItem("navbarUserType", "master");
      localStorage.setItem("specialKey", specialKey.password)
      localStorage.setItem("masterId", response.result.master_id);
      navigate("/masterhomepage");
    } else {
      setError("Master Not found or key provided is incorrect")
    }
  };

  const SuccessUserLogin = async (e) => {
    e.preventDefault();
    const response = await UserLogin(userData.email, userData.password);
    if (response.result.message.includes("User Logged in successfully") || response.result.message.includes("Reloging in user")) {
      localStorage.setItem("navbarUserType", "user");
      localStorage.setItem("isLoggedIn", true);
      localStorage.setItem("masterId", response.result.master_id);
      localStorage.setItem("userId", response.result.user_id);
      navigate("/userhomepage");
    } else if (response.result.message === "Password is incorrect") {
      setError("Password you provided is wrong please check your password");
    } else if (response.result.message === "User not found") {
      setError("User doesnt exist please check your email or password");
    } else {
      setError("Please check the password or email you provided");
    }
  };

  return (
    <>
      <section className="flex h-screen items-center justify-center">
        <div className="flex flex-col md:flex-row rounded-2xl p-5 md:p-8 items-center max-w-5xl w-full">
          <div className="w-full md:w-1/2 px-6 md:px-8 mr-10">
            <h2 className="font-bold text-3xl text-[#181B1E] text-center md:text-left">
              Login
            </h2>
            {userType === "user" ? (
              <form className="flex flex-col gap-4 mt-6">
                <input
                  className="p-2 rounded-xl border focus:outline-none focus:ring focus:border-blue-500"
                  type="email"
                  name="email"
                  placeholder="Email"
                  onChange={(e) =>
                    setUserData({
                      ...userData,
                      [e.target.name]: e.target.value,
                    })
                  }
                />
                <div className="relative">
                  <input
                    className="p-2 rounded-xl border w-full focus:outline-none focus:ring focus:border-blue-500"
                    type={showPassword ? "text" : "password"}
                    name="password"
                    placeholder="Password"
                    onChange={(e) =>
                      setUserData({
                        ...userData,
                        [e.target.name]: e.target.value,
                      })
                    }
                  />
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="gray"
                    className="bi bi-eye absolute top-1/2 right-3 -translate-y-1/2 cursor-pointer"
                    viewBox="0 0 16 16"
                    onClick={togglePasswordVisibility}
                  >
                    <path d="M16 8s-3-5.5-8-5.5S0 8 0 8s3 5.5 8 5.5S16 8 16 8zM1.173 8a13.133 13.133 0 0 1 1.66-2.043C4.12 4.668 5.88 3.5 8 3.5c2.12 0 3.879 1.168 5.168 2.457A13.133 13.133 0 0 1 14.828 8c-.058.087-.122.183-.195.288-.335.48-.83 1.12-1.465 1.755C11.879 11.332 10.119 12.5 8 12.5c-2.12 0-3.879-1.168-5.168-2.457A13.134 13.134 0 0 1 1.172 8z" />
                    <path d="M8 5.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5zM4.5 8a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0z" />
                  </svg>
                </div>
                {error && <p className="text-red-500 text-sm">{error}</p>}
                <button
                  onClick={SuccessUserLogin}
                  className="bg-[#AA7DFF] rounded-xl text-white py-2 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300"
                >
                  Login
                </button>

                <button
                  onClick={() => localStorage.setItem("userType", "master")}
                  className="bg-[#AA7DFF] rounded-xl text-white py-2 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300"
                >
                  Login As Master
                </button>
              </form>
            ) : (
              <form className="flex flex-col gap-4 mt-6">
                <div className="relative">
                  <input
                    className="p-2 rounded-xl border w-full focus:outline-none focus:ring focus:border-blue-500"
                    type={showPassword ? "text" : "password"}
                    name="password"
                    placeholder="Special Key"
                    onChange={(e) =>
                      setSpecialKey({
                        ...specialKey,
                        [e.target.name]: e.target.value,
                      })
                    }
                  />
                </div>
                {error && <p className="text-red-500 text-sm">{error}</p>}
                <button
                  onClick={SuucessMasterLogin}
                  className="bg-[#AA7DFF] rounded-xl text-white py-2 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300"
                >
                  Login
                </button>

                <button
                  onClick={() => localStorage.setItem("userType", "user")}
                  className="bg-[#AA7DFF] rounded-xl text-white py-2 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300"
                >
                  Login As User
                </button>
              </form>
            )}

            <Link to={"/resetpassword"}>
              {userType === "user" ? (
                <div className="mt-5 text-xs text-center md:text-left  py-4 text-[#181B1E]">
                  <span>Forgot your Password ?</span>
                </div>
              ) : (
                <div className="mt-5 text-xs text-center md:text-left  py-4 text-[#181B1E]">
                  <span>Forgot your API Key ?</span>
                </div>
              )}
            </Link>
          </div>
          <div className="hidden relative w-1/2 h-full lg:flex items-center justify-center">
            <div className="w-72 h-72 rounded-full bg-gradient-to-tr from-violet-500 to-pink-500 animate-spin" />
            <div className="w-full h-1/2 absolute bottom-0 bg-white/10 backdrop-blur-lg" />
          </div>
        </div>
      </section>
    </>
  );
};

export default LoginPage;
