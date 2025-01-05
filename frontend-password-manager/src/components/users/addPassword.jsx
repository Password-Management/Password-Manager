import React, { useEffect, useState } from "react";
import { RiNetflixFill } from "react-icons/ri";
import { FaFacebook } from "react-icons/fa";
import { BsInstagram } from "react-icons/bs";
import { BiLogoGmail } from "react-icons/bi";
import { RiLockPasswordFill } from "react-icons/ri";
import { FaGoogle } from "react-icons/fa";
import { FaDiscord } from "react-icons/fa6";
import { IoLogoGithub } from "react-icons/io";
import { SiPrime } from "react-icons/si";
import { FaApple } from "react-icons/fa";
import { FaSnapchatGhost } from "react-icons/fa";
import { SiPaytm } from "react-icons/si";
import { SiPhonepe } from "react-icons/si";
import { SiNike } from "react-icons/si";
import { MdOutlineAdd } from "react-icons/md";
import { DecryptAnimation } from "../animation";
import { DeleteAnimation } from "../animation";
import {
  AddWebsite,
  FetchAllWesbite,
  DeleteWebsite,
  VerifyAuthKey,
  DecryptPassword,
} from "../../services/userService/userService";

const AddPassword = () => {
  const [passwordData, setPasswordData] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isfeatureModelOpen, setIsFeatureModelOpen] = useState(false);
  const [isDecryptModalOpen, setIsDecryptModalOpen] = useState(false);
  const [password, setPassword] = useState("");
  const [DecryptModal, setDecryptModal] = useState(false);
  const [authKey, setAuthKey] = useState("");
  const [websiteName, setWebsiteName] = useState("");
  const [error, setError] = useState("");
  const [isDecryptAnimation, setIsDecryptAnimation] = useState(false);
  const [isAnimationPlaying, setIsAnimationPlaying] = useState(false);
  const [newCredentials, setNewCredentials] = useState({
    websiteName: "",
    userName: "",
    password: "",
  });
  const websiteData = {
    netflix: <RiNetflixFill />,
    facebook: <FaFacebook />,
    instagram: <BsInstagram />,
    gmail: <BiLogoGmail />,
    google: <FaGoogle />,
    discord: <FaDiscord />,
    github: <IoLogoGithub />,
    prime: <SiPrime />,
    apple: <FaApple />,
    snapchat: <FaSnapchatGhost />,
    paytm: <SiPaytm />,
    phonePe: <SiPhonepe />,
    nike: <SiNike />,
  };

  const handleAddClick = () => {
    setIsModalOpen(true);
  };

  const handleSave = async () => {
    if (newCredentials.websiteName && newCredentials.userName) {
      setNewCredentials({ websiteName: "", userName: "", password: "" });
      try {
        let websiteEntry = await AddWebsite(
          newCredentials.userName,
          newCredentials.password,
          newCredentials.websiteName
        );
        console.log("Website added successfully:", websiteEntry);
        FetchWebsiteData();
      } catch (error) {
        console.error("Failed to add website:", error);
      }
      setIsModalOpen(false);
    } else {
      alert("Please fill out all fields.");
    }
  };

  const handleDecryptAnimation = async () => {
    setError("");
    const response = await VerifyAuthKey(authKey);
    console.log(response.result);
    const passwordResponse = await DecryptPassword(websiteName);
    if (response.result.message === "Success") {
      setIsAnimationPlaying(true);
      setIsDecryptAnimation(true);
      setIsDecryptModalOpen(false);
      setIsFeatureModelOpen(false);
      setPassword(passwordResponse.result.message);
      setTimeout(() => {
        setIsAnimationPlaying(false);
        setDecryptModal(true)
      }, 5000);
    } else {
      console.log("Wrong key");
      setError("Invalid AuthKey !!!!");
    }
  };

  const handleDeletAnimation = () => {
    setIsAnimationPlaying(true);
    setIsDecryptAnimation(false);
    setIsDecryptModalOpen(false);
    setIsFeatureModelOpen(false);
    setTimeout(() => {
      const deleteWebsiteEntry = async () => {
        let response = await DeleteWebsite(websiteName);
        console.log(response.result);
      };
      deleteWebsiteEntry();
      window.location.reload();
      setIsAnimationPlaying(false);
    }, 5500);
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setNewCredentials((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleDecrypt = () => {
    setIsFeatureModelOpen(true);
    setIsDecryptModalOpen(true);
  };

  const handleDelete = () => {
    console.log("Handle delete is called");
    setIsFeatureModelOpen(true);
    setIsDecryptModalOpen(false);
  };

  const FetchWebsiteData = async () => {
    console.log("Fetch data is called");
    try {
      const response = await FetchAllWesbite();
      const validData = Array.isArray(response.result)
        ? response.result
            .filter(
              (item) =>
                item.website_name && typeof item.website_name === "string"
            )
            .map((item) => ({
              website_name: item.website_name,
              user_name: item.user_name,
            }))
        : [];
      if (validData.length === 0) {
        return (
          <>
            <h3 className="text-med font-semibold mb-3">Add the Password</h3>
          </>
        );
      }

      setPasswordData(validData);
    } catch (error) {
      console.error("Error fetching website data:", error);

      setPasswordData([]);
    }
  };

  useEffect(() => {
    FetchWebsiteData();
  }, []);

  return (
    <section className="min-h-screen flex flex-col justify-start pt-20 px-4 sm:px-10 bg-gray-200">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center">
        <h1 className="font-bold text-3xl sm:text-4xl md:text-5xl mb-4 sm:mb-0">
          Add Credentials
        </h1>
        <button
          onClick={handleAddClick}
          className="flex items-center justify-center md:w-auto bg-[#AA7DFF] hover:bg-[#C49DFF] py-2 px-4 text-white rounded-md transition-colors"
        >
          <MdOutlineAdd className="mr-2" /> Add Website
        </button>
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6 mt-8">
        {passwordData.map((item, index) => (
          <div
            key={index}
            className="bg-white shadow-lg rounded-lg p-6 flex flex-col items-center"
          >
            <div className="text-4xl mb-2">
              {websiteData[item.website_name.toLowerCase()] || (
                <RiLockPasswordFill />
              )}
            </div>
            <h2 className="text-lg font-semibold mb-4">{item.website_name}</h2>
            <h3 className="text-med font-semibold mb-3">{item.user_name}</h3>
            <div className="flex space-x-4 mt-4">
              <button
                onClick={() => {
                  setWebsiteName(item.website_name);
                  handleDecrypt();
                }}
                className="bg-[#AA7DFF] text-white px-4 py-2 rounded hover:bg-[#C49DFF] transition duration-300"
              >
                Decrypt
              </button>
              <button
                onClick={() => {
                  setWebsiteName(item.website_name);
                  handleDelete();
                }}
                className="bg-white text-purple-600 border border-purple-300 px-4 py-2 rounded hover:bg-[#C49DFF] hover:text-white transition duration-300"
              >
                Delete
              </button>
            </div>
          </div>
        ))}
      </div>
      {isModalOpen ? (
        <>
          <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
            <div className="bg-white p-6 rounded-md shadow-md w-full max-w-2xl">
              <h2 className="text-xl font-semibold mb-4">Add New Credential</h2>
              <div className="space-y-4">
                <input
                  type="text"
                  className="w-full py-2 px-4 border rounded-md"
                  placeholder="Website Name"
                  name="websiteName"
                  value={newCredentials.websiteName}
                  onChange={handleInputChange}
                  required
                />
                <input
                  type="text"
                  className="w-full py-2 px-4 border rounded-md"
                  placeholder="User Name"
                  name="userName"
                  value={newCredentials.userName}
                  onChange={handleInputChange}
                  required
                />
                <input
                  type="password"
                  className="w-full py-2 px-4 border rounded-md"
                  name="password"
                  placeholder="Password"
                  value={newCredentials.password}
                  onChange={handleInputChange}
                  required
                />
              </div>
              <div className="flex justify-end space-x-4 mt-4">
                <button
                  onClick={() => {
                    setNewCredentials({ websiteName: "", userName: "" });
                    setIsModalOpen(false);
                  }}
                  className="bg-white text-purple-600 border border-purple-300 px-4 py-2 rounded hover:bg-[#C49DFF] hover:text-white transition duration-300"
                >
                  Cancel
                </button>
                <button
                  onClick={handleSave}
                  className="py-2 px-4 bg-[#AA7DFF] hover:bg-[#C49DFF] text-white rounded-md"
                >
                  Save
                </button>
              </div>
            </div>
          </div>
        </>
      ) : (
        <></>
      )}
      {isfeatureModelOpen && (
        <div className="fixed inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center z-50">
          {isDecryptModalOpen ? (
            <>
              <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-70">
                <div className="bg-white p-6 rounded-md shadow-md  max-w-2xl">
                  <h2 className="text-xl font-semibold mb-4">Decrypt</h2>
                  <div className="mt-5 text-s text-center md:text-left  py-4 text-[#181B1E]">
                    <span>
                      For Decrypting the password for {websiteName} please enter
                      your AuthKey.
                    </span>
                  </div>
                  <input
                    type="text"
                    className="w-full py-2 px-4 border rounded-md"
                    placeholder="Auth Key"
                    onChange={(e) => setAuthKey(e.target.value)}
                    required
                  />
                  <div className="flex space-x-4 mt-4">
                    <button
                      onClick={handleDecryptAnimation}
                      className="bg-[#AA7DFF] text-white px-4 py-2 rounded hover:bg-[#C49DFF] transition duration-300"
                    >
                      Confirm
                    </button>
                    <button
                      onClick={() => {
                        setIsFeatureModelOpen(false);
                        setWebsiteName("");
                      }}
                      className="bg-white text-purple-600 border border-purple-300 px-4 py-2 rounded hover:bg-[#C49DFF] hover:text-white transition duration-300"
                    >
                      Cancel
                    </button>
                  </div>
                  {error && (
                    <p className="text-red-500 text-sm mt-2">{error}</p>
                  )}
                </div>
              </div>
            </>
          ) : (
            <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-70">
              <div className="bg-white p-6 rounded-md shadow-md  max-w-2xl">
                <h2 className="text-xl font-semibold mb-4">Delete</h2>
                <div className="mt-5 text-s text-center md:text-left  py-4 text-[#181B1E]">
                  <span>
                    Are you sure you want to delete entry for {websiteName} ?{" "}
                  </span>
                </div>
                <div className="flex space-x-4 mt-4">
                  <button
                    onClick={handleDeletAnimation}
                    className="bg-[#AA7DFF] text-white px-4 py-2 rounded hover:bg-[#C49DFF] transition duration-300"
                  >
                    Confirm
                  </button>
                  <button
                    onClick={() => {
                      setIsFeatureModelOpen(false);
                    }}
                    className="bg-white text-purple-600 border border-purple-300 px-4 py-2 rounded hover:bg-[#C49DFF] hover:text-white transition duration-300"
                  >
                    Cancel
                  </button>
                </div>
              </div>
            </div>
          )}
        </div>
      )}
      {isAnimationPlaying && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-80">
          <section className="flex items-center justify-center min-h-screen">
            {isDecryptAnimation ? <DecryptAnimation /> : <DeleteAnimation />}
          </section>
        </div>
      )}
      {DecryptModal ? (
      <>
      <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-70">
              <div className="bg-white p-6 rounded-md shadow-md  max-w-2xl">
                <h2 className="text-xl font-semibold mb-4">Decrytion</h2>
                <div className="mt-5 text-s text-center md:text-left  py-4 text-[#181B1E]">
                  <span>
                    The password for {websiteName} is {password}.
                  </span>
                </div>
                <div className="flex space-x-4 mt-4">
                <button
                      onClick={() => {
                        setDecryptModal(false)
                      }}
                      className="bg-white text-purple-600 border border-purple-300 px-4 py-2 rounded hover:bg-[#C49DFF] hover:text-white transition duration-300"
                    >
                      Done
                    </button>
                </div>
              </div>
            </div>
      </>): (<></>)}
    </section>
  );
};

export default AddPassword;
