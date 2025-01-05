import axios from "axios";

const API_URL = "http://localhost:8000/master";

export const GetAllUsers = async () => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
    };
    let key = localStorage.getItem("specialKey");
    const response = await axios.get(API_URL + `/listUsers?specialKey=` + key, {
      headers,
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const AddUser = async (name, email, isMaster) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
    };
    const response = await axios.post(
      API_URL + `/addUser`,
      {
        name: name,
        email: email,
        is_master: isMaster,
      },
      { headers }
    );
    return response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const EditAlgorithm = async (algorithm) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
    };
    let key = localStorage.getItem("specialKey");
    console.log("the request Body = ", algorithm);
    const resposne = await axios.patch(
      API_URL + `/algorithm`,
      {
        special_key: key,
        new_algorithm: algorithm,
      },
      { headers }
    );
    return resposne.data;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const GetMastersInfo = async () => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
    };
    const response = await axios.get(API_URL + `/getInfo`, {
      headers,
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const UpdateKey = async (key) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
    };
    const response = await axios.post(
      API_URL + `/editKey`,
      {
        special_key: localStorage.getItem("specialKey"),
        new_key: key,
      },
      { headers }
    );
    return response.data;
  } catch (error) {
    console.error("Error while updating the key:", error);
    throw error;
  }
};

export const GetUserIdByEmail = async (email) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
    };
    const response = await axios.get(API_URL + `/userbyId?email=${email}`, {
      headers,
    });
    return response.data;
  } catch (error) {
    console.error("Error while getting the user details by the email:", error);
    throw error;
  }
};

export const DeleteUser = async (id) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
    };
    const response = await axios.delete(API_URL + `/user?id=${id}`, {headers})
    return response.data
  } catch (error) {
    console.error("Error while deleting the user: ", error)
    throw error;
  }
}