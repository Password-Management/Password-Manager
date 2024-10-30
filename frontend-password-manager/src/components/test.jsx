import React from "react";

const Test = () => {
  const features = [
    { name: "Add User", description: "Description for Feature 1" },
    { name: "Edit Key", description: "Description for Feature 2" },
    { name: "Update Encryption Type", description: "Description for Feature 3" },
    { name: "List Users", description: "Description for Feature 4" },
    { name: "Get Information", description: "Description for Feature 5" },
  ];

  return (
    <>
      <section className="flex flex-col items-center justify-center bg-[#EADCF0] min-h-screen pt-14 mt-5">
        <h1 className="text-4xl font-bold mb-8">Features</h1>

        {/* Feature Blocks */}
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6 w-full max-w-6xl px-4">
          {features.map((feature, index) => (
            <div
              key={index}
              className="bg-white p-6 rounded-lg shadow-md flex flex-col items-center justify-between"
            >
              <h2 className="text-2xl font-bold mb-4">{feature.name}</h2>
              <p className="text-gray-600 mb-6 text-center">
                {feature.description}
              </p>
              <button className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition">
                Start {feature.name}
              </button>
            </div>
          ))}
        </div>
      </section>
    </>
  );
};

export default Test;
