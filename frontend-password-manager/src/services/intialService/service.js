import axios from "axios";

const API_URL = "http://localhost:8001";

export const CreateProduct = async (type, name, email) => {
    console.log("the request Body = ", type);
  try {
    const response = await axios.post(API_URL+`/productType`, {
        name: name,
        email: email,
        product_type: type

    });
    return response.data;
  } catch (error) {
    console.error("Error while creating the product:", error);
    throw error;
  }
};
