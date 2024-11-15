import { React, useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { FaCheck } from "react-icons/fa6";
import { CreateProduct } from "../services/intialService/service"; // Ensure this service is correctly set up.

const Request = () => {
  const location = useLocation();
  let navigate = useNavigate();
  const [type, setType] = useState("");
  const [formData, setFormData] = useState({
    email: "",
    username: "",
  });
  const [showDropdown, setShowDropdown] = useState(false);

  const options = ["Demo Product", "Basic Plan", "Pro Plan", "Premium Plan"];

  useEffect(() => {
    if (!location.state) {
      setType("Demo Product");
    } else {
      setType(location.state.planType);
    }
  }, [location.state]);

  const handleFormData = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleOptionClick = (option) => {
    setType(option);
    setShowDropdown(false);
  };

  const handlerSubmit = async () => {
    console.log("the form data = ", formData);
    let resp = await CreateProduct(type,formData.username, formData.email);
    console.log("the response form the api = ", resp);
    if (resp.result === "success") {
      navigate("/success")
    } else {
      navigate("/error")
    }
  };

  const isFormValid = () => {
    return formData.email.includes("@") && formData.username.trim() !== "";
  };

  return (
    <>
      <section className="min-h-screen flex items-center justify-center pt-16">
        <div className="max-w-[500px] px-10 py-10 rounded-3xl bg-white border-2 border-gray-100">
          <h1 className="text-5xl font-semibold text-center">Welcome</h1>
          <p className="font-medium text-lg text-gray-500 mt-4 text-center">
            Please provide the following information to set up a dedicated
            server tailored to your needs.
          </p>
          <div className="mt-8">
            <div className="flex flex-col">
              <label className="text-lg font-medium">Email</label>
              <input
                type="email"
                name="email" 
                value={formData.email}
                onChange={handleFormData}
                className="w-full border-2 border-gray-100 rounded-xl p-4 mt-1 bg-transparent"
                placeholder="Enter your Email"
                required 
              />
            </div>
            <div className="flex flex-col mt-4">
              <label className="text-lg font-medium">Username</label>
              <input
                name="username"
                value={formData.username}
                onChange={handleFormData}
                className="w-full border-2 border-gray-100 rounded-xl p-4 mt-1 bg-transparent"
                placeholder="Enter your username"
                required
              />
            </div>
            <div className="flex flex-col mt-4 relative">
              <label className="text-lg font-medium">Product Type</label>
              <input
                value={type}
                onClick={() => setShowDropdown(!showDropdown)}
                className="w-full border-2 border-gray-100 rounded-xl p-4 mt-1 bg-transparent cursor-pointer"
                readOnly
              />
              {/* Dropdown options */}
              {showDropdown && (
                <div className="absolute top-full mt-2 w-full border border-gray-100 bg-white rounded-lg shadow-lg z-10">
                  {options.map((option) => (
                    <div
                      key={option}
                      onClick={() => handleOptionClick(option)}
                      className="flex items-center justify-between px-4 py-2 cursor-pointer hover:bg-gray-100"
                    >
                      <span>{option}</span>
                      {option === type && (
                        <FaCheck className="text-purple-500" />
                      )}
                    </div>
                  ))}
                </div>
              )}
            </div>
            <div className="mt-8 flex flex-col gap-y-4">
              <button
                onClick={handlerSubmit}
                disabled={!isFormValid()}
                className={`active:scale-[.98] active:duration-75 transition-all hover:scale-[1.01] ease-in-out transform py-4 rounded-xl text-white font-bold text-lg ${isFormValid() ? 'bg-[#AA7DFF] hover:bg-[#C49DFF]' : 'bg-gray-300 cursor-not-allowed'}`}
              >
                Request a demo
              </button>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default Request;
