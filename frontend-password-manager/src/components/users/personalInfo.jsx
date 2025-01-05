import { React, useState, useEffect } from "react";
import { FaEye } from "react-icons/fa";
import { GetUserInfo } from "../../services/userService/userService";
import { useNavigate } from "react-router-dom";
import { CreateOTP } from "../../services/adminService/adminService";
import { FetchAllWesbite } from "../../services/userService/userService";
import { GetMastersInfo } from "../../services/masterService/masterservice";
const UserInfo = () => {
  const navigate = useNavigate();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isUpdateInfoModalOpen, setUpdateInfoModalOpen] = useState(false);
  const [productInfo, setProductInfo] = useState("");
  const [websiteList, setWebsiteList] = useState(0);
  const [userInformation, setUserInformation] = useState({
    userName: "",
    email: "",
  });

  const FetchWebsiteEntryOfUser = async () => {
    const response = await FetchAllWesbite();
    setWebsiteList(response.result.length)
  };

  const FetchMasterInformation = async () => {
    const response = await GetMastersInfo();
    setProductInfo(response.result.plan);
  };
  const FetchUserInfo = async () => {
    const response = await GetUserInfo();
    setUserInformation({
      userName: response.result.name,
      email: response.result.email,
    });
  };

  useEffect(() => {
    FetchUserInfo();
    FetchMasterInformation();
    FetchWebsiteEntryOfUser();
  }, []);
  const handleUpdateInfoClick = () => {
    setIsModalOpen(true);
    setUpdateInfoModalOpen(true);
  };
  const handleEyeClick = () => {
    setIsModalOpen(true);
    setUpdateInfoModalOpen(false);
  };

  const handleOTPCreation = async () => {
    const response = await CreateOTP();
    if (
      response.result.message === "OTP creation was successfull." ||
      response.result.message === "OTP already exists."
    ) {
      navigate("/otp");
    }
  };
  const handleCloseModal = () => setIsModalOpen(false);

  return (
    <>
      <section className="min-h-screen flex flex-col justify-start pt-20 px-4 sm:px-10">
        <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center">
          <h1 className="font-bold text-3xl sm:text-4xl md:text-5xl mb-4 sm:mb-0">
            Account Information
          </h1>
          <button
            onClick={handleUpdateInfoClick}
            className="bg-[#AA7DFF] text-white px-3 py-2 sm:px-4 sm:py-2 rounded-md hover:bg-[#C49DFF] mt-2 sm:mt-0"
          >
            Update Info
          </button>
        </div>

        {/* User Details Section */}
        <div className="flex flex-col md:flex-row justify-between items-center p-4 mt-10 border-b border-gray-200 space-y-4 md:space-y-0">
          <p className="text-lg font-medium break-words">
            {userInformation.userName}
          </p>
          <p className="text-lg font-medium break-words">{productInfo}</p>
          <button className="text-gray-500 hover:text-gray-700">
            <FaEye onClick={handleEyeClick} className="text-xl" />
          </button>
        </div>

        {/* Modal Section */}
        {isModalOpen && (
          <div className="fixed inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center z-50">
            {isUpdateInfoModalOpen ? (
              <div className="bg-white p-6 rounded-lg shadow-lg w-full max-w-lg">
                <div className="flex justify-between items-center mb-4">
                  <h2 className="text-2xl font-semibold">
                    Update User Information
                  </h2>
                  <button
                    onClick={handleCloseModal}
                    className="text-gray-500 hover:text-gray-700 text-lg"
                  >
                    ✕
                  </button>
                </div>
                <div className="space-y-6">
                  <div className="flex flex-col sm:flex-row justify-between items-center border-b border-gray-200 pb-4">
                    <div className="flex flex-col sm:flex-row items-start sm:items-center space-y-2 sm:space-y-0 sm:space-x-4">
                      <label className="block text-sm font-medium">
                        Password:
                      </label>
                      <input
                        type="password"
                        value="********"
                        readOnly
                        className="text-lg px-2 py-1 w-full sm:w-auto"
                      />
                    </div>
                    <button
                      onClick={() => {
                        localStorage.setItem("path", "resetPassword");
                        localStorage.setItem("updateType", "password");
                        handleOTPCreation();
                      }}
                      className="bg-[#AA7DFF] text-white px-4 py-2 rounded-md hover:bg-[#C49DFF] mt-4 sm:mt-0"
                    >
                      Update Password
                    </button>
                  </div>
                  <div className="flex flex-col sm:flex-row justify-between items-center border-b border-gray-200 pb-4">
                    <div className="flex flex-col sm:flex-row items-start sm:items-center space-y-2 sm:space-y-0 sm:space-x-4">
                      <label className="block text-sm font-medium">
                        Special Key:
                      </label>
                      <input
                        type="password"
                        value="********"
                        readOnly
                        className="text-lg px-2 py-1 w-full sm:w-auto"
                      />
                    </div>
                    <button
                      onClick={() => {
                        localStorage.setItem("path", "resetPassword");
                        localStorage.setItem("updateType", "key");
                        handleOTPCreation();
                      }}
                      className="bg-[#AA7DFF] text-white px-4 py-2 rounded-md hover:bg-[#C49DFF] mt-4 sm:mt-0"
                    >
                      Update Key
                    </button>
                  </div>
                </div>
              </div>
            ) : (
              <div className="bg-white p-6 rounded-lg shadow-lg w-full max-w-md">
                <div className="flex justify-between items-center mb-4">
                  <h2 className="text-2xl font-semibold">User Information</h2>
                  <button
                    onClick={handleCloseModal}
                    className="text-gray-500 hover:text-gray-700 text-lg"
                  >
                    ✕
                  </button>
                </div>
                <div className="space-y-4">
                  <div>
                    <label className="block text-sm font-medium">Name:</label>
                    <p className="text-lg">{userInformation.userName}</p>
                  </div>
                  <div>
                    <label className="block text-sm font-medium">Entry:</label>
                    <p className="text-lg">{websiteList}</p>
                  </div>
                  <div>
                    <label className="block text-sm font-medium">Email:</label>
                    <p className="text-lg">{userInformation.email}</p>
                  </div>
                </div>
              </div>
            )}
          </div>
        )}
      </section>
    </>
  );
};

export default UserInfo;
