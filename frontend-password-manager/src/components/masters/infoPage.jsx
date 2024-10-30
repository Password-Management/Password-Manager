import React, { useState } from "react";
import { FaEye } from "react-icons/fa";
import { useNavigate } from "react-router-dom";


const InfoPage = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isChecked, setIsChecked] = useState(false);
  const [isData, setIsData] = useState(false);
  const [isUpdateInfoModalOpen, setUpdateInfoModalOpen] = useState(false);
  const navigate = useNavigate();
  const handleCheckboxChange = () => {
    setIsChecked(!isChecked); // Toggle the checkbox state
  };

  const handleUpdateInfoClick = () => {
    setIsModalOpen(true);
    setUpdateInfoModalOpen(true); // Set for Update Info
  };

  const handleEyeClick = () => {
    setIsModalOpen(true);
    setUpdateInfoModalOpen(false); // Set for User Information
  };

  const handleCloseModal = () => setIsModalOpen(false);
  const handleUpdateAlgorithm = () => navigate("/master/editconfig");

  return (
    <>
      <section className="min-h-screen flex flex-col justify-start pt-20 px-4 sm:px-10">
        <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center">
          <h1 className="font-bold text-3xl sm:text-4xl md:text-5xl mb-4 sm:mb-0">
            Master Information
          </h1>
          <button
            onClick={handleUpdateInfoClick}
            className="bg-[#AA7DFF] text-white px-3 py-2 sm:px-4 sm:py-2 rounded-md hover:bg-[#C49DFF] mt-2 sm:mt-0"
          >
            Update Info
          </button>
        </div>

        {/* User Info Row with Eye Icon */}
        <div className="flex justify-between items-center p-4 mt-10 border-b border-gray-200">
          <p className="text-lg font-medium">Vivek Sharma</p>
          <p className="text-lg font-medium">Demo Product</p>
          <button
            onClick={handleEyeClick}
            className="text-gray-500 hover:text-gray-700"
          >
            <FaEye className="text-xl" />
          </button>
        </div>

        {/* Modal */}
        {isModalOpen && (
          <div className="fixed inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center z-65">
            {isUpdateInfoModalOpen ? (
              <>
                <div className="bg-white p-6 rounded-lg shadow-lg w-full max-w-lg h-[28rem]">
                  <div className="flex justify-between items-center mb-4">
                    <h2 className="text-2xl font-semibold">
                      Update Master Info ?
                    </h2>
                    <button
                      onClick={handleCloseModal}
                      className="text-gray-500 hover:text-gray-700 text-lg"
                    >
                      ✕
                    </button>
                  </div>

                  <div className="flex flex-col space-y-3 mb-4">
                    <div className="flex-grow">
                      <p>
                        Please note that if you choose to update the master
                        information, the existing data in our database will be
                        lost. However, rest assured that your subscription will
                        remain active. To preserve your database, kindly accept
                        the checkbox below.
                      </p>
                      <p className="mt-3">
                        If the 'Store Data' option is selected, it indicates
                        that either the Email or Username will be updated, but
                        not the Algorithm.
                      </p>
                    </div>

                    <div className="flex items-center space-x-5">
                      <input
                        type="checkbox"
                        checked={isChecked}
                        onChange={handleCheckboxChange}
                        className="form-checkbox h-5 w-5 text-[#AA7DFF] focus:ring-opacity-75"
                      />
                      <label className="text-gray-700 font-medium">
                        Accept Terms and Conditions
                      </label>
                    </div>

                    <div className="flex items-center space-x-5">
                      <input
                        type="checkbox"
                        className="form-checkbox h-5 w-5 text-[#AA7DFF] focus:ring-opacity-75"
                        onClick={() => {
                          setIsData(true);
                        }}
                      />
                      <label className="text-gray-700 font-medium">
                        Store Data
                      </label>
                    </div>
                    <div className="flex items-center justify-center mt-4">
                      <button
                        disabled={!isChecked}
                        className={`px-3 py-2 sm:px-4 sm:py-2 rounded-md mt-3 sm:mt-0 w-32 mt-4 ${
                          isChecked
                            ? "bg-[#AA7DFF] text-white hover:bg-[#C49DFF] "
                            : "bg-gray-300 text-gray-500 cursor-not-allowed"
                        }`}
                      >
                        Submit
                      </button>
                    </div>
                  </div>
                </div>
              </>
            ) : (
              <div>
                <div className="bg-white p-6 rounded-lg shadow-lg w-80">
                  <div className="flex justify-between items-center mb-4">
                    <h2 className="text-2xl font-semibold">
                      {isUpdateInfoModalOpen
                        ? "Update Master Info ?"
                        : "Master Info"}
                    </h2>
                    <button
                      onClick={handleCloseModal}
                      className="text-gray-500 hover:text-gray-700 text-lg"
                    >
                      ✕
                    </button>
                  </div>
                  <div className="mb-4">
                    <label className="block text-sm font-medium">Name:</label>
                    <p className="text-lg">User Name</p>
                  </div>
                  <div className="mb-4">
                    <label className="block text-sm font-medium">
                      Algorithm:
                    </label>
                    <p className="text-lg inline-block mr-2">RSA</p>
                    <button
                      onClick={handleUpdateAlgorithm}
                      className="bg-[#AA7DFF] text-white px-3 py-1 rounded hover:bg-[#C49DFF] text-sm ml-40"
                    >
                      Update Algorithm
                    </button>
                  </div>
                  <div className="mb-4">
                    <label className="block text-sm font-medium">Email:</label>
                    <p className="text-lg">user@example.com</p>
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

export default InfoPage;
