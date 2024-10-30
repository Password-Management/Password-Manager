import axios from "axios";

const API_URL = "http://localhost:8000";

export const GetAllUsers = async () => {
  try {
    let key = localStorage.getItem("specialKey");
    const response = await axios.get(API_URL + `/listUsers?specialKey=` + key);
    return response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const AddUser = async (name, email, password) => {
  try {
    let key = localStorage.getItem("specialKey");
    let masterId = "bc5c9780-9ca0-40fc-959a-3051dcbc3620";
    const response = await axios.post(API_URL + `/addUser`, {
      name: name,
      email: email,
      password: password,
      master_id: masterId,
      special_key: key,
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const EditAlgorithm = async (algorithm) => {
  try {
    let key = localStorage.getItem("specialKey");
    console.log("the request Body = ", algorithm);
    const resposne = await axios.patch(API_URL + `/algorithm`, {
      special_key: key,
      new_algorithm: algorithm,
    });
    return resposne.data;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const GetMastersInfo = async () => {
  try {
    let key = localStorage.getItem("specialKey");
    const response = await axios.get(API_URL + `/getInfo?specialKey=`+key)
    return response.data;
  } catch (error) {
    console.error("Error fetching users:", error);
    throw error;
  }
};

export const UpdateKey = async() => {
    
}