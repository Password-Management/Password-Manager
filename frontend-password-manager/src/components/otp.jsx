import React, { useState } from "react";
import ResetImage from "../assets/otp-main.jpg";
import { useNavigate } from "react-router-dom";
import { VerifyOTP } from "../services/adminService/adminService";

const Otp = () => {
  const [otpInput, setOtpInput] = useState("");
  const [message, setMessage] = useState(null); 
  const navigate = useNavigate();

  const handleOtpSubmit = async (e) => {
    e.preventDefault();
    const otpDetails = await VerifyOTP(otpInput)
    console.log(otpDetails)
    if (otpDetails.result.message == "SUCCESS") {
      localStorage.setItem("path", "resetPassword");
      navigate("/reset")
    } else {
      setMessage({ type: "error", text: "OTP verification failed!" });
    }
  };

  return (
    <section className="flex h-screen items-center justify-center">
      <div className="flex flex-col md:flex-row rounded-2xl p-5 md:p-8 items-center max-w-5xl w-full">
        {/* Form Section */}
        <div className="w-full md:w-1/2 px-6 md:px-8 mr-10">
          <form
            className="flex flex-col gap-4 mt-6 items-center"
            onSubmit={handleOtpSubmit}
          >
            <div className="text-lg font-medium mb-4 text-center">
              Enter OTP for Verification
            </div>
            <input
              className="p-2 rounded-xl border focus:outline-none focus:ring focus:border-blue-500"
              type="text"
              name="OTP"
              placeholder="Enter OTP"
              value={otpInput}
              onChange={(e) => setOtpInput(e.target.value)} // Update input state
            />
            <button
              className="bg-[#AA7DFF] rounded-xl text-white py-2 px-8 hover:scale-105 hover:bg-[#C49DFF] transition-transform duration-300 px-10"
              type="submit"
            >
              Confirm Details
            </button>
            <button
              className="bg-white text-purple-600 border border-purple-300 px-4 py-2 rounded-xl hover:bg-[#C49DFF] hover:text-white transition duration-300 px-10"
              type="submit"
            >
              Resend OTP
            </button>
          </form>
          {/* Feedback Message */}
          {message && (
            <div
              className={`mt-4 text-center ${
                message.type === "success" ? "text-green-500" : "text-red-500"
              }`}
            >
              {message.text}
            </div>
          )}
        </div>

        {/* Image Section */}
        <div className="hidden md:flex w-full md:w-1/2 h-full items-center justify-center">
          <img
            className="rounded-2xl object-cover w-full h-auto max-w-full"
            src={ResetImage}
            alt="Login"
          />
        </div>
      </div>
    </section>
  );
};

export default Otp;
