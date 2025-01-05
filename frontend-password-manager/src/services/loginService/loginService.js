import axios from "axios";

const API_URL = "http://localhost:8000/login";

export const UserLogin = async (email, password) => {
  try {
    const response = await axios.post(API_URL + `/user`, {
      email: email,
      password: password,
    });
    return response.data;
  } catch (error) {
    console.log("error while logging the user");
    console.log(error);
  }
};

export const UserLogout = async (userId) => {
  try {
    const response = await axios.put(API_URL + `/logout?id=${userId}`);
    return response.data;
  } catch (error) {
    console.log("error while logging out the user");
    throw error;
  }
};

export const MasterLogin = async (key) => {
  try {
    const response = await axios.post(API_URL + `/master`, {
      special_key: key,
    });
    return response.data;
  } catch (error) {
    console.log("error while logging in the master");
    throw error;
  }
};
