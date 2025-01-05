import React, { useState, useEffect } from "react";
import { FaEye } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import { GetMastersInfo } from "../../services/masterService/masterservice";

const InfoPage = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [originalData, setOriginalData] = useState({ name: "", email: "" });
  const [isModified, setIsModified] = useState(false);
  const [masterInfo, setMasterInfo] = useState({
    algorithm: "",
    email: "",
    userName: "",
    productType: "",
    customerId: "",
  });
  const [isUpdateInfoModalOpen, setUpdateInfoModalOpen] = useState(false);
  const navigate = useNavigate();
  const handleMasterInfo = async () => {
    const response = await GetMastersInfo();
    setMasterInfo({
      algorithm: response.result.algorithm || "",
      email: response.result.email || "",
      userName: response.result.name || "",
      productType: response.result.plan || "",
      customerId: response.result.customer_id || "",
    });
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setMasterInfo((prevData) => ({ ...prevData, [name]: value }));
    setIsModified(value !== originalData[name]);
  };

  useEffect(() => {
    handleMasterInfo();
    setOriginalData(masterInfo);
  }, []);
  const handleUpdateInfoClick = () => {
    setIsModalOpen(true);
    setUpdateInfoModalOpen(true);
  };

  const handleEyeClick = () => {
    setIsModalOpen(true);
    setUpdateInfoModalOpen(false);
  };

  const handleCloseModal = () => setIsModalOpen(false);
  return (
    <>
      <section className="min-h-screen flex flex-col justify-start pt-20 px-4 sm:px-10">
        <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center">
          <h1 className="font-bold text-3xl sm:text-4xl md:text-5xl mb-4 sm:mb-0">
            Personal Information
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
          <p className="text-lg font-medium">{masterInfo.userName}</p>
          <p className="text-lg font-medium">{masterInfo.productType}</p>
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
                        information, the existing data in our database will not be
                        lost. However, rest assured that your subscription will
                        remain active.
                      </p>
                      <p className="mt-3">
                        This Update box can update your{" "}
                        <strong>Email</strong>
                      </p>
                    </div>
                    <div className="flex flex-col items-start p-4  max-w-sm mx-auto">
                      
                      <div className="mb-4 w-full">
                        <label className="block text-gray-700 text-sm font-medium mb-1">
                          Email:
                        </label>
                        <input
                          type="email"
                          name="email"
                          value={masterInfo.email}
                          onChange={handleChange}
                          className="w-full px-3 py-2 border rounded-md bg-white text-gray-800"
                        />
                      </div>
                    </div>
                    <div className="flex items-center justify-center mt-4">
                      <button
                        disabled={!isModified}
                        className={`px-3 py-2 sm:px-4 sm:py-2 rounded-md  sm:mt-0 w-32 mt-3 ${
                          isModified
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
                      Master Information
                    </h2>
                    <button
                      onClick={handleCloseModal}
                      className="text-gray-500 hover:text-gray-700 text-lg"
                    >
                      ✕
                    </button>
                  </div>
                  <div className="mb-4">
                    <label className="block text-lg font-medium">Name:</label>
                    <p className="text-sm">{masterInfo.userName}</p>
                  </div>
                  <div className="mb-4">
                    <label className="block text-lg font-medium">
                      Algorithm:
                    </label>
                    <p className="text-sm inline-block mr-2">
                      {masterInfo.algorithm}
                    </p>
                  </div>
                  <div className="mb-4">
                    <label className="block text-lg font-medium">Email:</label>
                    <p className="text-sm">{masterInfo.email}</p>
                  </div>
                  <div className="mb-4">
                    <label className="block text-lg font-medium">
                      CustomerId:
                    </label>
                    <p className="text-sm">{masterInfo.customerId}</p>
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
