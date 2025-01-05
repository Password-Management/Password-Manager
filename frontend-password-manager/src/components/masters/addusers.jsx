import React, { useState, useEffect } from "react";
import {
  GetAllUsers,
  AddUser,
  GetUserIdByEmail,
  DeleteUser,
} from "../../services/masterService/masterservice";
import { AiOutlineUserAdd } from "react-icons/ai";
import { TbUserSearch } from "react-icons/tb";

const AddUsers = () => {
  const [showModal, setShowModal] = useState(false);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [isMaster, setIsMaster] = useState(false);
  const [users, setUsers] = useState([]);
  const [searchTerm, setSearchTerm] = useState("");

  // Fetch all users when the component mounts
  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await GetAllUsers();
        setUsers(response.result);
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    };
    fetchUsers();
  }, []);

  // Handle adding a new user
  const handleSubmit = async () => {
    try {
      await AddUser(name, email, isMaster);
      setShowModal(false);
      setName("");
      setEmail("");
      setIsMaster(false);
      const updatedUsers = await GetAllUsers();
      setUsers(updatedUsers.result);
    } catch (error) {
      console.error("Error adding user:", error);
    }
  };

  const handleDelete = async (indexToDelete, email) => {
    const resp = await GetUserIdByEmail(email);
    if (resp.status === "SUCCESS") {
      const deleteUserResponse = await DeleteUser(resp.result.user_id);
      console.log(deleteUserResponse);
    } else {
      console.log("fail");
    }

    const updatedUsers = users.filter((_, index) => index !== indexToDelete);
    setUsers(updatedUsers);
  };

  const filteredUsers = users.filter(
    (user) =>
      user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.email.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="min-h-[50vh] w-full flex flex-col items-center justify-center pt-16 px-4">
      <h1 className="text-2xl font-bold mb-6 text-center md:text-left">
        Add Users
      </h1>

      <div className="flex flex-col md:flex-row w-full max-w-2xl mb-4 space-y-4 md:space-y-0 md:space-x-4">
        <div className="relative flex-1">
          <TbUserSearch className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" />
          <input
            type="text"
            className="w-full py-3 pl-12 pr-4 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
            placeholder="Search Users"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
        <button
          onClick={() => setShowModal(true)}
          className="flex items-center justify-center w-full md:w-auto bg-[#AA7DFF] hover:bg-[#C49DFF] py-3 px-6 text-white rounded-md transition-colors"
        >
          <AiOutlineUserAdd className="mr-2" /> Add
        </button>
      </div>
      <div className="w-full max-w-2xl">
        {filteredUsers.length > 0 ? (
          filteredUsers.map((user, index) => (
            <div
              key={index}
              className="border p-4 mb-2 rounded-md shadow flex justify-between items-center"
            >
              <div>
                <p>
                  <strong>Name:</strong> {user.name}
                </p>
                <p>
                  <strong>Email:</strong> {user.email}
                </p>
              </div>
              <button
                onClick={() => handleDelete(index, user.email)}
                className="py-2 px-4 bg-[#FF8080] text-white rounded-md"
              >
                Delete
              </button>
            </div>
          ))
        ) : (
          <p className="text-gray-500">No users found</p>
        )}
      </div>

      {showModal && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white p-6 rounded-md shadow-md w-full max-w-2xl">
            <h2 className="text-xl font-semibold mb-4">Add New User</h2>
            <div className="space-y-4">
              <input
                type="text"
                className="w-full py-2 px-4 border rounded-md"
                placeholder="Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
              <input
                type="email"
                className="w-full py-2 px-4 border rounded-md"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
              <select
                className="w-full py-2 px-4 border rounded-md"
                value={isMaster}
                onChange={(e) => setIsMaster(e.target.value === "true")}
              >
                <option value="">Select Master Status</option>
                <option value="true">True</option>
                <option value="false">False</option>
              </select>
            </div>
            <div className="flex justify-end space-x-4 mt-4">
              <button
                onClick={() => setShowModal(false)}
                className="py-2 px-4 bg-gray-300 rounded-md"
              >
                Cancel
              </button>
              <button
                onClick={handleSubmit}
                className="py-2 px-4 bg-[#AA7DFF] hover:bg-[#C49DFF] text-white rounded-md"
              >
                Save
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default AddUsers;
