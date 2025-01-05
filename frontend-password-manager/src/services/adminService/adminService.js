import axios from "axios";

const API_URL = "http://localhost:8000";

export const CreateOTP = async () => {
  try {
    const headers = {
      "user-id": localStorage.getItem("userId"),
    };
    const response = await axios.post(API_URL + "/otp", null, { headers });
    return response.data;
  } catch (error) {
    console.log("eror while creating a otp for user: ", error);
    throw error;
  }
};

export const VerifyOTP = async (otp) => {
  try {
    const headers = {
      "user-id": localStorage.getItem("userId"),
    };
    const response = await axios.get(API_URL + `/verify?otp=${otp}`, {
      headers,
    });
    return response.data;
  } catch (error) {
    console.log("eror while verifying the otp for user: ", error);
    throw error;
  }
};

export const GetPlanDetails = async (id) => {
  try {
    const response = await axios.get(API_URL + `/plan?id=${id}`);
    return response.data;
  } catch (error) {
    console.log("eror while getting the plan details: ", error);
    throw error;
  }
};
