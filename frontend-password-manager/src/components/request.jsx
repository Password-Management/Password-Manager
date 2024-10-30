import React from "react";

const Request = () => {
  return (
    <>
      <section className="min-h-screen flex items-center justify-center pt-16">
        <div className="max-w-[500px] px-10 py-20 rounded-3xl bg-white border-2 border-gray-100">
          <h1 className="text-5xl font-semibold text-center">Welcome</h1>
          <p className="font-medium text-lg text-gray-500 mt-4 text-center">
          Please provide the following information to set up a temporary testing server for your product.
          </p>
          <div className="mt-8">
            <div className="flex flex-col">
              <label className="text-lg font-medium">Email</label>
              <input
                className="w-full border-2 border-gray-100 rounded-xl p-4 mt-1 bg-transparent"
                placeholder="Enter your Email"
              />
            </div>
            <div className="flex flex-col mt-4">
              <label className="text-lg font-medium">UserName</label>
              <input
                className="w-full border-2 border-gray-100 rounded-xl p-4 mt-1 bg-transparent"
                placeholder="Enter your username"
              />
            </div>
            <div className="mt-8 flex flex-col gap-y-4">
              <button className="active:scale-[.98] active:duration-75 transition-all hover:scale-[1.01]  ease-in-out transform py-4  rounded-xl text-white font-bold text-lg bg-[#AA7DFF] hover:bg-[#C49DFF]">
                Request a demo
              </button>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default Request;
