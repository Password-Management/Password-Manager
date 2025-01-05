import { React, useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
  EditAlgorithm,
  GetMastersInfo,
  UpdateKey
} from "../../services/masterService/masterservice";
import { LoaderAnimation } from "../animation";

const EditProductInfo = () => {
  let navigate = useNavigate();
  const [key, setKey] = useState("");
  const [algorithm, setAlgorithm] = useState("");
  const [isAnimating, setIsAnimating] = useState(false);
  const [count , setCount ] = useState(); 

  const handleKeyChange = (e) => {
    setKey(e.target.value);
  };

  useEffect(() => {
    const MastersAlgorithm = async () => {
      const algo = await GetMastersInfo();
      console.log("The database value = ", algo.result.algorithm);
      setAlgorithm(algo.result.algorithm); // Set algorithm from database as default
      setCount(algo.result.count);
    };
    MastersAlgorithm();
  }, []);

  const handleAlgorithmChange = (e) => {
    const selectedAlgorithm = e.target.value;
    setAlgorithm(selectedAlgorithm);
  };

  const handleUpdateAlgorithm = async () => {
    const resp = await EditAlgorithm(algorithm);
    console.log(resp.result);
  };

  const handleSubmitKey = async(e) => {
    e.preventDefault();
    console.log("Submitted Key: ", key);
    const response = await UpdateKey(key);
    if (response.status === "SUCCESS") {
      setIsAnimating(true);
      setTimeout(() => {
        localStorage.setItem("specialKey", key);
        window.location.reload();
        setIsAnimating(false);
      }, 5000);
    }
  };

  const navigateForgotPage = () => {
    navigate("/resetpassword");
  };
  const algorithms = ["RSA", "ASA"];

  return (
    <>
      <section className="min-h-screen flex flex-col items-center justify-center pt-16">
        <h2 className="text-2xl md:text-3xl lg:text-4xl font-bold mb-4">
          Configuration Options
        </h2>
        <div className="flex flex-col md:flex-row justify-center items-center gap-6 mt-10 p-4">
          <div className="border border-gray-300 rounded-lg shadow-lg p-6 w-full md:w-80 bg-white">
            <h2 className="text-xl font-bold mb-4">Edit Key</h2>
            {count > 0 ? (
              <div>
                <span>
                  You have already updated your special authorization key. If
                  you have forgotten it, please click on 'Forgot Authorization
                  Key' and follow the instructions.
                  <br />
                  <button
                    type="submit"
                    onClick={navigateForgotPage}
                    className="bg-purple-600 text-white w-1/2 py-2 rounded-lg hover:bg-purple-700 transition-colors mt-5"
                  >
                    Forgot AuthKey
                  </button>
                </span>
              </div>
            ) : (
              <>
                <p className="text-gray-600 mb-4">
                  Enter your SpecialKey below to update it.
                </p>
                <form onSubmit={handleSubmitKey}>
                  <input
                    type="text"
                    placeholder="Your Auth Key"
                    value={key}
                    onChange={handleKeyChange}
                    className="border border-gray-300 rounded-lg w-full p-2 mb-4 focus:outline-none focus:ring focus:ring-purple-200"
                  />
                  <button
                    type="submit"
                    className="bg-purple-600 text-white w-full py-2 rounded-lg hover:bg-purple-700 transition-colors"
                  >
                    Submit
                  </button>
                </form>
              </>
            )}
          </div>

          <div className="border border-gray-300 rounded-lg shadow-lg p-6 w-full md:w-80 bg-white">
            <h2 className="text-xl font-bold mb-4">Update Algorithm</h2>
            <p className="text-gray-600 mb-4">
              Select the encryption algorithm you want to use.
            </p>
            <form onSubmit={handleUpdateAlgorithm}>
              <select
                value={algorithm}
                onChange={handleAlgorithmChange}
                className="border border-gray-300 rounded-lg w-full p-2 mb-4 focus:outline-none focus:ring focus:ring-purple-200"
              >
                <option disabled>{"Update the Algorithm"}</option>
                {algorithms.map((algo) => (
                  <option key={algo} value={algo}>
                    {algo === algorithm ? `${algo}` : algo}
                  </option>
                ))}
              </select>
              <button
                type="submit"
                className="bg-purple-600 text-white w-full py-2 rounded-lg hover:bg-purple-700 transition-colors"
              >
                Update Algorithm
              </button>
            </form>
          </div>
        </div>
        {isAnimating ? (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-80">
          <section className="flex items-center justify-center min-h-screen">
            { <LoaderAnimation /> }
          </section>
        </div>
      ): (<></>)}
      </section>
    </>
  );
};

export default EditProductInfo;
