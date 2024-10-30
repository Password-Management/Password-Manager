import React, { useState } from "react";

const Banner = ({ message, type }) => {
  const [showBanner, setShowBanner] = useState(true);

  const handleClose = () => {
    setShowBanner(false);
  };

  return (
    <div
      className={`flex items-center justify-between p-4 rounded-lg ${
        type === "success"
          ? "bg-green-200 text-green-800"
          : "bg-red-200 text-red-800"
      } mb-4 ${showBanner ? '' : 'hidden'}`}
    >
      <p className="flex-grow">{message}</p>
      <button className="text-gray-500 hover:text-gray-700" onClick={handleClose}>
        X
      </button>
    </div>
  );
};

export default Banner;