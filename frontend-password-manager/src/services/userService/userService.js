import axios from "axios";

const API_URL = "http://localhost:8000/user";

export const AddWebsite = async (userName, password, websiteName) => {
  try {
    let masterId = localStorage.getItem("masterId");
    let userId = localStorage.getItem("userId");
    const headers = {
      "master-id": masterId,
      "user-id": userId,
    };
    console.log("the request Body = ", userName, password, websiteName);
    const response = await axios.post(
      API_URL + `/addwebiste`,
      {
        user_name: userName,
        website_name: websiteName,
        password: password,
      },
      { headers }
    );
    return response.data;
  } catch (error) {
    console.error("Error while adding website:", error);
    throw error;
  }
};

export const FetchAllWesbite = async () => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
      "user-id": localStorage.getItem("userId"),
    };
    const response = await axios.get(API_URL + "/listWebiste", { headers });
    return response.data;
  } catch (error) {
    console.error("Error while fetching website details:", error);
    throw error;
  }
};

export const DeleteWebsite = async (websiteName) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
      "user-id": localStorage.getItem("userId"),
    };
    const response = await axios.delete(
      API_URL + `/password?webisteName=${websiteName}`,
      { headers }
    );
    return response.data;
  } catch (error) {
    console.error("Error while deleting website details:", error);
    throw error;
  }
};

export const VerifyAuthKey = async (authKey) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
      "user-id": localStorage.getItem("userId"),
    };
    const response = await axios.get(API_URL + `/key?key=${authKey}`, {
      headers,
    });
    return response.data;
  } catch (error) {
    console.error("Error while verifying the authKey:", error);
    throw error;
  }
};

export const DecryptPassword = async (webisteName) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
      "user-id": localStorage.getItem("userId"),
    };
    const response = await axios.post(
      API_URL + `/getPassword`,
      {
        website_name: webisteName,
      },
      { headers }
    );
    return response.data;
  } catch (error) {
    console.error("Error while decrypting the password:", error);
    throw error;
  }
};

export const UpdateKey = async (key, value) => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
      "user-id": localStorage.getItem("userId"),
    };
    const response = await axios.put(
      API_URL + `/passKey`,
      {
        type: key,
        value: value,
      },
      { headers }
    );
    return response.data;
  } catch (error) {
    console.error("Error while updating the password:", error);
    throw error;
  }
};

export const GetUserInfo = async () => {
  try {
    const headers = {
      "master-id": localStorage.getItem("masterId"),
      "user-id": localStorage.getItem("userId"),
    };
    const response = await axios.get(API_URL + `/getInfo`, { headers });
    return response.data;
  } catch (error) {
    console.error("Error while getting the user info:", error);
    throw error;
  }
};
