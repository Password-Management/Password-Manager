import React, { useEffect, useState } from "react";
import { Player } from "@lottiefiles/react-lottie-player";
import { useNavigate } from "react-router-dom";
import * as assets from "../assets/Done.json";

const Reset = () => {
  const [isLoading, setLoading] = useState(true);
  const navigate = useNavigate();

  const path = localStorage.getItem("path");
  useEffect(() => {
    const timer = setTimeout(() => {
      setLoading(false);
      if (path == "logout") {
        localStorage.setItem("path", "");
        localStorage.removeItem("animationPlayed");
        localStorage.clear();
        navigate("/");
      } else if (path == "") {
        localStorage.setItem("path", "");
        navigate("/error");
      } else if (path == "resetPassword") {
        localStorage.setItem("path", "");
        navigate("/resetcreds");
      }
    }, 5000);
    return () => clearTimeout(timer);
  }, []);

  return (
    <>
      {isLoading ? (
        <>
          <Player
            autoplay
            loop
            src={assets}
            style={{
              height: "500px",
              width: "500px",
              margin: "auto",
              display: "block",
            }}
          />
          <div className="flex items-center justify-center">
            <span className="text-center text-xl text-gray-700">
              Redirecting you...
            </span>
          </div>
        </>
      ) : (
        <></>
      )}
    </>
  );
};

export default Reset;
